package main

import (
	"flag"
	"fmt"

	"github.com/suffus/vectest"
)

func main() {
	n := flag.Int("n", 10, "size of vector")
	d := flag.Int("d", 10, "number of documents")
	flag.Parse()
	list := vectest.NewVectorDocumentList()
	for i := 0; i < *d; i++ {
		list.Add(vectest.NewRandomVector(*n))
	}
	fmt.Println("Document set created:")

}
