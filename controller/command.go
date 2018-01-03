// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2017-2018 Canonical Ltd
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func commandHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	fmt.Fprintf(os.Stdout, "command: device request: devid: %s cmd: %s", vars["deviceId"], vars["cmd"])
	io.WriteString(w, "OK")
}

func commandsHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	fmt.Fprintf(os.Stdout, "command: all devices request: cmd: %s", vars["cmd"])
	io.WriteString(w, "OK")
}

func initCommand(r *mux.Router) {
	s := r.PathPrefix("/api/v1/device").Subrouter()
	s.HandleFunc("/{deviceId}/{cmd}", commandHandler)
	s.HandleFunc("/all/{cmd}", commandsHandler)
}
