package barbarasays_test

import (
	"testing"

	"github.com/waltervargas/barbarasays"
)

func TestSaysWhenColombian(t *testing.T) {
	want := "Malpariculilambido!"

	got := barbarasays.Say("co")

	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestSaysWhenGermany(t *testing.T) {
	want := "Ich liebe dich!"

	got := barbarasays.Say("de")

	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}
