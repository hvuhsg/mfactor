package mfactor

import (
	"math"
	"math/rand"

	"github.com/hvuhsg/mfactor/evolution"
)

type VectorsCreature[T Number] struct {
	goalMatrix *[][]T
	vec1, vec2 []T
	loss       float64
}

func NewVecCreature[T Number](vec1, vec2 []T, goalMetrix *[][]T) VectorsCreature[T] {
	vc := VectorsCreature[T]{
		goalMatrix: goalMetrix,
		vec1:       vec1,
		vec2:       vec2,
		loss:       -1.0,
	}

	return vc
}

func multiplyVectorsToMatrix[T Number](a, b []T) [][]T {
	n := len(a)
	m := len(b)

	// Create an n x m matrix
	result := make([][]T, n)
	for i := range result {
		result[i] = make([]T, m)
	}

	// Fill the matrix with the product of elements
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			result[i][j] = a[i] * b[j]
		}
	}

	return result
}

func calculateLoss[T Number](metric1, metric2 [][]T) float64 {
	loss := 0.0

	for i, row := range metric1 {
		for j := range row {
			cell1 := metric1[i][j]
			cell2 := metric2[i][j]

			// Negative value means there is no data
			if cell1 < 0 || cell2 < 0 {
				continue
			}

			loss += math.Abs(float64(cell1 - cell2))
		}
	}

	return loss
}

func (vc VectorsCreature[T]) GetLoss() float64 {
	if vc.loss >= 0 {
		return vc.loss
	}

	newMatrix := multiplyVectorsToMatrix(vc.vec1, vc.vec2)
	vc.loss = calculateLoss(newMatrix, *vc.goalMatrix)

	return vc.loss
}

func (vc VectorsCreature[T]) Merge(vc2 evolution.Creature) evolution.Creature {
	mergeVec := func(vec1, vec2 []T) []T {
		newVec := make([]T, len(vec1))
		for i := 0; i < len(vec1); i++ {
			newVec[i] = (vec1[i] + vec2[i]) / 2
		}

		return newVec
	}

	newVec1 := mergeVec(vc.vec1, vc2.(VectorsCreature[T]).vec1)
	newVec2 := mergeVec(vc.vec2, vc2.(VectorsCreature[T]).vec2)

	newVc := VectorsCreature[T]{
		vec1:       newVec1,
		vec2:       newVec2,
		goalMatrix: vc.goalMatrix,
		loss:       -1.0,
	}

	return newVc
}

func (vc VectorsCreature[T]) Mutate(gamma float64, sigma float64) evolution.Creature {
	mutateVec := func(vec []T) []T {
		for i := range vec {
			if rand.Intn(100) <= int(gamma)*100 {
				if rand.Intn(100)%2 == 0 {
					vec[i] += T(rand.Float64()) * T(sigma)
				} else {
					vec[i] -= T(rand.Float64()) * T(sigma)
				}

				if vec[i] < 0 {
					vec[i] = 0
				}
			}
		}
		return vec
	}

	newVec1 := mutateVec(vc.vec1)
	newVec2 := mutateVec(vc.vec2)

	newVc := VectorsCreature[T]{
		vec1:       newVec1,
		vec2:       newVec2,
		goalMatrix: vc.goalMatrix,
		loss:       -1.0,
	}

	return newVc
}
