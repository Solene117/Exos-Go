package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Produit struct {
	ID        int
	Nom       string
	Marque    string
	Prix      float64
	Stock     int
	Categorie string
	Actif     bool
}

type Catalogue struct {
	Produits []Produit
	NextID   int
}

func (c *Catalogue) AjouterProduit(p Produit) error {
	p.ID = c.NextID
	c.NextID++
	c.Produits = append(c.Produits, p)
	return nil
}

func (c Catalogue) TrouverParID(id int) (Produit, error) {
	for _, prod := range c.Produits {
		if prod.ID == id {
			return prod, nil
		}
	}
	return Produit{}, errors.New("erreur : produit introuvable")
}

func (c Catalogue) TrouverParCategorie(cat string) []Produit {
	var resultat []Produit
	for _, prod := range c.Produits {
		if strings.EqualFold(prod.Categorie, cat) {
			resultat = append(resultat, prod)
		}
	}
	return resultat
}

func (c *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	nbModifies := 0
	for i := range c.Produits {
		if strings.EqualFold(c.Produits[i].Categorie, categorie) {
			c.Produits[i].Prix = c.Produits[i].Prix * (1 - pct/100)
			nbModifies++
		}
	}
	return nbModifies
}

func (c *Catalogue) Vendre(id int, qte int) error {
	for i := range c.Produits {
		if c.Produits[i].ID == id {
			if c.Produits[i].Stock < qte {
				return fmt.Errorf("erreur : stock insuffisant (%d disponibles)", c.Produits[i].Stock)
			}
			c.Produits[i].Stock -= qte
			return nil
		}
	}
	return errors.New("erreur : produit introuvable")
}

func (c Catalogue) Rapport() string {
	nbProduits := len(c.Produits)
	valeurTotale := 0.0
	for _, prod := range c.Produits {
		valeurTotale += prod.Prix * float64(prod.Stock)
	}
	return fmt.Sprintf("Nombre de produits différents : %d\nValeur totale du stock : %.2f €", nbProduits, valeurTotale)
}

func main() {
	catalogue := Catalogue{NextID: 1}
	catalogue.AjouterProduit(Produit{Nom: "iPhone 15", Marque: "Apple", Prix: 950.00, Stock: 15, Categorie: "Smartphone", Actif: true})
	catalogue.AjouterProduit(Produit{Nom: "MacBook Air M3", Marque: "Apple", Prix: 1299.00, Stock: 8, Categorie: "Ordinateur", Actif: true})
	catalogue.AjouterProduit(Produit{Nom: "Galaxy S24", Marque: "Samsung", Prix: 899.00, Stock: 12, Categorie: "Smartphone", Actif: true})
	catalogue.AjouterProduit(Produit{Nom: "PlayStation 5", Marque: "Sony", Prix: 549.99, Stock: 20, Categorie: "Console", Actif: true})
	catalogue.AjouterProduit(Produit{Nom: "WH-1000XM5", Marque: "Sony", Prix: 329.00, Stock: 25, Categorie: "Audio", Actif: true})

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== BIENVENUE CHEZ TECHSHOP ===")

	for {
		fmt.Println("\nMenu Principal :")
		fmt.Println("[1] Ajouter un produit   [2] Chercher un produit   [3] Soldes / Réduction")
		fmt.Println("[4] Vendre un produit    [5] Rapport du stock      [0] Quitter")
		fmt.Print("Votre choix > ")

		if !scanner.Scan() {
			break
		}
		choix := strings.TrimSpace(scanner.Text())

		switch choix {
		case "1":
			fmt.Println("\n--- AJOUTER UN PRODUIT ---")
			var p Produit
			p.Actif = true

			fmt.Print("Nom : ")
			scanner.Scan()
			p.Nom = strings.TrimSpace(scanner.Text())

			fmt.Print("Marque : ")
			scanner.Scan()
			p.Marque = strings.TrimSpace(scanner.Text())

			fmt.Print("Prix (€) : ")
			scanner.Scan()
			fmt.Sscanf(scanner.Text(), "%f", &p.Prix)

			fmt.Print("Stock initial : ")
			scanner.Scan()
			fmt.Sscanf(scanner.Text(), "%d", &p.Stock)

			fmt.Print("Catégorie : ")
			scanner.Scan()
			p.Categorie = strings.TrimSpace(scanner.Text())

			if err := catalogue.AjouterProduit(p); err != nil {
				fmt.Println("Err", err)
			} else {
				fmt.Printf("Produit ajouté avec succès ! (ID auto-assigné : %d)\n", p.ID)
			}

		case "2":
			fmt.Println("\n--- CHERCHER ---")
			fmt.Println("[A] Par ID  [B] Par Catégorie")
			fmt.Print("Choix > ")
			scanner.Scan()
			subChoix := strings.ToUpper(strings.TrimSpace(scanner.Text()))

			switch subChoix {
			case "A":
				var id int
				fmt.Print("Entrez l'ID : ")
				scanner.Scan()
				fmt.Sscanf(scanner.Text(), "%d", &id)

				p, err := catalogue.TrouverParID(id)
				if err != nil {
					fmt.Println("Err", err)
				} else {
					fmt.Printf("Trouvé : %s (%s) - %.2f € | Stock : %d | Catégorie : %s\n", p.Nom, p.Marque, p.Prix, p.Stock, p.Categorie)
				}
			case "B":
				fmt.Print("Entrez la catégorie : ")
				scanner.Scan()
				cat := strings.TrimSpace(scanner.Text())

				prods := catalogue.TrouverParCategorie(cat)
				if len(prods) == 0 {
					fmt.Println("Aucun produit trouvé dans cette catégorie.")
				} else {
					fmt.Printf("Produits trouvés (%d) :\n", len(prods))
					for _, p := range prods {
						fmt.Printf(" - [ID %d] %s (%s) : %.2f € (Stock: %d)\n", p.ID, p.Nom, p.Marque, p.Prix, p.Stock)
					}
				}
			}

		case "3":
			fmt.Println("\n--- APPLIQUER DES SOLDES ---")
			fmt.Print("Catégorie ciblée : ")
			scanner.Scan()
			cat := strings.TrimSpace(scanner.Text())

			var pct float64
			fmt.Print("Pourcentage de réduction (ex: 15) : ")
			scanner.Scan()
			fmt.Sscanf(scanner.Text(), "%f", &pct)

			modifies := catalogue.AppliquerReduction(cat, pct)
			fmt.Printf("Réduction appliquée ! Nombre de produits mis à jour : %d\n", modifies)

		case "4":
			fmt.Println("\n--- VENDRE UN PRODUIT ---")
			var id, qte int
			fmt.Print("ID du produit à vendre : ")
			scanner.Scan()
			fmt.Sscanf(scanner.Text(), "%d", &id)

			fmt.Print("Quantité : ")
			scanner.Scan()
			fmt.Sscanf(scanner.Text(), "%d", &qte)

			if err := catalogue.Vendre(id, qte); err != nil {
				fmt.Println("Err", err)
			} else {
				fmt.Println("🛒 Vente enregistrée avec succès !")
			}

		case "5":
			fmt.Println("\n--- RAPPORT DU CATALOGUE ---")
			fmt.Println(catalogue.Rapport())

		case "0":
			fmt.Println("\nFermeture de l'application TechShop. Au revoir !")
			return

		default:
			fmt.Println("Choix invalide, veuillez sélectionner une option du menu.")
		}
		fmt.Println("\n-----------------------------------------------------")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur de lecture : %v\n", err)
	}
}
