package main

import (
	"utest/unittest"
	"testing"
)

type squareTestCase struct {
	n, want float64
}

type divideTestCase struct {
	x, y, want float64
	errExpected bool `default: false`
}

type additiobTestCase struct {
	x, y, want float64
}

type subtractionTestCase struct {
	x, y, want float64
}

func TestSquare(t *testing.T) {
	t.Parallel()

	var testCases = []squareTestCase {
		{ n: 0, want: 0 },
		{ n: 2.0, want: 4.0 },
		{ n: -2.0, want: 4.0 },
	}

	for _, tc := range testCases {
		got := unittest.Square(tc.n)
		if tc.want != got {
			t.Errorf("Square(%f): want %f, got %f", tc.n, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	var testCases = []divideTestCase {
		{ x: 2.0, y: 1.0, want: 2.0, errExpected: false, },
		{ x: 2.0, y: 2.0, want: 1.0, errExpected: false, },
		{ x: 4.0, y: 2.0, want: 2.0, errExpected: false, },
		// { x: 4.0, y: 0.0, want: 0.0, errExpected: true, },
	}

	for _, tc := range testCases {
		got, err := unittest.Divide(tc.x, tc.y)
		if !tc.errExpected != (err != nil) && tc.want != got {
			t.Fatalf("Divide(%f, %f): unexpected error status: %v", tc.x, tc.y, err != nil)
		}
	}
}

func TestAddition(t *testing.T) {
	t.Parallel()

	var testCases = []additiobTestCase {
		{ x: 10.0, y: 20.0, want: 30.0 },
		{ x: 30.0, y: 20.0, want: 50.0 },
		{ x: 10.0, y: 2.0, want: 12.0 },
	}

	for _, tc := range testCases {
		got := unittest.Addition(tc.x, tc.y)
		if tc.want != got {
			t.Errorf("Addition(%f, %f): want %f, got %f", tc.x, tc.y, tc.want, got)
		}
	}
}

func TestSubtraction(t *testing.T) {
	t.Parallel()

	var testCases = []subtractionTestCase {
		{ x: 30.0, y: 20.0, want: 10.0 },
		{ x: 30.0, y: 10.0, want: 20.0 },
		{ x: 10.0, y: 2.0, want: 8.0 },
	}

	for _, tc := range testCases {
		got := unittest.Subtraction(tc.x, tc.y)
		if tc.want != got {
			t.Errorf("Addition(%f, %f): want %f, got %f", tc.x, tc.y, tc.want, got)
		}
	}
}



