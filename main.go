package main

import "github.com/l-leniac-l/money-ordering-api/app/routes"

func main() {
	router := routes.SetupRouter()

	router.Run()
}
