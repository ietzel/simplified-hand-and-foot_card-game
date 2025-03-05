package main

import (
	"fmt"
)

var cards = [14]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
var names = [14]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jk", "Q", "K", "A", "Jr"}
var values = [14]int{20, -400, 5, 5, 5, 5, 10, 10, 10, 10, 10, 10, 20, 50}

var cardsA [14]int
var cardsB [14]int
var cardsC [14]int
var cardsD [14]int
var cardsE [14]int
var cardsF [14]int
var points [6]int

func countOccurrences(arr []int, x int) int {
	count := 0
	for _, num := range arr {
		if num == x {
			count++
		}
	}
	return count
}

func removeSameNumberValues(number int, slice []int) []int {
	/*
		if countOccurences(slice, number) > 7 {
			var result []int
			for _, num := range slice {
				if num != number {
					result = append(result, num)
				}
			}
			return result
		}
	*/
	return slice
}

/*
func bookcheck(player) {
	switch num := number; num {
	case 1:
		cardsA = cardsA.append(RandomIntegerwithinRange := rand.Intn(13))
		for i := 0; i <= 13; i++ {
			cardsA := removeSameNumberValues(j, cardsA)
		}
	case 2:
		cardsB = cardsB.append(RandomIntegerwithinRange := rand.Intn(13))
		for i := 0; i <= 13; i++ {
			cardsB := removeSameNumberValues(j, cardsB)
		}
    case 3:
		cardsC = cardsC.append(RandomIntegerwithinRange := rand.Intn(13))
		for i := 0; i <= 13; i++ {
			cardsC := removeSameNumberValues(j, cardsC)
		}
    case 4:
		cardsD = cardsD.append(RandomIntegerwithinRange := rand.Intn(13))
		for i := 0; i <= 13; i++ {
			cardsD := removeSameNumberValues(j, cardsD)
		}
	case 5:
		cardsE = cardsE.append(RandomIntegerwithinRange := rand.Intn(13))
		for i := 0; i <= 13; i++ {
			cardsE := removeSameNumberValues(j, cardsE)
		}
    case 6:
		cardsF = cardsF.append(RandomIntegerwithinRange := rand.Intn(13))
		for i := 0; i <= 13; i++ {
			cardsF := removeSameNumberValues(j, cardsF)
		}
	default:
		fmt.Printf("..........");
    }
}
*/

var outingplayer int = 0

func turn(number int) {

	/*
		    switch num := number; num {
			case 1:
				fmt.Printf(cardsA)
				cardsA = cardsA.append(RandomIntegerwithinRange := rand.Intn(13))
			        if((cardsA).len <= 2) {
						outingplayer = 1;
					}
			case 2:
				fmt.Printf(cardsB)
				cardsB = cardsB.append(RandomIntegerwithinRange := rand.Intn(13))
			        if((cardsA).len <= 2) {
						outingplayer = 2;
					}
		    case 3:
				fmt.Printf(cardsC)
				cardsC = cardsC.append(RandomIntegerwithinRange := rand.Intn(13))
			        if((cardsA).len <= 2) {
						outingplayer = 3;
					}
		    case 4:
				cardsD = cardsD.append(RandomIntegerwithinRange := rand.Intn(13))
				fmt.Printf(cardsD)
			        if((cardsA).len <= 2) {
						outingplayer = 4;
					}
			case 5:
				cardsE = cardsE.append(RandomIntegerwithinRange := rand.Intn(13))
				fmt.Printf(cardsE)
			        if((cardsA).len <= 2) {
						outingplayer = 5;
					}
		    case 6:
				cardsF = cardsF.append(RandomIntegerwithinRange := rand.Intn(13))
				fmt.Printf(cardsF)
			        if((cardsA).len <= 2) {
						outingplayer = 6;
					}
			default:
				fmt.Printf("No player has decided.")
		    }
	*/
}

func main() {

	for j := 0; j <= 31; j++ {
		/*
		   cardsA[RandomIntegerwithinRange := rand.Intn(13)]++
		   cardsB[RandomIntegerwithinRange := rand.Intn(13)]++
		   cardsC[RandomIntegerwithinRange := rand.Intn(13)]++
		   cardsD[RandomIntegerwithinRange := rand.Intn(13)]++
		   cardsE[RandomIntegerwithinRange := rand.Intn(13)]++
		   cardsF[RandomIntegerwithinRange := rand.Intn(13)]++
		*/
	}

	turn := 0
	for {
		for i := 1; i <= 6; i++ {
			//turn(i)
		}
		fmt.Println("round ", turn)
		turn++
	}
}
