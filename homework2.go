package main

import "fmt"

func main() {

	var enemyCapacity int = 60
	var targets [10]string = [10]string{"1.target", "2.target", "3.target", "4.target", "5.target", "6.target", "7.target", "8.target", "9.target", "10.target"}
	var enemies [10]int = [10]int{3, 5, 8, 3, 6, 10, 9, 5, 4, 7}
	totalEnemiesKilled := calculatingTheOperationStrategy(targets, enemies, enemyCapacity)
	fmt.Println(totalEnemiesKilled)

	var enemyCapacity2 int = 60
	var targets2 [10]string = [10]string{"11.target", "13.target", "15.target", "20.target", "28.target", "32.target", "35.target", "40.target", "45.target"}
	var enemies2 [10]int = [10]int{6, 3, 5, 7, 10, 8, 3, 9, 4, 5}
	totalEnemiesKilled2 := calculatingTheOperationStrategy(targets2, enemies2, enemyCapacity2)
	fmt.Println(totalEnemiesKilled2)

	fmt.Println("Operation Failed")

}

func calculatingTheOperationStrategy(targets [10]string, enemies [10]int, enemyCapacity int) int {

	fmt.Println("Operation Started")
	
	var existingEnemy int = 0

	for i := 0; i < len(targets); i++ {
		fmt.Println(targets[i], "hit")

	}

	var inComingEnemy int = enemies[9]
	if enemyCapacity == existingEnemy {

	} else if enemyCapacity-existingEnemy > 45 {
		existingEnemy += inComingEnemy
		fmt.Println("Kalan Düşman", enemyCapacity-existingEnemy)
		fmt.Println(existingEnemy)

	} else {
		

	}
	return existingEnemy
}
