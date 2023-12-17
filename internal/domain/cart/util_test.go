package cart

import (
	"testing"

	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
)

func TestNewMoney(t *testing.T) {
	expect := money.NewFromFloat(10, money.BRL)
	actual := newMoney(10)

	isEquals, _ := actual.Equals(expect)

	assert.True(t, isEquals)
}
