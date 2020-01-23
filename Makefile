DAY = 23

default:
	rm res.plt
	go run main.go $(DAY) > res.plt
	gnuplot plot.txt
	eog res.png
