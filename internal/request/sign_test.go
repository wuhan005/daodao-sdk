// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSign(t *testing.T) {
	for _, tc := range []struct {
		name  string
		nonce string
		time  string
		want  string
	}{
		{
			name:  "normal",
			nonce: "E99p1ant",
			time:  "1631370205310",
			want:  "c6a4fab63eec9f9aef76d6362f236948",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := Sign(tc.nonce, tc.time)
			assert.Equal(t, tc.want, got)
		})
	}
}
