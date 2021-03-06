package main

import (
	"flag"
	"fmt"
)

type Config struct {
	subject string
	isAwesome bool
	howAwesome int
	countTheWays CountTheWays
}

func (c *Config) Setup() {
	// you can set a flag directly like so:
	// var someVar = flag.String("flag_name", "default_val", "description")
	// but in practice putting it in a struct is generally better
	flag.StringVar(&c.subject, "subject", "", "subject is a straing, it defaults to empty")
	// --> shorthand version
	flag.StringVar(&c.subject, "s", "", "subject is a string, it defaults to empty (shorthand)")

	flag.BoolVar(&c.isAwesome, "isawesome", false, "is it awesome or what?")
	flag.IntVar(&c.howAwesome, "howawesome", 10, "how awesome out of 10?")

	// custom variable type
	flag.Var(&c.countTheWays, "c", "comma seperated list of integers")

}

func (c *Config) GetMessage() string {
	msg := c.subject
	if c.isAwesome {
		msg +=" is awesome"
	} else {
		msg += " is NOT awesome"
	}

	msg = fmt.Sprintf("%s with a certainty of %d out of 10. Let me count the ways %s", msg, c.howAwesome,
		c.countTheWays.String())
	return msg
}




