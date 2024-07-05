package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestValidateCoordinates tests the validateCoordinates function
func TestValidateCoordinates(t *testing.T) {
	tests := []struct {
		name         string
		coordinates  interface{}
		expectedErr  error
	}{
		{
			name:        "Valid 2D coordinates",
			coordinates: []interface{}{1.0, 2.0},
			expectedErr: nil,
		},
		{
			name:        "Valid 3D coordinates",
			coordinates: []interface{}{1.0, 2.0, 3.0},
			expectedErr: nil,
		},
		{
			name:        "Invalid coordinates (not an array)",
			coordinates: "invalid",
			expectedErr: ErrInvalidCoordinates,
		},
		{
			name:        "Invalid coordinates (wrong number of dimensions)",
			coordinates: []interface{}{1.0},
			expectedErr: ErrInvalidCoordinates,
		},
		{
			name:        "Invalid coordinates (non-numeric type)",
			coordinates: []interface{}{1.0, "invalid"},
			expectedErr: ErrInvalidCoordinates,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateCoordinates(tt.coordinates)
			if err != tt.expectedErr {
				t.Errorf("validateCoordinates() error = %v, expectedErr %v", err, tt.expectedErr)
			}
		})
	}
}

// TestEqualCoordinates tests the equalCoordinates function
func TestEqualCoordinates(t *testing.T) {
	tests := []struct {
		name     string
		c1       interface{}
		c2       interface{}
		expected bool
	}{
		{
			name:     "Equal coordinates",
			c1:       []interface{}{1.0, 2.0},
			c2:       []interface{}{1.0, 2.0},
			expected: true,
		},
		{
			name:     "Unequal coordinates",
			c1:       []interface{}{1.0, 2.0},
			c2:       []interface{}{2.0, 1.0},
			expected: false,
		},
		{
			name:     "Unequal length coordinates",
			c1:       []interface{}{1.0, 2.0},
			c2:       []interface{}{1.0, 2.0, 3.0},
			expected: false,
		},
		{
			name:     "Non-array coordinates",
			c1:       "invalid",
			c2:       []interface{}{1.0, 2.0},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := equalCoordinates(tt.c1, tt.c2)
			assert.Equal(t, tt.expected, result)
		})
	}
}
