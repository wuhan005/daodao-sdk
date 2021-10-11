// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package daodao

import (
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/wuhan005/daodao-sdk/internal/request"
)

type LoginResponse struct {
	ServerTime int    `json:"server_time"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Data       struct {
		Code         string `json:"code"`
		Message      string `json:"message"`
		UserID       int    `json:"user_id"`
		ShortUuid    string `json:"short_uuid"`
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	} `json:"data"`
	Encrypt int `json:"encrypt"`
}

// Login authenticate the Daodao API account.
func (c *Client) Login(phone, password string) (*LoginResponse, error) {
	form := c.DefaultParameters()
	form.Set("address", phone)
	form.Set("password", password)

	resp, err := request.Post("/login").Form(form).Do()
	if err != nil {
		return nil, errors.Wrap(err, "request")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("unexpected response code: %d", resp.StatusCode)
	}

	var loginResponse LoginResponse
	if err := resp.ToJSON(&loginResponse); err != nil {
		return nil, errors.Wrap(err, "to json")
	}

	c.accessToken = loginResponse.Data.AccessToken
	c.userID = strconv.Itoa(loginResponse.Data.UserID)

	return &loginResponse, nil
}
