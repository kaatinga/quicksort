package quicksort

import (
	"testing"
)

func Test_newArray(t *testing.T) {
	tests := []struct {
		name    string
		max     byte
		wantErr bool
	}{
		//{"!ok", 0, true},
		{"ok", 30, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numbers, err := newArray(tt.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("newArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			sort(&numbers)
		})
	}
}
