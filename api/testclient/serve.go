package testclient

import (
	"net/http"

	"github.com/khulnasoft-lab/gobaseline/api"
)

func init() {
	api.RegisterHandler("/test/", http.StripPrefix("/test/", http.FileServer(http.Dir("./api/testclient/root/"))))
}
