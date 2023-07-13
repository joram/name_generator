package name_generator

type NameGenerator struct {
	maleNames   *WeightedNameCollection
	femaleNames *WeightedNameCollection
}

func (ng *NameGenerator) MaleName(seed int64) string {
	return ng.maleNames.GetRandomName(seed)
}

func (ng *NameGenerator) FemaleName(seed int64) string {
	return ng.femaleNames.GetRandomName(seed)
}

func NewNameGenerator() *NameGenerator {
	ng := &NameGenerator{
		maleNames:   NewWeightedNameCollection("./name_generator/data/male_data.csv"),
		femaleNames: NewWeightedNameCollection("./name_generator/data/female_data.csv"),
	}
	return ng
}