package main

import (
	"fmt"
	"math/rand"
	"time"
	"math"
	"os"
	"bufio" //bufferedIO
	"./src"
	. "./src"
)
func TourDeC(demande chan src.AnnonceP, changement []chan ChangeurP, grid *[Columns][Rows]int){
	for {
		p := <- demande //les positions des avions //à clarifier SpOOd
		if grid[p.Next_X][p.Next_Y] == 2 || grid[p.Next_X][p.Next_Y] == 3 { // la prochine case est un avion
			changement[p.Train_Id] <- ChangeurP {
				Previous_X : p.Actual_X, //translater le X de 1 ou le Y à voir
				Previous_Y : p.Actual_Y,
				Next_X : p.Actual_X,
				Next_Y : p.Actual_Y,
				}
		} else {
			if grid[p.Next_X][p.Next_Y] != 1 {
				grid[p.Next_X][p.Next_Y] = 3 //on réserve la prochaine case le cas echant
			}
			changement[p.Train_Id] <- ChangeurP {
				//Garder le même Z
				Previous_X : p.Actual_X,
				Previous_Y : p.Actual_Y,
				Next_X : p.Next_X,
				Next_Y : p.Next_Y,
				}
		}
	}
}


	
}
func main() { 
	//ToA : commencer par faire un titre pas pas important
	grid := src.GenGridArray()
	rand.Seed(time.Now().UnixNano())
	aeroports := src.GenAeroport()
	avions := src.GenAvion()
	fmt.Print("Taper Entrer")
	bufio.NewScanner(os.Stdin).Scan()
	//gridIG := src.balalalal Interface graphique
	//la faire passe dans un channel du genre//
	//gridIGch := make(chan string,100)
	//puis y envoyer la gridIG, jsp comment faire en GO :(
	fini := make(chan bool, len(avions))
	mutex := make(chan bool, 1)
	//Y envoyer true d'abord :(
	//bref requests de la position de lavion pour traçage	:
	requetesPositions := make(chan src.AnnonceP )
	//Creer slice de len(avions):
	instuctions := make([]chan src.ChangeurP, len(avions)) //ToC
	for i := range instructions {
   	instructions[i] = make(chan src.ChangeurP, 10)
	}
	go TourDeC(requetesPositions,instructions, &grid)
	for i := 0 ; i < len(avions) ; i++ {
		go BougerAvion(avions[i], &grid, grid_view_channel_ToC, done, f, mutex, requetesPositions, instructions)
	}
	for i := 0 ; i < len(avions) ; i++ {
		<- fini
	}

}