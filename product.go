package vosfactures

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type (
	payloadProduct struct {
		APIToken string   `json:"api_token,omitempty"`
		Product  *Product `json:"product,omitempty"`
	}

	Product struct {
		ID                        int64         `json:"id,omitempty"`
		Name                      string        `json:"name,omitempty"`
		Code                      string        `json:"code,omitempty"`
		PriceGross                string        `json:"price_gross,omitempty"` // TTC
		PriceNet                  string        `json:"price_net,omitempty"`   // HT
		Tax                       string        `json:"tax,omitempty"`
		CategoryID                int64         `json:"category_id,omitempty"`
		Currency                  string        `json:"currency,omitempty"` // EUR
		Description               interface{}   `json:"description,omitempty"`
		CreatedAt                 time.Time     `json:"created_at,omitempty"`
		UpdatedAt                 time.Time     `json:"updated_at,omitempty"`
		AutomaticSales            bool          `json:"automatic_sales,omitempty"`
		Limited                   bool          `json:"limited,omitempty"`
		WarehouseQuantity         string        `json:"warehouse_quantity,omitempty"`
		AvailableFrom             interface{}   `json:"available_from,omitempty"`
		AvailableTo               interface{}   `json:"available_to,omitempty"`
		PaymentCallback           interface{}   `json:"payment_callback,omitempty"`
		PaymentURLOk              interface{}   `json:"payment_url_ok,omitempty"`
		PaymentURLError           interface{}   `json:"payment_url_error,omitempty"`
		Token                     string        `json:"token,omitempty"`
		Quantity                  string        `json:"quantity,omitempty"`
		QuantityUnit              interface{}   `json:"quantity_unit,omitempty"`
		AdditionalInfo            interface{}   `json:"additional_info,omitempty"`
		Disabled                  bool          `json:"disabled,omitempty"`
		PriceTax                  string        `json:"price_tax,omitempty"`
		FormFieldsHorizontal      bool          `json:"form_fields_horizontal,omitempty"`
		FormFields                interface{}   `json:"form_fields,omitempty"`
		FormName                  string        `json:"form_name,omitempty"`
		FormDescription           interface{}   `json:"form_description,omitempty"`
		QuantitySoldOutside       interface{}   `json:"quantity_sold_outside,omitempty"`
		FormKind                  string        `json:"form_kind,omitempty"`
		FormTemplate              interface{}   `json:"form_template,omitempty"`
		ElasticPrice              bool          `json:"elastic_price,omitempty"`
		NextProductID             interface{}   `json:"next_product_id,omitempty"`
		QuantitySoldInInvoices    string        `json:"quantity_sold_in_invoices,omitempty"`
		Deleted                   bool          `json:"deleted,omitempty"`
		Ecommerce                 bool          `json:"ecommerce,omitempty"`
		Period                    interface{}   `json:"period,omitempty"`
		ShowElasticPrice          bool          `json:"show_elastic_price,omitempty"`
		ElasticPriceDetails       interface{}   `json:"elastic_price_details,omitempty"`
		ElasticPriceDateTrigger   interface{}   `json:"elastic_price_date_trigger,omitempty"`
		Iid                       interface{}   `json:"iid,omitempty"`
		PurchasePriceNet          interface{}   `json:"purchase_price_net,omitempty"`
		PurchasePriceGross        interface{}   `json:"purchase_price_gross,omitempty"`
		UseFormula                bool          `json:"use_formula,omitempty"`
		Formula                   interface{}   `json:"formula,omitempty"`
		FormulaTestField          interface{}   `json:"formula_test_field,omitempty"`
		StockLevel                string        `json:"stock_level,omitempty"`
		Sync                      bool          `json:"sync,omitempty"`
		Kind                      string        `json:"kind,omitempty"`
		Package                   bool          `json:"package,omitempty"`
		PackageProductIds         interface{}   `json:"package_product_ids,omitempty"`
		DepartmentID              interface{}   `json:"department_id,omitempty"`
		UseProductWarehouses      bool          `json:"use_product_warehouses,omitempty"`
		PurchasePriceTax          interface{}   `json:"purchase_price_tax,omitempty"`
		PurchaseTax               interface{}   `json:"purchase_tax,omitempty"`
		Service                   bool          `json:"service,omitempty"`
		UseQuantityDiscount       bool          `json:"use_quantity_discount,omitempty"`
		QuantityDiscountDetails   interface{}   `json:"quantity_discount_details,omitempty"`
		PriceNetOnPayment         bool          `json:"price_net_on_payment,omitempty"`
		WarehouseNumbersUpdatedAt interface{}   `json:"warehouse_numbers_updated_at,omitempty"`
		EanCode                   interface{}   `json:"ean_code,omitempty"`
		Weight                    interface{}   `json:"weight,omitempty"`
		WeightUnit                interface{}   `json:"weight_unit,omitempty"`
		SizeHeight                interface{}   `json:"size_height,omitempty"`
		SizeWidth                 interface{}   `json:"size_width,omitempty"`
		Size                      interface{}   `json:"size,omitempty"`
		SizeUnit                  interface{}   `json:"size_unit,omitempty"`
		AutoPaymentDepartmentID   interface{}   `json:"auto_payment_department_id,omitempty"`
		AttachmentsCount          int           `json:"attachments_count,omitempty"`
		ImageURL                  interface{}   `json:"image_url,omitempty"`
		Tax2                      string        `json:"tax2,omitempty"`
		PurchaseTax2              string        `json:"purchase_tax2,omitempty"`
		SupplierCode              interface{}   `json:"supplier_code,omitempty"`
		PackageProductsDetails    interface{}   `json:"package_products_details,omitempty"`
		SiteorDisabled            bool          `json:"siteor_disabled,omitempty"`
		UseMoss                   bool          `json:"use_moss,omitempty"`
		SubscriptionID            interface{}   `json:"subscription_id,omitempty"`
		AccountingID              interface{}   `json:"accounting_id,omitempty"`
		Status                    interface{}   `json:"status,omitempty"`
		RestrictedToWarehouses    bool          `json:"restricted_to_warehouses,omitempty"`
		GtuCodes                  []interface{} `json:"gtu_codes,omitempty"`
		TagList                   []interface{} `json:"tag_list,omitempty"`
		ElectronicService         interface{}   `json:"electronic_service,omitempty"`
	}
)

func (vf *VosFactures) CreateProduct(product *Product) (*Product, error) {
	p := payloadProduct{
		APIToken: vf.APIToken,
		Product:  product,
	}

	d := new(bytes.Buffer)
	err := json.NewEncoder(d).Encode(p)
	if err != nil {
		return product, err
	}

	req, err := http.NewRequest("POST", vf.url("/products.json"), d)
	if err != nil {
		return product, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := clientHttp.Do(req)
	if err != nil {
		return product, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		dumpRequest(req)
		dumpPayload(p)
		err := fmt.Errorf("%s\n%s", resp.Status, d)
		return product, err
	}

	var newProduct Product
	if err := json.NewDecoder(resp.Body).Decode(&newProduct); err != nil {
		return product, err
	}

	return &newProduct, nil
}

func (vf *VosFactures) ListProducts(page int) ([]Product, error) {
	var products []Product
	req, err := http.NewRequest("GET", vf.url(fmt.Sprintf("/products.json?api_token=%s&page=%v", vf.APIToken, page)), nil)
	if err != nil {
		return products, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := clientHttp.Do(req)
	if err != nil {
		return products, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		err := fmt.Errorf("%s\n%v", resp.Status, resp)
		return products, err
	}

	var newProducts []Product
	if err := json.NewDecoder(resp.Body).Decode(&newProducts); err != nil {
		return products, err
	}

	return newProducts, nil
}

func (vf *VosFactures) GetProduct(productID int64) (*Product, error) {
	var product Product
	req, err := http.NewRequest("GET", vf.url(fmt.Sprintf("/products/%v.json?api_token=%s", productID, vf.APIToken)), nil)
	if err != nil {
		return &product, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := clientHttp.Do(req)
	if err != nil {
		return &product, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		err := fmt.Errorf("%s\n%v", resp.Status, resp)
		return &product, err
	}

	var newProduct Product
	if err := json.NewDecoder(resp.Body).Decode(&newProduct); err != nil {
		return &product, err
	}

	return &newProduct, nil
}

func (vf *VosFactures) UpdateProduct(product *Product) (*Product, error) {
	p := payloadProduct{
		APIToken: vf.APIToken,
		Product:  product,
	}

	d := new(bytes.Buffer)
	err := json.NewEncoder(d).Encode(p)
	if err != nil {
		return product, err
	}

	req, err := http.NewRequest("PUT", vf.url(fmt.Sprintf("/products/%v.json", product.ID)), d)
	if err != nil {
		return product, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := clientHttp.Do(req)
	if err != nil {
		return product, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		dumpRequest(req)
		dumpPayload(p)
		err := fmt.Errorf("%s\n%s", resp.Status, d)
		return product, err
	}

	var newProduct Product
	if err := json.NewDecoder(resp.Body).Decode(&newProduct); err != nil {
		return product, err
	}

	return &newProduct, nil
}

func (vf *VosFactures) DeleteProduct(productID int64) error {

	req, err := http.NewRequest("DELETE", vf.url(fmt.Sprintf("/products/%v.json?api_token=%s", productID, vf.APIToken)), nil)
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
		err := fmt.Errorf("%s\n%v", resp.Status, resp)
		return err
	}

	return nil
}
