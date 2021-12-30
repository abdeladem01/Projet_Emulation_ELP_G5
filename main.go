package main

import (
	"fmt"
	"math/rand"
	"time"
	"math"
	"os"
	"bufio" //bufferedIO
	"src"
	. "src"
)
func TourDeC(){


	
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