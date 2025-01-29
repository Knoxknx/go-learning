package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Knox"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello("Knox")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Knox") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloNameEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere un string vacio`, msg, err)
	}
}
func TestHellos(t *testing.T) {
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := Hellos(names)

	if err != nil {
		t.Fatal("Expected no error but got error")
	}

	if len(messages) != len(names) {
		t.Fatalf("Expected %d messages but got %d", len(names), len(messages))
	}

	for _, name := range names {
		if msg, ok := messages[name]; !ok {
			t.Errorf("Message not found for %v", name)
		} else if !regexp.MustCompile(`\b` + name + `\b`).MatchString(msg) {
			t.Errorf("Message '%v' doesn't contain name '%v'", msg, name)
		}
	}

	emptyMessages, err := Hellos([]string{})
	if err != nil {
		t.Fatal("Expected no error with empty slice")
	}
	if len(emptyMessages) != 0 {
		t.Errorf("Expected empty map but got %v messages", len(emptyMessages))
	}

	invalidNames := []string{"Carlos", ""}
	messages, err = Hellos(invalidNames)
	if err == nil {
		t.Fatal("Expected error with empty name but got none")
	}
	if messages != nil {
		t.Error("Expected nil map when error occurs")
	}
}
