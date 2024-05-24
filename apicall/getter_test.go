package apicall

import (
	"testing"
)

func TestNewHttpGetter(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHttpGetter()
			if got == nil {
				t.Errorf("NewHttpGetter() = %v, want not nil", got)
			}
			if got.(*httpGetter).Client == nil {
				t.Errorf("NewHttpGetter() = %v, client is nil", got)
			}
		})
	}
}
