package validator

// validatePoint validates coordinates for Point type
func validatePoint(coordinates interface{}) error {
	return validateCoordinates(coordinates)
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
