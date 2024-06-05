package api

import (
	"News/internal"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
)

func TestAPI_posts(t *testing.T) {
	type fields struct {
		r *mux.Router
		i internal.Inter
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				r: tt.fields.r,
				i: tt.fields.i,
			}
			api.posts(tt.args.w, tt.args.r)
		})
	}
}
