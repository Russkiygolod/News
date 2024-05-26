package api

import (
	"News/pkg/storage/postgres"
	"testing"

	"github.com/gorilla/mux"
)

func TestAPI_endpoints(t *testing.T) {
	type fields struct {
		r  *mux.Router
		db postgres.Store
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				r:  tt.fields.r,
				db: tt.fields.db,
			}
			api.endpoints()
		})
	}
}
