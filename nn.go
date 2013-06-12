package nn

import (
	"github.com/jcla1/matrix"
	"math"
)

func CostFunction(data [][]*matrix.Matrix, thetas []*matrix.Matrix, lambda float64) float64 {
	cost := float64(0)
	var estimation []float64
	var expected_output []float64

	// Cost

	for _, datum := range data {
		estimation = Hypothesis(thetas, datum[0]).Values()
		expected_output = datum[1].Values()

		for k, y := range expected_output {
			cost += y*math.Log(estimation[k]) + (1-y)*math.Log(1-estimation[k])
		}
	}

	// Regularization
	regularizationCost := float64(0)

	for _, theta := range thetas {
		for i, param := range theta.Values() {
			if i%theta.Columns() == 0 {
				continue
			}

			regularizationCost += param * param
		}
	}

	return cost/float64(len(data)) + (lambda/(2*float64(len(data))))*regularizationCost
}

func Hypothesis(thetas []*matrix.Matrix, input *matrix.Matrix) *matrix.Matrix {
	// Describes the current working values (a_1, a_2, ...)
	curValues := input

	// Is simply a 1 in a 1x1 matrix to b
	// inserted into a vector as the bias unit
	biasValueMatrix := matrix.Ones(1, 1)

	for _, theta := range thetas {
		// Insert the bias unit, multiply with theta and apply the sigmoid function
		curValues = theta.Mul(curValues.InsertRows(biasValueMatrix, 0)).Apply(sigmoidMatrix)
	}

	return curValues
}

// Helper functions

func sigmoidMatrix(index int, value float64) float64 {
	return sigmoid(value)
}

func sigmoid(z float64) float64 {
	return 1 / (1 + math.Pow(math.E, -z))
}
