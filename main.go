package main
import ("fmt")

func main() {
  var cards = [14]int{2,3,4,5,6,7,8,9,10,11,12,13,14,15}
  var names = [14]string{"2","3","4","5","6","7","8","9","10","Jk","Q","K","A","Jr"}
  var values = [14]int{20,-400,5,5,5,5,10,10,10,10,10,10,20,50}

  type Player struct {
    const NUMBER int
    var table = [32]int{}
    var hand = [16]int{}
    var foot = [16]int{}
    var points int
  }
  
  func (p Player) goDown() {  
      fmt.Printf("Player ", p.NUMBER, "has gone down!")
  }
  
  fmt.Println("Get 5 books and play all your cards to 'go out' - the player with the most points after everyone has gone out wins.")
  
}
