// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package request

import (
	"encoding/json"
	"io"
	"net/http"
)

// Response wraps the *http.Response and provides some helper methods.
type Response struct {
	*http.Response
}

// ToJSON binds the response body to given data structure.
func (r *Response) ToJSON(v interface{}) error {
	defer func() { _ = r.Body.Close() }()
	return json.NewDecoder(r.Body).Decode(&v)
}

// ToString reads the response body and return it as string.
func (r *Response) ToString() string {
	defer func() { _ = r.Body.Close() }()
	bodyBytes, _ := io.ReadAll(r.Body)
	return string(bodyBytes)
}
