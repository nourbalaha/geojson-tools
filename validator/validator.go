package validator

import (
	"encoding/json"
	"errors"
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