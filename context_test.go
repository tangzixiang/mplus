package mplus

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetContextValue(t *testing.T) {

	ctx := NewContext(context.Background())

	type args struct {
		ctx   context.Context
		key   string
		value interface{}
	}

	tests := []struct {
		name string
		args args
	}{
		{
			args: args{ctx: ctx, key: "a", value: "aa"},
		},
		{
			args: args{ctx: ctx, key: "b", value: "bb"},
		},
		{
			args: args{ctx: ctx, key: "c", value: "cc"},
		},
		{
			args: args{ctx: ctx, key: "d", value: "dd"},
		},
	}

	values := map[string]interface{}{}

	for _, tt := range tests {
		SetContextValue(tt.args.ctx, tt.args.key, tt.args.value)
		values[tt.args.key] = tt.args.value
	}

	for _, tt := range tests {
		assert.Equal(t, GetContextValue(ctx, tt.args.key), values[tt.args.key])
	}
}

func TestSetContextValueNotNew(t *testing.T) {

	ctx := context.Background()

	type args struct {
		ctx   context.Context
		key   string
		value interface{}
	}

	tests := []struct {
		name string
		args args
	}{
		{
			args: args{key: "a", value: "aa"},
		},
		{
			args: args{key: "b", value: "bb"},
		},
		{
			args: args{key: "c", value: "cc"},
		},
		{
			args: args{key: "d", value: "dd"},
		},
	}

	values := map[string]interface{}{}

	for _, tt := range tests {
		ctx = SetContextValue(ctx, tt.args.key, tt.args.value)
		values[tt.args.key] = tt.args.value
	}

	for _, tt := range tests {
		assert.Equal(t, GetContextValue(ctx, tt.args.key), values[tt.args.key])
	}
}

func TestGetContextValueString(t *testing.T) {

	ctx := NewContext(context.Background())

	type args struct {
		ctx context.Context
		key string
	}

	tests := []struct {
		args args
		want string
	}{
		{args{ctx: ctx, key: "a"}, "a"},
		{args{ctx: ctx, key: "b"}, "b"},
		{args{ctx: ctx, key: "c"}, "c"},
		{args{ctx: ctx, key: "d"}, "d"},
	}

	for _, tt := range tests {
		SetContextValue(tt.args.ctx, tt.args.key, tt.want)
	}

	for _, tt := range tests {
		assert.Equal(t, GetContextValueString(tt.args.ctx, tt.args.key), tt.want)
	}
}

func TestGetContextValuePanic(t *testing.T) {

	ctx := NewContext(context.Background())

	assert.Equal(t, GetContextValue(ctx, "a"), nil)
	assert.Panics(t, func() {
		_ = GetContextValue(ctx, "a").(string)
	})
}
