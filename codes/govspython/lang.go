package main

func main() {

	names := []string{"zeta", "chow", "world"}

	for i, n := range names {
		println(i, "Hello, "+n)
	}
}
