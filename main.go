package main

import (
	"Tmage/routers"
	"fmt"
)

// @title Tmage
// @version 1.0
// @description Tmage backend API testing.

// @host localhost:8088
// @BasePath /api/v1/
func main() {

	r := routers.SetupRouter()
	r.Run(":8088")
	err := r.Run(fmt.Sprintf(""))
	if err != nil {
		fmt.Printf("Run server failed, err: %s\n", err)
		return
	}
}
