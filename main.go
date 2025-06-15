package main

import "fmt"

//Будем хранить заказ в виде название:количество штук. Например: Кофе:2, Жареный рис:1
var customerOrder = make(map[string]uint8,0)

func main() {
	var userName string 
	fmt.Println("Введите ваше имя") //maybe fix that need only chars
	fmt.Scanln(&userName)
	greet(userName)

	//Logick 

	// printMenu()

	orderItems()

	//end 
	farewell(userName)
	
	

}