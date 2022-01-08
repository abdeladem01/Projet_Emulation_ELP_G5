package main;

import (
	"strings"
	"time"
	"math/rand"
	"fmt"
)
// MODIFICATION AUTORISÉE POUR LA PARTIE SUIVANTE :)
const nb_aero = 5
const long = 40
const larg = 7
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
func rien([]Aeroport){
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
	rien(aeroports)
}