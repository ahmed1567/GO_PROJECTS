package main

import "fmt"


type triangle struct{
	height float64
	base float64

}


type square struct {
	sideLength float64

}


func (s square) getArea() float64{
	return s.sideLength * s.sideLength
}


func (t triangle) getArea() float64{
	return 0.5 * t.base * t.height
}


type shape interface{
	getArea() float64
}

func printArea(s shape){
	fmt.Println(s.getArea())
}

func main(){
	s := square{sideLength: 5}
	t := triangle{height: 10, base: 5}
	printArea(s)
	printArea(t)
}
