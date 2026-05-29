package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func operer(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("erreur : division par zéro impossible")
		}
		return a / b, nil
	default:
		return 0, errors.New("erreur : opérateur inconnu")
	}
}

func creerOperation(op string) func(float64, float64) float64 {
	return func(a, b float64) float64 {
		res, err := operer(a, b, op)
		if err != nil {
			fmt.Println(err)
			return 0
		}
		return res
	}
}

func main() {
	fmt.Println("--- Calculatrice CLI (Écrivez 'quit' pour sortir) ---")
	fmt.Println("Format attendu : [nombre1] [nombre2] [opérateur] (ex: 10 5 +)")
	fmt.Println("-----------------------------------------------------")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		ligne := strings.TrimSpace(scanner.Text())

		if ligne == "quit" {
			break
		}

		var a, b float64
		var op string

		_, err := fmt.Sscanf(ligne, "%f %f %s", &a, &b, &op)

		if err != nil {
			if op == "quit" {
				break
			}
			fmt.Println("Erreur de saisie. Réessayez au format : 10 5 +")
			fmt.Println("-----------------------------------------------------")
			continue
		}

		resultat, errOp := operer(a, b, op)
		if errOp != nil {
			fmt.Println(errOp)
		} else {
			fmt.Printf("Résultat : %.2f\n", resultat)
		}
		fmt.Println("-----------------------------------------------------")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Erreur de lecture sur l'entrée standard:", err)
	}

	fmt.Println("Au revoir !")
}
