package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

type transportation interface{
	transportationDevice() string
}

func (s sedan) transportationDevice() string {
	return fmt.Sprintln("The sedan has", s.doors, "doors and can transport people.")
}

func (t truck) transportationDevice() string {
	return fmt.Sprintln("The truck has", t.doors, "doors and can transport packages.")
}

func report (t transportation) {
	fmt.Println(t.transportationDevice())
}

func main() {
	t1 := truck{
		vehicle{
			2,
			"red",
		},
		true,
	}
	fmt.Println(t1)
	fmt.Println(t1.color)
	fmt.Println(t1.transportationDevice())
	report(t1)
	
	s1 := sedan{
		vehicle{
			4,
			"white",
		},
		false,
	}
	fmt.Println(s1)
	fmt.Println(s1.luxury)
	fmt.Println(s1.transportationDevice())
	report(s1)


}
