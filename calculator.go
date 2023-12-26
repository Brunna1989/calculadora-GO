package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	Value1    float64 `json:"value1"`
	Value2    float64 `json:"value2"`
	Result    float64 `json:"result"`
	Operation string  `json:"operation"`
}

type ResponseBody struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

const (
	DefaultPort      = 8080
	SuccessStatus    = http.StatusOK
	BadRequestStatus = http.StatusBadRequest
)

var history []Operation

func jsonResponse(w http.ResponseWriter, message string, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(ResponseBody{Message: message, Data: data})
}

func calculate(value1, value2 float64, operation string) (float64, error) {
	switch operation {
	case "sum":
		return value1 + value2, nil
	case "sub":
		return value1 - value2, nil
	case "mul":
		return value1 * value2, nil
	case "div":
		if value2 == 0.0 {
			return 0, errors.New("cannot divide by zero")
		}
		return value1 / value2, nil
	default:
		return 0, errors.New(operation + " is not a valid operation")
	}
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 5 {
		jsonResponse(w, "missing path parameters in /calc/{operation}/{value1}/{value2}", nil, BadRequestStatus)
		return
	}

	operation := pathParts[2]
	value1, err1 := strconv.ParseFloat(pathParts[3], 64)
	value2, err2 := strconv.ParseFloat(pathParts[4], 64)

	if err1 != nil || err2 != nil {
		jsonResponse(w, "invalid input values", nil, BadRequestStatus)
		return
	}

	result, err := calculate(value1, value2, operation)
	if err != nil {
		jsonResponse(w, err.Error(), nil, BadRequestStatus)
		return
	}

	opInfo := Operation{value1, value2, result, operation}
	history = append(history, opInfo)

	jsonResponse(w, "operation performed successfully", opInfo, SuccessStatus)
}

func historyHandler(w http.ResponseWriter, r *http.Request) {
	if len(history) == 0 {
		jsonResponse(w, "empty history", nil, SuccessStatus)
		return
	}
	jsonResponse(w, "operations history", history, SuccessStatus)
}

func getServerPort() int {
	if value, err := strconv.Atoi(os.Getenv("SERVER_PORT")); err == nil {
		return value
	}
	return DefaultPort
}

func main() {
	fmt.Println("server running")
	http.HandleFunc("/calc/", calcHandler)
	http.HandleFunc("/calc/history", historyHandler)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", getServerPort()), nil)
}
