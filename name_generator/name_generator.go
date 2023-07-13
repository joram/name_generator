package name_generator

type NameGenerator struct {
	maleNames   *WeightedNameCollection
	femaleNames *WeightedNameCollection
	surenames   *WeightedNameCollection
}

func (ng *NameGenerator) MaleName(seed int64) string {
	return ng.maleNames.GetRandomName(seed)
}

func (ng *NameGenerator) FemaleName(seed int64) string {
	return ng.femaleNames.GetRandomName(seed)
}

func (ng *NameGenerator) Surename(seed int64) string {
	return ng.surenames.GetRandomName(seed)
}

func NewNameGenerator() *NameGenerator {
	ng := &NameGenerator{
		maleNames:   NewWeightedNameCollection(&MaleNames),
		femaleNames: NewWeightedNameCollection(&FemaleNames),
		surenames:   NewWeightedNameCollection(&Surnames),
	}
	return ng
}
