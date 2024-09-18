package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

type Request struct {
	ID      string      `json:"id"`
	Data    interface{} `json:"data"`
}

type Response struct {
	JobRunID string      `json:"jobRunID"`
	Data     interface{} `json:"data"`
	Result   interface{} `json:"result"`
	Status   string      `json:"status"`
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 调用你的 API
	apiURL := "https://your-api-url.com/endpoint"
	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var apiResponse interface{}
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 构建响应
	response := Response{
		JobRunID: req.ID,
		Data:     apiResponse,
		Result:   apiResponse,
		Status:   "success",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
