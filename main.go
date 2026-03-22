package main

import (
	"fmt"
	"git_study/feature1"
	"git_study/feature2"
	"git_study/feature_postgres/simple_connection"
)

func main() {
	fmt.Println("Hello Git!")
	feature1.Feature1()
	feature2.Feature2()
	simple_connection.CheckConnection()
}
