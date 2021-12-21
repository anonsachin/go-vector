package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeatedVector(t *testing.T){
	tt := []struct{
		Name string
		Size [2]int
		Value float64
		Output [][]float64
	}{
		{
			Name: "2X2 matrix of 1's",
			Size: [2]int{2, 2},
			Value: 1,
			Output: [][]float64{[]float64{1,1},[]float64{1,1}},
		},
		{
			Name: "2X1 vector of 1's",
			Size: [2]int{2, 1},
			Value: 1,
			Output: [][]float64{[]float64{1,1}},
		},
		{
			Name: "2X3 Matrix of 1's",
			Size: [2]int{2, 3},
			Value: 1,
			Output: [][]float64{{1,1},{1,1},{1,1}},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name,func(t *testing.T) {
			val := RepeatedValueVector(tc.Size,tc.Value)

			assert.Equal(t,tc.Size,val.Size)
			assert.Equal(t,tc.Output,val.Data)
		})
	}
}


func TestRandomVector(t *testing.T){
	tt := []struct{
		Name string
		Size [2]int
	}{
		{
			Name: "2X2 matrix",
			Size: [2]int{2, 2},
		},
		{
			Name: "2X1 vector",
			Size: [2]int{2, 1},
		},
		{
			Name: "2X3 Matrix",
			Size: [2]int{2, 3},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name,func(t *testing.T) {
			val := RandomVector(tc.Size)

			assert.Equal(t,tc.Size,val.Size)
		})
	}
}