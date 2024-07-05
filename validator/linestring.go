package validator

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
