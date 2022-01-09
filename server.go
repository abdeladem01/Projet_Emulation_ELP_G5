package main

//Importation des packages nécessaires
import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

//Definiton des types d'objets qui vont être utilisés
//Definition du type Avion
type Avion struct { //
	Id       int      //Un avion définit par
	posX     int      //sa position x actuelle
	posY     int      //sa position y actuelle
	Depart   Aeroport //Aeroport de départ
	Arrivee  Aeroport //Aeroport d'arrivée
	Tmps_vol int      //Temps de vol, sera initialisé à zero
}

//Définition du type Aeroport
type Aeroport struct { //définit par
	Id   int    //un identifiant
	Name string //le nom de la ville
	posX int    //position x de laeroport (fixe)
	posY int    //position y de laeroport (fixe)
}

//Defintion du type Annonce Position
//Ce type est un objet qui permet de selectionner le prochain
// couple (x,y) de lavion en fonction de ses actuels positions
type AnnonceP struct {
	IdentifAv int //Identifiant de l'avion concerné par cette annonce
	nowX      int
	nowY      int
	nextX     int
	nextY     int
}

//Definition du type changeur de position
// permettra de changer la position de lavion enfonction de ce
//qui a été annoncé (par AnnonceP)
//On y integre également un boolean qui change de
//true à false si la prochaine case est contenu par avion
type ChangeurP struct {
	CollisionPossible bool //utile pour gerer les collisions et indiquer la prescence prochaine d'un avion dans les parages
	precX             int
	precY             int
	nextX             int
	nextY             int
} //CollisionPossible nt usd faudra les z coordo

//Definition des generateur d'élement
//Fonction qui génére le slice sur le quel sera définit notre espace aerien
func GenGridSlice(long int, larg int) [][]int {
	grid := make([][]int, long)
	for i := 0; i < long; i++ {
		grid[i] = make([]int, larg)
	}
	return grid
}
//Code dans le slice de slice : code:   0=> Slot vide ; 1 => Slot aeroport ; 2 => Avion dans ce slot; 3 => intermidiate reservation

//Genere une visualitation graphique de cette slice of slice 
func GenIG(aeroports []Aeroport, nb_aero, long int, larg int) string {
	visu := "" //si probleme dans cette fonction, ajouter Obstacle_size
	for i := 0; i < larg; i++ {
		visu += strings.Repeat(".", long) + "\n" //remplir de point dans tous les cas
	}
	visu_lrg := long + 1 //largeur de l'IG
	for i := 0; i < nb_aero; i++ { 
		XIG := aeroports[i].posX
		YIG := aeroports[i].posY
		visu = visu[:visu_lrg*YIG+XIG] + aeroports[i].Name + visu[visu_lrg*YIG+XIG+1:]
		//puis remplacer les points par les lettres des aeroports dans les endroits données
	}
	return visu //l'objet string a renvoyer
}
//Code : . => vide, , {A->V} Airport, W => Avion

//Genere un nombre demandé d'aeroports
func GenAeroport(grid *[][]int, nb_aero int) []Aeroport {
	//fmt.Println("Creation de " + strconv.Itoa(nb_aero) + " aeroports.\n") //Error checking
	aeroports := make([]Aeroport, nb_aero) //création slice dobjets de type Aeroport
	for i := 0; i < nb_aero; i++ {
		x := rand.Intn(len(*grid)) //on choisit la x pos de aero en fonction de la taille de la grille
		y := rand.Intn(len((*grid)[0])) //de maniere aleatoire
		aeroports[i] = Aeroport{ //on cree un objet aeroport et on le met dans la liste
			Id:   i,
			Name: string(65 + i), //string 65 étant A
			posX: x,
			posY: y,
		}
		(*grid)[x][y] = 1 //1 aeroport, 2 avion etc selon le codage definit precedemment
	}
	return aeroports
}

func GenAvion(aeroports []Aeroport, conn net.Conn, nb_avion int, nb_aero int) []Avion {
	avions := make([]Avion, 0)
	avions = make([]Avion, 0)
	maj_nbavion := 0                      //nb avions déjà créé
	for j := 0; j < len(aeroports); j++ { 
		for k := 0; k < len(aeroports); k++ { 
			if k != j { //la destination != l'arrivee
				avions = append(avions, Avion{
					Id:       maj_nbavion,
					posX:     aeroports[j].posX,
					posY:     aeroports[j].posY,
					Depart:   aeroports[j],
					Arrivee:  aeroports[k],
					Tmps_vol: 0,
				})
				maj_nbavion++
			}
			if maj_nbavion > nb_avion-1 { //incrementer les deux
				k = nb_aero + 1
				j = nb_aero + 1
			}
		}
	}
	for i := 0; i < len(avions); i++ {

		io.WriteString(conn, "Le vol ELP0"+strconv.Itoa(i+100)+" va effectuer son départ de l'aéroport "+avions[i].Depart.Name+" à destination de l'aéroport "+avions[i].Arrivee.Name+".\n La température attendue à la ville "+avions[i].Arrivee.Name+" est de "+strconv.Itoa(rand.Intn(35))+"°C.\n")
	}
	io.WriteString(conn, "\n")
	io.WriteString(conn, "    ____ \n")
	io.WriteString(conn, "   / __ )____  ____               _   ______  __  ______ _____ ____ \n")
	io.WriteString(conn, "  / __  / __ \\/ __ \\             | | / / __ \\/ / / / __ `/ __ `/ _ \\\n")
	io.WriteString(conn, " / /_/ / /_/ / / / /             | |/ / /_/ / /_/ / /_/ / /_/ /  __/\n")
	io.WriteString(conn, "/_____/\\____/_/ /_/_ __   _____  |___/\\____/\\__, /\\__,_/\\__, /\\___/ \n")
	io.WriteString(conn, "               / __ `/ | / / _ \\/ ___/     /____/      /____/       \n")
	io.WriteString(conn, "              / /_/ /| |/ /  __/ /__                                \n")
	io.WriteString(conn, "              \\__,_/ |___/\\___/\\___/                                \n")
	io.WriteString(conn, "     ______ _      _____        _      \n")
	io.WriteString(conn, "    |  ____| |    |  __ \\ /\\   (_)            ____       _\n")
	io.WriteString(conn, "    | |__  | |    | |__) /  \\   _ _ __      |__\\\\_\\_o,___/ \\\n")
	io.WriteString(conn, "    |  __| | |    |  ___/ /\\ \\ | | '__|    ([___\\_\\_____-\\'\n")
	io.WriteString(conn, "    | |____| |____| |  / ____ \\| | |        | o'\n")
	io.WriteString(conn, "    |______|______|_| /_/    \\_\\_|_|  \n")
	io.WriteString(conn, "\n")
	return avions
}

//Programme Principal des go-routines
//Tour de controle, où sont diriger les avions por eviter collision
//en gros, ici les avions reserve les prochaines cases, ou detectent lapproche d'un aeroport
func TourDeC(demande chan AnnonceP, changement []chan ChangeurP, grid *[][]int) {
	for {
		p := <-demande                                                        //les positions des avions //à clarifier SpOOd
		if (*grid)[p.nextX][p.nextY] == 2 || (*grid)[p.nextX][p.nextY] == 3 { // la prochine case est un avion ou reservé
			changement[p.IdentifAv] <- ChangeurP{
				CollisionPossible: true, // la case prochaine est soit reserv soit y a un avion, possible collision
				precX:             p.nowX, //translater le X de 1 ou le Y à voir
				precY:             p.nowY,
				nextX:             p.nowX,
				nextY:             p.nowY,
			}
		} else {
			if (*grid)[p.nextX][p.nextY] != 1 { //si cest pas un aero
				(*grid)[p.nextX][p.nextY] = 3 //on réserve la prochaine case le cas echant
			}
			changement[p.IdentifAv] <- ChangeurP{
				CollisionPossible: false,
				precX:             p.nowX,
				precY:             p.nowY,
				nextX:             p.nextX,
				nextY:             p.nextY,
			}
		}
	}
}

//Mise à jour à chaque mvmt davion de la gui
func MaJIG(grid_view string, grid *[][]int, previous_x int, previous_y int, actual_x int, actual_y int, long int, larg int) string {
	visu_lrg := long + 1 //largeur de la grille visuellement
	preXIG := previous_x
	preYIG := previous_y
	if (*grid)[previous_x][previous_y] != 1 {
		grid_view = grid_view[:visu_lrg*preYIG+preXIG] + "." + grid_view[visu_lrg*preYIG+preXIG+1:] //la place laissé par lavion devient un pt
	}
	if (*grid)[actual_x][actual_y] != 1 {
		grid_view = grid_view[:visu_lrg*actual_y+actual_x] + "W" + grid_view[visu_lrg*actual_y+actual_x+1:] //et la place actuel contient un avion
	}
	return grid_view
}

func BougerAvion(avion Avion, grid *[][]int, IGchan chan string, fini chan bool, annonce chan AnnonceP, changementsP []chan ChangeurP, conn net.Conn, long int, larg int, tmps *string) {
	//la prochaine boucle continue jusqu a ce que lavion atteint larrivée
	for avion.posX != avion.Arrivee.posX || avion.posY != avion.Arrivee.posY {
		calculNextX := avion.posX
		calculNextY := avion.posY
		if avion.posX < avion.Arrivee.posX {
			calculNextX += 1
		} else if avion.posX > avion.Arrivee.posX {
			calculNextX -= 1
		} else {
			calculNextX = avion.Arrivee.posX
		}
		if avion.posY < avion.Arrivee.posY {
			calculNextY += 1
		} else if avion.posY > avion.Arrivee.posY {
			calculNextY -= 1
		} else {
			calculNextY = avion.Arrivee.posY
		}

		annonce <- AnnonceP{ //On annonce au canal de la tour de controle la position de lavion et sa prochaine postion
			IdentifAv: avion.Id, //Changer Id par Matricule
			nowX:      avion.posX,
			nowY:      avion.posY,
			nextX:     calculNextX,
			nextY:     calculNextY,
		}
		avion.Tmps_vol += 1 //le temps de vol est dnc incrémenté
		instruction := <-changementsP[avion.Id] //ToC
		if (*grid)[instruction.nextX][instruction.nextY] == 2 { //la prochaine case contient un avion (collision)
			io.WriteString(conn, "L'avion du vol ELP0"+strconv.Itoa(avion.Id+100)+" a changé d'altitude pour éviter un avion\n") //on ne gere pas la collision, mais deux avions peuvent coexister dans la même case
		}
		if (*grid)[instruction.nextX][instruction.nextY] != 1 && (*grid)[instruction.nextX][instruction.nextY] != 2 { //Ni aeroport ni station
			(*grid)[instruction.nextX][instruction.nextY] = 2 //on deplace notre avion vers la case
		}
		if (*grid)[instruction.nextX][instruction.nextY] == 1 && instruction.nextX == avion.Arrivee.posX && instruction.nextY == avion.Arrivee.posY { //l'avion est arrivée
			*tmps += "L'avion du vol ELP0" + strconv.Itoa(avion.Id+100) + " a mis " + strconv.Itoa(avion.Tmps_vol) + " heures pour arriver à destination.\n" //un string qui sera affiché à la fin de lexecuion du programme
			io.WriteString(conn, "L'avion du vol ELP0"+strconv.Itoa(avion.Id+100)+" va attérir à l'aeroport de la ville "+avion.Arrivee.Name+" au bout de "+strconv.Itoa(avion.Tmps_vol)+" heures de vol.\n")
		}
		if (*grid)[instruction.precX][instruction.precY] != 1 { //Si la position prec n'est pas aeroport
			(*grid)[instruction.precX][instruction.precY] = 0 //On dé-reserve la case reservé par avion
		}
		gridIG := <-IGchan //gridIG recupéré 
		gridIG = MaJIG(gridIG, grid, instruction.precX, instruction.precY, instruction.nextX, instruction.nextY, long, larg) //on effectue la mise a jour de la gui de la grille
		io.WriteString(conn, gridIG+"\n")//on envoie la grille au client
		IGchan <- gridIG //gridIG envoyé
		avion.posX = instruction.nextX
		avion.posY = instruction.nextY
	}
	fini <- true //permet de débloquer un pass dans la 3ème boucle for du maingo
}

// /!\ 2 FUNCTIONS MAIN (1 du serveur, 1 intercation avec le client) et un handler de connection tcp

//function maingo, ou tout se passe selon les données envoyés par le client
func maingo(conn net.Conn, nb_avion int, nb_aero int, larg int, long int) {
	tmps := "" //variable qui sera print à la fin de lexecution total avec le tmps de vol de tous les avions
	grid := GenGridSlice(long, larg) //slice of slice de lespace aerien en 2D
	rand.Seed(time.Now().UnixNano())
	aeroports := GenAeroport(&grid, nb_aero)//genera de nb_aero aeroport aleatoirement dans la grid
	avions := GenAvion(aeroports, conn, nb_avion, nb_aero) //gene davion dans les aeroports
	io.WriteString(conn, "######## Taper 'Entrée' pour lancer la simulation ########\n")
	//on attend maintenant que lutilisateur tape Entree et donc que le client envoie oui au sevreur
	read := bufio.NewReader(conn)
	for {
		result, err := read.ReadString('\n')
		if err != nil {
			fmt.Printf("DEBUG MAIN could not read from client")
			os.Exit(1)
		}
		if result == "oui\n" {
			break
		}
	}
	gridIG := GenIG(aeroports, nb_aero, long, larg) //on genere la gui de notre grid
	IGchan := make(chan string, 100) 
	IGchan <- gridIG //et on lenvoie par le canal pour quelle soit recupere par goroutine bougerAvion plus tard
	fini := make(chan bool, len(avions)) //permet lattente des avions mutuellement à la fin et debloquer la 3eme boucle for
	requetesPositions := make(chan AnnonceP) //canal denvoie des requetes (annonce) de position à la tour de controle
	changementsP := make([]chan ChangeurP, len(avions)) //canal de changements de positions davion à traiter
	for i := range changementsP {
		changementsP[i] = make(chan ChangeurP, 10) //un canal par avion
	}
	go TourDeC(requetesPositions, changementsP, &grid) //go routine ; tour de controle: centralisé pour tous les avions
	for i := 0; i < len(avions); i++ {
		go BougerAvion(avions[i], &grid, IGchan, fini, requetesPositions, changementsP, conn, long, larg, &tmps) //go routine: une par avion,permet de la deplacer et faire MaJ de la grid 
	}
	for i := 0; i < len(avions); i++ {
		<-fini
		if i == len(avions)-1 { // quand tous les avions ont fini
			io.WriteString(conn, tmps)
			io.WriteString(conn, "  \\      ___.   .__               __     /\\   __   \n")
			io.WriteString(conn, "_____    \\_ |__ |__| ____   _____/  |_  _____/  |_ \n")
			io.WriteString(conn, "\\__  \\    | __ \\|  |/ __ \\ /    \\   __\\/  _ \\   __\\\n")
			io.WriteString(conn, " / __ \\_  | \\_\\ \\  \\  ___/|   |  \\  | (  <_> )  |  \n")
			io.WriteString(conn, "(____  /  |___  /__|\\___  >___|  /__|  \\____/|__|  \n")
			io.WriteString(conn, "     \\/       \\/        \\/     \\/                  \n")

		}
	}

}
func getArgs() int {

	if len(os.Args) != 2 {
		fmt.Printf("Usage: go run server.go <portnumber>\n") // on affiche a l'utilisateur les paramètre requis pour la simulation si il ne les a pas tous mis
		os.Exit(1)
	} else {
		fmt.Printf("#DEBUG ARGS Port Number : %s\n", os.Args[1])
		portNumber, err := strconv.Atoi(os.Args[1]) // on récupère le numéro de port pour la connexion
		if err != nil {
			fmt.Printf("Usage: go run server.go <portnumber>\n")
			os.Exit(1)
		} else {
			return portNumber // on renvoie le numéro de port pour la connexion
		}

	}
	return -1
}

func main() {
	port := getArgs() // on récupère le numéro de port
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

		go handleConnection(conn, connum) // on lance une goroutine pour une connexion, possibilite de plusieurs clients en même temps
		connum += 1

	}
}

func handleConnection(connection net.Conn, connum int) {
	intu := []int{} // on crée la intu qui va contenir les paramètre envoyer par le client
	var u = 0       // on crée le compteur qui va nous permettre de sortir de la boucle infini d'écoute de paramètre
	defer connection.Close()
	connReader := bufio.NewReader(connection)
	for { // on attend que le client nous envoie tous les paramètres donné par l'utilisateur
		inputLine, err := connReader.ReadString('\n')
		resultString := strings.TrimSuffix(inputLine, "\n")
		intinput, err := strconv.Atoi(resultString)
		intu = append(intu, intinput) // on ajoute le paramètre recu à la liste de paramètre
		u = u + 1

		if u == 4 { // on sort de la boucle infini lorsque tous les paramètres ont été reçu
			break
		}

		if err != nil {
			os.Exit(1)
		}

		if err != nil {
			fmt.Printf("#DEBUG %d RCV ERROR no panic, just a client\n", connum)
			fmt.Printf("Error :|%s|\n", err.Error())
		}
	}
	nb_avion := intu[0] // on crée nos paramètre pour le serveur en fonction des paramètres reçu
	nb_aero := intu[1]
	larg := intu[2]
	long := intu[3]
	maingo(connection, nb_avion, nb_aero, larg, long) // on lance le programme avec les paramètre reçu par le client
}
