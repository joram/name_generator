package name_generator

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
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

func NewWeightedNameCollection(filename string) *WeightedNameCollection {
	ng := &WeightedNameCollection{
		names:   []string{},
		indices: []int{},
	}

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	type nameChance struct {
		name       string
		chances    int
		percentage float64
	}

	nameChances := []nameChance{}
	minPercentage := 100.0
	for _, record := range records {
		name := record[0]
		percentage, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			panic(err)
		}

		nameChances = append(nameChances, nameChance{
			name:       name,
			percentage: percentage,
		})

		if percentage < minPercentage {
			minPercentage = percentage
		}
	}

	for _, nameChance := range nameChances {
		nameChance.chances = int(nameChance.percentage / minPercentage)
		ng.names = append(ng.names, nameChance.name)
		nameIndex := len(ng.names) - 1
		for i := 0; i < nameChance.chances; i++ {
			ng.indices = append(ng.indices, nameIndex)
		}
	}
	return ng
}
