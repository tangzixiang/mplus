package mplus

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeForm(t *testing.T) {
	values := parseForm()
	user := User{}

	DecodeForm(&user, values)

	assert.Equal(t, user.Name, values.Get("Name"))
	assert.Equal(t, fmt.Sprint(user.Age), values.Get("Age"))
	assert.Equal(t, user.Gender, values.Get("Gender"))
	assert.Equal(t, user.Email, values.Get("Email"))
}

func TestDecodeForms(t *testing.T) {
	values01 := parseForm()
	values02 := parseForm()
	values02.Set("Name", "New name")
	user := User{}

	DecodeForm(&user, values01, values02)

	assert.Equal(t, user.Name, values02.Get("Name"))
	assert.NotEqual(t, user.Name, values01.Get("Name"))

	assert.Equal(t, fmt.Sprint(user.Age), values01.Get("Age"))
	assert.Equal(t, fmt.Sprint(user.Age), values02.Get("Age"))

	assert.Equal(t, user.Gender, values01.Get("Gender"))
	assert.Equal(t, user.Gender, values02.Get("Gender"))

	assert.Equal(t, user.Email, values01.Get("Email"))
	assert.Equal(t, user.Email, values02.Get("Email"))
}

// this simulates the results of http.Request's ParseForm() function
func parseForm() url.Values {
	return url.Values{
		"Name":   []string{"  joeybloggs  "},
		"Age":    []string{"3"},
		"Gender": []string{"Male"},
		"Email":  []string{"Dean.Karn@gmail.com  "},
	}
}
