package main

import (
	"fmt"
	"github.com/akashrzp/go-bootcamp/internal/firstTasks"
)

func main() {
	fmt.Printf("Day One tasks.....\n\n")

	fmt.Printf("Task - Matrices -- Start.....\n")
	matrix := firstTasks.NewMatrix(3, 3)

	rows := matrix.GetRows()
	fmt.Printf("ROWS: %d\n", rows)

	cols := matrix.GetCols()
	fmt.Printf("COLS: %d\n", cols)

	matrix.AddValue(5, 0, 0)

	matrixValue := make([]int, matrix.Cols)

	for i := 0; i < cols; i++ {
		matrixValue[i] = i * 2
	}

	matrix.AddMatrices(matrixValue)
	matrix.PrintMatrix()

	fmt.Printf("Task - Matrices -- End.....\n\n")
	fmt.Printf("Task - Binary Tree -- Start.....\n")

	binaryTree := &firstTasks.TreeNode{
		Value: "a",
		Left: &firstTasks.TreeNode{
			Value: "+",
			Left:  nil,
			Right: nil,
		},
		Right: &firstTasks.TreeNode{
			Value: "b",
			Left: &firstTasks.TreeNode{
				Value: "-",
				Left:  nil,
				Right: nil,
			},
			Right: &firstTasks.TreeNode{
				Value: "c",
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Printf("Printing Binary Tree:\n")
	binaryTree.PreOrderTraversal()
	fmt.Printf("Task - Binary Tree -- End.....\n\n")

	fmt.Printf("Task - Employee Salary -- Start.....\n")
	employees := []firstTasks.IEmployee{
		&firstTasks.Employee{
			Name:       "John",
			Type:       "FT",
			DailyWage:  500,
			DaysWorked: 30,
		},
		&firstTasks.Employee{
			Name:       "Jane",
			Type:       "CONTRACT",
			DailyWage:  100,
			DaysWorked: 30,
		},
		&firstTasks.FreeLancer{
			Name:        "Jon",
			HourlyPay:   100,
			HoursWorked: 10,
			Type:        "FREELANCER",
		},
	}

	for _, emp := range employees {
		firstTasks.PrintEmpInfo(emp)
		firstTasks.PrintSalary(emp)
	}
	fmt.Printf("Task - Employee Salary -- End.....\n\n")
}
