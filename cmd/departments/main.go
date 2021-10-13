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

	// Créer un département
	c := vosfactures.Department{
		Name: "Légumes companie",
	}
	vfDep, err := vf.CreateDepartment(&c)
	if err != nil {
		log.Println(err)
	}

	log.Println(vfDep.ID, vfDep.Name)

	// Voir un département
	vfDep, err = vf.GetDepartment(vfDep.ID)
	if err != nil {
		log.Println(err)
	}

	log.Println(vfDep.ID, vfDep.Name)

	// Lister les départements
	vfDeps, err := vf.ListDepartments(1)
	if err != nil {
		log.Println(err)
	}

	for _, dep := range vfDeps {
		log.Println(dep.ID, dep.Name)
	}

	// Mettre à jour un département
	nc := vosfactures.Department{
		ID:   vfDep.ID,
		Name: "Légumes companie",
	}
	vfDep, err = vf.UpdateDepartment(&nc)
	if err != nil {
		log.Println(err)
	}

	log.Println(vfDep.ID, vfDep.Name)

	// Supprimer un département
	if err := vf.DeleteDepartment(vfDep.ID); err != nil {
		log.Println(err)
	}

}
