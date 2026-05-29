package main

import "fmt"

const (
	Debug = iota
	Info
	Warn
	Error
)

func main() {
	fmt.Println("--- Exercice : Exploration des Tableaux, Slices et Syntaxes Go ---")

	var categories [4]string
	categories[Debug] = "DEBUG"
	categories[Info] = "INFO"
	categories[Warn] = "WARN"
	categories[Error] = "ERROR"

	alertes := make([]int, 0, 2)

	fmt.Printf("Initialisation de la slice -> Longueur : %d, Capacité : %d\n", len(alertes), cap(alertes))

	alertes = append(alertes, Info)
	alertes = append(alertes, Warn)

	alertes = append(alertes, Error)

	fmt.Printf("Après ajouts -> Longueur : %d, Capacité : %d\n", len(alertes), cap(alertes))
	fmt.Println("-----------------------------------------------------")

	fmt.Println("Analyse des alertes en cours :")
	for index, niveau := range alertes {
		fmt.Printf("Alerte n°%d [Niveau %s] : ", index+1, categories[niveau])

		switch niveau {
		case Error:
			fmt.Print("[CRITIQUE] Déclenchement de la sirène ! -> ")
			fallthrough
		case Warn:
			fmt.Print("Envoi d'un SMS de notification. -> ")
			fallthrough
		case Info, Debug:
			fmt.Println("Enregistrement dans le fichier de log.")
		}
	}
}
