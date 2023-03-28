package drone

import "testing"

type getterTest struct {
	in       string
	expected string
}

func TestGetter(t *testing.T) {
	tests := []getterTest{
		{in: "2-1", expected: "1"},
	}

	_getter := newGetter(New(newPlugin))
	for _, test := range tests {
		got := _getter.Get(test.in)
		if got != test.expected {
			t.Fatalf("期望：%v，实际：%v", test.expected, got)
		}
	}
}
