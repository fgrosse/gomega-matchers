package matchers_test

import (
	. "github.com/onsi/gomega"
	. "github.com/fgrosse/gomega-matchers"
)

func Example() {
	myCode := `
		package main

		import "fmt"

		func main() {
			fmt.Println("Hello, 世界")
		}
	`

	Expect(myCode).To(BeValidGoCode())
	Expect(myCode).To(DeclarePackage("main"))
	Expect(myCode).To(ImportPackage("fmt"))
	Expect(myCode).To(ContainCode(`fmt.Println("Hello, 世界")`))
}
