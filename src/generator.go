package src

import(
	"strconv"
	"math/rand"
	_ "reflect"
	"fmt"
	"os"
	"strings"
	"typesAndParam"
)

func GenGridArray() [Columns][Rows]int { //ToC
	var grid [Columns][Rows]int
	return grid
}

func GenGridSlice(R int, C int) [][]int { 
	slice := make([][]int, C)
	for i := range slice {
   		slice[i] = make([]int, R)
	}
	return slice
}


func GenAeroport(grid *[Columns][Rows]int) [nb_aero]Aeroport {

	//fmt.Println("Creation de " + strconv.Itoa(nb_aero) + " aeroports...\n") //Error checking
	aeroports := [nb_aero]Aeroport{}
	for i := 0 ; i < nb_aero ; i++ {
		x := rand.Intn(Columns)
		y := rand.Intn(Rows) //ToC
		aeroports[i] = Aeroport {
			Id: i,
			Name: string(65 + i),
			X_position: x, 
			Y_position: y,
		}
		
		grid[x][y] = 1 //1 aeroport, 2 avion etc
	}

	return aeroports
}

func GenAvion(aeroports [nb_aero]Aeroport) []Avion {

	avions := make([]Avion, 0)

/*	if nb_avion >= nb_aero * (nb_aero - 1) {
		fmt.Println("Creation de " + strconv.Itoa(nb_aero * (nb_aero - 1)) + " avions...\n")
	} else {
		fmt.Println("Creation de " + strconv.Itoa(nb_avion) + " avions...\n")
	}
*/ //Faire choix si 1 avion par aeroport et par voie ou pas?

	avions = make([]Avion, 0)

	maj_nbavion := 0 //déjà créé
//Si chaque avion a une trajection differente (C.A.)
	for j := 0 ; j < nb_aero ; j++ { //lignes
		for k := 0 ; k < nb_aero ; k++ { //colonnes
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
/* Pomper d'un site, à comprendre /!\ 
//TABLEAU DES DEPARTS D'AVIONS
	// Departures log file, save all the trajects
	f, err := os.Create("departures.log")

	if err != nil {
       	fmt.Println(err)
    }

    for i := 0; i < len(avions); i++ {	
    	
		f.WriteString("Avion " + strconv.Itoa(i) + " : From station " + avions[i].Departure.Name + " to " + avions[i].Arrival.Name + "\n")
	}
	
	f.Close()
*/

	return avions
}
