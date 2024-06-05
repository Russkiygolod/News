// Пакет для работы с RSS-потоками.
package rss

import (
	posts "News/pkg/model"
	"reflect"
	"testing"
)

func TestParseRss(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    []posts.Posts
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRss(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseRss() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseRss() = %v, want %v", got, tt.want)
			}
		})
	}
}
