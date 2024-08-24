package evolution

import (
	"fmt"
	"math/rand"
	"slices"
)

type Creature interface {
	GetLoss() float64
	Merge(Creature) Creature
	Mutate(gamma float64, sigma float64) Creature
}

const classSize = 100
const killRate = 0.92
const gamma = 0.01
const sigma = 10

func breed(creatures []Creature) []Creature {
	for len(creatures) < classSize {
		if len(creatures) == 1 {
			newCreature := creatures[rand.Intn(len(creatures))].Mutate(gamma, sigma)
			creatures = append(creatures, newCreature)
		} else {
			c1 := creatures[rand.Intn(len(creatures))]
			c2 := creatures[rand.Intn(len(creatures))]
			newCreature := c1.Merge(c2).Mutate(gamma, sigma)
			creatures = append(creatures, newCreature)
		}
	}

	return creatures
}

func sortCreatures(creatures []Creature) {
	slices.SortFunc(creatures, func(c1, c2 Creature) int {
		loss1 := c1.GetLoss()
		loss2 := c2.GetLoss()
		if loss1 < loss2 {
			return 1
		} else if loss1 == loss2 {
			return 0
		} else {
			return -1
		}
	})
}

func killWeak(creatures []Creature) []Creature {
	keep := int(classSize - classSize*killRate)

	if keep <= 0 || keep > len(creatures) {
		return nil
	}

	result := make([]Creature, keep)
	for i := 0; i < keep; i++ {
		result[i] = creatures[i]
	}
	sortCreatures(result)

	for _, creature := range creatures {
		if creature.GetLoss() < result[0].GetLoss() {
			result[0] = creature
			sortCreatures(result)
		}
	}

	return result
}

func StartEvolution(cycles int, creatures ...Creature) []Creature {

	fmt.Printf("Start loss: %f\n", creatures[len(creatures)-1].GetLoss())

	for i := 0; i < cycles; i++ {
		creatures = breed(creatures)
		creatures = killWeak(creatures)
		// fmt.Printf("Current loss %d: %f\n", i, creatures[len(creatures)-1].GetLoss())
	}

	fmt.Printf("Best loss: %f\n", creatures[len(creatures)-1].GetLoss())

	return creatures
}
