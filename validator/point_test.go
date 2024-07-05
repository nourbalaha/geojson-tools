package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePoint(t *testing.T) {
	validPoint := `{
        "type": "Feature",
        "geometry": {
            "type": "Point",
            "coordinates": [100.0, 0.0]
        },
        "properties": {}
    }`
	err := ValidateGeoJSON([]byte(validPoint))
	assert.Nil(t, err)

	validPoint2 := `{
        "type": "Feature",
        "geometry": {
            "type": "Point",
            "coordinates": ["100.0", "0.0"]
        },
        "properties": {}
    }`
	err1 := ValidateGeoJSON([]byte(validPoint2))
	assert.Nil(t, err1)

	invalidPoint := `{
        "type": "Feature",
        "geometry": {
            "type": "Point",
            "coordinates": [100.0]
        },
        "properties": {}
    }`
	err2 := ValidateGeoJSON([]byte(invalidPoint))
	assert.Error(t, err2)

	invalidPoint2 := `{
        "type": "Feature",
        "geometry": {
            "type": "Point",
            "coordinates": ["foo", "bar"]
        },
        "properties": {}
    }`
	err3 := ValidateGeoJSON([]byte(invalidPoint2))
	assert.Error(t, err3)
}

func TestValidateMultiPoint(t *testing.T) {
	validMultiPoint := `{
			"type": "Feature",
			"geometry": {
					"type": "MultiPoint",
					"coordinates": [
							[100.0, 0.0],
							[101.0, 1.0]
					]
			},
			"properties": {}
	}`
	err1 := ValidateGeoJSON([]byte(validMultiPoint))
	assert.Nil(t, err1)

	validMultiPoint2 := `{
			"type": "Feature",
			"geometry": {
					"type": "MultiPoint",
					"coordinates": [
							["100.0", "0.0"],
							["101.0", "1.0"]
					]
			},
			"properties": {}
	}`
	err2 := ValidateGeoJSON([]byte(validMultiPoint2))
	assert.Nil(t, err2)

	invalidMultiPoint := `{
			"type": "Feature",
			"geometry": {
					"type": "MultiPoint",
					"coordinates": [
							["foo", "bar"],
							["foo", "bar"]
					]
			},
			"properties": {}
	}`
	err3 := ValidateGeoJSON([]byte(invalidMultiPoint))
	assert.Error(t, err3)
}