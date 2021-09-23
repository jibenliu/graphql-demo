package main

import "graphql-demo/route"

func main() {
	r := router.Router

	router.SetRouter()

	r.Run(":8080")
}
