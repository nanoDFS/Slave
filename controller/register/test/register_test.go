package test

import (
	"testing"

	"github.com/nanoDFS/Slave/controller/register"
)

func TestRegister(t *testing.T) {
	registerer := register.NewRegister(":9800", ":8000")
	if err := registerer.Register(); err != nil {
		t.Errorf("%v", err)
	}
}
