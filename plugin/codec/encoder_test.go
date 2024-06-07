package codec

import (
	"testing"
)

type testCase struct {
	name string
	val  any
	want any
}

var testCases = []testCase{}

func TestEncoderEncode(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

		})
	}
}
