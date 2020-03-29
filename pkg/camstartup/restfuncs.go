package camstartup

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/edgex-camera/device-sdk-go/internal/common"
	"github.com/gorilla/mux"
)

const APIV1Prefix = "/api/v1"

func (ds *DService) appendRouter(r *mux.Router) {
	r.HandleFunc("/name", ds.handleName)

	r.PathPrefix("/driver/{dsname}/api/v1").HandlerFunc(ds.handleAPI)
	r.PathPrefix("/driver/{dsname}").HandlerFunc(ds.handleFrontend)
}

func (ds *DService) handleName(w http.ResponseWriter, r *http.Request) {
	name := []byte(ds.Name())
	if _, err := w.Write(name); err != nil {
		common.LoggingClient.Error(err.Error())
	}
}

func (ds *DService) checkDsname(r *http.Request) string {
	vars := mux.Vars(r)
	dsname := vars["dsname"]

	if ds.Name() != dsname {
		panic(fmt.Errorf("Wrong proxy: dsname %v != proxy name %v", ds.Name(), dsname))
	}
	return dsname
}

func (ds *DService) handleAPI(w http.ResponseWriter, r *http.Request) {
	dsname := ds.checkDsname(r)
	if ds.APIHandler != nil {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, APIV1Prefix+"/driver/"+dsname)
		ds.APIHandler.ServeHTTP(w, r)
	} else {
		http.Error(w, "not supported", http.StatusNotImplemented)
	}
}

func (ds *DService) handleFrontend(w http.ResponseWriter, r *http.Request) {
	dsname := ds.checkDsname(r)
	r.URL.Path = strings.TrimPrefix(r.URL.Path, APIV1Prefix+"/driver/"+dsname)
	ds.frontendHandler.ServeHTTP(w, r)
}
