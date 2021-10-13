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

	// Créer un client
	c := vosfactures.Client{
		Name: "Jean Bon",
	}
	vfCli, err := vf.CreateClient(&c)
	if err != nil {
		log.Println(err)
	}

	log.Println(vfCli.ID, vfCli.Name)

	// Voir un client
	vfCli, err = vf.GetClient(vfCli.ID)
	if err != nil {
		log.Println(err)
	}

	log.Println(vfCli.ID, vfCli.Name)

	// Lister les clients
	vfClis, err := vf.ListClients(1)
	if err != nil {
		log.Println(err)
	}

	for _, cli := range vfClis {
		log.Println(cli.ID, cli.Name)
	}

	// Mettre à jour un client
	nc := vosfactures.Client{
		ID:   vfCli.ID,
		Name: "Jean Bon",
	}
	vfCli, err = vf.UpdateClient(&nc)
	if err != nil {
		log.Println(err)
	}

	log.Println(vfCli.ID, vfCli.Name)

	// Supprimer un client
	if err := vf.DeleteClient(vfCli.ID); err != nil {
		log.Println(err)
	}

}
