package drone

import "testing"

type getterTest struct {
	in       string
	expected string
}

func TestGetter(t *testing.T) {
	tests := []getterTest{
		{in: "2-1", expected: "1"},
		{in: "file('testdata/file.txt')", expected: "test load file\n"},
		{in: `match(file('testdata/Dockerfile'), '.*(FROM onlyoffice/documentserver:(.*)\\s*).*')[2]`, expected: "load"},
	}

	_getter := newGetter(New(newPlugin))
	for index, test := range tests {
		got := _getter.Get(test.in)
		if got != test.expected {
			t.Fatalf("第%d个测试出错，期望：%v，实际：%v", index+1, test.expected, got)
		}
	}
}
