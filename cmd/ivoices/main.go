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

	// Créer une facture
	c := vosfactures.Document{
		DepartmentID: 1234,
		ClientID:     1234,
		Positions: []*vosfactures.Position{
			&vosfactures.Position{
				ProductID: 1234,
				Quantity:  2,
			},
		},
	}
	vfInv, err := vf.CreateInvoice(&c)
	if err != nil {
		log.Println(err)
	}
	log.Println(vfInv.ID, vfInv.Number)

	// Voir une facture
	vfInv, err = vf.GetInvoice(vfInv.ID)
	if err != nil {
		log.Println(err)
	}

	log.Println(vfInv.ID, vfInv.Number)

	// Lister les factures
	vfInvs, err := vf.ListInvoices(1)
	if err != nil {
		log.Println(err)
	}

	for _, inv := range vfInvs {
		log.Println(inv.ID, inv.Number)
	}

	// Mettre à jour une facture
	ni := vosfactures.Document{
		ID:           vfInv.ID,
		DepartmentID: 1234,
		ClientID:     1234,
		Positions: []*vosfactures.Position{
			&vosfactures.Position{
				ProductID: 1234,
				Quantity:  2,
			},
		},
	}
	vfInv, err = vf.UpdateInvoice(&ni)
	if err != nil {
		log.Println(err)
	}

	log.Println(vfInv.ID, vfInv.Number)

	// Supprimer une facture
	if err := vf.DeleteInvoice(vfInv.ID); err != nil {
		log.Println(err)
	}

}
