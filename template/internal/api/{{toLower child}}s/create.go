// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower child}}s

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/{{github}}/internal/api/render"
	"github.com/{{github}}/internal/logger"
	"github.com/{{github}}/internal/store"
	"github.com/{{github}}/types"
	"github.com/go-chi/chi"
)

// HandleCreate returns an http.HandlerFunc that creates
// the object and persists to the datastore.
func HandleCreate({{toLower parent}}s store.{{parent}}Store, {{toLower child}}s store.{{child}}Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projectID, err := strconv.ParseInt(chi.URLParam(r, "project"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse project id")
			return
		}

		{{toLower parent}}ID, err := strconv.ParseInt(chi.URLParam(r, "{{toLower parent}}"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse {{toLower parent}} id")
			return
		}

		in := new(types.{{child}}Input)
		err = json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("project", projectID).
				WithField("{{toLower parent}}", {{toLower parent}}ID).
				Debugln("cannot unmarshal json request")
			return
		}

		{{toLower parent}}, err := {{toLower parent}}s.Find(r.Context(), {{toLower parent}}ID)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("id", {{toLower parent}}ID).
				Debugln("{{toLower parent}} not found")
			return
		}

		if {{toLower parent}}.Project != projectID {
			render.NotFoundf(w, "Not Found")
			logger.FromRequest(r).
				WithField("{{toLower parent}}", {{toLower parent}}ID).
				WithField("project", projectID).
				Debugln("project id mismatch")
			return
		}

		{{toLower child}} := &types.{{child}}{
			{{parent}}:     {{toLower parent}}.ID,
			Name:    in.Name.String,
			Desc:    in.Desc.String,
			Created: time.Now().Unix(),
			Updated: time.Now().Unix(),
		}

		err = {{toLower child}}s.Create(r.Context(), {{toLower child}})
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("{{toLower parent}}.id", {{toLower parent}}.ID).
				WithField("{{toLower parent}}.name", {{toLower parent}}.Name).
				Errorln("cannot create {{toLower child}}")
		} else {
			render.JSON(w, {{toLower child}}, 200)
		}
	}
}
