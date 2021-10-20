package main

import (
	"fmt"
	"math/rand"
  "time"
  "strconv"
)
func genavion(n int){
  var l []string
  var a int 
  var b int
  var k int
  var p int
  for i:=0;i<n;i++{
    var lu []string
    for i:=0;i<2;i++{
      rand.Seed(time.Now().UnixNano())
      min:=0
      max:=500
      a:=(rand.Intn(max - min + 1) + min)
      b:=(rand.Intn(max -min +1) + min)
      k := strconv.Itoa(a)
      p:= strconv.Itoa(b)
      lu=append(lu,a)
      lu=append(lu,b)
    }
  l.append(lu)
  }
  fmt.Println(l)
}
func main() {
  genavion(5)
}
