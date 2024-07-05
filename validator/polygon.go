package validator

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
