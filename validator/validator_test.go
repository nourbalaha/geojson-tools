package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateGeoJSON(t *testing.T) {
	validGeoJSON := `{
        "type": "Feature",
        "geometry": {
            "type": "Polygon",
            "coordinates": [
                [
                    [100.0, 0.0],
                    [101.0, 0.0],
                    [101.0, 1.0],
                    [100.0, 1.0],
                    [100.0, 0.0]
                ]
            ]
        },
        "properties": {}
    }`
	err1 := ValidateGeoJSON([]byte(validGeoJSON))
	assert.Nil(t, err1)

	validGeoJSON2 := `{
        "type": "Feature",
        "geometry": {
            "type": "LineString",
            "coordinates": [
                [100.0, 0.0],
                [101.0, 1.0]
            ]
        },
        "properties": {}
    }`
	err2 := ValidateGeoJSON([]byte(validGeoJSON2))
	assert.Nil(t, err2)

	validGeoJSON3 := `{
        "type": "Feature",
        "geometry": {
            "type": "Point",
            "coordinates": [100.0, 0.0]
        },
        "properties": {}
    }`
	err3 := ValidateGeoJSON([]byte(validGeoJSON3))
	assert.Nil(t, err3)

	invalidGeoJSON := "{}"
	err4 := ValidateGeoJSON([]byte(invalidGeoJSON))
	assert.Error(t, err4)

	invalidGeoJSON2 := "invalid"
	err5 := ValidateGeoJSON([]byte(invalidGeoJSON2))
	assert.Error(t, err5)
}