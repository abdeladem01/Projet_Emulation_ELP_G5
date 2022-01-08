package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
  "io"
	"math/rand"
	"time"
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
func GenGridArray() [long][larg]int { //ToC
	var grid [long][larg]int
	return grid
}

func GenIG(aeroports [nb_aero]Aeroport) string { 
	visu := "" //si probleme dans cette fonction, ajouter Obstacle_size
	for i := 0 ; i < larg ; i++ {
		visu += strings.Repeat(".", long) + "\n"
	}
	visu_lrg := long + 1
	for i := 0 ; i < len(aeroports) ; i++ { //affichage en grille
		XIG := aeroports[i].X_position
		YIG := aeroports[i].Y_position
		visu = visu[:visu_lrg * YIG + XIG] + aeroports[i].Name + visu[visu_lrg * YIG + XIG + 1:]
	}
	return visu
}

func GenAeroport(grid *[long][larg]int) [nb_aero]Aeroport {
	//fmt.Println("Creation de " + strconv.Itoa(nb_aero) + " aeroports...\n") //Error checking
	aeroports := [nb_aero]Aeroport{}
	for i := 0 ; i < nb_aero ; i++ {
		x := rand.Intn(long)
		y := rand.Intn(larg) //ToC
		aeroports[i] = Aeroport {
			Id: i,
			Name: string(65 + i), //string 65 étant A
			X_position: x, 
			Y_position: y,
		}
		grid[x][y] = 1 //1 aeroport, 2 avion etc
	}
	return aeroports
}

func GenAvion(aeroports [nb_aero]Aeroport,conn net.Conn, nb_avion int) []Avion {
	avions := make([]Avion, 0)
 //Faire choix si 1 avion par aeroport et par voie ou pas?
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
    for i := 0; i < len(avions); i++ {	
    	

    io.WriteString(conn,"Le vol ELP010" + strconv.Itoa(i) + " va effectuer son départ de l'aéroport " + avions[i].Departure.Name + " à destination de l'aéroport " + avions[i].Arrival.Name + ".\n La température attendue à la ville "+ avions[i].Arrival.Name + " est de "+strconv.Itoa(rand.Intn(35))+"°C.\n")
		}
	io.WriteString(conn,"\n")
	io.WriteString(conn,"    ____ \n")
	io.WriteString(conn,"   / __ )____  ____               _   ______  __  ______ _____ ____ \n")
	io.WriteString(conn,"  / __  / __ \\/ __ \\             | | / / __ \\/ / / / __ `/ __ `/ _ \\\n")
	io.WriteString(conn," / /_/ / /_/ / / / /             | |/ / /_/ / /_/ / /_/ / /_/ /  __/\n")
	io.WriteString(conn,"/_____/\\____/_/ /_/_ __   _____  |___/\\____/\\__, /\\__,_/\\__, /\\___/ \n")
	io.WriteString(conn,"               / __ `/ | / / _ \\/ ___/     /____/      /____/       \n")
	io.WriteString(conn,"              / /_/ /| |/ /  __/ /__                                \n")
	io.WriteString(conn,"              \\__,_/ |___/\\___/\\___/                                \n")
	io.WriteString(conn,"     ______ _      _____        _      \n")
	io.WriteString(conn,"    |  ____| |    |  __ \\ /\\   (_)            ____       _\n")
	io.WriteString(conn,"    | |__  | |    | |__) /  \\   _ _ __      |__\\\\_\\_o,___/ \\\n")
	io.WriteString(conn,"    |  __| | |    |  ___/ /\\ \\ | | '__|    ([___\\_\\_____-\\'\n")
	io.WriteString(conn,"    | |____| |____| |  / ____ \\| | |        | o'\n")
	io.WriteString(conn,"    |______|______|_| /_/    \\_\\_|_|  \n")
	io.WriteString(conn,"\n")
	return avions
}
 //Programme Principal

func TourDeC(demande chan AnnonceP, changement []chan ChangeurP, grid *[long][larg]int){
	for {
		p := <- demande //les positions des avions //à clarifier SpOOd
		if grid[p.Next_X][p.Next_Y] == 2 || grid[p.Next_X][p.Next_Y] == 3 { // la prochine case est un avion
			changement[p.IdentifAv] <- ChangeurP {
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

func MaJIG(grid_view string, grid *[long][larg]int, previous_x int, previous_y int, actual_x int, actual_y int) string { //ToC

	visu_lrg := long + 1 //largeur de la grille visuellement

	preXIG := previous_x 
	preYIG := previous_y 
	if grid[previous_x][previous_y] != 1 {
		grid_view = grid_view[:visu_lrg * preYIG + preXIG] + "." + grid_view[visu_lrg * preYIG + preXIG + 1:]
	}
	if grid[actual_x][actual_y] != 1 {
		grid_view = grid_view[:visu_lrg * actual_y + actual_x] + "W" + grid_view[visu_lrg * actual_y + actual_x + 1:]
	}
	return grid_view
}

func BougerAvion(avion Avion, grid *[long][larg]int, IGchan chan string, fini chan bool, requests chan AnnonceP, instructions []chan ChangeurP,conn net.Conn) {
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
		if grid[instruction.Next_X][instruction.Next_Y] == 2  { 
				io.WriteString(conn,"L'avion du vol ELP0"+strconv.Itoa(avion.Id+100)+ " a changé d'altitude pour éviter un avion\n")
			}		
			if grid[instruction.Next_X][instruction.Next_Y] != 1 && grid[instruction.Next_X][instruction.Next_Y] != 2 { //Ni aeroport ni station
				grid[instruction.Next_X][instruction.Next_Y] = 2
			}
			if grid[instruction.Next_X][instruction.Next_Y] == 1  { 
				io.WriteString(conn,"L'avion du vol ELP0"+strconv.Itoa(avion.Id+100)+ " va attérir à l'aeroport de la ville "+avion.Arrival.Name+"\n")
			}
			if grid[instruction.Previous_X][instruction.Previous_Y] != 1 { //Si juste pas aeroport
				grid[instruction.Previous_X][instruction.Previous_Y] = 0 //On dé-reserve la case reservé par avion
			}
			

			gridIG := <- IGchan
			gridIG = MaJIG(gridIG, grid, instruction.Previous_X, instruction.Previous_Y, instruction.Next_X, instruction.Next_Y)

     	io.WriteString(conn,gridIG+"\n")

			IGchan <- gridIG
			avion.X_position = instruction.Next_X
			avion.Y_position = instruction.Next_Y
		
		}

	    
	    fini <- true
}
	


func maingo(conn net.Conn, nb_avion int) { 
	//ToA : commencer par faire un titre pas pas important
	grid := GenGridArray()
	rand.Seed(time.Now().UnixNano())
	aeroports := GenAeroport(&grid)
	avions := GenAvion(aeroports,conn,nb_avion)
  io.WriteString(conn,"######## Taper 'Entrée' pour lancer la simulation ########\n")
  read:= bufio.NewReader(conn)
  for {
      result, err := read.ReadString('\n')
        if (err != nil){
          fmt.Printf("DEBUG MAIN could not read from client")
          os.Exit(1)
            	}
        if result=="oui\n"{
          break
        }
  }
	gridIG := GenIG(aeroports)
	IGchan := make(chan string, 100) //grid_view_channel pour pas se perdre
	IGchan <- gridIG
	fini := make(chan bool, len(avions))
	//Y envoyer true d'abord :(
	//bref requests de la position de lavion pour traçage	:
	requetesPositions := make(chan AnnonceP )
	//Creer slice de len(avions):
	instructions := make([]chan ChangeurP, len(avions)) //ToC
	for i := range instructions {
   	instructions[i] = make(chan ChangeurP, 10)
	}
	go TourDeC(requetesPositions,instructions, &grid)
	for i := 0 ; i < len(avions) ; i++ {
		go BougerAvion(avions[i], &grid, IGchan, fini, requetesPositions, instructions,conn)
	}
	for i := 0 ; i < len(avions) ; i++ {
		<- fini
	}

}
func getArgs() int {

	if len(os.Args) != 2 {
		fmt.Printf("Usage: go run server.go <portnumber>\n")
		os.Exit(1)
	} else {
		fmt.Printf("#DEBUG ARGS Port Number : %s\n", os.Args[1])
		portNumber, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Usage: go run server.go <portnumber>\n")
			os.Exit(1)
		} else {
			return portNumber
		}

	}
	return -1
}

func main() {
	port := getArgs()
	fmt.Printf("#DEBUG MAIN Creating TCP Server on port %d\n", port)
	portString := fmt.Sprintf(":%s", strconv.Itoa(port))
	fmt.Printf("#DEBUG MAIN PORT STRING |%s|\n", portString)

	ln, err := net.Listen("tcp", portString)
	if err != nil {
		fmt.Printf("#DEBUG MAIN Could not create listener\n")
		panic(err)
	}

	//If we're here, we did not panic and ln is a valid listener

    connum := 1

	for {
		fmt.Printf("#DEBUG MAIN Accepting next connection\n")
		conn, errconn := ln.Accept()

		if errconn != nil {
			fmt.Printf("DEBUG MAIN Error when accepting next connection\n")
			panic(errconn)

		}

		//If we're here, we did not panic and conn is a valid handler to the new connection

		go handleConnection(conn, connum)
        connum +=1

	}
}

func handleConnection(connection net.Conn, connum int) {
	defer connection.Close()
	connReader:= bufio.NewReader(connection)
		inputLine, err := connReader.ReadString('\n')
resultString := strings.TrimSuffix(inputLine, "\n")
    intinput,err := strconv.Atoi(resultString)
    if err!= nil{
      os.Exit(1)
    }
		if err != nil {
			fmt.Printf("#DEBUG %d RCV ERROR no panic, just a client\n", connum)
            fmt.Printf("Error :|%s|\n", err.Error())
		}
    nb_avion:=intinput
    maingo(connection,nb_avion)
}
