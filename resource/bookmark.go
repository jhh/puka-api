package resource

import (
	"errors"
	"net/http"

	"github.com/jhh/puka/model"
	"github.com/jhh/puka/storage"
	"github.com/manyminds/api2go"
)

const typeErrMsg = "Invalid instance given"

// The BookmarkResource struct implements api2go routes
type BookmarkResource struct {
	BookmarkStorage storage.BookmarkStorage
}

// FindAll satisfies api2go.FindAll interface
func (s BookmarkResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	bookmarks, err := s.BookmarkStorage.GetAll(storage.NewQuery(r))
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusInternalServerError)
	}
	return &Response{Res: bookmarks, Code: http.StatusOK}, nil
}

// FindOne satisfies api2go.CRUD interface
func (s BookmarkResource) FindOne(id string, r api2go.Request) (api2go.Responder, error) {
	bookmark, err := s.BookmarkStorage.GetOne(id)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}
	return &Response{Res: bookmark, Code: http.StatusOK}, nil
}

// PaginatedFindAll satisfies teh api2go.PaginatedFindAll interface
func PaginatedFindAll(req api2go.Request) (uint, api2go.Responder, error) {
	// total count, response, error
	return 0, &Response{}, nil
}

// Create satisfies api2go.CRUD interface
func (s BookmarkResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	bookmark, ok := obj.(model.Bookmark)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New(typeErrMsg), typeErrMsg, http.StatusBadRequest)
	}

	id, err := s.BookmarkStorage.Insert(bookmark)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusInternalServerError)
	}
	bookmark.ID = id

	return &Response{Res: bookmark, Code: http.StatusCreated}, nil
}

// Delete satisfies api2go.CRUD interface
func (s BookmarkResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	if err := s.BookmarkStorage.Delete(id); err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}
	return &Response{Code: http.StatusNoContent}, nil
}

// Update satisfies api2go.CRUD interface
func (s BookmarkResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	bookmark, ok := obj.(model.Bookmark)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New(typeErrMsg), typeErrMsg, http.StatusBadRequest)
	}

	if err := s.BookmarkStorage.Update(bookmark); err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusInternalServerError)
	}

	return &Response{Res: bookmark, Code: http.StatusOK}, nil
}
