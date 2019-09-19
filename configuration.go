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
	"log"
	"os"
	"strings"
)

type Configuration struct {
	VerificationKey string
	WebhookRoute    string
	ListenPort      string
}

var (
	config = Configuration{}
)

func getEnvVar(key, defaultValue string) string {
	value := os.Getenv(key)
	if strings.TrimSpace(value) == "" {
		return defaultValue
	}
	return strings.TrimSpace(value)
}

func loadEnvironment() {
	log.Println("== Configuration ==")
	config.VerificationKey = getEnvVar("REVOPS_VERIFICATION_KEY", "")
	config.WebhookRoute = getEnvVar("REVOPS_WEBHOOK_ROUTE", "/")
	config.ListenPort = getEnvVar("REVOPS_LISTEN_PORT", "8080")

	if !strings.HasPrefix(config.ListenPort, ":") {
		config.ListenPort = ":" + config.ListenPort
	}

	if config.VerificationKey == "" {
		log.Println(" - WARNING: Verification Key has not been configured.  Webhook events will not be verified")
	} else {
		log.Println(" - Verification Key loaded")
	}
	log.Println(" - WebhookRoute: ", config.WebhookRoute)
	log.Println(" - ListenPort: ", config.ListenPort)
	log.Println("== End Configuration ==")
}
