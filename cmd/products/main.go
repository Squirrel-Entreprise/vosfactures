package main

import (
	"log"
	"os"

	"github.com/Squirrel-Entreprise/vosfactures"
)

func main() {
	t := os.Getenv("VF_TOKEN")
	a := os.Getenv("VF_ACCOUNT")
	vf := vosfactures.New(t, a)

	// Créer un produit
	c := vosfactures.Product{
		Name: "Pomme",
	}
	vfProd, err := vf.CreateProduct(&c)
	if err != nil {
		log.Println(err)
	}

	log.Println(vfProd.ID, vfProd.Name)

	// Voir un produit
	vfProd, err = vf.GetProduct(vfProd.ID)
	if err != nil {
		log.Println(err)
	}

	log.Println(vfProd.ID, vfProd.Name)

	// Lister les produits
	vfProds, err := vf.ListProducts(1)
	if err != nil {
		log.Println(err)
	}

	for _, pro := range vfProds {
		log.Println(pro.ID, pro.Name)
	}

	// Mettre à jour un produit
	nc := vosfactures.Product{
		ID:   vfProd.ID,
		Name: "Pomme",
	}
	vfProd, err = vf.UpdateProduct(&nc)
	if err != nil {
		log.Println(err)
	}

	log.Println(vfProd.ID, vfProd.Name)

	// Supprimer un produit
	if err := vf.DeleteProduct(vfProd.ID); err != nil {
		log.Println(err)
	}

}
