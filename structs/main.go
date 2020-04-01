package main

import (
	"encoding/json"
	"fmt"
)

type Rectangle struct {
	Length float64 `json:"Length"`
	Width  float64 `json:"hEiGhT"`
}

type Cuboid struct {
	Length float64
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (c *Cuboid) Volume() float64 {
	return c.Length * c.Width * c.Height
}

//func (r *Rectangle) String() string {
//	return fmt.Sprintf("Rectangle with Area: %f", r.Area())
//}
//
//func (r *Cuboid) String() string {
//	return fmt.Sprintf("Cuboid with Volume: %f", r.Volume())
//}

func main() {

	rekt := &Rectangle{
		Length: 30,
		Width:  20,
	}

	kube := &Cuboid{
		Length: 30,
		Width:  30,
		Height: 20,
	}

	fmt.Printf("%#v\n", rekt)
	fmt.Printf("%#v\n", kube)

	fmt.Printf("%f\n", rekt.Area())
	fmt.Printf("%f\n", kube.Volume())

	fmt.Printf("%s\n", rekt)
	fmt.Printf("%s\n", kube)

	r, _ := json.Marshal(rekt)
	k, _ := json.Marshal(kube)

	fmt.Printf("%s\n%s\n", string(r), string(k))

}
