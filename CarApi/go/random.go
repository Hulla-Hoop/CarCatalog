package swagger

import (
	"math/rand"
)

func RandomCar(regNum string) Car {

	car := generateRandomCar()
	car.RegNum = regNum
	return car
}

func generateRandomPeople() People {
	names := []string{"Alice", "Bob", "Charlie", "David", "Eve"}
	surnames := []string{"Smith", "Johnson", "Williams", "Jones", "Brown"}
	patronymics := []string{"James", "Liam", "Olivia", "Sophia", "Lucas"}

	return People{
		Name:       names[rand.Intn(len(names))],
		Surname:    surnames[rand.Intn(len(surnames))],
		Patronymic: patronymics[rand.Intn(len(patronymics))],
	}
}

func generateRandomCar() Car {
	marks := []string{"Toyota", "Honda", "Ford", "Chevrolet", "BMW"}
	models := []string{"Camry", "Civic", "Focus", "Malibu", "X5"}
	years := rand.Int31n(30) + 1990 // Random year between 1990 and 2019
	p := generateRandomPeople()

	return Car{
		Mark:  marks[rand.Intn(len(marks))],
		Model: models[rand.Intn(len(models))],
		Year:  years,
		Owner: &p,
	}
}
