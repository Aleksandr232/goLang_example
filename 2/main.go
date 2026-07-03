package main;

import (
	"fmt"
	"errors"
	"log"
)


func main() {
	message, canEnter, err := enterTheClub(14)
	fmt.Println(message)
	fmt.Println(canEnter)
	if err != nil {
		fmt.Println("Ошибка:", err);
		log.Fatal(err);
	}
}



func enterTheClub(age int) (string, bool, error) {
	if age >= 19 {
		return "Ты можешь купить алкоголь", true, nil
	} else if age >= 18 {
		return "Ты можешь войти в клуб", true, nil	} else {
		return "Ты не можешь войти в клуб", false, errors.New("Ты не можешь войти в клуб")
	}
}