package main

import "fmt"

type Point struct {X, Y int}
type Circle struct {
	 Point
	 Redius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main()  {
	var w Wheel

	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	w = Wheel{
		Circle: Circle{
			Point:Point{X:8, Y:8},
			Redius:5,
		},
		Spokes:20,
	}

	fmt.Printf("%#v\n", w)
	w.X = 42

	fmt.Printf("%#v\n", w)

}
