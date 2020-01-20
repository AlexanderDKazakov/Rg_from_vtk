package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Point is useful struct
type Point struct {
	x int
	y int
	z int
	v float64
}

func (p Point) SubPoint(other Point) Point {
	return Point{
		p.x - other.x,
		p.y - other.y,
		p.z - other.z,
		p.v + 0,
	}
}


func NewPoint(x, y, z int, v float64) *Point {
	p := Point{x: x, y: y, z: z, v: v}
	return &p
}

func main() {
	fmt.Println("==== vtk =====")
	if len(os.Args) == 1 {
		fmt.Println("Please specify path to vtk file!")
		return
	}
	inputPathFilename := os.Args[1]

	var (
		lineNum int
		x, y, z int
		boxSize int
		points  []Point
	)

	file, err := os.Open(inputPathFilename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var cs string = scanner.Text()

		matched, err := regexp.MatchString("DIMENSIONS", cs)
		check(err)
		reInt := regexp.MustCompile("[0-9]+")

		if matched {
			var _boxSize []string
			_boxSize = reInt.FindAllString(cs, -1)
			boxSize, err = strconv.Atoi(_boxSize[0])
			check(err)
			fmt.Printf("Box size %d %d %d\n", boxSize, boxSize, boxSize)
		}
		
		// Skip header
		if lineNum > 9 {
			f, err := strconv.ParseFloat(cs, 32)
			check(err)

			if z == boxSize {
				z = z % boxSize
				y++
			}
			if y == boxSize {
				y = y % boxSize
				x++
			}

			p := NewPoint(x, y, z, f)
			points = append(points, *p)
			z++
		}
		lineNum++
	}
	
	err = scanner.Err()
	check(err)

	rg2 := 0.0
	mass := 0.0
	for _, v := range points { mass += v.v }
	fmt.Println("mass:", mass)

	// cm
	var xtemp, ytemp, ztemp float64 = 0.0, 0.0, 0.0
	for _, point := range points {
		xtemp += (float64(point.x) * point.v)
		ytemp += (float64(point.y) * point.v)
		ztemp += (float64(point.z) * point.v)
	}
	cm := Point{int(xtemp / mass), int(ytemp / mass), int(ztemp / mass), 0.0}
	fmt.Println("Center of mass: ", cm)

	var cdist Point
	for _, point := range points {
		cdist = point.SubPoint(cm)
		rg2 += point.v * (math.Pow(float64(cdist.x), 2) + math.Pow(float64(cdist.y), 2) + math.Pow(float64(cdist.z), 2))
	}
	rg2 /= mass
	fmt.Println("Rg2:", rg2)
}
