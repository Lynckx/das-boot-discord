package games

import "testing"

func TestGame(t *testing.T) {
	tests := []struct {
		want  string
		param string
	}{
		{
			want:  "You choose to play mexxen",
			param: "mexxen",
		},
		{
			want:  "Unknown game type",
			param: "u-boten",
		},
	}
	for _, test := range tests {
		if got := Game(test.param); got != test.want {
			t.Errorf("Game() = %q, want %q", got, test.want)
		}
	}
}
