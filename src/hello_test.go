package src

import "testing"

func TestHello(t *testing.T){
	got := Hello()
	want := "Hi there!"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
