/* Exercise Twenty three
23. Create an interface `Shape` with `Area()` and `Perimeter()` methods. Implement for `Circle` and `Rectangle`.
*/

package exercise

import "fmt"

const PI = 3.1459

// chape interface
type Shape interface {
	Area() float32
	Perimeter() float32
}

// circle struct implementing all the Shape methods
type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return PI * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return PI * (c.Radius * 2)
}

// Rect struct implementing all Shape methods
type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Width + r.Height)
}

func DescribeShape(s Shape) {
	fmt.Println("Area of the Shape is: ", s.Area())
	fmt.Println("Perimeter of the Shape is: ", s.Perimeter())
}
