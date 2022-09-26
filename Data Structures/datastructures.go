package main

import (
	"fmt"
	// "log"
	// "net/http"
)


// Datastructurer
func main() {

	// formatteirng med verbs

	// fmt.Printf("This is a type: %t \n", true)
	// fmt.Printf("Hello, %v \n", "World")
	

	// Arrays med bestämnda värden

	// integerArray := [5]int{10, 20, 30, 40, 50}
	// stringArray := [4]string{"first", "second", "third", "fourth"}

	// fmt.Println(integerArray)
	// fmt.Println(stringArray)

	//slice som är dynamiska med undefined antal elements
	// integerSlice := []int{10, 20, 30, 40, 50}
	// stringSlice := []string{"first", "second", "third", "fourth"}
	
	// fmt.Println("integer slice: ", integerSlice)
	// fmt.Println("string slice: ", stringSlice)

	//slice, 4:an är antalet elements
	// s := make([]int, 4)
	
	// s[0] = 10
	// s[1] = 20
	// s[2] = 30
	// s[3] = 40

	// fmt.Println("Slice created with 'make': ", s)


	// CRUD på slice:
	// https://medium.com/@victorsteven/understanding-data-structures-in-golang-f55205afdcaa

	// Lägg till ett element
	// s = append(s, 60)
	// fmt.Println(s)

	// Lägg till fler elements
	// s = append(s, 60, 70)
	// fmt.Println(s)

	// Ta bort från slice
	// Tar bort 20
	// s = append(s[:1], s[1+1:]...)
	// fmt.Println(s)

	// Byt värde på ett element
	// s[3] = 23
	// fmt.Println(s)

	//Ta fram specifik index
	//Tar fram index mellan 1 och 3
	// fmt.Println(s[1:3])

	//Ta fram längden på slice
	// fmt.Println(len(s))

	//Ta fram capacity AKA antalet elements som kan finnas
	// fmt.Println(cap(s))	

	//Kopiera en slice till en annan
	// d := make([]int, len(s))
	// copy(d, s)
	// fmt.Println(d)

	//Loopar på slice
	// myName := make([]string, 4)
	// myName[0] = "Hej"
	// myName[1] = "jag"
	// myName[2] = "heter"
	// myName[3] = "jesper"

	// for key, value := range myName{
	// 	fmt.Println(key, value)
	// }

	// for _, value := range myName {
	// 	fmt.Printf("%v \n", value)
	// }

	// for i := 0; i < len(myName); i++ {
	// 	fmt.Println(myName[i])
	// }

	//Nested slice

	// sliceOfSlice := make([][]int, 4)
	// sliceOfSlice[0] = []int{1, 2, 3, 4}
	// sliceOfSlice[1] = []int{1, 2, 3, 4}
	// sliceOfSlice[2] = []int{1, 2, 3, 4}
	// sliceOfSlice[3] = []int{1, 2, 3, 4}
	
	// fmt.Println(sliceOfSlice)

	// nested := make([][]int, 0, 3)
	// for i := 0; i < 3; i++ {
	// 	out := make([]int, 0, 4)
	// 	for j := 0; j < 4; j++ {
	// 		out = append(out, j)
	// 	}
	// 	nested = append(nested, out)
	// }
	// fmt.Println(nested)	

	// appleLaptops := []string{"MacbookPro", "MacbookAir"}
	// hpLaptops := []string{"hp650", "hpEliteBook"}
	// laptops := [][]string{appleLaptops, hpLaptops}
	// for i, v := range laptops {
	// 	fmt.Println("Record: ", i)
	// 	for _, name := range v {
	// 		fmt.Printf("\t Laptop name: %v \n", name)
	// 	}
	// }

	// Maps!
	sets := map[int]string{
		3: "Bänk",
		4: "Mark",
	}
	for key, value := range sets {
		fmt.Printf("Övning: %v Sets: %d\n", value, key)
	}
}