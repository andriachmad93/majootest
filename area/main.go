package main

import (
	"area/models"
	repositories "area/repository"
	"fmt"
)

func main() {
	repositories.InsertArea(10, 10, "persegi", &models.Area{})

	fmt.Println("success")
}
