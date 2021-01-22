package main

// Variadic Functions
func main() {
	sayHelloVariadic("Ozan", "Zeynep", "Abdullah")
}

func sayHelloVariadic(names ...string) {
	for _, name := range names {
		println("Hello", name)
	}
}
