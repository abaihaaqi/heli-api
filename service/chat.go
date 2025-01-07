package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ijaybaihaqi/heli-api/model"
)

type ChatService interface {
	AnalyzeData(table map[string][]string, query string) (string, error)
	ChatWithAI(query string) (model.ChatResponse, error)
}

type chatServiceImpl struct {
	HuggingfaceToken string
}

func NewChatService(HuggingfaceToken string) *chatServiceImpl {
	return &chatServiceImpl{HuggingfaceToken: HuggingfaceToken}
}

func (c *chatServiceImpl) AnalyzeData(table map[string][]string, query string) (string, error) {
	if len(table) == 0 {
		return "", fmt.Errorf("table is empty")
	}

	data := model.AnalyzePayload{
		Inputs: model.Inputs{
			Query: query,
			Table: table,
		},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error on json.Marshal()")
		return "", err
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api-inference.huggingface.co/models/google/tapas-base-finetuned-wtq", body)
	if err != nil {
		fmt.Println("error on http.NewRequest()")
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-wait-for-model", "true")
	req.Header.Set("Authorization", "Bearer "+c.HuggingfaceToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error on s.Client.Do()")
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error on io.ReadAll()")
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%s", string(respBody))
	}

	result := model.TapasResponse{}
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return "", err
	}

	return result.Answer, nil
}

func (c *chatServiceImpl) ChatWithAI(query string) (model.ChatResponse, error) {
	data := model.ChatPayload{
		Inputs:    query,
		MaxTokens: 1000,
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error on json.Marshal()")
		return model.ChatResponse{}, err
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api-inference.huggingface.co/models/microsoft/Phi-3.5-mini-instruct", body)
	if err != nil {
		fmt.Println("error on http.NewRequest()")
		return model.ChatResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-wait-for-model", "true")
	req.Header.Set("Authorization", "Bearer "+c.HuggingfaceToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return model.ChatResponse{}, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.ChatResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return model.ChatResponse{}, fmt.Errorf("error http response")
	}

	result := []model.ChatResponse{}
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return model.ChatResponse{}, err
	}

	return result[0], err
}
