package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func getHTTPRouterParamID(r *http.Request, param string) (int, error) {
	params := httprouter.ParamsFromContext(r.Context())
	pid := params.ByName(param)
	id, err := strconv.Atoi(pid)
	if err != nil {
		return 0, err
	}
	return id, err
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("error: %v: %v\n", r.URL, err)
	msg := ErrorHTTP{Error: err.Error()}
	w.WriteHeader(http.StatusBadRequest)
	encodeJSON(w, r, msg)
}

func encodeJSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	e := json.NewEncoder(w)
	// e.SetIndent(" ", " ")
	if err := e.Encode(v); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func decodeJSON(w http.ResponseWriter, r *http.Request, v interface{}) error {
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := d.Decode(v); err != nil {
		return err
	}
	return nil
}
