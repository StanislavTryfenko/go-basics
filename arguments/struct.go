package arguments

import (
	"fmt"
)

// Questa struct rappresenta un auto (Le struct sono le classi di go poró definiscono solo le proprietà)
type car struct {
	Brand      string
	Model      string
	Color      string
	Year       int
	FrontWheel Wheel
	BackWheel  Wheel
}

// Questa struct rappresenta una ruota (Le struct sono le classi di go porò definiscono solo le proprietà)
type Wheel struct {
	Radius   int
	Material string
}

// Go permette di usare struct esistenti per crearne di nuove
type truck struct {
	car
	TrunkCapacity int
}

// Go anche se non é un linguaggio orientato a oggettipermette di avere funzioni che ricevono dati dalle struct
type rect struct {
	width, height int
}

// Funzione basata su una struct
// La funzione sta ricevendo una struct come valori
func (r rect) area() int {
	return r.width * r.height
}

// Questa funzione stampa una struct
func printCar(c car) {
	fmt.Println("Brand:", c.Brand)
	fmt.Println("Model:", c.Model)
	fmt.Println("Color:", c.Color)
	fmt.Println("Year:", c.Year)
}

func Structs() {
	//fmt.Println("Se vedi questo messaggio non c' è nulla di attivo in struct.go")

	// Questo é un array di struct
	cars := [5]car{
		{
			Brand: "Toyota",
			Model: "Camry",
			Color: "Blue",
			Year:  2020,
		},
		{
			Brand: "Honda",
			Model: "Civic",
			Color: "Black",
			Year:  2018,
		},
		{
			Brand: "Tesla",
			Model: "Model S",
			Color: "White",
			Year:  2022,
		},
		{
			Brand: "Chevrolet",
			Model: "Impala",
			Color: "Silver",
			Year:  2019,
		},
		{
			Brand: "BMW",
			Model: "X5",
			Color: "Gray",
			Year:  2021,
		},
	}

	// Questo é il foreach in go
	for car := range cars {
		printCar(cars[car])
		fmt.Println("==================")
	}

	// Testo la struct dentro una struct
	myCar := car{
		Brand: "Ford",
		Model: "Mustang",
		Color: "Red",
		Year:  2022,
		FrontWheel: Wheel{
			Radius:   17,
			Material: "Steel",
		},
		BackWheel: Wheel{
			Radius:   17,
			Material: "Steel",
		},
	}

	fmt.Println("Brand:", myCar.Brand, "Model:", myCar.Model, "Color:", myCar.Color, "Year:", myCar.Year, "FrontWheel:", myCar.FrontWheel.Radius, myCar.FrontWheel.Material, "BackWheel:", myCar.BackWheel.Radius, myCar.BackWheel.Material)
	fmt.Println("==================")

	trucks := [2]truck{
		{
			car: car{
				Brand: "Ford",
				Model: "F150",
				Color: "Red",
				Year:  2022,
				FrontWheel: Wheel{
					Radius:   17,
					Material: "Steel",
				},
				BackWheel: Wheel{
					Radius:   17,
					Material: "Steel",
				},
			},
			TrunkCapacity: 1000,
		},
		{
			car: car{
				Brand: "Chevrolet",
				Model: "Cruze",
				Color: "Blue",
				Year:  2020,
				FrontWheel: Wheel{
					Radius:   17,
					Material: "Steel",
				},
				BackWheel: Wheel{
					Radius:   17,
					Material: "Steel",
				},
			},
			TrunkCapacity: 500,
		},
	}

	for truck := range trucks {
		printCar(trucks[truck].car)
		fmt.Println("FrontWheel:", trucks[truck].car.FrontWheel.Radius, trucks[truck].car.FrontWheel.Material)
		fmt.Println("BackWheel:", trucks[truck].car.BackWheel.Radius, trucks[truck].car.BackWheel.Material)
		fmt.Println("Trunk Capacity:", trucks[truck].TrunkCapacity)
		fmt.Println("==================")
	}

	// stampo un truck custom in tutte le sue proprietà
	myTruck := truck{
		car: car{
			Brand: "Ford",
			Model: "F150",
			Color: "Red",
			Year:  2022,
			FrontWheel: Wheel{
				Radius:   17,
				Material: "Steel",
			},
			BackWheel: Wheel{
				Radius:   17,
				Material: "Steel",
			},
		},
		TrunkCapacity: 1000,
	}

	// Brutta
	// fmt.Printf("My truck details: \nBrand: %v \nModel: %v \nColor: %v \nYear: %v \nFrontWheel: %vcm, %v \nBackWheel: %vcm, %v \nTrunk Capacity: %v \n", myTruck.Brand, myTruck.Model, myTruck.Color, myTruck.Year, myTruck.FrontWheel.Radius, myTruck.FrontWheel.Material, myTruck.BackWheel.Radius, myTruck.BackWheel.Material, myTruck.TrunkCapacity)

	// Formattata
	fmt.Printf(
		"My truck details: \n"+
			"Brand: %v \n"+
			"Model: %v \n"+
			"Color: %v \n"+
			"Year: %v \n"+
			"FrontWheel: %vcm, %v \n"+
			"BackWheel: %vcm, %v \n"+
			"Trunk Capacity: %v \n",
		myTruck.Brand, myTruck.Model, myTruck.Color, myTruck.Year,
		myTruck.FrontWheel.Radius, myTruck.FrontWheel.Material,
		myTruck.BackWheel.Radius, myTruck.BackWheel.Material,
		myTruck.TrunkCapacity,
	)
	fmt.Println("==================")

	// Dimostazione di funzione basata su una struc
	r := rect{
		width:  10,
		height: 5,
	}

	// r.area é una funzione che si basa sulla struct e in pratica moltiplica 2 valori di una variabile dichiarata come struct rect in questo caso
	// Attenzione puó causare bug in quanto modificando le dimenrioni di r la sua area non si aggiornerá dinamicamente se lo storiamo in una variabile
	fmt.Println("Area:", r.area())
}
