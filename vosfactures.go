package vosfactures

import (
	"fmt"
	"net/http"
	"time"
)

type (
	VosFactures struct {
		APIToken string
		Account  string
	}
)

const (
	protocol = "https:"
	suffix   = "vosfactures.fr"
)

var (
	clientHttp = &http.Client{
		Timeout: 15 * time.Second,
	}
)

func (vf *VosFactures) url(path string) string {
	return fmt.Sprintf("%s//%s.%s%s", protocol, vf.Account, suffix, path)
}

// New create new VosFactures{}
func New(t, a string) *VosFactures {
	return &VosFactures{
		APIToken: t,
		Account:  a,
	}
}
