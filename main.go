package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main(){
	fmt.Println("I am main function")

	app := fiber.New()

	app.Listen("localhost:9000")
}