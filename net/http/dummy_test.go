package http_test

import (
	"net/http"
)

type dotFileHidingFileSystem struct{ any }

func (fs dotFileHidingFileSystem) Open(name string) (http.File, error)

type countHandler struct{}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)

type apiHandler struct{}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)

func newPeopleHandler() http.Handler

type wantRange struct{ a, b int }

const testFileLen = 100
