package mplus

import (
	"github.com/tangzixiang/mplus/query"
)

type Query = query.Query
type DynamicQueryValue = query.DynamicQueryValue
type DynamicQueryPairs = query.DynamicQueryPairs

var (
	ParseQuery   = query.ParseQuery
	Queries      = query.Queries
	NewQuery     = query.New
	NewQueryWith = query.NewWith
)
