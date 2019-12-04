package handler

import (
	"reflect"
	"testing"
)

func TestNewSample(t *testing.T) {
	tests := []struct {
		name string
		want *Sample
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSample(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSample() = %v, want %v", got, tt.want)
			}
		})
	}
}
