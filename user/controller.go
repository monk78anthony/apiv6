package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/monk78anthony/apiv6/domain"
	"github.com/monk78anthony/apiv6/package/storage"
)

type Controller struct {
	Storage storage.UserStorer
}

// POST /api/v1/users/create
func (c Controller) Create(w http.ResponseWriter, r *http.Request) {
	var req User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id := uuid.New().String()

	err := c.Storage.Insert(r.Context(), storage.User{
		UUID:      id,
		Name:      req.Name,
		Grade:     req.Grade,
		IsBlocked: req.IsBlocked,
		CreatedAt: req.CreatedAt,
		Roles:     req.Roles,
	})
	if err != nil {
		switch err {
		case domain.ErrConflict:
			w.WriteHeader(http.StatusConflict)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(id))
}

// GET /api/v1/users/find?id={UUID}
func (c Controller) Find(w http.ResponseWriter, r *http.Request) {
	res, err := c.Storage.Find(r.Context(), r.URL.Query().Get("id"))
	if err != nil {
		switch err {
		case domain.ErrNotFound:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	user := User{
		UUID:      res.UUID,
		Name:      res.Name,
		Grade:     res.Grade,
		IsBlocked: res.IsBlocked,
		CreatedAt: res.CreatedAt,
		Roles:     res.Roles,
	}

	data, err := json.Marshal(user)
	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write(data)
}

// DELETE /api/v1/users/delete?id={UUID}
func (c Controller) Delete(w http.ResponseWriter, r *http.Request) {
	err := c.Storage.Delete(r.Context(), r.URL.Query().Get("id"))
	if err != nil {
		switch err {
		case domain.ErrNotFound:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// PATCH /api/v1/users/update?id={UUID}
func (c Controller) Update(w http.ResponseWriter, r *http.Request) {
	var req User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := c.Storage.Update(r.Context(), storage.User{
		UUID:  r.URL.Query().Get("id"),
		Name:  req.Name,
		Grade: req.Grade,
		Roles: req.Roles,
	})
	if err != nil {
		switch err {
		case domain.ErrNotFound:
			w.WriteHeader(http.StatusConflict)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
