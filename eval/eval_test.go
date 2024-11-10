package eval

import (
	"go_type_inference/object"
	"go_type_inference/parser"
	"testing"
)

func TestEval(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input    string
		expected string
	}{
		"identifier": {
			input:    "x",
			expected: "10",
		},
		"integer": {
			input:    "2",
			expected: "2",
		},
		"boolean": {
			input:    "true",
			expected: "true",
		},
		"binary operator(PLUS)": {
			input:    "2 + 3",
			expected: "5",
		},
		"binary operator(PLUS) identifier": {
			input:    "i + x",
			expected: "11",
		},
		"binary operator(ASTERISK)": {
			input:    "2 * 3",
			expected: "6",
		},
		"binary operator(LT)": {
			input:    "2 < 3",
			expected: "true",
		},
		"if true": {
			input:    "if true then 2 else 3",
			expected: "2",
		},
		"if else": {
			input:    "if false then 2 else 3",
			expected: "3",
		},
		"parentheses integer": {
			input:    "(2 + 3)",
			expected: "5",
		},
		"nested if true": {
			input:    "if (if false then true else false) then 2 else 3",
			expected: "3",
		},
		"nested if identifier": {
			input:    "if (if x < v then true else false) then 2 else i",
			expected: "1",
		},
		"let declaration": {
			input:    "let x = 2",
			expected: "2",
		},
		"let in": {
			input:    "let x = 2 in x + 3",
			expected: "5",
		},
		"nested let in": {
			input:    "let x = 2 in let y = 3 in x + y",
			expected: "5",
		},
	}

	for name, tt := range tests {
		tt := tt

		env := object.NewEnvironment()
		env.Set("i", &object.Integer{Value: 1})
		env.Set("v", &object.Integer{Value: 5})
		env.Set("x", &object.Integer{Value: 10})

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := Eval(parser.Parse(tt.input), env).Inspect()

			if got != tt.expected {
				t.Errorf("expect: %s, but got: %s", tt.expected, got)
			}
		})
	}
}
