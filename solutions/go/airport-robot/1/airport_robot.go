package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	languageName() string
	Greet(a string) string
}

func SayHello(visitor string, myGreeter Greeter) string {

	return fmt.Sprintf("I can speak %s: %s", myGreeter.languageName(), myGreeter.Greet(visitor))

}

type Italian struct {
}

func (i Italian) languageName() string {
	return "Italian"
}
func (i Italian) Greet(a string) string {
	return "Ciao " + a + "!"
}

type Portuguese struct {
}

func (i Portuguese) languageName() string {
	return "Portuguese"
}
func (i Portuguese) Greet(a string) string {
	return "Olá" + " " + a + "!"
}
