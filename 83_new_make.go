package main

type Foo map[string]string
type Bar struct {
	thingOne string
	thingTwo int
}

func main() {
	// ok
	y := new(Bar)
	(*y).thingOne = "hello"
	(*y).thingTwo = 1

	// not ok
	// make is only for slice, map, and channel
	// z:=make(Bar)// compile error: cannot make Bar
	// (*z).thingOne = "hello"
	// (*z).thingTwo = 1

	// ok
	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"

	// not ok
	// u := new(Foo)
	// (*u)["x"] = "goodbye" // run error,panic: assignment to entry in nil map
	// (*u)["y"] = "world"
}
