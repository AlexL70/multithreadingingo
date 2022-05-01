package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Point2D struct {
	x int
	y int
}

var (
	r = regexp.MustCompile(`\((\d*),(\d*)\)`)
)

func findArea(pointsString string) {
	var points []Point2D
	for _, p := range r.FindAllStringSubmatch(pointsString, -1) {
		x, _ := strconv.Atoi(p[1])
		y, _ := strconv.Atoi(p[2])
		points = append(points, Point2D{x: x, y: y})
	}
	area := 0.0
	for i := 0; i < len(points); i++ {
		a, b := points[i], points[(i+1)%len(points)]
		area += float64(a.x*b.y) - float64(a.y*b.x)
	}
	fmt.Println(math.Abs(area) / 2.0)
}

func main() {
	dat, _ := ioutil.ReadFile("polygons.txt")
	text := string(dat)
	start := time.Now()
	for _, line := range strings.Split(text, "\n") {
		findArea(line)
	}
	elapsed := time.Since(start)
	fmt.Printf("Finding all areas took %s.\n", elapsed)
}
