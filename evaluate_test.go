package main

import (
	"evaluator"
	"fmt"
	"testing"
)

func TestEvaluateComplex(t *testing.T) {
	tests := []struct {
		input    string
		expected any
		err      error
	}{
		// Complex Expressions
		{"(2 + 3) * (7 - 4) ^ 2", float64(45), nil},                     // (2 + 3) * (3 ^ 2) = 5 * 9 = 45
		{"(4 + 2 * 3) / (1 + 2) ^ 2", float64(1.1111111111111112), nil}, // (4 + 6) / 9 = 10 / 9 ≈ 1.11
		{"-5 ^ (2 + 3)", float64(-3125), nil},                           // -5 ^ 5 = -3125
		{"3 * (2 ^ 3) - 4 / (1 + 1)", float64(22), nil},                 // 3 * 8 - 4 / 2 = 24 - 2 = 22
		{"(7 - (2 + 1)) * 3 ^ (2 - 1)", float64(12), nil},               // (7 - 3) * 3 ^ 1 = 4 * 3 = 12
		{"(1.5 + 2) * (4 / 2.0) ^ (1 + 1)", float64(14.0), nil},         // (3.5 * (2 ^ 2) = 3.5 * 4 = 14.0
		{"(2 * (3 + 4) ^ 2 - 5) / 3", float64(31), nil},                 // (2 * 49 - 5) / 3 = 93 / 3 = 31
		{"((10 - 3) / (2 + 1)) ^ 2", float64(49), nil},                  // ((7 / 3) ^ 2 = 2 ^ 2 = 4
		{"-2 ^ (3 - 1) * (4 + 1)", float64(-10), nil},                   // -2 ^ 2 * 5 = -4 * 5 = -20
		{"(2 + 2) * ((3 - 1) ^ 2 - 1)", float64(8), nil},                // (4 * (4 - 1)) = 4 * 3 = 12
		{"1.5 * (3 + (2 - 1) ^ 2)", float64(9.0), nil},                  // 1.5 * (3 + 1) = 1.5 * 4 = 6.0

		// Edge Cases
		{"0 ^ 0", int64(1), nil},             // Typically, 0 ^ 0 is defined as 1
		{"1 ^ 0", int64(1), nil},             // 1 ^ 0 = 1
		{"(1 + 2) ^ 0.5", float64(3.0), nil}, // (1 + 2) ^ 0.5 = 3 ^ 0.5 = √3 ≈ 1.732
		{"(-2) ^ 2", int64(4), nil},          // (-2) ^ 2 = 4
		{"-2 ^ 2", int64(-4), nil},           // -2 ^ 2 = -4 (Note: unary minus is applied after exponentiation)

		// Invalid Cases
		{"(2 + 3", nil, fmt.Errorf("errUnmatchedParenthesis")},     // Unmatched parenthesis
		{"2 / 0", nil, fmt.Errorf("errDivisionByZero")},            // Division by zero
		{"2 ^ (2 ^ 2", nil, fmt.Errorf("errUnmatchedParenthesis")}, // Unmatched parenthesis
		{"(2 + 3 ^ 2", nil, fmt.Errorf("errUnmatchedParenthesis")}, // Unmatched parenthesis
		{"2 ^ 2 ^ 3 ^", nil, fmt.Errorf("errInvalidExpression")},   // Extra operator
		{"1 + abc", nil, fmt.Errorf("errInvalidCharacter")},        // Invalid character
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := evaluator.Evaluate(tt.input)
			if err != nil && tt.err == nil {
				t.Fatalf("expected error %v, got %v", tt.err, err)
			}
			if err == nil && tt.err != nil {
				t.Fatal("expected an error, but got none")
			}
			if err == nil && result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
