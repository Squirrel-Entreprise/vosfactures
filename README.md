# VosFactures

Implémentation sommaire de l'API VosFactures en Golang.
Ce projet est en cours de réalisation, il est suffisamment complet pour nos besoins actuellement, pour toutes corrections ou améliorations je serai ravis de merge vos pull requests

### Doc officiel
[https://github.com/vosfactures/api](https://github.com/vosfactures/api)

### Point de départ
```json
t := "xxxyyyzz"
a := "votrecompte"
vf := vosfactures.New(t, a)
```

### Catégories VosFactures
`cmd/categories/main.go`

#### Créer une catégorie

```go
c := vosfactures.Category{
    Name: "Fruits",
}
vfCat, err := vf.CreateCategory(&c)
if err != nil {
    log.Println(err)
}
```

#### Voir une catégorie

```go
vfCat, err := vf.GetCategory(1234)
if err != nil {
    log.Println(err)
}
```

#### Lister les catégories

```go
vfCats, err := vf.ListCategories(1)
if err != nil {
    log.Println(err)
}
```

#### Mettre à jour une catégorie

```go
c := vosfactures.Category{
    ID:   1234,
    Name: "Fruits",
}
vfCat, err := vf.UpdateCategory(&c)
if err != nil {
    log.Println(err)
}
```

#### Supprimer une catégorie

```go
if err := vf.DeleteCategory(1234); err != nil {
    log.Println(err)
}
```

### Produits VosFactures
`cmd/products/main.go`

#### Créer un produit

```go
p := vosfactures.Product{
    Name: "Pomme",
}
vfProd, err := vf.CreateProduct(&p)
if err != nil {
    log.Println(err)
}
```

#### Voir un produit

```go
vfProd, err := vf.GetProduct(1234)
if err != nil {
    log.Println(err)
}
```

#### Lister les produits

```go
vfProds, err := vf.ListProducts(1)
if err != nil {
    log.Println(err)
}
```

#### Mettre à jour un produit

```go
p := vosfactures.Product{
    ID:   1234,
    Name: "Pomme",
}
vfProd, err := vf.UpdateProduct(&p)
if err != nil {
    log.Println(err)
}
```

#### Supprimer un produit

```go
if err := vf.DeleteProduct(1234); err != nil {
    log.Println(err)
}
```

### Clients VosFactures
`cmd/clients/main.go`

#### Créer un client

```go
p := vosfactures.Client{
    Name: "Jean Bon",
}
vfClient, err := vf.CreateClient(&p)
if err != nil {
    log.Println(err)
}
```

#### Voir un client

```go
vfClient, err := vf.GetClient(1234)
if err != nil {
    log.Println(err)
}
```

#### Lister les clients

```go
vfClients, err := vf.ListClients(1)
if err != nil {
    log.Println(err)
}
```

#### Mettre à jour un client

```go
p := vosfactures.Client{
    ID:   1234,
    Name: "Jean Bon",
}
vfClient, err := vf.UpdateClient(&p)
if err != nil {
    log.Println(err)
}
```

#### Supprimer un client

```go
if err := vf.DeleteClient(1234); err != nil {
    log.Println(err)
}
```


### Département VosFactures
`cmd/clients/main.go`

#### Créer un département

```go
p := vosfactures.Department{
    Name: "Légumes companie",
}
vfDep, err := vf.CreateDepartment(&p)
if err != nil {
    log.Println(err)
}
```

#### Voir un département

```go
vfDep, err := vf.GetDepartment(1234)
if err != nil {
    log.Println(err)
}
```

#### Lister les départements

```go
vfDeps, err := vf.ListDepartments(1)
if err != nil {
    log.Println(err)
}
```

#### Mettre à jour un département

```go
p := vosfactures.Department{
    ID:   1234,
    Name: "Légumes companie",
}
vfDep, err := vf.UpdateDepartment(&p)
if err != nil {
    log.Println(err)
}
```

#### Supprimer un département

```go
if err := vf.DeleteDepartment(1234); err != nil {
    log.Println(err)
}
```

### Factures VosFactures
`cmd/ivoices/main.go`

#### Créer une facture

```go
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
```

#### Voir une facture

```go
vfInv, err = vf.GetInvoice(vfInv.ID)
if err != nil {
    log.Println(err)
}
```

#### Lister les factures

```go
vfInvs, err := vf.ListInvoices(1)
if err != nil {
    log.Println(err)
}
```

#### Mettre à jour une facture

```go
p := vosfactures.Document{
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
vfInv, err = vf.UpdateInvoice(&p)
if err != nil {
    log.Println(err)
}

```

#### Supprimer une facture

```go
if err := vf.DeleteInvoice(vfInv.ID); err != nil {
    log.Println(err)
}
```

#### Envoyer une facture par email

```go
if err := vf.SendInvoiceByMail(vfInv.ID); err != nil {
    log.Println(err)
}
```
