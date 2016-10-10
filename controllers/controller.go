// controller
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"servertest/common"
	"time"
)

var Delay time.Duration
var IsPrintVal bool
var Status int
var genericMessage interface{}

func SetDelay(delay time.Duration) {
	Delay = delay
}

func SetStatus(status int) {
	Status = status
}

func SetIsPrintVal(is bool) {
	IsPrintVal = is
}

//handler for post object
//post and return object
func ServerTest(w http.ResponseWriter, r *http.Request) {
	delay := Delay
	status := Status
	isPrintVal := IsPrintVal
	time.Sleep(delay)
	err := json.NewDecoder(r.Body).Decode(&genericMessage)
	if err != nil {
		common.DisplayAppError(w,
			err,
			"Invalid Object Data",
			500,
		)
		return
	}
	message := genericMessage.(map[string]interface{})
	if isPrintVal {
		for k, v := range message {
			switch vv := v.(type) {
			case int:
				fmt.Println(">>>>>field", k, ":", vv)
			case string:
				fmt.Println(">>>>>field", k, ":", vv)
			case []interface{}:
				fmt.Println(">>>>>field", k, "is an array:")
				for i, u := range vv {
					fmt.Println("     ", i, u)
				}
			default:
				fmt.Println(">>>>>field", k, "is unknown format")
			}
		}
	}

	if j, err := json.Marshal(message); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write(j)
	}
}
