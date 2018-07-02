package ginrest

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNew(t *testing.T) {
	type args struct {
		path   string
		object string
	}
	tests := []struct {
		name string
		args args
		want *IO
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.path, tt.args.object); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIO_SetGin(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		s    *IO
		args args
		want *IO
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetGin(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IO.SetGin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIO_Res(t *testing.T) {
	type args struct {
		code     int
		elements Payload
		msg      string
	}
	tests := []struct {
		name string
		s    *IO
		args args
		want *IO
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Res(tt.args.code, tt.args.elements, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IO.Res() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIO_httpCodes(t *testing.T) {
	type args struct {
		code int
	}
	tests := []struct {
		name string
		s    *IO
		args args
		want *IO
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.httpCodes(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IO.httpCodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
