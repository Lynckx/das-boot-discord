package score

type Score struct {
	wins  int
	loses int
}

func (s Score) GetScore() (int, int) {
	return s.wins, s.loses
}
