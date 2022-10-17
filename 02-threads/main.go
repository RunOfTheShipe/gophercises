package main

import "fmt"

func main() {
	done := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["name"] = "data race"
		done <- true
	}()
	m["name"] = "world"
	fmt.Println("Hello,", m["name"])
	<-done
}
