package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
//	"./blib"
)

func main() {

	d21 := rdRes(21)
	
	parser(d21[0], d21[1])
}

func parser(l1, l2 string) (int, int, int, int, int, float64) {

	l1s := strings.Split(l1, " ")
	l2s := strings.Split(l2, " ")

	d, _ := strconv.Atoi(l1s[0])
	h, _ := strconv.Atoi(l1s[1])
	m, _ := strconv.Atoi(l1s[2])
	s, _ := strconv.Atoi(l1s[3])

	for i := 0; i < len(l2s); i++ {
		fmt.Println(l2s[i])
	}

	return d, h, m, s, 0, 0.0
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









