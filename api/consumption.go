package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ijaybaihaqi/heli-api/model"
)

func (api *API) FetchAllConsumptions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit /consumption/get-all")

	user, err := api.userService.FetchByUsername(api.userService.GetCurrentUsername())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	consumptions, err := api.consumptionService.FetchAll(user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(consumptions)
}

func (api *API) StoreConsumption(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit /consumption/add")

	var consumption model.Consumption

	err := json.NewDecoder(r.Body).Decode(&consumption)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	err = api.consumptionService.Store(&consumption)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(consumption)
}

func (api *API) ResetConsumption(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit /consumption/reset")

	user, err := api.userService.FetchByUsername(api.userService.GetCurrentUsername())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = api.consumptionService.Reset(user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Message: "User Appliance berhasil direset"})
}
