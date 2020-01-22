package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"./blib"
)

func main() {

	day, _ := strconv.Atoi(os.Args[1])
	d21 := rdRes(day)	

	for i := 0; i < len(d21); i += 2 {
		nb := blib.MkNbench(parser(d21[i], d21[i + 1]))
		if 0.0 < nb.Quembps() {
			fmt.Printf("%.2f %.2f\n", nb.Quedtime(), nb.Quembps())
		}
	}
}

func parser(l1, l2 string) (int, int, int, int, int, float64) {

	l1s := strings.Split(l1, " ")
	l2s := strings.Split(l2, " ")

	e := strings.Index(l2, "error")

	d, _ := strconv.Atoi(l1s[0])
	h, _ := strconv.Atoi(l1s[1])
	m, _ := strconv.Atoi(l1s[2])
	s, _ := strconv.Atoi(l1s[3])
	t := -1
	ds := 0.0

	if e < 0 {
		for i := 0; i < len(l2s); i ++ {
			if strings.Index(l2s[i], "in") < 0 {
				continue
			}
			t = i
		}
	}

	if 0 < t {
		sdsize := strings.Split(l2s[t - 1], "M")[0]
		ds, _ = strconv.ParseFloat(sdsize, 64)

		ts := strings.Split(l2s[t + 1], ":")
		mt, _ := strconv.Atoi(ts[0])
		ms, _ := strconv.Atoi(ts[1])
		t = mt * 60 + ms
	}

	return d, h, m, s, t, ds
}

func rdRes(day int) []string {
	fp, _ := os.Open("RES")
	defer fp.Close()

	rtxs := make([]string, 0)
	dtxs := make([]string, 0)

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		rtx := scanner.Text()
		if rtx != " " {
			rtxs = append(rtxs, rtx)
		}
	}

	var sday string
	if day < 10 {
		sday = "00" + strconv.Itoa(day)
	} else if day < 100 {
		sday = "0" + strconv.Itoa(day)
	} else {
		sday = strconv.Itoa(day)
	}

	for i := 0; i < len(rtxs); i += 2 {
		if strings.Index(rtxs[i], sday) < 0 {
			continue
		}
		dtxs = append(dtxs, rtxs[i], rtxs[i + 1])
	}

	return dtxs
}









