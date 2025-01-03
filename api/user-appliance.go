package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ijaybaihaqi/heli-api/model"
)

func (api *API) FetchAllUserAppliance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit /user-appliance/get-all")

	user, err := api.userService.FetchByUsername(api.userService.GetCurrentUsername())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userAppliances, err := api.userApplianceService.FetchAll(user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userAppliances)
}

func (api *API) FetchUserApplianceByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit /user-appliance/get")

	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userAppliance, err := api.userApplianceService.FetchByID(idInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userAppliance)
}

func (api *API) StoreUserAppliance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit /user-appliance/add")

	var req model.UserApplianceRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	user, err := api.userService.FetchByUsername(api.userService.GetCurrentUsername())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	var newUserAppliance model.UserAppliance
	newUserAppliance.UserID = user.ID
	newUserAppliance.ApplianceID = req.ApplianceID
	newUserAppliance.Room = req.Room

	err = api.userApplianceService.Store(&newUserAppliance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	message := fmt.Sprintf("Successfully add appliance to room %s", req.Room)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Message: message})
}

func (api *API) UpdateUserAppliance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit /user-appliance/update")

	var req model.UserApplianceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	var updatedUserAppliance model.UserAppliance
	updatedUserAppliance.Room = req.Room

	err = api.userApplianceService.Update(req.ID, &updatedUserAppliance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Message: "User Appliance berhasil diubah"})
}

func (api *API) DeleteUserAppliance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit /user-appliance/delete")

	id := r.URL.Query().Get("id")

	err := api.userApplianceService.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Message: "User Appliance berhasil dihapus"})
}
