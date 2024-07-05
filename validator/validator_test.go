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

	invalidPoint := `{
        "type": "Feature",
        "geometry": {
            "type": "Point",
            "coordinates": [100.0]
        },
        "properties": {}
    }`

	err = ValidateGeoJSON([]byte(invalidPoint))
	assert.Error(t, err)
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

	err := ValidateGeoJSON([]byte(validMultiPoint))
	assert.Nil(t, err)

	// invalidMultiPoint := `{
	// 		"type": "Feature",
	// 		"geometry": {
	// 				"type": "MultiPoint",
	// 				"coordinates": [
	// 						[100.0, 0.0]
	// 				]
	// 		},
	// 		"properties": {}
	// }`

	// err = ValidateGeoJSON([]byte(invalidMultiPoint));
	// assert.Error(t, err)
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

	err := ValidateGeoJSON([]byte(validLineString))
	assert.Nil(t, err)

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

	err = ValidateGeoJSON([]byte(invalidLineString))
	assert.Error(t, err)
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

	err := ValidateGeoJSON([]byte(validMultiLineString))
	assert.Nil(t, err)

	// invalidMultiLineString := `{
	//     "type": "Feature",
	//     "geometry": {
	//         "type": "MultiLineString",
	//         "coordinates": [
	//             [
	//                 [100.0, 0.0]
	//             ]
	//         ]
	//     },
	//     "properties": {}
	// }`

	// err = ValidateGeoJSON([]byte(invalidMultiLineString));
	// assert.Error(t, err)
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

	err := ValidateGeoJSON([]byte(validPolygon))
	assert.Nil(t, err)

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
	err = ValidateGeoJSON([]byte(invalidPolygon))
	assert.Error(t, err)
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

	err := ValidateGeoJSON([]byte(validMultiPolygon))
	assert.Nil(t, err)

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

	err = ValidateGeoJSON([]byte(invalidMultiPolygon))
	assert.Error(t, err)
}
