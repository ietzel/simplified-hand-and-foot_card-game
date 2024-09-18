package main
import ("fmt")

func main() {
  cards := [14]int{2,3,4,5,6,7,8,9,10,11,12,13,14,15}
  names := [14]string{"2","3","4","5","6","7","8","9","10","Jk","Q","K","A","Jr"}
  values := [14]int{20,-400,5,5,5,5,10,10,10,10,10,10,20,50}

  type Player struct {
    number int
    table := [100]int{}
    hand := [20]int{}
    foot := [20]int{}
    points int
  }
  
  func (p Player) goDown() {  
      fmt.Printf("Player ", number, "has gone down!")
  }
  
  fmt.Println("Get 5 books and play all your cards to 'go out' - the player with the most points after everyone has gone out wins.")
  
}
