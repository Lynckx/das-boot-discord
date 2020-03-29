package mexxen

import "testing"

func TestGameType(t *testing.T) {
	want := gameName
	if got := GameType(); got != want {
		t.Errorf("Game() = %q, want %q", got, want)
	}
}
