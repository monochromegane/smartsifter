package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/monochromegane/smartsifter"
)

var (
	useCategorical = flag.Bool("x", false, "Use categorical input X.")
)

func main() {
	flag.Parse()
	dir := "tmp"
	file := "faithful.csv"

	// Load data.
	if _, err := os.Stat(file); err != nil {
		downloadFaithful(file)
	}
	data := loadData(file)

	cellNum := 0
	if *useCategorical {
		cellNum = 4
	}
	ss := smartsifter.NewSmartSifter(0.1, 1.5, 0.2, cellNum, 2, 2)

	// Online outlier by SmartSifter
	os.Mkdir(dir, 0777)
	linspace := linspace(-2.5, 2.5, 100)
	for i, d := range data {
		out, _ := os.OpenFile(filepath.Join(dir, fmt.Sprintf("tmp_%03d.csv", i)), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)

		for _, x := range linspace {
			for _, y := range linspace {
				score := ss.Input(categoricalData(x, y), []float64{x, y}, false)
				out.WriteString(fmt.Sprintf("%f\n", score))
			}
		}
		out.Close()
		ss.Input(categoricalData(d[0], d[1]), d, true)
	}
}

func loadData(file string) [][]float64 {
	var data [][]float64
	var xs []float64
	var ys []float64

	f, _ := os.Open(file)
	defer f.Close()

	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		e, _ := strconv.ParseFloat(record[0], 64)
		w, _ := strconv.ParseFloat(record[1], 64)
		xs = append(xs, e)
		ys = append(ys, w)
	}
	sxs := standardScale(xs)
	sys := standardScale(ys)

	for i, _ := range sxs {
		data = append(data, []float64{sxs[i], sys[i]})
	}
	return data
}

func standardScale(xs []float64) []float64 {
	var mean float64
	for _, x := range xs {
		mean += x
	}
	mean = mean / float64(len(xs))

	var std float64
	for _, x := range xs {
		std += math.Pow(x-mean, 2)
	}
	std = math.Sqrt(std / float64(len(xs)))
	standardXs := make([]float64, len(xs))
	for i, x := range xs {
		standardXs[i] = (x - mean) / std
	}
	return standardXs
}

func downloadFaithful(file string) {
	url := "http://www.stat.cmu.edu/~larry/all-of-statistics/=data/faithful.dat"

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	lines := []string{}
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	lines = lines[26:]

	out, _ := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	defer out.Close()

	s := regexp.MustCompile(` +`)
	for _, line := range lines {
		words := strings.Split(s.ReplaceAllString(line, " "), " ")
		out.WriteString(words[1] + "," + words[2] + "\n")
	}
}

func linspace(from, to float64, num int) []float64 {
	size := math.Abs(from) + math.Abs(to)
	span := float64(size) / float64(num)

	n := int(size / span)
	l := make([]float64, n)
	for i := 0; i < n; i++ {
		l[i] = from + span*float64(i)
	}
	return l
}

func categoricalData(x, y float64) []int {
	if !*useCategorical {
		return nil
	}
	c := 0
	if x > 0 {
		if y > 0 {
			c = 0
		} else {
			c = 1
		}
	} else {
		if y > 0 {
			c = 2
		} else {
			c = 3
		}

	}
	return []int{c}
}
