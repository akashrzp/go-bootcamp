package firstTasks

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	Rows     int
	Cols     int
	Elements [][]int
}

func NewMatrix(rows, cols int) *Matrix {
	data := make([][]int, rows)
	for i := range data {
		data[i] = make([]int, cols)
	}
	return &Matrix{Rows: rows, Cols: cols, Elements: data}
}

func (m *Matrix) GetRows() int {
	return m.Rows
}

func (m *Matrix) GetCols() int {
	return m.Cols
}

func (m *Matrix) AddValue(value, i, j int) {
	m.Elements[i][j] = value
}

func (m *Matrix) AddMatrices(arr []int) {
	matrix := make([][]int, m.Rows)
	for i := range matrix {
		matrix[i] = make([]int, m.Cols)
	}

	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			// Assigning values to the matrix
			matrix[i][j] = arr[j]
		}
	}

	m.Elements = matrix
}

func (m *Matrix) PrintMatrix() {
	jsonData, err := json.Marshal(m.Elements)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Print the JSON data
	fmt.Printf("Matrix: %v\n", string(jsonData))
}
