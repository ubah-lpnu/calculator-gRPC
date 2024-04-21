package calculate

import (
	"testing"
)

func TestMultiplyAndDivide(t *testing.T) {
	testCases := []struct {
		a, b        float64
		expectedMul float64
		expectedDiv float64
		expectError bool
	}{
		{4, 2, 8.0, 2.0, false},
		{5, 0.0, 0.0, 0.0, true},
	}

	for _, tc := range testCases {

		result, err := MultiplyAndDivide(tc.a, tc.b)

		if (err != nil) != tc.expectError {
			t.Errorf("Test failed for case: a=%f, b=%f, error: %v", tc.a, tc.b, err)
			continue
		}

		if err != nil {
			continue
		}

		if result.ResultMul != tc.expectedMul {
			t.Errorf("Multiplication result incorrect for case: a=%f, b=%f, got: %f, expected: %f", tc.a, tc.b, result.ResultMul, tc.expectedMul)
		}

		if result.ResultDiv != tc.expectedDiv {
			t.Errorf("Division result incorrect for case: a=%f, b=%f, got: %f, expected: %f", tc.a, tc.b, result.ResultDiv, tc.expectedDiv)
		}
	}
}
