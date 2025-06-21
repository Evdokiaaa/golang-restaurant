package main

import (
	"fmt"
	"strings"
	"time"
)

func displayBillTitle() {
	billText:= "+----------------------------------------Формурируем чек----------------------------------------+"
	for _, val := range billText {
		fmt.Printf("%c", val)
		time.Sleep(time.Millisecond * 5) //одна миллисекунда 
		
	}
	fmt.Println()
	

}

func printBill(){
	fmt.Println()
	fmt.Printf("+%s+\n", strings.Repeat("-",100))
	fmt.Printf("%-7s %20s %20s %25s\n", "Название продукта","Цена(₽)","Количество", "Итоговая Цена(₽)")
	fmt.Printf("+%s+\n", strings.Repeat("-",100))

	printBillData()
	fmt.Printf("+%s+\n", strings.Repeat("-",100))
}

func printBillData(){
	//Будем брать из массива customerOrder и выводлить
	var sum float32 = 0.0
	for name,quantity:= range customerOrder {
		for _,el := range menu {
			if(el.name == name){
				totalPrice := el.price * float32(quantity)
				sum += totalPrice
				fmt.Printf("%-30s %-20.2f %-20v %-20.2f\n",name,el.price,quantity,totalPrice)
			}
		}
		//fmt.Println(name,quantity) //name quantity
	}
	
	fmt.Printf("+%s+\n", strings.Repeat("-",100))
	fmt.Println()
	fmt.Printf("%90s %-40.2f\n", "Итоговая сумма:", sum)
}


