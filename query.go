package mplus

import (
	"net/http"
	"net/url"
	"strings"
)

// Queries 获取 URL 上的请求字段
func Queries(r *http.Request) url.Values {

	values, err := ParseQuery(r)
	if err != nil {
		return make(url.Values)
	}

	return values
}

// Query url.Values 装载对象
type Query struct {
	v url.Values
}

// NewQuery 获取一个 Query 实例
func NewQuery() *Query {
	return &Query{v: url.Values{}}
}

// NewQueryWith 获取一个 Query 实例,同时初始化
func NewQueryWith(v url.Values) *Query {
	return &Query{v: v}
}

// Set sets the key to value. It replaces any existing
// values.
func (q *Query) Set(key, value string) *Query {
	q.v.Set(key, value)
	return q
}

// SetPairs sets the key to value. It replaces any existing
// values.
// For example:
//  - q.SetPairs(key1,value1,key2,value2,...)
func (q *Query) SetPairs(key, value string, pairs ... string) *Query {
	q.v.Set(key, value)

	if len(pairs) <= 0 || len(pairs)%2 != 0 {
		return q
	}

	for i := 0; i < len(pairs); i += 2 {
		q.v.Set(pairs[i], pairs[i+1])
	}

	return q
}

// SetByM sets the key to value. It replaces any existing
// values.
// For example:
//  - q.SetByM(map[string]string{"key1":value1})
func (q *Query) SetByM(pairs map[string]string) *Query {

	if pairs == nil {
		return q
	}

	for key, value := range pairs {
		q.v.Set(key, value)
	}

	return q
}

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (q *Query) Get(key string) string {
	return q.v.Get(key)
}

// Add adds the value to key. It appends to any existing
// values associated with key.
func (q *Query) Add(key, value string) *Query {
	q.v.Add(key, value)
	return q
}

// AddPairs adds the value to key. It appends to any existing
// values associated with key.
// For example:
//  - q.AddPairs(key1,value1,key2,value2,...)
func (q *Query) AddPairs(key, value string, pairs ... string) *Query {
	q.v.Add(key, value)
	if len(pairs) <= 0 || len(pairs)%2 != 0 {
		return q
	}

	for i := 0; i < len(pairs); i += 2 {
		q.v.Add(pairs[i], pairs[i+1])
	}

	return q
}

// AddByM adds the value to key. It appends to any existing
// values associated with key.
// For example:
//  - q.AddByM(map[string]string{"key1":value1})
func (q *Query) AddByM(pairs map[string]string) *Query {

	if pairs == nil {
		return q
	}

	for key, value := range pairs {
		q.v.Add(key, value)
	}

	return q
}

// Del deletes the values associated with key.
func (q *Query) Del(key string) *Query {
	q.v.Del(key)
	return q
}

// Encode encodes the values into ``URL encoded'' form
// ("bar=baz&foo=quux") sorted by key.
func (q *Query) Encode() string {
	return q.v.Encode()
}

// With 将 v 同步合并
func (q *Query) With(v url.Values) *Query {
	for key, valueArr := range v {
		for _, valueItem := range valueArr {
			if q.v.Get(key) == "" {
				q.v.Set(key, valueItem)
			} else {
				q.v.Add(key, valueItem)
			}
		}
	}

	return q
}

// ParseForm 解析指定 URL-encoded query string
// 底层调用的是 url.ParseQuery
func (q *Query) ParseForm(query string) *Query {
	values, err := url.ParseQuery(query)
	if err != nil {
		return q
	}

	return q.With(values)
}

// AppendTo 将请求字段追加至 URI 上
func (q *Query) AppendToURI(uri string) string {

	if strings.LastIndex(uri, "?") == 0 {
		return uri + q.Encode()
	}

	return uri + "?" + q.Encode()
}
