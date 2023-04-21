package logdis

import (
	"testing"
)

func Test_String(t *testing.T) {
	tests := []struct {
		name       Level
		reqmessage string
	}{
		{
			name:       InfoLevel,
			reqmessage: "info",
		},
	}

	for _, tt := range tests {
		if tt.reqmessage != tt.name.String() {
			t.Error("err")
		}
	}
}
