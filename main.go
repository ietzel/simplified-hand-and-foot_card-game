package main
import ("fmt" "runtime")

func () round() {

}

function (array []int) goOut {
    for i := 0; i < len(array); i++ {
        fmt.Println(array[i])
    }
}

func (number) turn {
    switch num := number; num {
	case 1:
		cardsA = cardsA.append(RandomIntegerwithinRange := rand.Intn(13));
	case 2:
		cardsB = cardsB.append(RandomIntegerwithinRange := rand.Intn(13));
    	case 3:
		cardsC = cardsC.append(RandomIntegerwithinRange := rand.Intn(13));
    	case 4:
		cardsD = cardsD.append(RandomIntegerwithinRange := rand.Intn(13));
	case 5:
		cardsE = cardsE.append(RandomIntegerwithinRange := rand.Intn(13));
    	case 6:
		cardsF = cardsF.append(RandomIntegerwithinRange := rand.Intn(13));
	default:
		fmt.Printf("No player has decided.");
    }
}

func main() {
  fmt.Printf("Less than 7 of a type of card is considered a hand book, otherwise a foot book.")
  var cards = [14]int{2,3,4,5,6,7,8,9,10,11,12,13,14,15}
  var names = [14]string{"2","3","4","5","6","7","8","9","10","Jk","Q","K","A","Jr"}
  var values = [14]int{20,-400,5,5,5,5,10,10,10,10,10,10,20,50}

  var cardsA [14]int
  var cardsB [14]int
  var cardsC [14]int
  var cardsD [14]int
  var cardsE [14]int
  var cardsF [14]int
  var points [6]int

  for j := 0; j <= 31; j++ {
    cardsA[RandomIntegerwithinRange := rand.Intn(13)]++
    cardsB[RandomIntegerwithinRange := rand.Intn(13)]++
    cardsC[RandomIntegerwithinRange := rand.Intn(13)]++
    cardsD[RandomIntegerwithinRange := rand.Intn(13)]++
    cardsE[RandomIntegerwithinRange := rand.Intn(13)]++
    cardsF[RandomIntegerwithinRange := rand.Intn(13)]++
  }

  turn := 0
  for {
    for i := 1; i <= 6; i++ {
    	turn(i);
    }
    fmt.Println("turn ", turn);
    turn++
  } 
}
