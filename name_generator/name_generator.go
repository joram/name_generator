package name_generator

import "github.com/joram/name_generator/name_generator/data"

type NameGenerator struct {
	maleNames   *WeightedNameCollection
	femaleNames *WeightedNameCollection
	surnames    *WeightedNameCollection
}

func (ng *NameGenerator) MaleName(seed int64) string {
	return ng.maleNames.GetRandomName(seed)
}

func (ng *NameGenerator) FemaleName(seed int64) string {
	return ng.femaleNames.GetRandomName(seed)
}

func (ng *NameGenerator) Surname(seed int64) string {
	return ng.surnames.GetRandomName(seed)
}

func NewNameGenerator() *NameGenerator {
	ng := &NameGenerator{
		maleNames:   NewWeightedNameCollection(&data.MaleNames),
		femaleNames: NewWeightedNameCollection(&data.FemaleNames),
		surnames:    NewWeightedNameCollection(&data.Surnames),
	}
	return ng
}
