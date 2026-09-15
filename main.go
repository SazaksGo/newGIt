package main

import (
	"GITpro/feature1"
	"GITpro/feature2"
	"GITpro/feature_postgres/simple_connection"
	"fmt"
)

func main() {
	fmt.Println("Hi Git")

	feature1.HiGIT()

	feature2.Feature2()
	simple_connection.CheckConnection()
}
