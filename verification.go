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
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
)

const (
	RevOpsContentHMACHeader = "X-Revops-Content-Hmac"
)

func VerifyContentHMAC(r *http.Request, buf []byte) error {
	if config.VerificationKey == "" {
		log.Println(" WARNING: Verification Key is not configured.  Webhook event NOT verified")
		return nil
	}
	revopsHash := r.Header.Get(RevOpsContentHMACHeader)
	if revopsHash == "" {
		return fmt.Errorf("Missing RevOps HTTP Header: X-Revops-Content-Hmac")
	}
	// sha256 HMAC message verification
	h := hmac.New(sha256.New, []byte(config.VerificationKey))
	h.Write(buf)
	hash := hex.EncodeToString(h.Sum(nil))
	if hash != revopsHash {
		return fmt.Errorf("Invalid Content HMAC: %v (Expected: %v)", hash, revopsHash)
	}
	return nil
}
