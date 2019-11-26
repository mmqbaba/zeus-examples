package handler

import (
	"reflect"
	"testing"
)

func TestNewHelloDemo(t *testing.T) {
	tests := []struct {
		name string
		want *HelloDemo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHelloDemo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHelloDemo() = %v, want %v", got, tt.want)
			}
		})
	}
}
