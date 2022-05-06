package main

import "testing"

func TestHello(t *testing.T) {
	tt := []struct {
		input string
		want  string
	}{
		{
			input: "world",
			want:  "Hello, world",
		},
		{
			input: "Chris",
			want:  "Hello, Chris",
		},
		{
			input: "",
			want:  "Hello",
		},
	}

	for _, test := range tt {
		got := Hello(test.input)
		if got != test.want {
			t.Errorf("Got %s, wanted %s", got, test.want)
		}
	}
}
