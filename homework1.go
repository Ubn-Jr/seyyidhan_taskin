package main

import "fmt"

//TODO
//kod içerisinde koşul yapılarınıda kullanalım ve fonksiyonlara bir kaç tane daha ekleyelim 
//çeşitliliği artıralım

func main() {

	//TODO
	//nesne isimlerini tasarlarken go dilinde en iyi kullanım
	//camelCase yöntemidir araştıralım
	visaExamScore := 80
	finalExamScore := 60
	bring := average(visaExamScore, finalExamScore)

	fmt.Println(bring)

}

func average(x int, y int) int {

	calculatex := (x * 40) / 100
	calculatey := (y * 60) / 100

	return calculatex + calculatey

}
