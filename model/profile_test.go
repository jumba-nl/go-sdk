package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProfile_Name(t *testing.T) {
	tests := []struct {
		input  *Profile
		output string
	}{
		{input: &Profile{Firstname: "A", Infix: "van der", Lastname: "A"}, output: "A van der A"},
		{input: &Profile{Firstname: "Jan", Infix: "", Lastname: "Brouwers"}, output: "Jan Brouwers"},
		{input: &Profile{Firstname: "\t Jan\n", Infix: "", Lastname: "Brouwers\n"}, output: "\t Jan\n Brouwers\n"},
		{input: &Profile{Firstname: "0000x01", Infix: "0000x01", Lastname: "0000x01"}, output: "0000x01 0000x01 0000x01"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			assert.New(t).Equal(test.output, test.input.Name(), "should be equal")
		})
	}
}
