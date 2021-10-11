package main

import (
	"fmt"

	"github.com/market-tracker/market-tracker/config"
	"github.com/market-tracker/market-tracker/server"
)

func init() {
	fmt.Println("First of all")
}

func main() {
	c := config.GetConfiguration()
	s := server.InitServer(c.Port)
	s.Start()
}
