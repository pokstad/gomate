package main

type sample struct{}

func (s sample) Do() {}

func main() {
	sample{}.Do()
}
