package validator

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// GeoJSON represents the main structure of GeoJSON data
type GeoJSON struct {
	Type       string      `json:"type"`
	Geometry   *Geometry   `json:"geometry,omitempty"`
	Properties interface{} `json:"properties,omitempty"`
}

// Geometry represents the geometry of a GeoJSON object
type Geometry struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates"`
	Geometries  []*Geometry `json:"geometries,omitempty"` // For GeometryCollection
}

// Error messages
var (
	ErrInvalidType         = errors.New("invalid type")
	ErrInvalidCoordinates  = errors.New("invalid coordinates")
	ErrInvalidGeometryType = errors.New("invalid geometry type")
	ErrInvalidPolygon      = errors.New("invalid polygon coordinates")
)

// ValidateGeoJSON validates the GeoJSON data
func ValidateGeoJSON(data []byte) error {
	var geo GeoJSON
	if err := json.Unmarshal(data, &geo); err != nil {
		return err
	}
	return validateGeoJSONObject(&geo)
}

// validateGeoJSONObject validates a GeoJSON object
func validateGeoJSONObject(geo *GeoJSON) error {
	switch geo.Type {
	case "Feature":
		return validateFeature(geo)
	case "FeatureCollection":
		return validateFeatureCollection(geo)
	default:
		return validateGeometry(geo.Geometry)
	}
}

// validateFeature validates a GeoJSON Feature
func validateFeature(geo *GeoJSON) error {
	if geo.Geometry == nil {
		return errors.New("feature must have a geometry")
	}
	return validateGeometry(geo.Geometry)
}

// validateFeatureCollection validates a GeoJSON FeatureCollection
func validateFeatureCollection(geo *GeoJSON) error {
	features, ok := geo.Properties.([]interface{})
	if !ok {
		return errors.New("feature collection properties must be an array of features")
	}
	for _, f := range features {
		featureMap, ok := f.(map[string]interface{})
		if !ok {
			return errors.New("each feature must be an object")
		}
		featureJSON, err := json.Marshal(featureMap)
		if err != nil {
			return err
		}
		if err := ValidateGeoJSON(featureJSON); err != nil {
			return err
		}
	}
	return nil
}

// validateGeometry validates a GeoJSON Geometry
func validateGeometry(geo *Geometry) error {
	if geo == nil {
		return errors.New("geometry cannot be nil")
	}
	switch geo.Type {
	case "Point":
		return validatePoint(geo.Coordinates)
	case "MultiPoint":
		return validateMultiPoint(geo.Coordinates)
	case "LineString":
		return validateLineString(geo.Coordinates)
	case "MultiLineString":
		return validateMultiLineString(geo.Coordinates)
	case "Polygon":
		return validatePolygon(geo.Coordinates)
	case "MultiPolygon":
		return validateMultiPolygon(geo.Coordinates)
	case "GeometryCollection":
		if geo.Geometries == nil {
			return errors.New("geometry collection must have geometries")
		}
		for _, g := range geo.Geometries {
			if err := validateGeometry(g); err != nil {
				return err
			}
		}
	default:
		return ErrInvalidGeometryType
	}
	return nil
}

// validatePoint validates coordinates for Point type
func validatePoint(coordinates interface{}) error {
	return validateCoordinates(coordinates, 2)
}

// validateMultiPoint validates coordinates for MultiPoint type
func validateMultiPoint(coordinates interface{}) error {
	coords, ok := coordinates.([]interface{})
	if !ok {
		return ErrInvalidCoordinates
	}
	for _, point := range coords {
		if err := validatePoint(point); err != nil {
			return err
		}
	}
	return nil
}

// validateLineString validates the coordinates for a LineString type
func validateLineString(coordinates interface{}) error {
	coords, ok := coordinates.([]interface{})
	if !ok || len(coords) < 2 {
		return ErrInvalidCoordinates
	}
	for _, c := range coords {
		if err := validatePoint(c); err != nil {
			return err
		}
	}
	return nil
}

// validateMultiLineString validates coordinates for MultiLineString type
func validateMultiLineString(coordinates interface{}) error {
	coords, ok := coordinates.([]interface{})
	if !ok {
		return ErrInvalidCoordinates
	}
	for _, lineString := range coords {
		if err := validateLineString(lineString); err != nil {
			return err
		}
	}
	return nil
}

// validatePolygon validates the coordinates for a Polygon type
func validatePolygon(coordinates interface{}) error {
	coords, ok := coordinates.([]interface{})
	if !ok {
		return ErrInvalidCoordinates
	}
	for _, ring := range coords {
		ringCoords, ok := ring.([]interface{})
		if !ok || len(ringCoords) < 4 {
			return ErrInvalidPolygon
		}
		// The first and last position must be the same
		if !equalCoordinates(ringCoords[0], ringCoords[len(ringCoords)-1]) {
			return ErrInvalidPolygon
		}
		for _, c := range ringCoords {
			if err := validatePoint(c); err != nil {
				return err
			}
		}
	}
	return nil
}

// validateMultiPolygon validates the coordinates for a MultiPolygon type
func validateMultiPolygon(coordinates interface{}) error {
	coords, ok := coordinates.([]interface{})
	if !ok {
		return ErrInvalidCoordinates
	}
	for _, polygon := range coords {
		if err := validatePolygon(polygon); err != nil {
			return err
		}
	}
	return nil
}

// validateCoordinates validates coordinates for a given dimension
func validateCoordinates(coordinates interface{}, dim int) error {
	coords, ok := coordinates.([]interface{})
	if !ok || len(coords) != dim {
		return ErrInvalidCoordinates
	}
	for _, c := range coords {
		switch c.(type) {
		case float64, float32, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, json.Number:
			continue
		case string:
			if _, err := strconv.ParseFloat(fmt.Sprintf("%v", c), 64); err != nil {
				return ErrInvalidCoordinates
			}
		default:
			return ErrInvalidCoordinates
		}
	}
	return nil
}

// equalCoordinates checks if two coordinates are equal
func equalCoordinates(c1, c2 interface{}) bool {
	coords1, ok1 := c1.([]interface{})
	coords2, ok2 := c2.([]interface{})
	if !ok1 || !ok2 || len(coords1) != len(coords2) {
		return false
	}
	for i := range coords1 {
		if coords1[i] != coords2[i] {
			return false
		}
	}
	return true
}
