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
			changement[p.Train_Id] <- src.ChangeurP {
				CollisionPossible: true,
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
				CollisionPossible: false,
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

func BougerAvion(avion Avion, grid *[Columns][Rows]int, IGchan chan string, fini chan bool, mutex chan bool, requests chan Requested_Position, instructions []chan Instruction_Position) {
	for avion.X_position != avion.Arrival.X_position || avion.Y_position != avion.Arrival.Y_position {
		request_new_x := avion.X_position //ToC
		request_new_y := avion.Y_position //ToC
		if avion.X_position < avion.Arrival.X_position{
			request_new_x += 1
		} else if avion.X_position > avion.Arrival.X_position {
			request_new_x -= 1
		} else { request_new_x = avion.Arrival.X_position }
		if avion.Y_position < avion.Arrival.Y_position {
			request_new_y += 1
		} else if avion.Y_position > avion.Arrival.Y_position {
			request_new_y -= 1
		} else {request_new_y = avion.Arrival.Y_position}

		requests <- Requested_Position { //ToC
						Avion_Id : avion.Id, //Changer Id par Matricule
						Actual_X : avion.X_position,
						Actual_Y : avion.Y_position,
						Next_X : request_new_x,
						Next_Y : request_new_y,
					}
		instruction := <- instructions[avion.Id] //ToC			
			if grid[instruction.Next_X][instruction.Next_Y] != 1 && grid[instruction.Next_X][instruction.Next_Y] != 2 { //Ni aeroport ni station
				grid[instruction.Next_X][instruction.Next_Y] = 2
			}
			if grid[instruction.Previous_X][instruction.Previous_Y] != 1 { //Si juste pas aeroport
				grid[instruction.Previous_X][instruction.Previous_Y] = 0 //On dé-reserve la case reservé par avion
			}

			gridIG := <- IGchan
			gridIG = UpdateGridView(gridIG, grid, instruction.Previous_X, instruction.Previous_Y, instruction.Next_X, instruction.Next_Y)

			fmt.Println(gridIG) //Faudra peut etre lenlever et montrer letat que quand tt les avions finissent de bouger
			IGchan <- gridIG
			avion.X_position = instruction.Next_X
			avion.Y_position = instruction.Next_Y
		}
	    //time.Sleep(time.Millisecond * 5000) //Pour voir le lancement plus longtemps
	}
	fini <- true
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
	instructions := make([]chan src.ChangeurP, len(avions)) //ToC
	for i := range instructions {
   	instructions[i] = make(chan src.ChangeurP, 10)
	}
	go TourDeC(requetesPositions,instructions, &grid)
	for i := 0 ; i < len(avions) ; i++ {
		go BougerAvion(avions[i], &grid, IGchan, fini, mutex, requetesPositions, instructions)
	}
	for i := 0 ; i < len(avions) ; i++ {
		<- fini
	}

}

