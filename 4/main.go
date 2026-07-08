package main

import (
	"errors"
	"fmt"
)

// var msg string
// var count int

// func init(){
// 	msg = "Hello, World!"
// 	count = 0
// }


// func main(){
// 	num := 10
// 	var p *int
// 	// msg := "Hello, World!"
	
// 	p = &num
// 	fmt.Println(p)
// 	fmt.Println(&num)

// 	*p = 20
// 	fmt.Println(num)
// 	// printMessage(&msg);
// }

// func printMessage(msg *string){
// 	*msg += " (from printMessage)"
// 	fmt.Println(msg)
// }

func main(){
	
	// messages := [3]string{"hello", "world", "golang"}

	// messages := make([]string, 3)
	// messages = append(messages, "hellow")
	// messages[0] = "hello"
	// messages[1] = "world"
	// messages[2] = "golang"

	// printMessages(messages)

	
	// fmt.Println(len(messages))
	// fmt.Println(cap(messages))

	// matrix := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }

	// fmt.Println(matrix)
	// matrix := make([][]int, 10)
	// counter := 0
	// for i := 0; i < 10; i++ {
	// 	matrix[i] = make([]int, 10)
	// 	for j := 0; j < 10; j++ {
	// 		counter++
	// 		matrix[i][j] = counter
	// 	}

	// 	fmt.Println(matrix[i])
	// }

	messages := []string{"hello", "world", "golang"}

	for index, message := range messages {
		fmt.Println(index, message)
	}
}

func printMessages(messages []string) error{

	if len(messages) == 0{
		return errors.New("messages is empty")
	}

	
	fmt.Println(messages)
	return nil
}