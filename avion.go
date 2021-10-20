.gpackage avion

import (
	"fmt"
	"math/rand"
	"os"
  "time"
)
def genavion(nb){
  rand.Seed(time.Now().UnixNano())
  min:=0
  max:=100
  fmt.Println(rand.Intn(max - min + 1) + min)
}