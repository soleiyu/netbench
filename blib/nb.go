package blib

type Nbench struct {
	Day int
	H int
	M int
	S int
	Dsize float64
	T int
}

func MkNbench(day, h, m, s, t int, dsize float64) Nbench {
	var nb Nbench

	nb.Day = day
	nb.H = h
	nb.M = m
	nb.S = s
	nb.T = t
	nb.Dsize = dsize

	return nb
}

func (this *Nbench) Quembps() float64 {
	return this.Dsize * 8.0 / float64(this.T)
}

func (this *Nbench) Quedtime() float64 {
	return float64(this.H) + float64(this.M) / 60.0
}
