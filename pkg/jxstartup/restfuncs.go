package jxstartup

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/edgexfoundry/device-sdk-go/internal/common"
	"github.com/gorilla/mux"
)

const APIV1Prefix = "/api/v1"

func (js *JxService) appendRouter(r *mux.Router) {
	sr := r.PathPrefix(APIV1Prefix).Subrouter()

	sr.HandleFunc("/name", js.handleName)

	sr.HandleFunc("/driver/{dsname}/api/v1/{path:.*}", js.handleAPI)
	sr.HandleFunc("/driver/{dsname}{path:.*}", js.handleFrontend)
}

func (js *JxService) handleName(w http.ResponseWriter, r *http.Request) {
	name := []byte(js.Name())
	if _, err := w.Write(name); err != nil {
		common.LoggingClient.Error(err.Error())
	}
}

func (js *JxService) checkDsname(r *http.Request) string {
	vars := mux.Vars(r)
	dsname := vars["dsname"]

	if js.Name() != dsname {
		panic(fmt.Errorf("Wrong proxy: dsname %v != proxy name %v", js.Name(), dsname))
	}
	return dsname
}

func (js *JxService) handleAPI(w http.ResponseWriter, r *http.Request) {
	dsname := js.checkDsname(r)
	if js.APIHandler != nil {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, APIV1Prefix+"/driver/"+dsname)
		js.APIHandler.ServeHTTP(w, r)
	} else {
		http.Error(w, "not supported", http.StatusNotImplemented)
	}
}

func (js *JxService) handleFrontend(w http.ResponseWriter, r *http.Request) {
	js.checkDsname(r)
	js.frontendHandler.ServeHTTP(w, r)
}
