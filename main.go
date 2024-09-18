package main
import ("fmt")

type Player struct {
  var number int
  var table = [32]int{}
  var hand = [16]int{}
  var foot = [16]int{}
  var points int
}

func (p Player) goDown() {  
    fmt.Printf("Player ", p.number, "has gone down!")
}

func main() {
  var cards = [14]int{2,3,4,5,6,7,8,9,10,11,12,13,14,15}
  var names = [14]string{"2","3","4","5","6","7","8","9","10","Jk","Q","K","A","Jr"}
  var values = [14]int{20,-400,5,5,5,5,10,10,10,10,10,10,20,50}

  var p1 Player
  var p2 Player
  var p3 Player
  var p4 Player
  var p5 Player
  var p6 Player

  p1.number = 1
  p1.table = [32]int{}
  p1.hand = [16]int{}
  p1.foot = [16]int{}
  p1.points = 0

  p1.number = 2
  p1.table = [32]int{}
  p1.hand = [16]int{}
  p1.foot = [16]int{}
  p1.points = 0

  p1.number = 3
  p1.table = [32]int{}
  p1.hand = [16]int{}
  p1.foot = [16]int{}
  p1.points = 0

  p1.number = 4
  p1.table = [32]int{}
  p1.hand = [16]int{}
  p1.foot = [16]int{}
  p1.points = 0

  p1.number = 5
  p1.table = [32]int{}
  p1.hand = [16]int{}
  p1.foot = [16]int{}
  p1.points = 0

  p1.number = 6
  p1.table = [32]int{}
  p1.hand = [16]int{}
  p1.foot = [16]int{}
  p1.points = 0
  
  fmt.Println("Get 5 books and play all your cards to 'go out' - the player with the most points after everyone has gone out wins.")
}
