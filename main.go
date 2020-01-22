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

	rdRes(22)

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

	for i := 0; i < len(dtxs); i ++ {
		fmt.Println(dtxs[i])
	}

	return dtxs
}









