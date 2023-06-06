package main

import (
	"testing"

	"reflect"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	// Verificar el tipo de dato del campo Name

	var checkName = assert.IsType(t, "", InfoUser.Name)

	if checkName != true {
		t.Error("Name must be string, but got: ", reflect.TypeOf(InfoUser.Name))
		return
	}

	// Verificar el tipo de dato del campo Email
	var chekEmail = assert.IsType(t, "", InfoUser.Email)
	if chekEmail != true {
		t.Error("Name must be string, but got: ", reflect.TypeOf(InfoUser.Email))
		return
	}

	// Verificar el tipo de dato del campo Phone
	var chekPhone = assert.IsType(t, 0, InfoUser.Phone)

	if chekPhone != true {
		t.Error("Name must be int, but got: ", reflect.TypeOf(InfoUser.Phone))
		return
	}
}
