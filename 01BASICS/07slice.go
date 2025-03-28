package main

import (
	"fmt"
)

func main() {
	/*	var things = []string{"maleta", "ropa", "reloj"}
		fmt.Println(things)

		things = append(things, "bolso")
		fmt.Println(things)

		athings := append(things[1:])
		fmt.Println(athings)

		smallThings := append(things[1 : len(things)-1])
		fmt.Println(smallThings)

		heros := make([]string, 3, 3)
		heros[0] = "thor"
		heros[1] = "ironman"
		heros[2] = "spiderman"
		fmt.Println(heros)

		heros = append(heros, "deadpool")
		fmt.Println(heros)
		fmt.Println(cap(heros))

		myints := []int{4, 7, 5, 8, 2}
		isSorted := sort.IntsAreSorted(myints)
		fmt.Println("Are int sorted: ", isSorted)
		sort.Ints(myints)
		fmt.Println(myints)

		highScore := make([]int, 4)
		highScore[0] = 234
		highScore[1] = 434
		highScore[2] = 435
		highScore[3] = 644
		//highScore[4] = 545

		fmt.Printf("Type of highScore is %T\n", highScore)

		highScore = append(highScore, 555, 545, 123)
		fmt.Println(highScore)

		var courses = []string{"java", "c", "ruby", "go", "python"}
		fmt.Println(courses)
		var index = 2
		courses = append(courses[:index], courses[index+1:]...)
		fmt.Println(courses)
	*/
	var x = []int{1, 2, 3, 4, 5}
	var p = x[1:3]
	fmt.Println(p)
	p[0] = 8
	//	p[3] = 88
	fmt.Println(p)
	fmt.Println("cap :", cap(p))
	fmt.Println(x)

	a := make([]int, 1024)
	fmt.Println(a, len(a), cap(a))

	a = append(a, 103)
	fmt.Println(a, len(a), cap(a))

}
