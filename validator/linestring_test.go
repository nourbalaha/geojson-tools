package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateLineString(t *testing.T) {
	validLineString := `{
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
	err1 := ValidateGeoJSON([]byte(validLineString))
	assert.Nil(t, err1)

	invalidLineString := `{
        "type": "Feature",
        "geometry": {
            "type": "LineString",
            "coordinates": [
                [100.0, 0.0]
            ]
        },
        "properties": {}
    }`
	err2 := ValidateGeoJSON([]byte(invalidLineString))
	assert.Error(t, err2)
}

func TestValidateMultiLineString(t *testing.T) {
	validMultiLineString := `{
        "type": "Feature",
        "geometry": {
            "type": "MultiLineString",
            "coordinates": [
                [
                    [100.0, 0.0],
                    [101.0, 1.0]
                ],
                [
                    [102.0, 2.0],
                    [103.0, 3.0]
                ]
            ]
        },
        "properties": {}
    }`
	err1 := ValidateGeoJSON([]byte(validMultiLineString))
	assert.Nil(t, err1)

	invalidMultiLineString := `{
	    "type": "Feature",
	    "geometry": {
	        "type": "MultiLineString",
	        "coordinates": [
	            [
	                [100.0, 0.0]
	            ]
	        ]
	    },
	    "properties": {}
	}`

	err2 := ValidateGeoJSON([]byte(invalidMultiLineString))
	assert.Error(t, err2)
}
