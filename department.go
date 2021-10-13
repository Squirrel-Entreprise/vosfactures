package vosfactures

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type (
	payloadDepartment struct {
		APIToken   string      `json:"api_token,omitempty"`
		Department *Department `json:"department,omitempty"`
	}

	Department struct {
		ID                           int64            `json:"id,omitempty"`
		Shortcut                     string           `json:"shortcut,omitempty"`
		Name                         string           `json:"name,omitempty"`
		TaxNo                        string           `json:"tax_no,omitempty"`
		PostCode                     string           `json:"post_code,omitempty"`
		City                         string           `json:"city,omitempty"`
		Street                       string           `json:"street,omitempty"`
		Person                       string           `json:"person,omitempty"`
		Country                      string           `json:"country,omitempty"`
		Email                        string           `json:"email,omitempty"`
		Phone                        string           `json:"phone,omitempty"`
		Www                          string           `json:"www,omitempty"`
		Fax                          string           `json:"fax,omitempty"`
		LogoFileName                 interface{}      `json:"logo_file_name,omitempty"`
		LogoContentType              interface{}      `json:"logo_content_type,omitempty"`
		LogoFileSize                 interface{}      `json:"logo_file_size,omitempty"`
		LogoUpdatedAt                interface{}      `json:"logo_updated_at,omitempty"`
		UsePattern                   bool             `json:"use_pattern,omitempty"`
		InvoicePattern               string           `json:"invoice_pattern,omitempty"`
		Bank                         string           `json:"bank,omitempty"`
		BankAccount                  string           `json:"bank_account,omitempty"`
		CreatedAt                    time.Time        `json:"created_at,omitempty"`
		UpdatedAt                    time.Time        `json:"updated_at,omitempty"`
		InvoicePatternBill           string           `json:"invoice_pattern_bill,omitempty"`
		InvoicePatternProforma       string           `json:"invoice_pattern_proforma,omitempty"`
		InvoicePatternCorrection     string           `json:"invoice_pattern_correction,omitempty"`
		WarehouseID                  interface{}      `json:"warehouse_id,omitempty"`
		RestrictWarehouses           bool             `json:"restrict_warehouses,omitempty"`
		Deleted                      bool             `json:"deleted,omitempty"`
		InvoicePatternReceipt        string           `json:"invoice_pattern_receipt,omitempty"`
		InvoicePatternAdvance        string           `json:"invoice_pattern_advance,omitempty"`
		InvoicePatternFinal          string           `json:"invoice_pattern_final,omitempty"`
		PatternVatRr                 string           `json:"pattern_vat_rr,omitempty"`
		PatternVatMp                 string           `json:"pattern_vat_mp,omitempty"`
		PatternVatMargin             string           `json:"pattern_vat_margin,omitempty"`
		TaxNoKind                    string           `json:"tax_no_kind,omitempty"`
		Main                         bool             `json:"main,omitempty"`
		StampFileName                interface{}      `json:"stamp_file_name,omitempty"`
		StampContentType             interface{}      `json:"stamp_content_type,omitempty"`
		StampFileSize                interface{}      `json:"stamp_file_size,omitempty"`
		StampUpdatedAt               interface{}      `json:"stamp_updated_at,omitempty"`
		PatternPz                    string           `json:"pattern_pz,omitempty"`
		PatternWz                    string           `json:"pattern_wz,omitempty"`
		PatternMm                    string           `json:"pattern_mm,omitempty"`
		PatternKp                    string           `json:"pattern_kp,omitempty"`
		PatternKw                    string           `json:"pattern_kw,omitempty"`
		PatternInvoiceOther          string           `json:"pattern_invoice_other,omitempty"`
		PatternEstimate              string           `json:"pattern_estimate,omitempty"`
		UseMassPayment               bool             `json:"use_mass_payment,omitempty"`
		MassPaymentPattern           interface{}      `json:"mass_payment_pattern,omitempty"`
		Kind                         string           `json:"kind,omitempty"`
		MobilePhone                  string           `json:"mobile_phone,omitempty"`
		UseCorrespondenceAddress     bool             `json:"use_correspondence_address,omitempty"`
		CorrespondenceAddress        interface{}      `json:"correspondence_address,omitempty"`
		CapitalKind                  string           `json:"capital_kind,omitempty"`
		Capital                      string           `json:"capital,omitempty"`
		CapitalCurrency              string           `json:"capital_currency,omitempty"`
		TaxKind                      string           `json:"tax_kind,omitempty"`
		Register1Status              string           `json:"register1_status,omitempty"`
		Register1Number1             string           `json:"register1_number1,omitempty"`
		Register1Number2             string           `json:"register1_number2,omitempty"`
		Register1Number3             string           `json:"register1_number3,omitempty"`
		Register2Status              string           `json:"register2_status,omitempty"`
		Register2Number1             string           `json:"register2_number1,omitempty"`
		Register2Number2             string           `json:"register2_number2,omitempty"`
		Register2Number3             interface{}      `json:"register2_number3,omitempty"`
		Register3Status              string           `json:"register3_status,omitempty"`
		Register3Number1             string           `json:"register3_number1,omitempty"`
		Register3Number2             string           `json:"register3_number2,omitempty"`
		Register3Number3             string           `json:"register3_number3,omitempty"`
		Register4Status              string           `json:"register4_status,omitempty"`
		Register4Number1             interface{}      `json:"register4_number1,omitempty"`
		Register4Number2             interface{}      `json:"register4_number2,omitempty"`
		Register4Number3             interface{}      `json:"register4_number3,omitempty"`
		Register5Status              interface{}      `json:"register5_status,omitempty"`
		Register5Number1             interface{}      `json:"register5_number1,omitempty"`
		Register5Number2             interface{}      `json:"register5_number2,omitempty"`
		Register5Number3             interface{}      `json:"register5_number3,omitempty"`
		BankAccountName              interface{}      `json:"bank_account_name,omitempty"`
		BankIban                     interface{}      `json:"bank_iban,omitempty"`
		BankSwift                    string           `json:"bank_swift,omitempty"`
		OwnEmailSettings             bool             `json:"own_email_settings,omitempty"`
		EmailFrom                    string           `json:"email_from,omitempty"`
		EmailCc                      string           `json:"email_cc,omitempty"`
		EmailSubject                 string           `json:"email_subject,omitempty"`
		EmailTemplate                interface{}      `json:"email_template,omitempty"`
		EmailTemplateKind            string           `json:"email_template_kind,omitempty"`
		EmailPdf                     bool             `json:"email_pdf,omitempty"`
		OwnOverdueEmailSettings      bool             `json:"own_overdue_email_settings,omitempty"`
		OverdueEmailSubject          string           `json:"overdue_email_subject,omitempty"`
		OverdueEmailTemplate         interface{}      `json:"overdue_email_template,omitempty"`
		OverdueEmailTemplateKind     string           `json:"overdue_email_template_kind,omitempty"`
		OverdueEmailPdf              bool             `json:"overdue_email_pdf,omitempty"`
		InvoiceLang                  interface{}      `json:"invoice_lang,omitempty"`
		InvoiceDescription           string           `json:"invoice_description,omitempty"`
		BankAccountCurrency          string           `json:"bank_account_currency,omitempty"`
		PatternZt                    string           `json:"pattern_zt,omitempty"`
		InvoicePatternCorrectionNote string           `json:"invoice_pattern_correction_note,omitempty"`
		PatternClientOrder           string           `json:"pattern_client_order,omitempty"`
		InvoicePatternAccountingNote string           `json:"invoice_pattern_accounting_note,omitempty"`
		OwnFooter                    bool             `json:"own_footer,omitempty"`
		FooterContent                string           `json:"footer_content,omitempty"`
		CalculateNumberByPatternOnly bool             `json:"calculate_number_by_pattern_only,omitempty"`
		FooterKind                   string           `json:"footer_kind,omitempty"`
		UseInvoiceIssuer             bool             `json:"use_invoice_issuer,omitempty"`
		CashInitState                interface{}      `json:"cash_init_state,omitempty"`
		InvoiceTemplateID            interface{}      `json:"invoice_template_id,omitempty"`
		AdditionalFields             AdditionalFields `json:"additional_fields,omitempty"`
		DefaultTax                   string           `json:"default_tax,omitempty"`
		LumpSumTaxed                 bool             `json:"lump_sum_taxed,omitempty"`
		DefaultLumpSumTax            interface{}      `json:"default_lump_sum_tax,omitempty"`
		RegisterNumber               interface{}      `json:"register_number,omitempty"`
		SecStampFileName             interface{}      `json:"sec_stamp_file_name,omitempty"`
		SecStampContentType          interface{}      `json:"sec_stamp_content_type,omitempty"`
		SecStampFileSize             interface{}      `json:"sec_stamp_file_size,omitempty"`
		SecStampUpdatedAt            interface{}      `json:"sec_stamp_updated_at,omitempty"`
	}

	AdditionalFields struct {
		OverduePrintTitle      string      `json:"overdue_print_title"`
		OverduePrintTextKind   string      `json:"overdue_print_text_kind"`
		OverduePrintTemplateID string      `json:"overdue_print_template_id"`
		OverduePrintText       interface{} `json:"overdue_print_text"`
		PatternPaymentReceipt  string      `json:"pattern_payment_receipt"`
	}
)

func (vf *VosFactures) CreateDepartment(department *Department) (*Department, error) {
	p := payloadDepartment{
		APIToken:   vf.APIToken,
		Department: department,
	}

	d := new(bytes.Buffer)
	err := json.NewEncoder(d).Encode(p)
	if err != nil {
		return department, err
	}

	req, err := http.NewRequest("POST", vf.url("/departments.json"), d)
	if err != nil {
		return department, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := clientHttp.Do(req)
	if err != nil {
		return department, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		dumpRequest(req)
		dumpPayload(p)
		err := fmt.Errorf("%s\n%s", resp.Status, d)
		return department, err
	}

	var newDepartment Department
	if err := json.NewDecoder(resp.Body).Decode(&newDepartment); err != nil {
		return department, err
	}

	return &newDepartment, nil
}

func (vf *VosFactures) ListDepartments(page int) ([]Department, error) {
	var departments []Department
	req, err := http.NewRequest("GET", vf.url(fmt.Sprintf("/departments.json?api_token=%s&page=%v", vf.APIToken, page)), nil)
	if err != nil {
		return departments, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := clientHttp.Do(req)
	if err != nil {
		return departments, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		err := fmt.Errorf("%s\n%v", resp.Status, resp)
		return departments, err
	}

	var newDepartments []Department
	if err := json.NewDecoder(resp.Body).Decode(&newDepartments); err != nil {
		return departments, err
	}

	return newDepartments, nil
}

func (vf *VosFactures) GetDepartment(departmentID int64) (*Department, error) {
	var department Department
	req, err := http.NewRequest("GET", vf.url(fmt.Sprintf("/departments/%v.json?api_token=%s", departmentID, vf.APIToken)), nil)
	if err != nil {
		return &department, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := clientHttp.Do(req)
	if err != nil {
		return &department, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		err := fmt.Errorf("%s\n%v", resp.Status, resp)
		return &department, err
	}

	var newDepartment Department
	if err := json.NewDecoder(resp.Body).Decode(&newDepartment); err != nil {
		return &department, err
	}

	return &newDepartment, nil
}

func (vf *VosFactures) UpdateDepartment(department *Department) (*Department, error) {
	p := payloadDepartment{
		APIToken:   vf.APIToken,
		Department: department,
	}

	d := new(bytes.Buffer)
	err := json.NewEncoder(d).Encode(p)
	if err != nil {
		return department, err
	}

	req, err := http.NewRequest("PUT", vf.url(fmt.Sprintf("/departments/%v.json", department.ID)), d)
	if err != nil {
		return department, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := clientHttp.Do(req)
	if err != nil {
		return department, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		dumpRequest(req)
		dumpPayload(p)
		err := fmt.Errorf("%s\n%s", resp.Status, d)
		return department, err
	}

	var newDepartment Department
	if err := json.NewDecoder(resp.Body).Decode(&newDepartment); err != nil {
		return department, err
	}

	return &newDepartment, nil
}

func (vf *VosFactures) DeleteDepartment(departmentID int64) error {

	req, err := http.NewRequest("DELETE", vf.url(fmt.Sprintf("/departments/%v.json?api_token=%s", departmentID, vf.APIToken)), nil)
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
