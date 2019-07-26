// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package {{toLower child}}s

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/{{github}}/internal/api/render"
	"github.com/{{github}}/internal/logger"
	"github.com/{{github}}/internal/store"
	"github.com/{{github}}/types"

	"github.com/go-chi/chi"
)

// HandleUpdate returns an http.HandlerFunc that processes http
// requests to update the object details.
func HandleUpdate({{toLower parent}}s store.{{parent}}Store, {{toLower child}}s store.{{child}}Store) http.HandlerFunc {
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

		{{toLower child}}ID, err := strconv.ParseInt(chi.URLParam(r, "{{toLower child}}"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse {{toLower child}} id")
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
				WithField("{{toLower child}}", {{toLower child}}ID).
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

		{{toLower child}}, err := {{toLower child}}s.Find(r.Context(), {{toLower child}}ID)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("id", {{toLower child}}ID).
				Debugln("{{toLower parent}} not found")
			return
		}

		if {{toLower parent}}.Project != projectID {
			render.NotFoundf(w, "Not Found")
			logger.FromRequest(r).
				WithField("{{toLower parent}}", {{toLower parent}}ID).
				WithField("{{toLower child}}", {{toLower child}}ID).
				WithField("project", projectID).
				Debugln("project id mismatch")
			return
		}

		if {{toLower parent}}.ID != {{toLower child}}.{{parent}} {
			render.NotFoundf(w, "Not Found")
			logger.FromRequest(r).
				WithField("{{toLower parent}}.id", {{toLower parent}}.ID).
				WithField("{{toLower child}}.id", {{toLower child}}.ID).
				WithField("project", projectID).
				Debugln("{{toLower parent}} id mismatch")
			return
		}

		if in.Name.IsZero() == false {
			{{toLower child}}.Name = in.Name.String
		}
		if in.Desc.IsZero() == false {
			{{toLower child}}.Desc = in.Desc.String
		}

		err = {{toLower child}}s.Update(r.Context(), {{toLower child}})
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("{{toLower child}}.name", {{toLower child}}.Name).
				WithField("{{toLower child}}.id", {{toLower child}}.ID).
				WithField("{{toLower parent}}.id", {{toLower parent}}.ID).
				WithField("{{toLower parent}}.name", {{toLower parent}}.Name).
				Errorln("cannot update {{toLower parent}}")
		} else {
			render.JSON(w, {{toLower child}}, 200)
		}
	}
}
