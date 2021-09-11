// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package daodao

import (
	"net/url"
)

func (c *Client) DefaultForm() url.Values {
	data := map[string]string{
		"access_token":       c.accessToken,
		"app_channel":        "AppStore",
		"app_client_version": "1.1.7",
		"app_name":           "ddjz_lite",
		"app_real_name":      "ddjz_lite",
		"client_id":          "daodao_ios",
		"client_secret":      "daodao2018",
		"grant_type":         "access_token",
		"hardware":           "iPhone%252012",
		"height":             "2532",
		"is_new_user":        "0",
		"latitude":           "0",
		"longitude":          "0",
		"mp_code":            "86",
		"os_version":         "14.7.1",
		"platform":           "iOS",
		"redirect_uri":       "",
		"timezone":           "8",
		"version":            "v1",
		"width":              "1170",
	}

	u := url.Values{}
	for k, v := range data {
		u.Set(k, v)
	}
	return u
}
