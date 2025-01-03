package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (api *API) FetchAllAppliance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit /appliance/get-all")

	appliances, err := api.applianceService.FetchAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(appliances)
}
