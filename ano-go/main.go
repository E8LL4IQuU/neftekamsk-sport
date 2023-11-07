package main

import (
	"github.com/E8LL4IQuU/ano-go/routes"
	"github.com/E8LL4IQuU/ano-go/model"
)

func main() {

	model.InitializeDB()
	routes.InitializeFiber()
	
}
