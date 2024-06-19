package vectest

import (
	"math"
	"math/rand"
	"math/sort"
	"sort"
)

type Vector struct {
	vec []float64
}

func NewVector(n int) *Vector {
	return &Vector{make([]float64, n)}
}

func NewRandomVector(n int) *Vector {
	v := NewVector(n)
	for i := range v.vec {
		v.vec[i] = rand.Float64()
	}
	return v
}

func (v *Vector) Set(i int, x float64) {
	v.vec[i] = x
}

func (v *Vector) Get(i int) float64 {
	return v.vec[i]
}

// generates new vector u = c1v + c2w
func (v *Vector) AddWeighted(w *Vector, c1 float64, c2 float64) *Vector {
	// = c1u + c2w
	if len(v.vec) != len(w.vec) {
		panic("different dimensions")
	}
	u := NewVector(len(v.vec))
	for i := range v.vec {
		u.vec[i] = c1*v.vec[i] + c2*w.vec[i]
	}
	return u
}

func (v *Vector) AddWeightedInPlace(w *Vector, c1 float64, c2 float64) {
	if len(v.vec) != len(w.vec) {
		panic("different dimensions")
	}
	for i := range v.vec {
		v.vec[i] = c1*v.vec[i] + c2*w.vec[i]
	}
}

func (v *Vector) Norm() float64 {
	sum := 0.0
	for _, x := range v.vec {
		sum += x * x
	}
	return math.Sqrt(sum)
}

func (v *Vector) InnerProduct(w *Vector) float64 {
	if len(v.vec) != len(w.vec) {
		panic("different dimensions")
	}
	sum := 0.0
	for i := range v.vec {
		sum += v.vec[i] * w.vec[i]
	}
	return sum
}

func (v *Vector) Equal(w *Vector) bool {
	if len(v.vec) != len(w.vec) {
		return false
	}
	for i := range v.vec {
		if v.vec[i] != w.vec[i] {
			return false
		}
	}
	return true
}

func (v *Vector) Copy() *Vector {
	u := NewVector(len(v.vec))
	copy(u.vec, v.vec)
	return u
}

func (v *Vector) Add(w *Vector) *Vector {
	return v.AddWeighted(w, 1.0, 1.0)
}

func (v *Vector) AddInPlace(w *Vector) {
	v.AddWeightedInPlace(w, 1.0, 1.0)
}

func (v *Vector) Sub(w *Vector) *Vector {
	return v.AddWeighted(w, 1.0, -1.0)
}

func (v *Vector) SubInPlace(w *Vector) {
	v.AddWeightedInPlace(w, 1.0, -1.0)
}

func (v *Vector) Scale(c float64) *Vector {
	return v.AddWeighted(v, c, 0.0)
}

func (v *Vector) ScaleInPlace(c float64) {
	v.AddWeightedInPlace(v, c, 0.0)
}

type VectorDocument struct {
	DocId  int
	Vector *Vector
}

func NewVectorDocument(docId int, vec *Vector) *VectorDocument {
	return &VectorDocument{docId, vec}
}

type VectorDocumentList struct {
	Docs []*VectorDocument
}

func NewVectorDocumentList() *VectorDocumentList {
	return &VectorDocumentList{}
}

func (l *VectorDocumentList) Add(doc *VectorDocument) {
	l.Docs = append(l.Docs, doc)
}

func (l *VectorDocumentList) Get(i int) *VectorDocument {
	return l.Docs[i]
}

func (l *VectorDocumentList) Search(v *Vector, k int) []*VectorDocument {
	if len(l.Docs) == 0 {
		return nil
	}
	if len(l.Docs) < k {
		k = len(l.Docs)
	}
	scores := make([]float64, len(l.Docs))
	for i, doc := range l.Docs {
		scores[i] = doc.Vector.InnerProduct(v)
	}
	indices := make([]int, len(l.Docs))
	for i := range indices {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		return scores[i] > scores[j]
	})

	result := make([]*VectorDocument, k)
	for i := 0; i < k; i++ {
		result[i] = l.Docs[indices[i]]
	}
	return result
}
