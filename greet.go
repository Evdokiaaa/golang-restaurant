package main

import "fmt"


func greet(name string ){
	fmt.Printf("%52s %s%s\n", "Здравствуйте", name,"!")
	fmt.Printf("%70s", "Добро пожаловать в наш магазин Дальничи!")
	fmt.Println()
}

func farewell(name string){
	fmt.Println()
	fmt.Printf("%30s", " ")
	fmt.Printf("%s%s\n","Спасибо за посещение нашего магазина Дальничи, ", name)
	fmt.Printf("%30s\n","Счастливого дня")
	fmt.Printf("%31s\n","До новых встреч!")
}