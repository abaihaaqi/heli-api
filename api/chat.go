package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ijaybaihaqi/heli-api/model"
)

func (api *API) AnalyzeData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit /chat/analyze-data")

	var req model.Inputs

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	result, err := api.chatService.AnalyzeData(req.Table, req.Query)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Error getting response from AI"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (api *API) ChatWithAI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit /chat/new")

	var req model.ChatRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	result, err := api.chatService.ChatWithAI(req.Query)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Error getting response from AI"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result.GeneratedText)
}
