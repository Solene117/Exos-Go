package main

import "fmt"

func main() {
	const Nom = "Solène"

	var poids float64 = 70.5
	var taille float64 = 1.75

	const (
		IMCMaigreur = 18.5
		IMCNormal   = 25.0
		IMCSurpoids = 30.0
	)

	imc := poids / (taille * taille)

	fmt.Printf("Bonjour %s !\n", Nom)
	fmt.Printf("Votre IMC est de : %.2f\n", imc)

	fmt.Print("Catégorie : ")
	if imc < IMCMaigreur {
		fmt.Println("Maigreur")
	} else if imc < IMCNormal {
		fmt.Println("Normal")
	} else if imc < IMCSurpoids {
		fmt.Println("Surpoids")
	} else {
		fmt.Println("Obésité")
	}
}
