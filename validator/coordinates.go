package validator

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// validateCoordinates validates coordinates for a given dimension (2 or 3)
func validateCoordinates(coordinates interface{}) error {
	coords, ok := coordinates.([]interface{})
	if !ok || (len(coords) != 2 && len(coords) != 3) {
		return ErrInvalidCoordinates
	}
	for _, c := range coords {
		switch c.(type) {
		case float64, float32, int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64, json.Number:
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
