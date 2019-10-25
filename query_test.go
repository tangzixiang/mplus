package mplus

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestQuery_Encode(t *testing.T) {

	urlValues := make(url.Values)
	urlValues.Add("name", "tom")
	urlValues.Add("age", "15")
	queryString := urlValues.Encode()

	assert.Equal(t, queryString, NewQuery().Add("name", "tom").Add("age", "15").Encode())
	assert.Equal(t, queryString, NewQuery().AddPairs("name", "tom", "age", "15").Encode())
	assert.Equal(t, queryString, NewQuery().AddByM(map[string]string{"name": "tom", "age": "15"}).Encode())

}

func TestQuery_AppendToURI(t *testing.T) {

	urlValues := make(url.Values)
	urlValues.Add("name", "tom")
	urlValues.Add("age", "15")
	path := "http://localhost?" + urlValues.Encode()

	assert.Equal(t, path, NewQuery().Add("name", "tom").Add("age", "15").AppendToURI("http://localhost?"))
	assert.Equal(t, path, NewQuery().AddPairs("name", "tom", "age", "15").AppendToURI("http://localhost"))
}

func TestQuery_AppendToURIFormat(t *testing.T) {
	urlValues := make(url.Values)
	urlValues.Add("name", "tom")
	urlValues.Add("age", "15")
	path := "http://localhost/users/1?" + urlValues.Encode()

	assert.Equal(t, path, NewQuery().Add("name", "tom").Add("age", "15").AppendToURIFormat("http://localhost/users/%v?", 1))
	assert.Equal(t, path, NewQuery().AddPairs("name", "tom", "age", "15").AppendToURIFormat("http://localhost/users/%v", 1))
}

func TestParseQuery(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://127.0.0.1?name=tom&age=15", nil)
	values, err := ParseQuery(r)

	assert.Nil(t, err)
	assert.Equal(t, "tom", values.Get("name"))
	assert.Equal(t, "15", values.Get("age"))
	assert.Len(t, values, 2)

	r.URL = nil
	values, err = ParseQuery(r)

	assert.NotNil(t, err)
	assert.Nil(t, values)
}

func TestQueries(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://127.0.0.1?name=tom&age=15", nil)
	values := Queries(r)

	assert.Equal(t, "tom", values.Get("name"))
	assert.Equal(t, "15", values.Get("age"))
	assert.Len(t, values, 2)

	r.URL = nil
	values = Queries(r)

	assert.NotNil(t, values)
	assert.Len(t, values, 0)
}

func TestNewWith(t *testing.T) {
	urlValues := make(url.Values)

	m := map[string]string{
		"name": "tom", "age": "15",
	}

	for key, value := range m {
		urlValues.Set(key, value)
	}

	q := NewQueryWith(urlValues)

	for key, value := range m {
		assert.Equal(t, value, q.Get(key))
	}
}

func TestQuery_Set(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		args args
	}{
		{args: args{key: "name", value: "tom"}},
		{args: args{key: "age", value: "15"}},
	}

	for _, tt := range tests {
		q := NewQuery()
		t.Run("test set "+tt.args.key, func(t *testing.T) {
			q.Set(tt.args.key, tt.args.value)

			assert.Equal(t, tt.args.value, q.Get(tt.args.key))
		})
	}
}

func TestQuery_SetD(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		args args
	}{
		{args: args{key: "name", value: "tom"}},
		{args: args{key: "age", value: "15"}},
	}

	for _, tt := range tests {
		q := NewQuery()

		t.Run("test set dynamically"+tt.args.key, func(t *testing.T) {
			q.SetD(tt.args.key, func() string {
				return tt.args.value + "D"
			})

			assert.Equal(t, tt.args.value+"D", q.Get(tt.args.key))
		})
	}
}

func TestQuery_SetIf(t *testing.T) {
	values := map[string]interface{}{
		"name": nil, // will be continue
		"age":  "15",
	}

	urlValues := make(url.Values)
	for key, value := range values {
		if value != nil {
			urlValues.Set(key, value.(string))
		}
	}

	path := "http://localhost/users/1?" + urlValues.Encode()
	generatePath := NewQuery().
		SetIf(values["name"] != nil, "name", "tom").
		SetIf(values["age"] != nil, "age", "15").
		AppendToURIFormat("http://localhost/users/%v?", 1)

	assert.Equal(t, path, generatePath)
}

func TestQuery_SetIfD(t *testing.T) {
	values := map[string]interface{}{
		"name": nil, // will be continue
		"age":  "15",
	}

	urlValues := make(url.Values)
	for key, value := range values {
		if value != nil {
			urlValues.Set(key, value.(string))
		}
	}

	path := "http://localhost/users/1?" + urlValues.Encode()
	generatePath := NewQuery().
		SetIfD(values["name"] != nil, "name", func() string { return values["name"].(string) }). // will not panic because callback is lazy evaluation
		SetIfD(values["age"] != nil, "age", func() string { return values["age"].(string) }).
		AppendToURIFormat("http://localhost/users/%v?", 1)

	assert.Equal(t, path, generatePath)
}

func TestQuery_SetPairs(t *testing.T) {

	type args struct {
		key   string
		value string
		pairs []string
	}
	tests := []struct {
		args args
	}{
		{args: args{key: "name", value: "tom", pairs: []string{"age", "15"}}},
		{args: args{key: "name", value: "lili", pairs: []string{"age", "16"}}},
	}

	for _, tt := range tests {
		q := NewQuery()

		t.Run("test SetPairs", func(t *testing.T) {
			q.SetPairs(tt.args.key, tt.args.value, tt.args.pairs...)

			assert.Equal(t, tt.args.value, q.Get(tt.args.key))
			assert.Equal(t, tt.args.pairs[1], q.Get(tt.args.pairs[0]))
		})
	}
}

func TestQuery_SetPairsD(t *testing.T) {

	values := map[string]int{
		"first":  1,
		"second": 2,
		"third":  3,
	}

	urlValues := make(url.Values)

	for key, value := range values {
		urlValues.Set(key, fmt.Sprint(value))
	}

	path := "http://localhost/users/1?" + urlValues.Encode()
	generatePath := NewQuery().SetPairsD(func() []string { // lazy dynamic
		var pairs []string

		for key, value := range values {
			pairs = append(pairs, key, fmt.Sprint(value))
		}

		return pairs
	}).AppendToURIFormat("http://localhost/users/%v?", 1)

	assert.Equal(t, path, generatePath)
}

func TestQuery_SetPairsIf(t *testing.T) {

	type args struct {
		ensure bool
		key    string
		value  string
		pairs  []string
	}
	tests := []struct {
		args args
	}{
		{args: args{key: "name", value: "tom", pairs: []string{"age", "15"}, ensure: true}},
		{args: args{key: "name", value: "lili", pairs: []string{"age", "16"}, ensure: false}},
	}

	for _, tt := range tests {
		q := NewQuery()

		t.Run("test SetPairsIf", func(t *testing.T) {
			q.SetPairsIf(tt.args.ensure, tt.args.key, tt.args.value, tt.args.pairs...)

			if tt.args.ensure {
				assert.Equal(t, tt.args.value, q.Get(tt.args.key))
				assert.Equal(t, tt.args.pairs[1], q.Get(tt.args.pairs[0]))
			} else {
				assert.NotEqual(t, tt.args.value, q.Get(tt.args.key))
				assert.NotEqual(t, tt.args.pairs[1], q.Get(tt.args.pairs[0]))
			}
		})
	}
}

func TestQuery_SetPairsIfD(t *testing.T) {
	type args struct {
		ensure bool
		pairs  []string
	}
	tests := []struct {
		args args
	}{
		{args: args{pairs: []string{"age", "15"}, ensure: true}},
		{args: args{pairs: []string{"age", "16"}, ensure: false}},
	}

	for _, tt := range tests {
		q := NewQuery()

		t.Run("test SetPairsIf", func(t *testing.T) {
			q.SetPairsIfD(tt.args.ensure, func() []string {
				return append([]string{"name", "tom"}, tt.args.pairs...)
			})

			if tt.args.ensure {
				assert.Equal(t, "tom", q.Get("name"))
				assert.Equal(t, tt.args.pairs[1], q.Get(tt.args.pairs[0]))
			} else {
				assert.NotEqual(t, "tom", q.Get("name"))
				assert.NotEqual(t, tt.args.pairs[1], q.Get(tt.args.pairs[0]))
			}
		})
	}
}

func TestQuery_SetByM(t *testing.T) {
	type args struct {
		pairs map[string]string
	}
	tests := []struct {
		args args
	}{
		{args: args{pairs: map[string]string{"name": "tom", "age": "15"}}},
		{args: args{pairs: map[string]string{"name": "lili", "age": "15"}}},
	}
	for _, tt := range tests {
		q := NewQuery()
		t.Run("test SetByM", func(t *testing.T) {
			q.SetByM(tt.args.pairs)
			for key, value := range tt.args.pairs {
				assert.Equal(t, value, q.Get(key))
			}
		})
	}
}

func TestQuery_SetByMIf(t *testing.T) {

	type args struct {
		ensure bool
		pairs  map[string]string
	}
	tests := []struct {
		args args
	}{
		{args: args{pairs: map[string]string{"name": "tom", "age": "15"}, ensure: true}},
		{args: args{pairs: map[string]string{"name": "lili", "age": "15"}, ensure: false}},
	}
	for _, tt := range tests {
		q := NewQuery()
		t.Run("test SetByM", func(t *testing.T) {
			q.SetByMIf(tt.args.ensure, tt.args.pairs)
			if tt.args.ensure {
				for key, value := range tt.args.pairs {
					assert.Equal(t, value, q.Get(key))
				}
			} else {
				for key, value := range tt.args.pairs {
					assert.NotEqual(t, value, q.Get(key))
				}
			}
		})
	}
}

func TestQuery_Add(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		args args
	}{
		{args: args{key: "name", value: "tom"}},
		{args: args{key: "age", value: "15"}},
	}

	for _, tt := range tests {
		q := NewQuery()
		t.Run("test add "+tt.args.key, func(t *testing.T) {
			q.Add(tt.args.key, tt.args.value)

			assert.Equal(t, tt.args.value, q.Get(tt.args.key))
		})
	}
}

func TestQuery_AddD(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		args args
	}{
		{args: args{key: "name", value: "tom"}},
		{args: args{key: "age", value: "15"}},
	}

	for _, tt := range tests {
		q := NewQuery()

		t.Run("test add dynamically"+tt.args.key, func(t *testing.T) {
			q.AddD(tt.args.key, func() string {
				return tt.args.value + "D"
			})

			assert.Equal(t, tt.args.value+"D", q.Get(tt.args.key))
		})
	}
}

func TestQuery_AddIf(t *testing.T) {
	values := map[string]interface{}{
		"name": nil, // will be continue
		"age":  "15",
	}

	urlValues := make(url.Values)
	for key, value := range values {
		if value != nil {
			urlValues.Set(key, value.(string))
		}
	}

	path := "http://localhost/users/1?" + urlValues.Encode()
	generatePath := NewQuery().
		AddIf(values["name"] != nil, "name", "tom").
		AddIf(values["age"] != nil, "age", "15").
		AppendToURIFormat("http://localhost/users/%v?", 1)

	assert.Equal(t, path, generatePath)
}

func TestQuery_AddIfD(t *testing.T) {
	values := map[string]interface{}{
		"name": nil, // will be continue
		"age":  "15",
	}

	urlValues := make(url.Values)
	for key, value := range values {
		if value != nil {
			urlValues.Set(key, value.(string))
		}
	}

	path := "http://localhost/users/1?" + urlValues.Encode()
	generatePath := NewQuery().
		AddIfD(values["name"] != nil, "name", func() string { return values["name"].(string) }). // will not panic because callback is lazy evaluation
		AddIfD(values["age"] != nil, "age", func() string { return values["age"].(string) }).
		AppendToURIFormat("http://localhost/users/%v?", 1)

	assert.Equal(t, path, generatePath)
}

func TestQuery_AddPairs(t *testing.T) {
	type args struct {
		key   string
		value string
		pairs []string
	}
	tests := []struct {
		args args
	}{
		{args: args{key: "name", value: "tom", pairs: []string{"age", "15"}}},
		{args: args{key: "name", value: "lili", pairs: []string{"age", "16"}}},
	}

	for _, tt := range tests {
		q := NewQuery()

		t.Run("test AddPairs", func(t *testing.T) {
			q.AddPairs(tt.args.key, tt.args.value, tt.args.pairs...)

			assert.Equal(t, tt.args.value, q.Get(tt.args.key))
			assert.Equal(t, tt.args.pairs[1], q.Get(tt.args.pairs[0]))
		})
	}
}

func TestQuery_AddPairsD(t *testing.T) {
	values := map[string]int{
		"first":  1,
		"second": 2,
		"third":  3,
	}

	urlValues := make(url.Values)

	for key, value := range values {
		urlValues.Set(key, fmt.Sprint(value))
	}

	path := "http://localhost/users/1?" + urlValues.Encode()
	generatePath := NewQuery().AddPairsD(func() []string { // lazy dynamic
		var pairs []string

		for key, value := range values {
			pairs = append(pairs, key, fmt.Sprint(value))
		}

		return pairs
	}).AppendToURIFormat("http://localhost/users/%v?", 1)

	assert.Equal(t, path, generatePath)
}

func TestQuery_AddPairsIf(t *testing.T) {
	type args struct {
		ensure bool
		key    string
		value  string
		pairs  []string
	}
	tests := []struct {
		args args
	}{
		{args: args{key: "name", value: "tom", pairs: []string{"age", "15"}, ensure: true}},
		{args: args{key: "name", value: "lili", pairs: []string{"age", "16"}, ensure: false}},
	}

	for _, tt := range tests {
		q := NewQuery()

		t.Run("test AddPairsIf", func(t *testing.T) {
			q.AddPairsIf(tt.args.ensure, tt.args.key, tt.args.value, tt.args.pairs...)

			if tt.args.ensure {
				assert.Equal(t, tt.args.value, q.Get(tt.args.key))
				assert.Equal(t, tt.args.pairs[1], q.Get(tt.args.pairs[0]))
			} else {
				assert.NotEqual(t, tt.args.value, q.Get(tt.args.key))
				assert.NotEqual(t, tt.args.pairs[1], q.Get(tt.args.pairs[0]))
			}
		})
	}
}

func TestQuery_AddPairsIfD(t *testing.T) {
	type args struct {
		ensure bool
		pairs  []string
	}
	tests := []struct {
		args args
	}{
		{args: args{pairs: []string{"age", "15"}, ensure: true}},
		{args: args{pairs: []string{"age", "16"}, ensure: false}},
	}

	for _, tt := range tests {
		q := NewQuery()

		t.Run("test AddPairsIfD", func(t *testing.T) {
			q.AddPairsIfD(tt.args.ensure, func() []string {
				return append([]string{"name", "tom"}, tt.args.pairs...)
			})

			if tt.args.ensure {
				assert.Equal(t, "tom", q.Get("name"))
				assert.Equal(t, tt.args.pairs[1], q.Get(tt.args.pairs[0]))
			} else {
				assert.NotEqual(t, "tom", q.Get("name"))
				assert.NotEqual(t, tt.args.pairs[1], q.Get(tt.args.pairs[0]))
			}
		})
	}
}

func TestQuery_AddByM(t *testing.T) {
	type args struct {
		pairs map[string]string
	}
	tests := []struct {
		args args
	}{
		{args: args{pairs: map[string]string{"name": "tom", "age": "15"}}},
		{args: args{pairs: map[string]string{"name": "lili", "age": "15"}}},
	}
	for _, tt := range tests {
		q := NewQuery()
		t.Run("test AddByM", func(t *testing.T) {
			q.AddByM(tt.args.pairs)
			for key, value := range tt.args.pairs {
				assert.Equal(t, value, q.Get(key))
			}
		})
	}
}

func TestQuery_AddByMIf(t *testing.T) {
	type args struct {
		ensure bool
		pairs  map[string]string
	}
	tests := []struct {
		args args
	}{
		{args: args{pairs: map[string]string{"name": "tom", "age": "15"}, ensure: true}},
		{args: args{pairs: map[string]string{"name": "lili", "age": "15"}, ensure: false}},
	}
	for _, tt := range tests {
		q := NewQuery()
		t.Run("test AddByMIf", func(t *testing.T) {
			q.AddByMIf(tt.args.ensure, tt.args.pairs)
			if tt.args.ensure {
				for key, value := range tt.args.pairs {
					assert.Equal(t, value, q.Get(key))
				}
			} else {
				for key, value := range tt.args.pairs {
					assert.NotEqual(t, value, q.Get(key))
				}
			}
		})
	}
}

func TestQuery_Del(t *testing.T) {
	urlValues := make(url.Values)

	m := map[string]string{
		"name": "tom", "age": "15",
	}

	for key, value := range m {
		urlValues.Set(key, value)
	}

	q := NewQueryWith(urlValues)

	for key, value := range m {
		assert.Equal(t, value, q.Get(key))
		q.Del(key)
	}

	for key := range m {
		assert.Equal(t, "", q.Get(key))
	}
}

func TestQuery_DelIf(t *testing.T) {
	urlValues := make(url.Values)

	m := map[string]string{
		"name": "tom", "age": "15",
	}

	for key, value := range m {
		urlValues.Set(key, value)
	}

	q := NewQueryWith(urlValues)

	for key, value := range m {
		assert.Equal(t, value, q.Get(key))
		q.DelIf(false, key)
	}

	for key, value := range m {
		assert.Equal(t, value, q.Get(key))
	}
}

func TestQuery_With(t *testing.T) {
	urlValues := make(url.Values)

	m := map[string]string{
		"name": "tom", "age": "15",
	}

	for key, value := range m {
		urlValues.Set(key, value)
	}

	q := NewQuery().With(urlValues)

	for key, value := range m {
		assert.Equal(t, value, q.Get(key))
	}
}

func TestQuery_WithIf(t *testing.T) {
	urlValues := make(url.Values)

	m := map[string]string{
		"name": "tom", "age": "15",
	}

	for key, value := range m {
		urlValues.Set(key, value)
	}

	q := NewQuery().WithIf(false, urlValues)

	for key := range m {
		assert.Equal(t, "", q.Get(key))
	}
}

func TestQuery_Len(t *testing.T) {
	urlValues := make(url.Values)

	m := map[string]string{
		"name": "tom", "age": "15",
	}

	for key, value := range m {
		urlValues.Set(key, value)
	}

	assert.Equal(t, 2, NewQuery().WithIf(true, urlValues).Len())
}

func TestQuery_ParseForm(t *testing.T) {
	m := map[string]string{
		"name": "tom", "age": "15",
	}

	q := NewQuery().ParseForm("name=tom&age=15")

	for key, value := range m {
		assert.Equal(t, value, q.Get(key))
	}
}

func TestQuery_Values(t *testing.T) {
	m := map[string]string{
		"name": "tom", "age": "15",
	}
	urlValues := make(url.Values)
	for key, value := range m {
		urlValues.Set(key, value)
	}

	values := NewQuery().With(urlValues).Values()

	for key, value := range m {
		assert.Equal(t, value, values.Get(key))
	}
}
