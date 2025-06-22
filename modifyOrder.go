package main

import (
	"fmt"
	"strings"
)

func modifyOrder() {
	var changeOption string
	var changeNumber uint
	fmt.Println("Хотите изменить заказ? Введите y или n")
	fmt.Scan(&changeOption)
	if(changeOption != "y" && changeOption != "n"){
		fmt.Println("Введите y или n")
		modifyOrder()
	}
	if(changeOption == "y"){
		fmt.Println()
		fmt.Println("Введите соотвествующий номер для продолжения:")
		fmt.Println("Введите '1' для изменения количества товара")
		fmt.Println("Введите '2' чтобы удалить товар из списка заказов")
		fmt.Println("Введите '3' чтобы добавить товар в список заказа")
		fmt.Printf("%s", "Ваш выбор: ")
		fmt.Scan(&changeNumber)

		switch changeNumber {
			case 1:
				updateItemQuantity()
				
				//TODO: Возможно в будущем как то массив отсортировать
				printOrder()
				modifyOrder()
				
			case 2:
			
				deleteItem()
				printOrder()
				modifyOrder()
			case 3:
				//Туту будут функции
				addItem()
				printOrder()
				modifyOrder()
				
			default:
				fmt.Println("Неверный выбор")
				modifyOrder()

		}


	}
}

func printModifyOrderMenu(){
	fmt.Printf("%20s\n","Меню")
	fmt.Printf("%s\n", strings.Repeat("-",50))
	fmt.Printf("%-7s %13s %23s\n", "Номер","Название","Количество")
	fmt.Printf("%s\n", strings.Repeat("-",50))
	fmt.Println()
	id := 1
	for name, quantity := range customerOrder {
		fmt.Printf("%3d %20s %20d\n",
			id, name, quantity,
		)
	id++

	}
	fmt.Println()
	//После того как пользователь ввел номер, мы берем его key и  ищем в customerOrder элемент с таким key, после чего мы обновляем quantity
}

func updateItemQuantity(){
	var choice uint8;
	var newQuantity uint8;
	printModifyOrderMenu()
	fmt.Printf("%s","Введите номер товара для изменения: ")
	fmt.Scan(&choice)
	fmt.Printf("%s","Введите новое количество товара: ")
	fmt.Scan(&newQuantity)
	
	item:=menu[choice-1].name
	customerOrder[item] = newQuantity
}
func deleteItem(){
	var choice uint8;
	fmt.Println("Введите номер товара для удаления:")
	fmt.Scan(&choice)
	delete(customerOrder, menu[choice-1].name)
}

func addItem(){
	var choice uint8;
	var productQuantity uint8;
	fmt.Println("Введите номер предмет для добавления:")
	fmt.Scan(&choice)
	fmt.Printf("%s","Введите количество товара: ")
	fmt.Scan(&productQuantity)
	customerOrder[menu[choice-1].name] = productQuantity


}