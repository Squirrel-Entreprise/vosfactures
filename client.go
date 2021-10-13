package vosfactures

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	payloadClient struct {
		APIToken string  `json:"api_token,omitempty"`
		Client   *Client `json:"client,omitempty"`
	}

	Client struct {
		ID          int64  `json:"id,omitempty"`
		Name        string `json:"name,omitempty"`
		TaxNo       string `json:"tax_no,omitempty"`
		Bank        string `json:"bank,omitempty"`
		BankAccount string `json:"bank_account,omitempty"`
		City        string `json:"city,omitempty"`
		Country     string `json:"country,omitempty"`
		Email       string `json:"email,omitempty"`
		Person      string `json:"person,omitempty"`
		PostCode    string `json:"post_code,omitempty"`
		Phone       string `json:"phone,omitempty"`
		MobilePhone string `json:"mobile_phone,omitempty"`
		Street      string `json:"street,omitempty"`
		PanelURL    string `json:"panel_url,omitempty"`
	}
)

func (vf *VosFactures) CreateClient(client *Client) (*Client, error) {
	p := payloadClient{
		APIToken: vf.APIToken,
		Client:   client,
	}

	d := new(bytes.Buffer)
	err := json.NewEncoder(d).Encode(p)
	if err != nil {
		return client, err
	}

	req, err := http.NewRequest("POST", vf.url("/clients.json"), d)
	if err != nil {
		return client, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := clientHttp.Do(req)
	if err != nil {
		return client, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		dumpRequest(req)
		dumpPayload(p)
		err := fmt.Errorf("%s\n%s", resp.Status, d)
		return client, err
	}

	var newClient Client
	if err := json.NewDecoder(resp.Body).Decode(&newClient); err != nil {
		return client, err
	}

	return &newClient, nil
}

func (vf *VosFactures) ListClients(page int) ([]Client, error) {
	var clients []Client
	req, err := http.NewRequest("GET", vf.url(fmt.Sprintf("/clients.json?api_token=%s&page=%v", vf.APIToken, page)), nil)
	if err != nil {
		return clients, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := clientHttp.Do(req)
	if err != nil {
		return clients, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		dumpRequest(req)
		err := fmt.Errorf("%s", resp.Status)
		return clients, err
	}

	var newClients []Client
	if err := json.NewDecoder(resp.Body).Decode(&newClients); err != nil {
		return clients, err
	}

	return newClients, nil
}

func (vf *VosFactures) GetClient(clientID int64) (*Client, error) {
	var client Client
	req, err := http.NewRequest("GET", vf.url(fmt.Sprintf("/clients/%v.json?api_token=%s", clientID, vf.APIToken)), nil)
	if err != nil {
		return &client, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := clientHttp.Do(req)
	if err != nil {
		return &client, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		dumpRequest(req)
		err := fmt.Errorf("status code : %v", resp.Status)
		return &client, err
	}

	var newClient Client
	if err := json.NewDecoder(resp.Body).Decode(&newClient); err != nil {
		return &client, err
	}

	return &newClient, nil
}

func (vf *VosFactures) UpdateClient(client *Client) (*Client, error) {
	p := payloadClient{
		APIToken: vf.APIToken,
		Client:   client,
	}

	d := new(bytes.Buffer)
	err := json.NewEncoder(d).Encode(p)
	if err != nil {
		return client, err
	}
	req, err := http.NewRequest("PUT", vf.url(fmt.Sprintf("/clients/%v.json", client.ID)), d)
	if err != nil {
		return client, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := clientHttp.Do(req)
	if err != nil {
		return client, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		dumpRequest(req)
		dumpPayload(p)
		err := fmt.Errorf("%s\n%s", resp.Status, d)
		return client, err
	}

	var newClient Client
	if err := json.NewDecoder(resp.Body).Decode(&newClient); err != nil {
		return client, err
	}

	return &newClient, nil
}

func (vf *VosFactures) DeleteClient(clientID int64) error {

	req, err := http.NewRequest("DELETE", vf.url(fmt.Sprintf("/clients/%v.json?api_token=%s", clientID, vf.APIToken)), nil)
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
