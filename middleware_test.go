package httprouter

import (
	"time"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleware(t *testing.T) {
	a := 0
	mwf1 := func (next Handle) Handle {
		return func (w http.ResponseWriter, r *http.Request, ps Params) {
			a ++
			t.Log("second")
			next(w, r, ps)
		}
	}
	mwf2 := func (next Handle) Handle {
		return func (w http.ResponseWriter, r *http.Request, ps Params) {
			a ++
			t.Log("first")
			next(w, r, ps)
		}
	}
	r := New()
	r.GET("/add", func(w http.ResponseWriter, r *http.Request, ps Params) {
		t.Log("get /add\n")
	})

	r.Use(mwf1)
	r.Use(mwf2)

	req := httptest.NewRequest("GET", "/add", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	time.Sleep(time.Second * 1)

	if a < 1 {
		t.Fail()
	}
}
