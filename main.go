package main

import (
	"fmt"
	"math/rand"
	"time"
	"math"
	"os"
	"bufio" //bufferedIO
	"src" //regler probleme dimport du local ou modifier le go env
	. "src"
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
			changement[p.Train_Id] <- src.ChangeurP {
				Previous_X : p.Actual_X,
				Previous_Y : p.Actual_Y,
				Next_X : p.Next_X,
				Next_Y : p.Next_Y,
				}
		}
	}
}

func MaJIG(grid_view string, grid *[Columns][Rows]int, previous_x int, previous_y int, actual_x int, actual_y int) string { //ToC

	visu_lrg := Columns + 1 //largeur de la grille visuellement

	preXIG := previous_x * (Obstacle_size + 1)
	preYIG := previous_y * (Obstacle_size + 1)
	if grid[previous_x][previous_y] != 1 {
		grid_view = grid_view[:visu_lrg * preYIG + preXIG] + "." + grid_view[visu_lrg * preYIG + preXIG + 1:]
	}
	if grid[actual_x][actual_y] != 1 {
		grid_view = grid_view[:visu_lrg * actual_y + actual_x] + "W" + grid_view[visu_lrg * actual_y + actual_x + 1:]
	}
	return grid_view
}

	

func main() { 
	//ToA : commencer par faire un titre pas pas important
	grid := src.GenGridArray()
	rand.Seed(time.Now().UnixNano())
	aeroports := src.GenAeroport()
	avions := src.GenAvion()
	fmt.Print("Taper Entrer")
	bufio.NewScanner(os.Stdin).Scan()
	gridIG := src.GenIG(aeroports)
	IGchan := make(chan string, 100) //grid_view_channel pour pas se perdre
	IGchan <- gridIG
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

