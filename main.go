package main

import (
	"fmt"

	"github.com/alustau/issuer/router"
)

func init() {
	fmt.Print("Starting server")
}
func main() {
	fmt.Print("[OK]")
	router.Routes()
}
