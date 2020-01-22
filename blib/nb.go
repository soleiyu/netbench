package nbench

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
