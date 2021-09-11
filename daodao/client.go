// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package daodao

// Client is the Daodao API request client.
type Client struct {
	accessToken string
	userID      string
}

// NewClient creates and return a new Daodao API client.
func NewClient() *Client {
	return &Client{}
}
