package parseurl

import (
	posts "News/pkg/model"
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	type args struct {
		addres string
	}
	tests := []struct {
		name       string
		args       args
		wantUrls   []string
		wantPerion int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUrls, gotPerion := Read(tt.args.addres)
			if !reflect.DeepEqual(gotUrls, tt.wantUrls) {
				t.Errorf("Read() gotUrls = %v, want %v", gotUrls, tt.wantUrls)
			}
			if gotPerion != tt.wantPerion {
				t.Errorf("Read() gotPerion = %v, want %v", gotPerion, tt.wantPerion)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		url    string
		posts  chan<- []posts.Posts
		errs   chan<- error
		period int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Parse(tt.args.url, tt.args.posts, tt.args.errs, tt.args.period)
		})
	}
}
