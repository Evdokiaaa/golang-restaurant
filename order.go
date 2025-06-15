package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

func orderItems() {
	printMenu()
	var quantity uint8
	var itemId uint8
	for {
		fmt.Println("Введите 0 чтобы выйти из меню")
		fmt.Printf("Введите номер товара из меню чтобы заказать: ")
		fmt.Scan(&itemId)
		if itemId == 0 {
			break
		}
		
		//Если введеный Id больше чем массив
		if (itemId > uint8(len(menu))) {
			fmt.Println("Неверный номер товара")
			fmt.Println()
			continue
		}
		//Надо еще проверять, есть ли имя в мапе
		if(customerOrder[menu[itemId-1].name] > 0){
			fmt.Println("У вас уже есть товар с таким названием")
			fmt.Println()
			continue
		}
		order:=menu[itemId-1] //Фулл заказ {id name price}
		fmt.Println("Сколько",order.name,"хотите заказать?")
		fmt.Scan(&quantity)
		if(quantity == 0){
			continue
		}
		customerOrder[order.name] = quantity
		
		fmt.Println("Спасибо за заказ, вы только что заказали",order.name,"который стоил",order.price * float32(quantity))
		printOrder()


		
	
		//break nado
	}
}
func printOrder(){
	w:=tabwriter.NewWriter(os.Stdout,30,20,10,' ',0)
	fmt.Println("На данный момент ваш заказ выглядит вот так:")
	fmt.Printf("%s\n", strings.Repeat("-",40))
	fmt.Printf("%10s %25s\n", "Название","Количество")
	fmt.Printf("%s\n", strings.Repeat("-",40))
	for name,quantity := range customerOrder{
		//   fmt.Printf("%7s %14d\n",name,quantity)
		  fmt.Fprintf(w, "%s\t%d\t\n",name,quantity)
	}
	w.Flush()
	fmt.Printf("%s\n", strings.Repeat("-",40))
	
}