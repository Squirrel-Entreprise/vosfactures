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

	// Créer une catégorie
	c := vosfactures.Category{
		Name: "Fruits",
	}
	vfCat, err := vf.CreateCategory(&c)
	if err != nil {
		log.Println(err)
	}

	log.Println(vfCat.ID, vfCat.Name)

	// Voir une catégorie
	vfCat, err = vf.GetCategory(vfCat.ID)
	if err != nil {
		log.Println(err)
	}

	log.Println(vfCat.ID, vfCat.Name)

	// Lister les catégories
	vfCats, err := vf.ListCategories(1)
	if err != nil {
		log.Println(err)
	}

	for _, cat := range vfCats {
		log.Println(cat.ID, cat.Name)
	}

	// Mettre à jour une catégorie
	nc := vosfactures.Category{
		ID:   vfCat.ID,
		Name: "Fruits",
	}
	vfCat, err = vf.UpdateCategory(&nc)
	if err != nil {
		log.Println(err)
	}

	log.Println(vfCat.ID, vfCat.Name)

	// Supprimer une catégorie
	if err := vf.DeleteCategory(vfCat.ID); err != nil {
		log.Println(err)
	}

}
