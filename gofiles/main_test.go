package main

import (
	"testing"

	P "github.com/rtseuztz/ApartmentGenerator/gofiles/pages"
)

func tGetSummoner(t *testing.T) {
	result := P.GetSummoner("afk ff 15")
	if result.Name != "afk ff 15" {
		t.Error("name is incorrect")
	}
}
func TestAll(t *testing.T) {
	t.Run("a=0", tGetSummoner)
}
