package main

import "fmt"

//TODO
//kod içerisinde koşul yapılarınıda kullanalım ve fonksiyonlara bir kaç tane daha ekleyelim 
//çeşitliliği artıralım

func main() {

	//TODO
	//nesne isimlerini tasarlarken go dilinde en iyi kullanım
	//camelCase yöntemidir araştıralım
	visaexamscore := 80
	finalexamscore := 60
	bring := average(visaexamscore, finalexamscore)

	fmt.Println(bring)

}

func average(x int, y int) int {

	calculatex := (x * 40) / 100
	calculatey := (y * 60) / 100

	return calculatex + calculatey

}
