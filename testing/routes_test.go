package testing

import (
	"api/routes"
	"testing"
	
)

func TestEmptyData(t *testing.T) {
	empty := routes.Hello("")
	if empty != "bye" {
		t.Error("empty value")
	}

}
func TestNotEmptyData(t *testing.T) {
	data := routes.Hello("hhye")
	if data != "hello" {
		t.Error("invalid input")
	}
}
