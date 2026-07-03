package main;

import (
	"fmt"
	// "errors"
	// "log"
	// "time"
)

// func main() {
// 	message, err := prediction(time.Now().Weekday())
// 	if err != nil {
// 		fmt.Println("Ошибка:", err)
// 		log.Fatal(err)
// 	}
// 	fmt.Println(message)
// }


// func prediction(dayOfWeek time.Weekday) (string, error) {
// 	switch dayOfWeek {
// 	case time.Monday:
// 		return "Хорошей рабочей недели! Начни день с маленькой победы.", nil
// 	case time.Tuesday:
// 		return "Держи темп — ты уже на втором круге. Удачи сегодня!", nil
// 	case time.Wednesday:
// 		return "Середина недели — самое время не сдаваться. Ты справишься!", nil
// 	case time.Thursday:
// 		return "Ещё немного до выходных. Заверши дела с хорошим настроением.", nil
// 	case time.Friday:
// 		return "Пятница! Отлично поработай и порадуй себя вечером.", nil
// 	case time.Saturday:
// 		return "Хороших выходных! Отдыхай и делай то, что приносит радость.", nil
// 	case time.Sunday:
// 		return "Воскресенье — время восстановиться. Новая неделя будет удачной!", nil
// 	default:
// 		return "", errors.New("неизвестный день недели")
// 	}
// }

// func main(){
// 	fmt.Println(findMin(1, 2, 3, 4, 5, 11313, -2332, 232352323, 24523523112, -232323))
// }

// func findMin(numbers ...int) int {
// 	if len(numbers) == 0 {
// 		return 0
// 	}
// 	min := numbers[0]
// 	for _, number := range numbers {
// 		if number < min {
// 			min = number
// 		}
// 	}
// 	return min
// }

func main(){
	// func(){
	// 	fmt.Println("Hello")
	// }()
	counter := innerFunction()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func innerFunction() func() int{
 count := 0
 return func() int{
	count++
	return count
 }
}

