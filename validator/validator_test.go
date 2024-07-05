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
