package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestValidatePolygon(t *testing.T) {
	validPolygon := `{
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
	err1 := ValidateGeoJSON([]byte(validPolygon))
	assert.Nil(t, err1)

	invalidPolygon := `{
        "type": "Feature",
        "geometry": {
            "type": "Polygon",
            "coordinates": [
                [
                    [100.0, 0.0],
                    [101.0, 0.0],
                    [101.0, 1.0],
                    [100.0, 1.0]
                ]
            ]
        },
        "properties": {}
    }`
	err2 := ValidateGeoJSON([]byte(invalidPolygon))
	assert.Error(t, err2)
}

func TestValidateMultiPolygon(t *testing.T) {
	validMultiPolygon := `{
        "type": "Feature",
        "geometry": {
            "type": "MultiPolygon",
            "coordinates": [
                [
                    [
                        [102.0, 2.0],
                        [103.0, 2.0],
                        [103.0, 3.0],
                        [102.0, 3.0],
                        [102.0, 2.0]
                    ]
                ],
                [
                    [
                        [100.0, 0.0],
                        [101.0, 0.0],
                        [101.0, 1.0],
                        [100.0, 1.0],
                        [100.0, 0.0]
                    ]
                ]
            ]
        },
        "properties": {}
    }`
	err1 := ValidateGeoJSON([]byte(validMultiPolygon))
	assert.Nil(t, err1)

	invalidMultiPolygon := `{
        "type": "Feature",
        "geometry": {
            "type": "MultiPolygon",
            "coordinates": [
                [
                    [
                        [102.0, 2.0],
                        [103.0, 2.0],
                        [103.0, 3.0],
                        [102.0, 3.0]
                    ]
                ]
            ]
        },
        "properties": {}
    }`
	err2 := ValidateGeoJSON([]byte(invalidMultiPolygon))
	assert.Error(t, err2)
}
