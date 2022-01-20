package handlers

import (
	"encoding/json"
	"github.com/enspzr/getir-go-assigment/model"
	"log"
	"net/http"
)

// These functions were written to respond the requests from one place.

func successRecordResponse(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	resp, err := createRecordResponse(data, "Success", 0)
	if err != nil {
		log.Println("Create response error => " + err.Error())
		return
	}
	_, err = w.Write(resp)
	if err != nil {
		log.Println("Create response error => " + err.Error())
	}
}

func internalError(w http.ResponseWriter, err error) {
	errorResponse(w, 500, 1, err.Error())
}

func badRequestError(w http.ResponseWriter, err error) {
	errorResponse(w, 400, 2, err.Error())
}

func methodNotAllowedError(w http.ResponseWriter, err error) {
	errorResponse(w, 415, 3, err.Error())
}

func errorResponse(w http.ResponseWriter, status, code int, message string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	resp, err := createRecordResponse(nil, message, code)
	if err != nil {
		log.Println("Create response error => " + err.Error())
		return
	}
	_, err = w.Write(resp)
	if err != nil {
		log.Println("Response send error => " + err.Error())
	}
}

func createRecordResponse(data interface{}, msg string, code int) ([]byte, error) {
	resp := model.Response{
		Code:    code,
		Msg:     msg,
		Records: data,
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	return jsonResp, err
}

func successInMemoryResponse(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(data)
	if err != nil {
		log.Println("Json marshal error => " + err.Error())
		return
	}
	_, err = w.Write(resp)
	if err != nil {
		log.Println("Response send error => " + err.Error())
	}
}

func errorResponseInMemory(w http.ResponseWriter, message error, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte(message.Error()))
	if err != nil {
		log.Println("Response send error => " + err.Error())
	}
}
