package vosfactures

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type (
	payloadCategory struct {
		APIToken string    `json:"api_token,omitempty"`
		Category *Category `json:"category,omitempty"`
	}

	Category struct {
		ID          int64     `json:"id,omitempty"`
		Name        string    `json:"name,omitempty"`
		Description string    `json:"description,omitempty"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		AltName     string    `json:"alt_name"`
		ProductTags string    `json:"product_tags"`
	}
)

func (vf *VosFactures) CreateCategory(category *Category) (*Category, error) {
	p := payloadCategory{
		APIToken: vf.APIToken,
		Category: category,
	}

	d := new(bytes.Buffer)
	err := json.NewEncoder(d).Encode(p)
	if err != nil {
		return category, err
	}

	req, err := http.NewRequest("POST", vf.url("/categories.json"), d)
	if err != nil {
		return category, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := clientHttp.Do(req)
	if err != nil {
		return category, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		dumpRequest(req)
		dumpPayload(p)
		err := fmt.Errorf("%s\n%s", resp.Status, d)
		return category, err
	}

	var newCategory Category
	if err := json.NewDecoder(resp.Body).Decode(&newCategory); err != nil {
		return category, err
	}

	return &newCategory, nil
}

func (vf *VosFactures) ListCategories(page int) ([]Category, error) {
	var categories []Category
	req, err := http.NewRequest("GET", vf.url(fmt.Sprintf("/categories.json?api_token=%s&page=%v", vf.APIToken, page)), nil)
	if err != nil {
		return categories, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := clientHttp.Do(req)
	if err != nil {
		return categories, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		err := fmt.Errorf("%s\n%v", resp.Status, resp)
		return categories, err
	}

	var newCategories []Category
	if err := json.NewDecoder(resp.Body).Decode(&newCategories); err != nil {
		return categories, err
	}

	return newCategories, nil
}

func (vf *VosFactures) GetCategory(categoryID int64) (*Category, error) {
	var category Category
	req, err := http.NewRequest("GET", vf.url(fmt.Sprintf("/categories/%v.json?api_token=%s", categoryID, vf.APIToken)), nil)
	if err != nil {
		return &category, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := clientHttp.Do(req)
	if err != nil {
		return &category, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		err := fmt.Errorf("%s\n%v", resp.Status, resp)
		return &category, err
	}

	var newCategory Category
	if err := json.NewDecoder(resp.Body).Decode(&newCategory); err != nil {
		return &category, err
	}

	return &newCategory, nil
}

func (vf *VosFactures) UpdateCategory(category *Category) (*Category, error) {
	p := payloadCategory{
		APIToken: vf.APIToken,
		Category: category,
	}

	d := new(bytes.Buffer)
	err := json.NewEncoder(d).Encode(p)
	if err != nil {
		return category, err
	}

	req, err := http.NewRequest("PUT", vf.url(fmt.Sprintf("/categories/%v.json", category.ID)), d)
	if err != nil {
		return category, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := clientHttp.Do(req)
	if err != nil {
		return category, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		dumpRequest(req)
		dumpPayload(p)
		err := fmt.Errorf("%s\n%s", resp.Status, d)
		return category, err
	}

	var newCategory Category
	if err := json.NewDecoder(resp.Body).Decode(&newCategory); err != nil {
		return category, err
	}

	return &newCategory, nil
}

func (vf *VosFactures) DeleteCategory(categoryID int64) error {

	req, err := http.NewRequest("DELETE", vf.url(fmt.Sprintf("/categories/%v.json?api_token=%s", categoryID, vf.APIToken)), nil)
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
