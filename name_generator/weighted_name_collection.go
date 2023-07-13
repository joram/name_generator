package name_generator

import (
	"math/rand"
	"sort"
)

type WeightedNameCollection struct {
	names   []string
	indices []int
}

func (ng *WeightedNameCollection) GetRandomName(seed int64) string {
	rand.Seed(seed)
	i := rand.Intn(len(ng.indices))
	index := ng.indices[i]
	return ng.names[index]
}

func NewWeightedNameCollection(data *map[string]int) *WeightedNameCollection {
	ng := &WeightedNameCollection{
		names:   []string{},
		indices: []int{},
	}

	type nameChance struct {
		name       string
		chances    int
		percentage float64
	}

	names := []string{}
	for name, _ := range *data {
		names = append(names, name)
	}
	sort.Strings(names)

	nameChances := []nameChance{}
	setMinChance := false
	minChance := 0
	for _, name := range names {
		chance := (*data)[name]
		nameChances = append(nameChances, nameChance{
			name:    name,
			chances: chance,
		})

		if !setMinChance || chance < minChance {
			minChance = chance
			setMinChance = true
		}
	}

	for _, nameChance := range nameChances {
		ng.names = append(ng.names, nameChance.name)
		nameIndex := len(ng.names) - 1
		for i := 0; i < nameChance.chances; i++ {
			ng.indices = append(ng.indices, nameIndex)
		}
	}
	return ng
}
