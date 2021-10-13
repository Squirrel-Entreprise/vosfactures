package vosfactures

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
)

func dumpRequest(req *http.Request) error {
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		return err
	}
	fmt.Println(string(requestDump))
	return nil
}

func dumpPayload(s interface{}) error {
	b, err := json.MarshalIndent(s, "", "  ")

	if err != nil {
		return err
	}

	fmt.Println("Body :", string(b))
	return nil
}
