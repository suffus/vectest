package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/suffus/vectest"
)

func main() {
	n := flag.Int("n", 1024, "size of vector")
	d := flag.Int("d", 10000, "number of documents")
	flag.Parse()
	list := vectest.NewVectorDocumentList()
	for i := 0; i < *d; i++ {
		list.Add(vectest.NewVectorDocument(i+1, vectest.NewRandomVector(*n)))
	}
	fmt.Println("Document set created:")

	fmt.Println("Query vector created:")

	fmt.Println("Calculating cosine similarity:")

	t0 := time.Now()
	N := 100
	for i := 0; i < N; i++ {
		qVec := vectest.NewRandomVector(*n)
		docs := list.Search(qVec, 3)

		for i, doc := range docs {
			fmt.Printf("Rank %d: DocId %d, Score %f\n", i+1, doc.DocId, doc.Vector.InnerProduct(qVec))
		}
	}
	fmt.Printf("Avg Time Per Search: %v\n", float64(time.Since(t0))/float64(N)/1e9)

}
