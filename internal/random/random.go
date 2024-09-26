package random

import "math/rand"

func RandomName() string {
	names := []string{"Beto", "Alphonso", "David", "Daniel", "Jonathan", "Julio Cesar", "Carlos", "Migel", "TestUsername",
		"TiredDeveloper", "TeavaroTestDeveloper"}

	name := names[rand.Intn(len(names))]
	return name
}

func RandomOccupation() string {
	occupations := []string{"Designer", "Backend Developer", "Frontend Developer", "Golang Developer", "JavaScript Developer", "Architect", "Senior Developer", "Junior Developer"}

	occupation := occupations[rand.Intn(len(occupations))]
	return occupation
}
