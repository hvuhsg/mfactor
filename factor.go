package mfactor

import (
	"fmt"
	"math/rand"

	"github.com/hvuhsg/mfactor/evolution"
)

func generateRandomVectors[T Number](matrix [][]T) ([]T, []T) {
	vec1 := make([]T, len(matrix))
	for i := 0; i < len(vec1); i++ {
		vec1[i] = T(rand.Float64())
	}

	vec2 := make([]T, len(matrix[0]))
	for i := 0; i < len(vec2); i++ {
		vec2[i] = T(rand.Float64())
	}

	return vec1, vec2
}

func printResultMatrix[T Number](vec1, vec2 []T) {
	resultMatrix := multiplyVectorsToMatrix(vec1, vec2)

	for _, vec := range resultMatrix {
		for _, i := range vec {
			fmt.Printf("%v, ", i)
		}
		fmt.Println()
	}
}

func MFactor[T Number](matrix *[][]T) ([]T, []T) {
	vec1, vec2 := generateRandomVectors(*matrix)

	creature := VectorsCreature[T]{
		vec1:       vec1,
		vec2:       vec2,
		goalMatrix: matrix,
		loss:       -1.0,
	}

	bestCreatures := evolution.StartEvolution(5000, creature)
	bestCreature := bestCreatures[len(bestCreatures)-1]

	vec1 = bestCreature.(VectorsCreature[T]).vec1
	vec2 = bestCreature.(VectorsCreature[T]).vec2

	fmt.Printf("Loss to Size ratio: %f\n", bestCreature.GetLoss()/float64(len(vec1)*len(vec2)/100))

	printResultMatrix(vec1, vec2)

	return vec1, vec2
}
