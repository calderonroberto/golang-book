package main

import ("fmt";"math") // a shorthand to importing several modules

func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}

// Struct is a type containing named fields. Type introduces a new type
// followed by a name (Circle) and the keyword struct indicating we 
// are creating a struct type. Initializing structs can be done:
// var c Circle
// which will create a local Circle variable that is default set to zero.
// (O for ints, 0.0 for floats, "" for strings and null for pointers)
// We can also use the new function, allocating memory for all the fields:
// c:= new(Circle)
// which returns a pointer (*Circle), we can do either ways:
// c := Circle{x:0, y:0, r:5}
// c := Circle{0,0,5}

type Circle struct {
	x float64
	y float64
	r float64
	//x,y,r float64 // another way to shortand declare fields
}

type Rectangle struct{
	x1,y1,x2,y2 float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r*c.r
}

func circleAreaPtr(c *Circle) float64{
	return math.Pi * c.r*c.r
}

/*
* Methods
*/

// Methods are special kinds of functions that add a "receiver", below
// (c *Circle), receivers allow calling the function using the . operator
func (c *Circle) area() float64 {
	return math.Pi * c.r*c.r
}

func (r *Rectangle) area() float64 {
        l := distance(r.x1, r.y1, r.x1, r.y2)
        w := distance(r.x1, r.y1, r.x2, r.y1)
        return l * w
}

/*
* Embedded types (Something IS Something else relationship)
*/

type Person struct{
	Name string
}

// An embedded type allows us to say that Android IS a Person. This is 
// done using and embedded type (Person). See main.
type Android struct {
	Person
	Model string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is",p.Name)
}

/*
* Interfaces
*/

// Interfaces are created similar to structs, but instead of defining
// fields they define a method set, or a list of METHODS that a type
// must have in order to "implement" such interface. (NOT FUNCTIONS)
type Shape interface{
	area() float64
}

// Interfaces can also be used as fields:
type MultiShape struct {
	shapes []Shape
}

// And we can even turn MultiShape into a Shape by giving it an area method
func (m *MultiShape) area() float64{
	var area float64
	for _,s := range m.shapes{
		area += s.area()
	}
	return area
}

// We can use interface types as arguments to functions:
func totalArea(shapes ...Shape) float64{
	var area float64
	for _,s := range shapes{
		area += s.area()
	}
	return area
} //see main on how to call this

func main() {
	c := Circle{0,0,5}
	// Accessing the fields is done with the . operator:
	//c.x = 10
	//c.y = 5
	fmt.Println(c.x, c.y, c.r)
	fmt.Println(circleArea(c))
	// Because arguments are always copied in Go we might
	// want to write the function using pointers.
	fmt.Println(circleAreaPtr(&c))

	// Defining a method allows us to modify the Circle type
	// directly, removing the need to pass the pointer like above.
	fmt.Println(c.area())
	r := Rectangle{0,0,10,10}
	fmt.Println(r.area())

	// When using an IS relationship the Person struct can be
	// accessed using the type name
	a := new(Android)
	a.Person.Talk()
	// and we can also call any Person methods directly on the Android
	a.Talk()

	// Calling an interface-designed function:
	fmt.Println(totalArea(&c, &r))

}
