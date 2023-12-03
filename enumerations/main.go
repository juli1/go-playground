package main

// Language enumerations with a string
type Language string

const (
	Java       Language = "java"
	JavaScript Language = "javascript"
)

// Operating system enumerations encoded with an integer
type OperatingSystem int

const (
	Unix    OperatingSystem = iota
	Windows OperatingSystem = iota
	MacOs   OperatingSystem = iota
)

func main() {

	println("Java:    ", Java)
	println("Windows: ", Windows)
}
