package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/unprofession-al/mystrom"
)

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	d := flag.String("d", "", "IP or hostname of the myStrom Switch device")
	a := flag.String("a", "toggle", "action to execute, can be [on, off, toggle, report, temp]")
	flag.Parse()

	s, err := mystrom.NewSwitch(*d)
	must(err)

	switch *a {
	case "toggle":
		err := s.Toggle()
		must(err)
	case "on":
		err := s.On()
		must(err)
	case "off":
		err := s.Off()
		must(err)
	case "report":
		r, err := s.Report()
		must(err)
		if r.Relay {
			fmt.Printf("The Switch is turned on, the current power consumption is %f\n", r.Power)
		} else {
			fmt.Println("The Switch is turned off")
		}
	case "temp":
		t, err := s.Temperature()
		must(err)
		fmt.Printf("The current Switch temperature is %f°C\n", t.Compensated)
	default:
		err := fmt.Errorf("Action '%s' is not defined\n", *a)
		must(err)
	}
}
