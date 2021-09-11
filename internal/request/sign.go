// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package request

import (
	"github.com/wuhan005/gadget"
)

// appSecret is the Daodao application secret key.
const appSecret = "853a0bb675aa143e6fa2dc607d55a9bb"

// Sign signs the requests and return the parameter `sign` for request.
func Sign(nonce, timestamp string) string {
	return gadget.Md5(appSecret + nonce + timestamp)
}
