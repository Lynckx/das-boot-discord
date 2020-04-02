package mexxen

import "testing"

func TestGameType(t *testing.T) {
	want := "mexxen"
	if got := GameType(); got != want {
		t.Errorf("Game() = %q, want %q", got, want)
	}
}
