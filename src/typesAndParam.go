package src
// MODIFICATION AUTORISÉE POUR LA PARTIE SUIVANTE :)
const nb_aero = 2
const nb_avion = 3
const long = 10
const larg = 10


// NE PAS MODIFIER LES PARTIES SUIVANTES /!\ //ToC
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
	//Ajouter Z 0 ou 1 atteri ou en vol
}

type ChangeurP struct {
	CollisionPossible bool
	Previous_X int
	Previous_Y int
	Next_X int
	Next_Y int
}