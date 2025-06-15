package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

type Menu struct {
	id    uint8
	name  string
	price float32
}

// id name price
var menu = []Menu{
	{1, "Кофе", 100},
	{2, "Кофе с молоком", 200},
	{3, "Жареный рис", 300},
	{4, "Чай", 400},
	{5, "Эспрессо", 120},
	{6, "Латте", 250},
	{7, "Капучино", 230},
	{8, "Сэндвич с ветчиной", 180},
	{9, "Блинчики с джемом", 150},
	{10, "Круассан шоколадный", 170},
}

func printMenu() {
	w:=tabwriter.NewWriter(os.Stdout,11,8,2,' ',0)
	fmt.Printf("%15s\n", "Меню")
	fmt.Printf("%s\n", strings.Repeat("-",40))
	fmt.Printf("%-7s %7s %15s\n", "Номер","Цена","Название")
	fmt.Printf("%s\n", strings.Repeat("-",40))
	
	for _, item := range menu {
		 fmt.Fprintf(w, "%d\t%.2f\t%s\t\n",
            item.id, item.price, item.name,
        )
	}
	w.Flush()
	fmt.Printf("%s\n", strings.Repeat("-",45))
	fmt.Println()

}