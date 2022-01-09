package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

func getArgs() []int {
	var s []int
	if len(os.Args) != 6 {
		fmt.Printf("Usage: go run client.go <portnumber>  <planenumber>  <airportnumber>  <mapwide>  <mapheight>\n") // on affiche a l'utilisateur les paramètre requis pour la simulation si il ne                                                                                                               les a pas tous mis
		os.Exit(1)
	} else {
		fmt.Printf("#DEBUG ARGS Port Number : %s\n", os.Args[1])
		portNumber, err := strconv.Atoi(os.Args[1]) // on récupère le port qui a été donné en premier paramètre de la fonction
		s = append(s, portNumber)                   // on ajoute s le numéro de port
		if err != nil {
			fmt.Printf("Usage: go run client.go <portnumber> <planenumber>  <airportnumber>  <mapwide>  <mapheight>\n") // on affiche a l'utilisateur les paramètre requis pour la simulation
			os.Exit(1)
		}
		for i := 2; i < 6; i++ {
			para, err := strconv.Atoi(os.Args[i]) // on récupère tout les paramètres donné en lancant la fonction
			if err != nil {
				os.Exit(1)
			}
			s = append(s, para) // on ajoute le paramètre à s
		}
	}
	return s
}

func main() {
	para := getArgs() // on récupère les paramètre
	fmt.Printf("#DEBUG DIALING TCP Server on port %d\n", para[0])
	portString := fmt.Sprintf("127.0.0.1:%s", strconv.Itoa(para[0]))
	fmt.Printf("#DEBUG MAIN PORT STRING |%s|\n", portString)

	conn, err := net.Dial("tcp", portString)
	if err != nil {
		fmt.Printf("#DEBUG MAIN could not connect\n")
		os.Exit(1)
	} else {
		defer conn.Close()
		fmt.Printf("#DEBUG MAIN connected\n")
		for i := 1; i < 5; i++ {
			io.WriteString(conn, strconv.Itoa(para[i])+"\n") // on envoie les 4 paramètre donné en lancant la fonction au serveur un par un
		}
		reader := bufio.NewReader(conn)
		for { // on lance une boucle infini pour écouter le serveur
			resultString, err := reader.ReadString('\n') // on récupère le message du serveur dans la variable resultString
			if err != nil {
				fmt.Printf("\nDéconnexion du serveur.")
				os.Exit(1)
			}
			fmt.Print(resultString)                                                             // on affiche le message recu par le client
			if resultString == "######## Taper 'Entrée' pour lancer la simulation ########\n" { // on crée une condition pour le lancement de la simulation lorsque le serveur demande si il                                                                                            peut lancer la conditon
				bufio.NewScanner(os.Stdin).Scan() // on attend que l'utilisateur appui sur Entrée
				io.WriteString(conn, "oui\n")     // on envoie un message spécifique au serveur pour lancer la simulation
			}
		}
	}
}
