// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package user

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/{{github}}/internal/api/request"
	"github.com/{{github}}/types"

	"github.com/google/go-cmp/cmp"
)

func TestFind(t *testing.T) {
	mockUser := &types.User{
		ID:    1,
		Email: "octocat@github.com",
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/user", nil)
	r = r.WithContext(
		request.WithUser(r.Context(), mockUser),
	)

	HandleFind()(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := &types.User{}, mockUser
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}
