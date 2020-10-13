package tests

import (
	"github.com/godvlpr/template/server/http"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	srv := httptest.NewServer()
	defer srv.Close()
}