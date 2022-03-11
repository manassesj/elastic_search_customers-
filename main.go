package main

import "elastic_search_customers/routes"

func main() {
	e := routes.Routes()
	e.Logger.Fatal(e.Start(":8080"))
}
