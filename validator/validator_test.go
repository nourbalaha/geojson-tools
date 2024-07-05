package validator

import "testing"

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

    if err := ValidateGeoJSON([]byte(validPolygon)); err != nil {
        t.Errorf("Expected no error, got %v", err)
    }

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

    if err := ValidateGeoJSON([]byte(invalidPolygon)); err == nil {
        t.Errorf("Expected error, got nil")
    }
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

    if err := ValidateGeoJSON([]byte(validLineString)); err != nil {
        t.Errorf("Expected no error, got %v", err)
    }

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

    if err := ValidateGeoJSON([]byte(invalidLineString)); err == nil {
        t.Errorf("Expected error, got nil")
    }
}

func TestValidatePoint(t *testing.T) {
    validPoint := `{
        "type": "Feature",
        "geometry": {
            "type": "Point",
            "coordinates": [100.0, 0.0]
        },
        "properties": {}
    }`

    if err := ValidateGeoJSON([]byte(validPoint)); err != nil {
        t.Errorf("Expected no error, got %v", err)
    }

    invalidPoint := `{
        "type": "Feature",
        "geometry": {
            "type": "Point",
            "coordinates": [100.0]
        },
        "properties": {}
    }`

    if err := ValidateGeoJSON([]byte(invalidPoint)); err == nil {
        t.Errorf("Expected error, got nil")
    }
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

	if err := ValidateGeoJSON([]byte(validMultiPoint)); err != nil {
			t.Errorf("Expected no error, got %v", err)
	}

	invalidMultiPoint := `{
			"type": "Feature",
			"geometry": {
					"type": "MultiPoint",
					"coordinates": [
							[100.0, 0.0]
					]
			},
			"properties": {}
	}`

	if err := ValidateGeoJSON([]byte(invalidMultiPoint)); err == nil {
			t.Errorf("Expected error, got nil")
	}
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

    if err := ValidateGeoJSON([]byte(validMultiLineString)); err != nil {
        t.Errorf("Expected no error, got %v", err)
    }

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

    if err := ValidateGeoJSON([]byte(invalidMultiLineString)); err == nil {
        t.Errorf("Expected error, got nil")
    }
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

    if err := ValidateGeoJSON([]byte(validMultiPolygon)); err != nil {
        t.Errorf("Expected no error, got %v", err)
    }

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

    if err := ValidateGeoJSON([]byte(invalidMultiPolygon)); err == nil {
        t.Errorf("Expected error, got nil")
    }
}
