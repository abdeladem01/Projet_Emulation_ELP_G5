package main

import (
	"math/rand"
  "time"
  "fmt"
)
func genavion(n int)[][]int{
  var l [][]int
  for i:=0;i<n;i++{
    var lu []int
    for i:=0;i<1;i++{
      rand.Seed(time.Now().UnixNano())
      min:=0
      max:=500
      a:=(rand.Intn(max - min + 1) + min)
      a=a*100
      b:=(rand.Intn(max -min +1) + min)
      b=b*100
      lu=append(lu,a)
      lu=append(lu,b)
      min2:=100
      max2:=400
      m:=(rand.Intn(max2 - min2 + 1) + min2)
      lu=append(lu,m)
    }
  l=append(l,lu)
  }
  return(l)
}
func main() {
  u:=genavion(5)
  fmt.Println(u)
}
