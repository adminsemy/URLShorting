package shorten

import (
	"context"
	"reflect"
	"testing"

	"github.com/adminsemy/URLShorting/internal/model"
)

func TestService_Shorten(t *testing.T) {
	type args struct {
		ctx   context.Context
		input model.ShortenInput
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		want    *model.Shortening
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			got, err := s.Shorten(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Shorten() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Shorten() = %v, want %v", got, tt.want)
			}
		})
	}
}
