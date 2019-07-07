package mplus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery_Encode(t *testing.T) {

	queryString := NewQuery().Add("name","tom").Add("age","15").Encode()

	if !assert.Equal(t,"age=15&name=tom",queryString){
		return
	}
}
