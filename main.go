// Example RevOps webhook event server written in Golang
//
// The full list of events that a webhook can receive are available at:
//   https://www.revops.io/docs/webhooks

/**
 * The MIT License
 * 
 * Copyright (c) 2019 RevOps, Inc https://www.revops.io
 * 
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 * 
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 * 
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type WebhookResponse struct {
	HttpStatus int    `json:"http_status"`
	Message    string `json:"message,omitempty"`
	Error      string `json:"error,omitempty"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	// prepare webhook response
	resp := &WebhookResponse{}

	defer func() {
		obj, _ := json.Marshal(resp)
		w.WriteHeader(resp.HttpStatus)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(obj))
	}()

	body, err := ioutil.ReadAll(r.Body)
	log.Printf("Processing request: %v", string(body))

	if err != nil {
		log.Printf("Error reading body: %v", err)
		resp.Error = fmt.Sprintf("%v", err)
		resp.HttpStatus = http.StatusBadRequest
		return
	}
	if err = VerifyContentHMAC(r, body); err != nil {
		resp.Error = fmt.Sprintf("%v", err)
		resp.HttpStatus = http.StatusUnauthorized
		return
	}

	log.Printf("Message verified!")
	resp.HttpStatus = http.StatusOK
	resp.Message = "Message verified!"
}

func main() {
	log.Println("Starting revops-webhook-server-go")
	loadEnvironment()

	http.HandleFunc(config.WebhookRoute, handler)
	log.Fatal(http.ListenAndServe(config.ListenPort, nil))
}
