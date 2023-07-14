# Name Generator

## Example
```golang
ng := name_generator.NewNameGenerator()
seed := rand.Int64()

firstNameMale   := ng.MaleName(seed)
firstNameFemale := ng.FemaleName(seed)
lastName        := ng.Surame(seed)

fmt.Printf("Welcome Mr. %s %s and Ms. %s %s\n", firstNameMale, lastName, firstNameFemale, lastName)
```