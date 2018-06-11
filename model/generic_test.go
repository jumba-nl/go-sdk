package model

import "testing"

func TestFormatIdentifier(t *testing.T) {
	tests := []struct {
		input  string
		output string
		ok     bool
	}{
		{
			input:  "4641Sp2",
			output: "4641SP2",
			ok:     true,
		},
		{
			input:  "4641Sp 2",
			output: "4641SP2",
			ok:     true,
		},
		{
			input:  "4641sp2",
			output: "4641SP2",
			ok:     true,
		},
		{
			input:  "4641Sp2-a",
			output: "4641SP2-a",
			ok:     true,
		},
		{
			input:  "4641Kl 2 -aB",
			output: "4641KL2-aB",
			ok:     true,
		},
		{
			input:  "4641Kl 2a",
			output: "4641KL2a",
			ok:     true,
		},
		{
			input:  "benk",
			output: "benk",
			ok:     false,
		},
		{
			input:  "0000xx",
			output: "0000xx",
			ok:     false,
		},
		{
			input:  "\t  46  41 Sp2-a\n",
			output: "4641SP2-a",
			ok:     true,
		},
		{
			input:  "0000x0",
			output: "0000x0",
			ok:     false,
		},
		{
			input:  "0000x01",
			output: "0000X01",
			ok:     false,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			output, ok := FormatIdentifier(test.input)
			if ok != test.ok {
				t.Errorf("expected [%v], got [%v] instead", test.ok, ok)
			}
			if output != test.output {
				t.Errorf("expected [%v], got [%v] instead", test.output, output)
			}
		})
	}
}
