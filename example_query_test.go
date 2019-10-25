package mplus

import (
	"fmt"
	"net/url"
)

func ExampleNewQuery() {
	urlValues := make(url.Values)
	urlValues.Add("name", "tom")
	urlValues.Add("age", "15")
	path := "http://localhost?" + urlValues.Encode()

	fmt.Printf("%v\n", path)
	fmt.Printf("%v\n", NewQuery().AddPairs("name", "tom", "age", "15").AppendToURI("http://localhost"))

	// OutPut:
	// http://localhost?age=15&name=tom
	// http://localhost?age=15&name=tom
}

func ExampleQuery_SetIf() {

	values := map[string]bool{
		"name": false, // will be continue
		"age":  true,
	}

	fmt.Printf("%v\n", NewQuery().
		SetIf(values["name"], "name", "tom").
		SetIf(values["age"], "age", "15").
		AppendToURIFormat("http://localhost/users/%v", 1))

	// OutPut:
	// http://localhost/users/1?age=15
}

func ExampleQuery_SetIfD() {

	values := map[string]interface{}{
		"name": nil, // will be continue
		"age":  "15",
	}

	query :=  NewQuery().
		SetIfD(values["name"] != nil, "name", func() string { return values["name"].(string) }). // will not panic because callback is lazy evaluation
		SetIfD(values["age"] != nil, "age", func() string { return values["age"].(string) })

	fmt.Printf("%v\n",query.AppendToURI("http://localhost/users/1"))

	// OutPut:
	// http://localhost/users/1?age=15
}

func ExampleQuery_SetPairsD() {

	query :=  NewQuery().SetPairsD(func() []string { // lazy dynamic
		values := map[string]int{
			"first":  1,
			"second": 2,
			"third":  3,
		}

		var pairs []string

		for key, value := range values {
			pairs = append(pairs, key, fmt.Sprint(value))
		}

		return pairs
	})

	fmt.Printf("%v\n",query.AppendToURI("http://localhost/users/1"))

	// OutPut:
	// http://localhost/users/1?first=1&second=2&third=3
}