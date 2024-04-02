package main

type gasEngine struct {
	x int32
}

type eletricEngine struct {
	y int32
}

type car[T gasEngine | eletricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func main() {
	var gasCar = car[gasEngine]{
		carMake:  "Honda",
		carModel: "civic",
		engine: gasEngine{
			x: 20,
		},
	}

	var eletricCar = car[eletricEngine]{
		carMake:  "Honda",
		carModel: "civic",
		engine: eletricEngine{
			y: 30,
		},
	}
}
