package vosfactures

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	PaymentTypeTransfer PaymentType = "transfer" // - virement bancaire
	PaymentTypeCard     PaymentType = "card"     // - carte bancaire
	PaymentTypeCash     PaymentType = "cash"     // -  espèce
	PaymentTypeCheque   PaymentType = "cheque"   // - chèque
	PaymentTypePaypal   PaymentType = "paypal"   // - PayPal
	PaymentTypeLcr      PaymentType = "lcr"      // - LCR Lettre de Change Relevé
	PaymentTypeOff      PaymentType = "off"      // - aucun (ne pas afficher)
)

const (
	IncomeSale     Income = true  // - vente
	IncomePurchase Income = false // - achat
)

const (
	LangEN Lang = "en" // - Anglais
	LangDE Lang = "de" // - Allemand
	LangFR Lang = "fr" // - Français
	LangHE Lang = "he" // - Grec
	LangES Lang = "es" // - Espagnol
	LangIT Lang = "it" // - Italien
	LangNL Lang = "nl" // - Hollandais
	LangCZ Lang = "cz" // - Tchèque
	LangHR Lang = "hr" // - Croate
	LangPL Lang = "pl" // - Polonais
	LangHU Lang = "hu" // - Hongrois
	LangSK Lang = "sk" // - Slovaque
	LangSL Lang = "sl" // - Slovène
	LangET Lang = "et" // - Estonien
	LangRU Lang = "ru" // - Russe
	LangCN Lang = "cn" // - Chinois
	LangAR Lang = "ar" // - Arabe
	LangTR Lang = "tr" // - Turc
	LangFA Lang = "fa" // - Persan
)

const (
	KindVat          Kind = "vat"           // - facture
	KindProforma     Kind = "proforma"      // - facture Proforma
	KindAdvance      Kind = "advance"       // - facture d'acompte
	KindFinal        Kind = "final"         // - facture de solde
	KindCorrection   Kind = "correction"    // - facture d'avoir
	KindEstimate     Kind = "estimate"      // - devis
	KindClientOrder  Kind = "client_order"  // - bon de commande
	KindReceipt      Kind = "receipt"       // - reçu
	KindKp           Kind = "kp"            // - bon d'entrée de caisse
	KindKw           Kind = "kw"            // - bon de sortie de caisse
	KindInvoiceOther Kind = "invoice_other" // - Autre
)

const (
	DiscountKindPercentUnit  DiscountKind = "percent_unit"  // - % calculé sur le prix unitaire
	DiscountKindPercentTotal DiscountKind = "percent_total" // - % calculé sur le montant total
	DiscountKindAmount       DiscountKind = "amount"        // - montant
)

const (
	StatusIssued   Status = "issued"   // - Créé
	StatusSent     Status = "sent"     // - Envoyé
	StatusPaid     Status = "paid"     // - Payé
	StatusPartial  Status = "partial"  // - Payé en partie
	StatusRejected Status = "rejected" // - Refusé
	StatusAccepted Status = "accepted" // - Accepté
)

type (
	payloadInvoice struct {
		APIToken string    `json:"api_token,omitempty"`
		Invoice  *Document `json:"invoice,omitempty"`
	}

	PaymentType string

	Income bool

	Lang string

	Kind string

	DiscountKind string

	Status string

	Document struct {
		ID                      int64                `json:"id,omitempty"`
		Number                  string               `json:"number,omitempty"`              // numéro du document (généré automatiquement si non indiqué)
		Kind                    Kind                 `json:"kind,omitempty"`                // type du document : "vat" pour facture,
		Correction              string               `json:"correction,omitempty"`          // pour avoir,
		Receipt                 string               `json:"receipt,omitempty"`             // pour reçu,
		Advance                 string               `json:"advance,omitempty"`             // pour facture d'acompte,
		Final                   string               `json:"final,omitempty"`               // pour facture de solde,
		InvoiceOther            string               `json:"invoice_other,omitempty"`       // pour autre type de document comptable,
		Estimate                string               `json:"estimate,omitempty"`            // pour devis,
		Proforma                string               `json:"proforma,omitempty"`            // pour facture proforma,
		ClientOrder             string               `json:"client_order,omitempty"`        // pour bon de commande client,
		MaintenanceRequest      string               `json:"maintenance_request,omitempty"` // pour bon d'intervention,
		PaymentReceipt          string               `json:"payment_receipt,omitempty"`     // pour reçu de paiment,
		Kw                      string               `json:"kw,omitempty"`                  // pour versements en espèces,
		Kp                      string               `json:"kp,omitempty"`                  // pour reçus en espèces.
		Income                  Income               `json:"income,omitempty"`              // revenu (1) ou dépense (0)
		UserID                  int64                `json:"user_id,omitempty"`             // numéro ID de l'utilisateur ayant créé le document (en cas de compte Multi-utilisateurs : https: //aide.vosfactures.fr/2703898-Multi-utilisateurs-cr-ation-fonctions-et-restrictions)
		IssueDate               string               `json:"issue_date,omitempty"`          // date de création
		Place                   string               `json:"place,omitempty"`               // lieu de création
		SellDate                string               `json:"sell_date,omitempty"`           // date additionnelle (ex: date de vente) : date complète ou juste mois et année:YYYY-MM. Pour ne pas faire apparaître cette date, indiquez "off" (ou décochez l'option "Afficher la Date additionnelle" depuis vos paramètres du compte).
		Test                    bool                 `json:"test,omitempty"`                // ou "false" document test ou non (en savoir plus ici: http: //aide.vosfactures.fr/15399051-Cr-er-un-Document-ou-Paiement-Test)
		CategoryID              int64                `json:"category_id,omitempty"`         // ID ou Nom de la catégorie : le système va d'abord regarder si la valeur renseignée correspond à un n° ID d'une catégorie existante, et ensuite à un Nom d'une catégorie existante. Si aucune valeur ne correspond, le système va créer une  nouvelle catégorie.
		DepartmentID            int64                `json:"department_id,omitempty"`       // ID du département vendeur (depuis Paramètres > Compagnies/Départments, cliquer sur le nom de la compagnie/département pour visualiser l'ID dans l'url affiché). Le système affichera alors automatiquement les coordonnées du département vendeur (nom, adresse...) sur le document (les autres champs "seller_" ne sont plus nécessaires).
		SellerName              string               `json:"seller_name,omitempty"`         // Nom du département vendeur. Si ce champ n'est pas renseigné, le département principal est sélectionné par défaut. Préférez plutôt "department_id". Si vous utilisez toutefois "seller_name", le système tentera d'identifier le département portant ce nom, sinon il créera un nouveau département.
		SellerTaxNo             string               `json:"seller_tax_no,omitempty"`       // numéro d'identification fiscale du vendeur (ex: n° TVA)
		SellerTaxNoKind         string               `json:"seller_tax_no_kind,omitempty"`  // initulé du numéro d'identification du vendeur : si non renseigné, il s'agit de "Numéro TVA", sinon il faut spécifier l'intitulé préalablement listé dans vos paramètres du compte, comme par exemple "SIREN" ou "CIF" (en savoir plus ici: http: //aide.vosfactures.fr/1802938-Num-ro-d-identification-fiscale-de-votre-entreprise-TVA-SIREN-IDE-CIF-)
		SellerBankAccount       string               `json:"seller_bank_account,omitempty"` // coordonnées bancaires du vendeur
		SellerBank              string               `json:"seller_bank,omitempty"`         // domiciliation bancaire
		SellerBankSwift         string               `json:"seller_bank_swift,omitempty"`   // code bancaire BIC. Attention, en json vous devez envoyer ce paramètre comme ceci: "additional_fields": {"seller_bank_swift": "BIC"} lors de la création d'un document de facturation.
		SellerPostCode          string               `json:"seller_post_code,omitempty"`    // code postal du vendeur
		SellerCity              string               `json:"seller_city,omitempty"`         // ville du vendeur
		SellerStreet            string               `json:"seller_street,omitempty"`       // numéro et nom de rue du vendeur
		SellerCountry           string               `json:"seller_country,omitempty"`      // pays du vendeur (ISO 3166)
		SellerEmail             string               `json:"seller_email,omitempty"`        // email du vendeur
		SellerWww               string               `json:"seller_www,omitempty"`          // site internet du vendeur
		SellerFax               string               `json:"seller_fax,omitempty"`          // numéro de fax du vendeur
		SellerPhone             string               `json:"seller_phone,omitempty"`        // numéro de tel du vendeur
		SellerPerson            string               `json:"seller_person,omitempty"`       // Nom du vendeur (figurant en bas de page des documents)
		ClientID                int64                `json:"client_id,omitempty"`           // ID du contact (si la valeur est -1 alors le contact sera créé et ajouté à la liste des contacts)
		BuyerName               string               `json:"buyer_name,omitempty"`          // nom du contact (acheteur en cas de vente ou fournisseur en cas d'achat)
		BuyerFirstName          string               `json:"buyer_first_name,omitempty"`    // du contact
		BuyerLastName           string               `json:"buyer_last_name,omitempty"`     // du contact
		BuyerCompany            bool                 `json:"buyer_company,omitempty"`       // si le contact est un professionnel, "0" si c'est un particulier
		BuyerTitle              string               `json:"buyer_title,omitempty"`         // Civilité du contact. Attention, en json vous devez envoyer ce paramètre comme ceci: "additional_fields": {"buyer_title"":"Mme"} lors de la création d'un document de facturation.
		BuyerTaxNo              string               `json:"buyer_tax_no,omitempty"`        // numéro d'identification fiscale du contact (ex: n° TVA)
		BuyerTaxNoKind          string               `json:"buyer_tax_no_kind,omitempty"`   // intitulé du numéro d'identification du contact : si non renseigné, il s'agit de "Numéro TVA", sinon il faut spécifier l'intitulé préalablement listé dans vos paramètres du compte, comme par exemple "SIREN" ou "CIF" (en savoir plus ici: https: //aide.vosfactures.fr/19032497-Num-ro-d-identification-fiscale-des-contacts)
		DisableTaxNoValidation  string               `json:"disable_tax_no_validation,omitempty"`
		UseMoss                 bool                 `json:"use_moss,omitempty"`                 // document sous le régime "Moss" (1) ou non (0) : régime de l'Autoliquidation de la TVA. En savoir plus ici: http: //vosfactures.fr/tva-moss
		BuyerPostCode           string               `json:"buyer_post_code,omitempty"`          // code postal du contact
		BuyerCity               string               `json:"buyer_city,omitempty"`               // ville du contact
		BuyerStreet             string               `json:"buyer_street,omitempty"`             // numéro et nom de rue du contact
		BuyerCountry            string               `json:"buyer_country,omitempty"`            // pays du contact (ISO 3166)
		BuyerNote               string               `json:"buyer_note,omitempty"`               // description additionnelle du contact
		DeliveryAddress         string               `json:"delivery_address,omitempty"`         // contenu du champ "Adresse supplémentaire" du contact
		UseDeliveryAddress      bool                 `json:"use_delivery_address,omitempty"`     // ou "false" afficher ou non le champ "Adresse supplémentaire" du contact sur le document
		BuyerEmail              string               `json:"buyer_email,omitempty"`              // email du contact
		BuyerPhone              string               `json:"buyer_phone,omitempty"`              // numéro de tel du contact
		BuyerMobilePhone        string               `json:"buyer_mobile_phone,omitempty"`       // numéro de portable du contact
		AdditionalInfo          bool                 `json:"additional_info,omitempty"`          // afficher (1) ou non (0) la colonne aditionnelle sur le document de facturation (dont l'intitulé est à définir dans Paramètres du compte > Options par défaut)
		AdditionalInfoDesc      string               `json:"additional_info_desc,omitempty"`     // contenu de la colonne aditionnelle (contenu que vous retrouvez sur la fiche du produit correspondant)
		AdditionalInvoiceField  string               `json:"additional_invoice_field,omitempty"` // contenu du champ additionnel (dont l'intitulé est à définir dans Paramètres du compte > Options par défaut). Attention, en json vous devez envoyer ce paramètre comme ceci: "additional_fields": {"additional_invoice_field": "contenu"} lors de la création d'un document de facturation.
		ShowDiscount            bool                 `json:"show_discount,omitempty"`            // afficher (1) ou non (0) la colonne réduction
		Discount                string               `json:"discount,omitempty"`
		DiscountKind            DiscountKind         `json:"discount_kind,omitempty"`             // type de réduction: "amount" (pour un montant ttc),
		PercentUnit             string               `json:"percent_unit,omitempty"`              // (pour un % sur le prix unitaire), ou  "percent_total" (pour un % calculé sur le prix total)
		PaymentType             PaymentType          `json:"payment_type,omitempty"`              // "chèque" - mode de règlement
		PaymentToKind           string               `json:"payment_to_kind,omitempty"`           // date limite de règlement (parmi les options proposées). Si l'option est "Autre" ("other_date"), vous pouvez définir une date spécifique grâce au champ "payment_to". Si vous indiquez "5", la date d'échéance est de 5 jours. Pour ne pas afficher ce champ, indiquez "off".
		PaymentTo               string               `json:"payment_to,omitempty"`                // date limite de règlement
		SumRecovery             string               `json:"sum_recovery,omitempty"`              // afficher (client_professionnel) ou non (client_particulier) la mention "Indemnité forfaitaire de recouvrement". Attention, en json vous devez envoyer ce paramètre comme ceci: "additional_fields": {"sum_recovery": "client_professionnel"} lors de la création d'un document de facturation.
		InterestRate            string               `json:"interest_rate,omitempty"`             // Taux de pénalité en cas de retard de paiement (attention, en json vous devez envoyer ce paramètre comme ceci: "additional_fields": {"interest_rate": "10%"} lors de la création d'un document de facturation.
		AdvancedPaymentDiscount string               `json:"advanced_payment_discount,omitempty"` // Escompte en % (attention, en json vous devez envoyer ce paramètre comme ceci: "additional_fields": { "advanced_payment_discount": "10"} lors de la création d'un document de facturation)
		Status                  Status               `json:"status,omitempty"`                    // état du document
		Paid                    string               `json:"paid,omitempty"`                      // montant payé
		Oid                     string               `json:"oid,omitempty"`                       // numéro de commande (ex: numéro généré par une application externe)
		OidUnique               string               `json:"oid_unique,omitempty"`                // si la valeur est «yes», alors il ne sera pas permis au système de créer 2 factures avec le même OID (cela peut être utile en cas de synchronisation avec une boutique en ligne)
		WarehouseID             int64                `json:"warehouse_id,omitempty"`              // numéro d'identification de l'entrepôt
		Description             string               `json:"description,omitempty"`               // Informations spécifiques
		PaidDate                string               `json:"paid_date,omitempty"`                 // Date du paiement ("Paiement reçu le")
		Currency                string               `json:"currency,omitempty"`                  // devise
		Lang                    Lang                 `json:"lang,omitempty"`                      // langue du document
		ExchangeCurrency        string               `json:"exchange_currency,omitempty"`         // convertir en (la conversion en une autre devise du montant total et du montant de la taxe selon taux de change du jour)
		ExchangeKind            string               `json:"exchange_kind,omitempty"`             // Source du taux de change utilisé en cas de conversion ("ecb" pour la Banque Centrale Européenne, "nbp" pour la Banque Nationale de Pologne, "cbr" pour la Banque Centrale de Russie, "nbu" pour la Banque Nationale d'Ukraine, "nbg" pour la Banque Nationale de Géorgie, "nbt" Banque Nationale Tchèque, "own" pour un taux propre)
		ExchangeCurrencyRate    string               `json:"exchange_currency_rate,omitempty"`    // taux de change propre (à utiliser uniquement si le paramètre "exchange_kind" est égal à "own") "title": "" Objet (attention, en json vous devez envoyer ce paramètre comme ceci: "additional_fields": {"title": "contenu de l'objet"} lors de la création d'un document de facturation).
		InternalNote            string               `json:"internal_note,omitempty"`             // Notes privées
		InvoiceTemplateID       int64                `json:"invoice_template_id,omitempty"`       // format d'impression
		DescriptionLong         string               `json:"description_long,omitempty"`          // Texte additionnel (imprimé sur la page suivante)
		FromInvoiceID           int64                `json:"from_invoice_id,omitempty"`           // ID du document à partir duquel le document a été généré (utile par ex quand une facture est générée depuis un devis)
		InvoiceID               int64                `json:"invoice_id,omitempty"`                // ID du document de référence ayant un lien fonctionnel avec le document (ex: le devis de référence pour un acompte).
		Positions               []*Position          `json:"positions,omitempty"`
		HideTax                 string               `json:"hide_tax,omitempty"` // Montant TTC uniquement (ne pas afficher de montant HT ni de taxe) (attention, en json vous devez envoyer ce paramètre comme ceci: "additional_fields": {"hide_tax": "1"    } lors de la création d'un document de facturation)
		CalculatingStrategy     *CalculatingStrategy `json:"calculating_strategy,omitempty"`
		SplitPayment            string               `json:"split_payment,omitempty"` // 1 ou 0 selon que la facture fait ou non l'objet d'un paiement partiel
	}

	Position struct {
		ID              int64  `json:"id,omitempty"`                // ID du produit
		ProductID       int64  `json:"product_id,omitempty"`        // ID du produit
		Name            string `json:"name,omitempty"`              // nom du produit
		Description     string `json:"description,omitempty"`       // description du produit
		Code            string `json:"code,omitempty"`              // Référence du produit
		AdditionalInfo  bool   `json:"additional_info,omitempty"`   // contenu de la colonne additionnelle
		DiscountPercent string `json:"discount_percent,omitempty"`  // % de la réduction
		Discount        string `json:"discount,omitempty"`          // montant ttc de la réduction
		Quantity        int64  `json:"quantity,omitempty"`          // quantité
		QuantityUnit    string `json:"quantity_unit,omitempty"`     // unité
		PriceNet        string `json:"price_net,omitempty"`         // prix unitaire HT (calculé automatiquement si non indiqué)
		Tax             string `json:"tax,omitempty"`               // % de taxe (les valeurs "disabled" ou "np" rendent la taxe inactive)
		PriceGross      string `json:"price_gross,omitempty"`       // prix unitaire TTC (calculé automatiquement si non indiqué)
		TotalPriceNet   string `json:"total_price_net,omitempty"`   // total HT (calculé automatiquement si non indiqué)
		TotalPriceGross string `json:"total_price_gross,omitempty"` // total TTC
		Kind            string `json:"kind,omitempty"`              // pour insérer une ligne de texte (voir exemple plus bas) "kind": "subtotal", // - pour insérer un sous-total (voir exemple plus bas)
		Destroy         bool   `json:"_destroy,omitempty"`          // pour supprimer un article sur la facture, entrez l'ID du produit avec le paramètre "_destroy" égal à 1.

	}

	CalculatingStrategy struct {
		Position             string `json:"position,omitempty"`                // ou "keep_gross" - Comment se calcule le total de chaque ligne
		Sum                  string `json:"sum,omitempty"`                     // ou "keep_gross" ou "keep_net" - Comment se calcule le total des colonnes
		InvoiceFormPriceKind string `json:"invoice_form_price_kind,omitempty"` // ou "gross" - prix unitaire (HT ou TTC)
	}
)

func (vf *VosFactures) CreateInvoice(document *Document) (*Document, error) {
	p := payloadInvoice{
		APIToken: vf.APIToken,
		Invoice:  document,
	}

	d := new(bytes.Buffer)
	if err := json.NewEncoder(d).Encode(p); err != nil {
		return document, err
	}
	req, err := http.NewRequest("POST", vf.url("/invoices.json"), d)
	if err != nil {
		return document, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := clientHttp.Do(req)
	if err != nil {
		return document, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		dumpRequest(req)
		dumpPayload(p)
		err := fmt.Errorf("%s\n%s", resp.Status, d)
		return document, err
	}

	var newDocument Document
	if err := json.NewDecoder(resp.Body).Decode(&newDocument); err != nil {
		return document, err
	}

	return &newDocument, nil
}

func (vf *VosFactures) ListInvoices(page int) ([]Document, error) {
	var invoices []Document
	req, err := http.NewRequest("GET", vf.url(fmt.Sprintf("/invoices.json?api_token=%s&page=%v", vf.APIToken, page)), nil)
	if err != nil {
		return invoices, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := clientHttp.Do(req)
	if err != nil {
		return invoices, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		dumpRequest(req)
		err := fmt.Errorf("status code : %v", resp.Status)
		return invoices, err
	}

	var newInvoices []Document
	if err := json.NewDecoder(resp.Body).Decode(&newInvoices); err != nil {
		return invoices, err
	}

	return newInvoices, nil
}

func (vf *VosFactures) GetInvoice(invoiceID int64) (*Document, error) {
	var document Document
	req, err := http.NewRequest("GET", vf.url(fmt.Sprintf("/invoices/%v.json?api_token=%s", invoiceID, vf.APIToken)), nil)
	if err != nil {
		return &document, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := clientHttp.Do(req)
	if err != nil {
		return &document, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		dumpRequest(req)
		err := fmt.Errorf("status code : %v", resp.Status)
		return &document, err
	}

	var newInvoice Document
	if err := json.NewDecoder(resp.Body).Decode(&newInvoice); err != nil {
		return &document, err
	}

	return &newInvoice, nil
}

func (vf *VosFactures) UpdateInvoice(document *Document) (*Document, error) {
	p := payloadInvoice{
		APIToken: vf.APIToken,
		Invoice:  document,
	}

	d := new(bytes.Buffer)
	err := json.NewEncoder(d).Encode(p)
	if err != nil {
		return document, err
	}
	req, err := http.NewRequest("PUT", vf.url(fmt.Sprintf("/invoices/%v.json", document.ID)), d)
	if err != nil {
		return document, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := clientHttp.Do(req)
	if err != nil {
		return document, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		dumpRequest(req)
		dumpPayload(p)
		err := fmt.Errorf("%s\n%s", resp.Status, d)
		return document, err
	}

	var newInvoice Document
	if err := json.NewDecoder(resp.Body).Decode(&newInvoice); err != nil {
		return document, err
	}

	return &newInvoice, nil
}

func (vf *VosFactures) DeleteInvoice(invoiceID int64) error {

	req, err := http.NewRequest("DELETE", vf.url(fmt.Sprintf("/invoices/%v.json?api_token=%s", invoiceID, vf.APIToken)), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := clientHttp.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		dumpRequest(req)
		err := fmt.Errorf("status code : %v", resp.Status)
		return err
	}

	return nil
}

func (vf *VosFactures) SendInvoiceByMail(invoiceID int64) error {

	req, err := http.NewRequest("POST", vf.url(fmt.Sprintf("/invoices/%v/send_by_email.json?api_token=%s", invoiceID, vf.APIToken)), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := clientHttp.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		dumpRequest(req)
		err := fmt.Errorf("status code : %v", resp.Status)
		return err
	}

	return nil
}
