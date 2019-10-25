package query

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

// ParseQuery 解析 URL 上的 query string
func ParseQuery(r *http.Request) (url.Values, error) {
	var newValues url.Values

	if r.URL == nil {
		return nil, errors.New("url is empty")
	}

	newValues = r.URL.Query()
	if newValues != nil {
		return newValues, nil
	}

	return url.ParseQuery(r.URL.RawQuery)
}

// Queries 获取 URL 上的请求字段,若不存在则返回空的 url.Value
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

// DynamicQueryValue 回调计算 value 值的 handler,可以用于延迟计算 value
type DynamicQueryValue func() string

// DynamicQueryPairs 回调计算 pair 对的 handler,可以用于延迟计算多对 key value
type DynamicQueryPairs func() []string

// New 获取一个 Query 实例
func New() *Query {
	return &Query{v: url.Values{}}
}

// NewWith 获取一个 Query 实例,同时初始化
func NewWith(v url.Values) *Query {
	return &Query{v: v}
}

// Set set the key to value. It replaces any existing
// values.
func (q *Query) Set(key, value string) *Query {
	q.v.Set(key, value)
	return q
}

// SetD set the key to value but dynamic. It replaces any existing
// values.
func (q *Query) SetD(key string, value DynamicQueryValue) *Query {
	q.v.Set(key, value())
	return q
}

// SetIf set the key to value if ensure is true. It replaces any existing
// values.
func (q *Query) SetIf(ensure bool, key, value string) *Query {
	if ensure {
		q.Set(key, value)
	}
	return q
}

// SetIfD set the key to value dynamically if ensure is true. It replaces any existing
// values.
func (q *Query) SetIfD(ensure bool, key string, value DynamicQueryValue) *Query {
	if ensure {
		q.Set(key, value())
	}
	return q
}

// SetPairs set the key to value. It replaces any existing
// values.
//  q.SetPairs(key1,value1,key2,value2,...)
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

// SetPairsD set the key to value but dynamic. It replaces any existing
// values.
//  q.SetPairsD(func ()[]string{return []string{key1,value1,key2,value2,...}})
func (q *Query) SetPairsD(pairs DynamicQueryPairs) *Query {
	_pairs := pairs()
	l := len(_pairs)

	if l <= 0 || l%2 != 0 {
		return q
	}

	for i := 0; i < l; i += 2 {
		q.v.Set(_pairs[i], _pairs[i+1])
	}

	return q
}

// SetPairsIf set the key to value if ensure is true. It replaces any existing
// values.
//  q.SetPairsIf(GetTrue(),key1,value1,key2,value2,...)
func (q *Query) SetPairsIf(ensure bool, key, value string, pairs ... string) *Query {
	if ensure {
		q.SetPairs(key, value, pairs...)
	}
	return q
}

// SetPairsIfD set the key to value dynamically if ensure is true. It replaces any existing
// values.
//  q.SetPairsIfD(GetTrue(),func ()[]string{return []string{key1,value1,key2,value2,...}})
func (q *Query) SetPairsIfD(ensure bool, pairs DynamicQueryPairs) *Query {
	if ensure {
		q.SetPairsD(pairs)
	}
	return q
}

// SetByM sets the key to value by map. It replaces any existing
// values.
//  q.SetByM(map[string]string{"key1":value1})
func (q *Query) SetByM(pairs map[string]string) *Query {

	if pairs == nil {
		return q
	}

	for key, value := range pairs {
		q.v.Set(key, value)
	}

	return q
}

// SetByMIf set the key to value by map if ensure is true. It replaces any existing
// values.
//  q.SetByMIf(GetTrue(),map[string]string{"key1":value1})
func (q *Query) SetByMIf(ensure bool, pairs map[string]string) *Query {
	if ensure {
		q.SetByM(pairs)
	}

	return q
}

// Get get the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (q *Query) Get(key string) string {
	return q.v.Get(key)
}

// Add add the value to key. It appends to any existing
// values associated with key.
func (q *Query) Add(key, value string) *Query {
	q.v.Add(key, value)
	return q
}

// AddD add the value to key but dynamic. It appends to any existing
// values associated with key.
func (q *Query) AddD(key string, value DynamicQueryValue) *Query {
	q.v.Add(key, value())
	return q
}

// AddIf add the value to key if ensure is true. It appends to any existing
// values associated with key.
func (q *Query) AddIf(ensure bool, key, value string) *Query {
	if ensure {
		q.Add(key, value)
	}
	return q
}

// AddIf add the value to key dynamically if ensure is true. It appends to any existing
// values associated with key.
func (q *Query) AddIfD(ensure bool, key string, value DynamicQueryValue) *Query {
	if ensure {
		q.Add(key, value())
	}
	return q
}

// AddPairs add the value to key. It appends to any existing
// values associated with key.
//  q.AddPairs(key1,value1,key2,value2,...)
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

// AddPairsD add the value to key but dynamic. It appends to any existing
// values associated with key.
//  q.AddPairsD(func()[]string{return []string{key1,value1,key2,value2,...}})
func (q *Query) AddPairsD(pairs DynamicQueryPairs) *Query {
	_pairs := pairs()
	l := len(_pairs)

	if l <= 0 || l%2 != 0 {
		return q
	}

	for i := 0; i < l; i += 2 {
		q.v.Add(_pairs[i], _pairs[i+1])
	}

	return q
}

// AddPairsIf add the value to key if ensure is true. It appends to any existing
// values associated with key.
//  q.AddPairsIf(GetTrue(),key1,value1,key2,value2,...)
func (q *Query) AddPairsIf(ensure bool, key, value string, pairs ... string) *Query {
	if ensure {
		q.AddPairs(key, value, pairs...)
	}

	return q
}

// AddPairsIfD add the value to key dynamically if ensure is true. It appends to any existing
// values associated with key.
//  q.AddPairsIfD(GetTrue(),func()[]string{return []string{key1,value1,key2,value2,...}})
func (q *Query) AddPairsIfD(ensure bool, pairs DynamicQueryPairs) *Query {
	if ensure {
		q.AddPairsD(pairs)
	}

	return q
}

// AddByM add the value to key by map. It appends to any existing
// values associated with key.
//  q.AddByM(map[string]string{"key1":value1})
func (q *Query) AddByM(pairs map[string]string) *Query {

	if pairs == nil {
		return q
	}

	for key, value := range pairs {
		q.v.Add(key, value)
	}

	return q
}

// AddByMIf add the value to key by map if ensure is true. It appends to any existing
// values associated with key.
//  q.AddByMIf(GetTrue(),map[string]string{"key1":value1})
func (q *Query) AddByMIf(ensure bool, pairs map[string]string) *Query {
	if ensure {
		q.AddByM(pairs)
	}

	return q
}

// Del delete the values associated with key.
func (q *Query) Del(key string) *Query {
	q.v.Del(key)
	return q
}

// DelIf delete the values associated with key if ensure is true.
func (q *Query) DelIf(ensure bool, key string) *Query {
	if ensure {
		q.Del(key)
	}
	return q
}

// Encode encode the values into ``URL encoded'' form
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

// WithIf 将 v 同步合并,当且 ensure 为 true
func (q *Query) WithIf(ensure bool, v url.Values) *Query {
	if ensure {
		q.With(v)
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

// AppendToURI 将请求字段追加至 URI 上
func (q *Query) AppendToURI(uri string) string {

	if strings.Index(uri, "?") != -1 {
		return uri + q.Encode()
	}

	return uri + "?" + q.Encode()
}

// AppendToURIf 将请求字段追加至 URI 上
func (q *Query) AppendToURIFormat(formatURI string, a ... interface{}) string {
	return q.AppendToURI(fmt.Sprintf(formatURI, a...))
}

// Values 获取内部持有的 Values
func (q *Query) Values() url.Values {
	return q.v
}

// Len 获取键值对数量
func (q *Query) Len() int {
	return len(q.v)
}
