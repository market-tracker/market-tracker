package main

import (
	"fmt"
	"log"

	"github.com/market-tracker/market-tracker/server"
)

func init() {
	fmt.Println("First of all")
}

func main() {
	fmt.Println("Hello market-tracker project")
	s := server.InitServer(3000)
	s.Start(func() {
		log.Panic("App crash")
	})
}
