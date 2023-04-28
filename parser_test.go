package goarg

import "testing"

func TestArgParserOption(t *testing.T) {
	parser := NewParser()

	optionFoo := NewOption("foo", "--foo", "-f", false)
	optionBar := NewOption("bar", "--bar", "-b", false)

	parser.AddOption(optionFoo, optionBar)

	arg1 := []string{"program", "nonce1", "nonce2", "--foo", "FOO_VALUE", "nonce3", "-b", "BAR_VALUE"}
	expectedFooValue := "FOO_VALUE"
	expectedBarValue := "BAR_VALUE"

	err := parser.Parse(arg1)

	if err != nil {
		t.Error(err)
	}

	fooValue, ok := parser.Value("foo")

	if !ok {
		t.Error("foo value not parsed")
	}

	if fooValue != expectedFooValue {
		t.Errorf("foo value mismatch: expected '%s', but parsed '%s'", expectedFooValue, fooValue)
	}

	barValue, ok := parser.Value("bar")

	if !ok {
		t.Error("bar value not parsed")
	}

	if barValue != expectedBarValue {
		t.Errorf("bar value mismatch: expected '%s', but parsed '%s'", expectedBarValue, barValue)
	}
}

func TestArgParserPostionalOption(t *testing.T) {
	parser := NewParser()

	optionFoo := NewPostionalOption("foo", 1, true)
	optionBar := NewPostionalOption("bar", 2, true)

	parser.AddOption(optionFoo, optionBar)

	arg1 := []string{"program", "FOO_VALUE", "BAR_VALUE"}
	expectedFooValue := "FOO_VALUE"
	expectedBarValue := "BAR_VALUE"

	err := parser.Parse(arg1)

	if err != nil {
		t.Error(err)
	}

	fooValue, ok := parser.Value("foo")

	if !ok {
		t.Error("foo value not parsed")
	}

	if fooValue != expectedFooValue {
		t.Errorf("foo value mismatch: expected '%s', but parsed '%s'", expectedFooValue, fooValue)
	}

	barValue, ok := parser.Value("bar")

	if !ok {
		t.Error("bar value not parsed")
	}

	if barValue != expectedBarValue {
		t.Errorf("bar value mismatch: expected '%s', but parsed '%s'", expectedBarValue, barValue)
	}
}

func TestArgParserRequiredPostionalOption(t *testing.T) {
	parser := NewParser()

	optionFoo := NewPostionalOption("foo", 1, true)
	optionBar := NewPostionalOption("bar", 2, true)

	parser.AddOption(optionFoo, optionBar)

	arg1 := []string{"program", "FOO_VALUE"}

	err := parser.Parse(arg1)

	if err == nil {
		t.Error("missing required argument should trigger error")
	}
}

func TestArgParserRequiredOption(t *testing.T) {
	parser := NewParser()

	optionFoo := NewOption("foo", "--foo", "-f", false)
	optionBar := NewOption("bar", "--bar", "-b", true)

	parser.AddOption(optionFoo, optionBar)

	arg1 := []string{"program", "nonce1", "nonce2", "-f", "FOO_VALUE", "nonce3", "-b", "BAR_VALUE"}
	arg2 := []string{"program", "-f", "FOO_VALUE"}
	expectedFooValue := "FOO_VALUE"
	expectedBarValue := "BAR_VALUE"

	err := parser.Parse(arg1)

	if err != nil {
		t.Error(err)
	}

	fooValue, ok := parser.Value("foo")

	if !ok {
		t.Error("foo value not parsed")
	}

	if fooValue != expectedFooValue {
		t.Errorf("foo value mismatch: expected '%s', but parsed '%s'", expectedFooValue, fooValue)
	}

	barValue, ok := parser.Value("bar")

	if !ok {
		t.Error("bar value not parsed")
	}

	if barValue != expectedBarValue {
		t.Errorf("bar value mismatch: expected '%s', but parsed '%s'", expectedBarValue, barValue)
	}

	err = parser.Parse(arg2)

	if err == nil {
		t.Error("missing required argument should trigger error")
	}
}
