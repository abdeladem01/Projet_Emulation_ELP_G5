package main;

import (
	"strings"
	"time"
	"math/rand"
	"fmt"
	"strconv"
)

// MODIFICATION AUTORISÉE POUR LA PARTIE SUIVANTE :)
var s = []int{3, 5 , 40 , 7}
var nb_avions = s[0]
var nb_aero = s[1]
var long = s[2]
var larg = s[3]
// NE PAS MODIFIER LES PARTIES SUIVANTES /!\ //ToC
//Definiton des types
type Avion struct {
	Id int
	X_position int
	Y_position int
	Departure Aeroport
	Arrival Aeroport
}

type Aeroport struct {
	Id int
	Name string
	X_position int
	Y_position int
}
type AnnonceP struct {
	IdentifAv int
	Actual_X int
	Actual_Y int
	Next_X int
	Next_Y int
}

type ChangeurP struct {
	CollisionPossible bool
	Previous_X int
	Previous_Y int
	Next_X int
	Next_Y int
}

//Definition des generateur d'élement
func GenGridSlice() [][]int { //ToC
	grid := make([][]int , long)
	 for i := 0; i < long; i++ {
        grid[i] = make([]int, larg)
	
	 }
	 return grid
}

func GenIG(aeroports []Aeroport) string { 
	visu := "" //si probleme dans cette fonction, ajouter Obstacle_size
	for i := 0 ; i < larg ; i++ {
		visu += strings.Repeat(".", long) + "\n"
	}
	visu_lrg := long + 1
	for i := 0 ; i < nb_aero ; i++ { //affichage en grille
		XIG := aeroports[i].X_position
		YIG := aeroports[i].Y_position
		visu = visu[:visu_lrg * YIG + XIG] + aeroports[i].Name + visu[visu_lrg * YIG + XIG + 1:]
	}
	return visu
}

func GenAeroport(grid *[][]int) []Aeroport {
	//fmt.Println("Creation de " + strconv.Itoa(nb_aero) + " aeroports...\n") //Error checking
	aeroports := make([]Aeroport, nb_aero)
	for i := 0 ; i < nb_aero ; i++ {
		x := rand.Intn(len(*grid))
		y := rand.Intn(len((*grid)[0])) //ToC
		aeroports[i] = Aeroport {
			Id: i,
			Name: string(65 + i), //string 65 étant A
			X_position: x, 
			Y_position: y,
		}
		(*grid)[x][y] = 1 //1 aeroport, 2 avion etc
	}
	return aeroports
}

func GenAvion(aeroports []Aeroport, nb_avion int) []Avion {
	avions := make([]Avion, 0)
	avions = make([]Avion, 0)
	maj_nbavion := 0 //déjà créé
	for j := 0 ; j < len(aeroports) ; j++ { //lignes
		for k := 0 ; k < len(aeroports) ; k++ { //colonnes
			if k != j {	
				avions = append(avions, Avion{
					Id: maj_nbavion,
					X_position: aeroports[j].X_position,
					Y_position: aeroports[j].Y_position,
					Departure: aeroports[j], //ToC
					Arrival: aeroports[k], //ToC
				})
				maj_nbavion++
			}
			if maj_nbavion > nb_avion - 1 { //incrementer les deux
				k = nb_aero + 1
				j = nb_aero + 1
			}
		}
	}
    for i := 0; i < len(avions); i++ {	
    fmt.Print("Le vol ELP010" + strconv.Itoa(i) + " va effectuer son départ de l'aéroport " + avions[i].Departure.Name + " à destination de l'aéroport " + avions[i].Arrival.Name + ".\n La température attendue à la ville "+ avions[i].Arrival.Name + " est de "+strconv.Itoa(rand.Intn(35))+"°C.\n")
		}
	return avions
}

 //Programme Principal

func TourDeC(demande chan AnnonceP, changement []chan ChangeurP, grid *[][]int){
	for {
		p := <- demande //les positions des avions //à clarifier SpOOd
		if (*grid)[p.Next_X][p.Next_Y] == 2 || (*grid)[p.Next_X][p.Next_Y] == 3 { // la prochine case est un avion
			changement[p.IdentifAv] <- ChangeurP {
				CollisionPossible: true,
				Previous_X : p.Actual_X, //translater le X de 1 ou le Y à voir
				Previous_Y : p.Actual_Y,
				Next_X : p.Actual_X,
				Next_Y : p.Actual_Y,
				}
		} else {
			if (*grid)[p.Next_X][p.Next_Y] != 1 {
				(*grid)[p.Next_X][p.Next_Y] = 3 //on réserve la prochaine case le cas echant
			}
			changement[p.IdentifAv] <- ChangeurP {
				CollisionPossible: false,
				Previous_X : p.Actual_X,
				Previous_Y : p.Actual_Y,
				Next_X : p.Next_X,
				Next_Y : p.Next_Y,
				}
		}
	}
}
func MaJIG(grid_view string, grid *[][]int, previous_x int, previous_y int, actual_x int, actual_y int) string { //ToC
	visu_lrg := long + 1 //largeur de la grille visuellement
	preXIG := previous_x 
	preYIG := previous_y 
	if (*grid)[previous_x][previous_y] != 1 {
		grid_view = grid_view[:visu_lrg * preYIG + preXIG] + "." + grid_view[visu_lrg * preYIG + preXIG + 1:]
	}
	if (*grid)[actual_x][actual_y] != 1 {
		grid_view = grid_view[:visu_lrg * actual_y + actual_x] + "W" + grid_view[visu_lrg * actual_y + actual_x + 1:]
	}
	return grid_view
}
func BougerAvion(avion Avion, grid *[][]int, IGchan chan string, fini chan bool, requests chan AnnonceP, instructions []chan ChangeurP) {
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

		requests <- AnnonceP { //ToC
						IdentifAv : avion.Id, //Changer Id par Matricule
						Actual_X : avion.X_position,
						Actual_Y : avion.Y_position,
						Next_X : request_new_x,
						Next_Y : request_new_y,
					}
		instruction := <- instructions[avion.Id] //ToC	
		if (*grid)[instruction.Next_X][instruction.Next_Y] == 2  { 
				fmt.Print("L'avion du vol ELP0"+strconv.Itoa(avion.Id+100)+ " a changé d'altitude pour éviter un avion\n")
			}		
			if (*grid)[instruction.Next_X][instruction.Next_Y] != 1 && (*grid)[instruction.Next_X][instruction.Next_Y] != 2 { //Ni aeroport ni station
				(*grid)[instruction.Next_X][instruction.Next_Y] = 2
			}
			if (*grid)[instruction.Next_X][instruction.Next_Y] == 1  { 
				fmt.Print("L'avion du vol ELP0"+strconv.Itoa(avion.Id+100)+ " va attérir à l'aeroport de la ville "+avion.Arrival.Name+"\n")
			}
			if (*grid)[instruction.Previous_X][instruction.Previous_Y] != 1 { //Si juste pas aeroport
				(*grid)[instruction.Previous_X][instruction.Previous_Y] = 0 //On dé-reserve la case reservé par avion
			}
			

			gridIG := <- IGchan
			gridIG = MaJIG(gridIG, grid, instruction.Previous_X, instruction.Previous_Y, instruction.Next_X, instruction.Next_Y)

     	fmt.Print(gridIG+"\n")

			IGchan <- gridIG
			avion.X_position = instruction.Next_X
			avion.Y_position = instruction.Next_Y
		
		}

	    
	    fini <- true
}

func rien([]Avion){
	i := 3
switch i {
case 3:
case 0:
    fmt.Println("Hello, playground")
}
}
func main(){
	grid := GenGridSlice()
	rand.Seed(time.Now().UnixNano())
	aeroports := GenAeroport(&grid)
	avions := GenAvion(aeroports,nb_avions)
	rien(avions)
		gridIG := GenIG(aeroports)
	IGchan := make(chan string, 100) //grid_view_channel pour pas se perdre
	IGchan <- gridIG
	fini := make(chan bool, len(avions))
	requetesPositions := make(chan AnnonceP )
	instructions := make([]chan ChangeurP, len(avions)) //ToC
	for i := range instructions {
   	instructions[i] = make(chan ChangeurP, 10)
	}
	time.Sleep(150 * time.Second)
	go TourDeC(requetesPositions,instructions, &grid)
	for i := 0 ; i < len(avions) ; i++ {
		go BougerAvion(avions[i], &grid, IGchan, fini, requetesPositions, instructions)
	}
	for i := 0 ; i < len(avions) ; i++ {
		<- fini
	}
}