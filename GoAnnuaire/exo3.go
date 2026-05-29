package main

import (
	"fmt"
)

type Personne struct {
	Prenom string
	Nom    string
	Age    int
	Email  string
}

func (p Personne) NomComplet() string {
	return fmt.Sprintf("%s %s", p.Prenom, p.Nom)
}

func (p Personne) Presentation() string {
	return fmt.Sprintf("%s, %d ans (Email: %s)", p.NomComplet(), p.Age, p.Email)
}

type Adresse struct {
	Rue        string
	Ville      string
	CodePostal string
}

func (a Adresse) Format() string {
	return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville)
}

type Employe struct {
	Personne
	Adresse
	Poste   string
	Salaire float64
}

func (e Employe) FicheEmploye() string {
	return fmt.Sprintf(
		"=== FICHE EMPLOYÉ ===\nStatut : %s\nPoste : %s\nSalaire : %.2f €\nAdresse : %s\n=====================",
		e.Presentation(),
		e.Poste,
		e.Salaire,
		e.Format(),
	)
}

func (e *Employe) AugmenterSalaire(pct float64) {
	e.Salaire = e.Salaire * (1 + pct/100)
}

type Etudiant struct {
	Personne
	Promo   string
	Moyenne float64
}

func (e Etudiant) MentionObtenue() string {
	switch {
	case e.Moyenne >= 16:
		return "Très Bien (TB)"
	case e.Moyenne >= 14:
		return "Bien (B)"
	case e.Moyenne >= 12:
		return "Assez Bien (AB)"
	case e.Moyenne >= 10:
		return "Passable (P)"
	default:
		return "Insuffisant"
	}
}

func main() {
	fmt.Println("--- Exercice 3 : Système de contacts ---")
	fmt.Println()

	emp1 := Employe{
		Personne: Personne{Prenom: "Jean", Nom: "Dupont", Age: 34, Email: "j.dupont@company.com"},
		Adresse:  Adresse{Rue: "12 Rue de la Paix", Ville: "Paris", CodePostal: "75002"},
		Poste:    "Développeur Go",
		Salaire:  3500.0,
	}

	emp2 := Employe{
		Personne: Personne{Prenom: "Sophie", Nom: "Martin", Age: 29, Email: "s.martin@company.com"},
		Adresse:  Adresse{Rue: "45 Avenue de la République", Ville: "Lyon", CodePostal: "69003"},
		Poste:    "Product Owner",
		Salaire:  4100.0,
	}

	etu1 := Etudiant{
		Personne: Personne{Prenom: "Lucas", Nom: "Bernard", Age: 21, Email: "lucas.b@esgi.fr"},
		Promo:    "M2 Architecture Logicielle",
		Moyenne:  16.5,
	}

	etu2 := Etudiant{
		Personne: Personne{Prenom: "Emma", Nom: "Petit", Age: 23, Email: "emma.p@esgi.fr"},
		Promo:    "M2 Cloud computing",
		Moyenne:  11.2,
	}

	fmt.Printf("Salaire initial de %s : %.2f €\n", emp1.NomComplet(), emp1.Salaire)
	emp1.AugmenterSalaire(10)
	fmt.Printf("Nouveau salaire après augmentation : %.2f €\n\n", emp1.Salaire)

	fmt.Println(emp1.FicheEmploye())
	fmt.Println()
	fmt.Println(emp2.FicheEmploye())
	fmt.Println()

	fmt.Println("=== FICHES ÉTUDIANTS ===")
	fmt.Printf("Étudiant : %s | Promo : %s\nMoyenne : %.2f/20 -> Mention : %s\n",
		etu1.NomComplet(), etu1.Promo, etu1.Moyenne, etu1.MentionObtenue())
	fmt.Println("---------------------------------")
	fmt.Printf("Étudiant : %s | Promo : %s\nMoyenne : %.2f/20 -> Mention : %s\n",
		etu2.NomComplet(), etu2.Promo, etu2.Moyenne, etu2.MentionObtenue())
	fmt.Println("========================")
}
