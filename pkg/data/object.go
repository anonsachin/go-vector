package data

import (
	"fmt"
	"math/rand"
)

type Vector struct {
	// First element in the size array is number of rows
	// the second one is the columns
	Size [2]int
	// It can be visulized as an array of columns
	// all associated functions are written with this 
	// as the main view
	Data [][]float64
}

func FromArrays(arrays ...[]float64) (*Vector, error){
	if (len(arrays) == 0) {
		return nil, fmt.Errorf("The array is empty")
	}
	var size [2]int
	size[0] = len(arrays)

	for i, arr := range arrays{

		if i == 0 {
			size[1] = len(arr)
		} else if size[1] != len(arr){
			return nil, fmt.Errorf("The arrays are not of the same length %d & %d",size[1],len(arr))
		}

	}

	// assigning the values
	return &Vector{
		Size: size,
		Data: arrays,
	}, nil
}

// RandomVector is function to generate a vector or matrix
// of an arbitary size with random float values
func RandomVector(size [2]int) *Vector {

	data := make([][]float64, 0)

	// Going through the columns and filling them out
	for i := 0; i < size[1]; i++ {
		data = append(data, randomFloat(size[0]))
	}

	return &Vector{
		Size: size,
		Data: data,
	}
}

// randomFloat is for getting an array of random float values
func randomFloat(size int) []float64 {
	data := make([]float64, 0)
	// Going through the columns and filling them out
	for i := 0; i < size; i++ {
		data = append(data, rand.Float64())
	}

	return data
}

// RepeatedValueVector will create a vector with all values with same values
// in every position
func RepeatedValueVector(size [2]int, value float64) *Vector {
	data := make([][]float64, 0)
	// Creating and array of a given length
	for i := 0; i < size[1]; i++ {
		data = append(data, repeatedFloat(size[0],value))
	}

	return &Vector{
		Size: size,
		Data: data,
	}
}

// repeatedFloat is for getting an array of same float values in all positions
func repeatedFloat(size int, value float64) []float64 {
	data := make([]float64,0)
	// Creating and array of a given length
	for i := 0; i < size; i++ {
		data = append(data, value)
	}

	return data
}
