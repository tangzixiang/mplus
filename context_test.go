package mplus

import (
	"context"
	"math"
	"reflect"
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
	"github.com/tangzixiang/mplus/util"
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
		ctx          context.Context
		key          string
		defaultValue []string
		setValue     func() (string, bool)
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{ctx: ctx, key: "a", defaultValue: []string{"1"}, setValue: func() (string, bool) {
				return "", false
			}}, want: "1", name: "GetStringValueWithDefault",
		},
		{
			args: args{ctx: ctx, key: "b", defaultValue: []string{"1"}, setValue: func() (string, bool) {
				return "", true
			}}, want: "", name: "GetStringValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "c", defaultValue: []string{"1"}, setValue: func() (string, bool) {
				return "2", true
			}}, want: "2", name: "GetStringValueWithoutDefault",
		},
	}

	for _, item := range tests {
		value, set := item.args.setValue()
		if set {
			SetContextValue(item.args.ctx, item.args.key, value)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetContextValueString(tt.args.ctx, tt.args.key, tt.args.defaultValue...); got != tt.want {
				t.Errorf("TestGetContextValueString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetContextValuePanic(t *testing.T) {

	ctx := NewContext(context.Background())

	assert.Equal(t, GetContextValue(ctx, "a"), nil)
	assert.Panics(t, func() {
		_ = GetContextValue(ctx, "a").(string)
	})
}

func TestGetContextValue(t *testing.T) {
	ctx := NewContext(context.Background())
	data := map[string]interface{}{
		"map":   map[string]interface{}{"key1": "value1"},
		"slice": []string{"item1"},
		"chan":  make(chan struct{}),
	}

	for key, value := range data {
		SetContextValue(ctx, key, value)
	}

	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{args: args{ctx: ctx, key: "map"}, name: "getMapFromCtx", want: data["map"]},
		{args: args{ctx: ctx, key: "slice"}, name: "getSliceFromCtx", want: data["slice"]},
		{args: args{ctx: ctx, key: "chan"}, name: "getChanFromCtx", want: data["chan"]},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetContextValue(tt.args.ctx, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetContextValue() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestNewContext(t *testing.T) {
	assert.NotPanics(t, func() {
		_ = NewContext(context.Background())
	})

	assert.NotNil(t, NewContext(context.Background()))

	timeoutBGCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ctx := NewContext(timeoutBGCtx)

	success := false
	select {
	case <-ctx.Done():
		success = true
	case <-time.After(time.Second * 2):
		success = false
	}

	assert.True(t, success)
}

func TestCopyContext(t *testing.T) {
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

	for _, tt := range tests {
		SetContextValue(tt.args.ctx, tt.args.key, tt.args.value)
	}

	copyCtx := CopyContext(ctx)

	assert.NotEqual(t, copyCtx, ctx)

	for _, tt := range tests {
		assert.Equal(t, GetContextValue(ctx, tt.args.key), GetContextValue(copyCtx, tt.args.key))
	}
}

func TestGetContextValueInt(t *testing.T) {
	ctx := NewContext(context.Background())

	max, min := 0, 0
	if util.IsSystem64Bit() {
		max = math.MaxInt64
		min = math.MinInt64
	} else {
		max = math.MaxInt32
		min = math.MinInt32
	}

	type args struct {
		ctx          context.Context
		key          string
		defaultValue []int
		setValue     func() (int, bool)
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{ctx: ctx, key: "a", defaultValue: []int{1}, setValue: func() (int, bool) {
				return 0, false
			}}, want: 1, name: "GetIntValueWithDefault",
		},
		{
			args: args{ctx: ctx, key: "b", defaultValue: []int{1}, setValue: func() (int, bool) {
				return 0, true
			}}, want: 0, name: "GetIntValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "max", defaultValue: []int{1}, setValue: func() (int, bool) {
				return max, true
			}}, want: max, name: "GetMaxIntValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "min", defaultValue: []int{1}, setValue: func() (int, bool) {
				return min, true
			}}, want: min, name: "GetMinIntValueWithoutDefault",
		},
	}

	for _, item := range tests {
		value, set := item.args.setValue()
		if set {
			SetContextValue(item.args.ctx, item.args.key, value)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetContextValueInt(tt.args.ctx, tt.args.key, tt.args.defaultValue...); got != tt.want {
				t.Errorf("GetContextValueInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetContextValueInt8(t *testing.T) {
	ctx := NewContext(context.Background())

	type args struct {
		ctx          context.Context
		key          string
		defaultValue []int8
		setValue     func() (int8, bool)
	}

	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			args: args{ctx: ctx, key: "a", defaultValue: []int8{1}, setValue: func() (int8, bool) {
				return 0, false
			}}, want: 1, name: "GetInt8ValueWithDefault",
		},
		{
			args: args{ctx: ctx, key: "b", defaultValue: []int8{1}, setValue: func() (int8, bool) {
				return 0, true
			}}, want: 0, name: "GetInt8ValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "max", defaultValue: []int8{1}, setValue: func() (int8, bool) {
				return math.MaxInt8, true
			}}, want: math.MaxInt8, name: "GetMaxInt8ValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "min", defaultValue: []int8{1}, setValue: func() (int8, bool) {
				return math.MinInt8, true
			}}, want: math.MinInt8, name: "GetMinInt8ValueWithoutDefault",
		},
	}

	for _, item := range tests {
		value, set := item.args.setValue()
		if set {
			SetContextValue(item.args.ctx, item.args.key, value)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetContextValueInt8(tt.args.ctx, tt.args.key, tt.args.defaultValue...); got != tt.want {
				t.Errorf("GetContextValueInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetContextValueInt16(t *testing.T) {
	ctx := NewContext(context.Background())

	type args struct {
		ctx          context.Context
		key          string
		defaultValue []int16
		setValue     func() (int16, bool)
	}

	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			args: args{ctx: ctx, key: "a", defaultValue: []int16{1}, setValue: func() (int16, bool) {
				return 0, false
			}}, want: 1, name: "GetInt16ValueWithDefault",
		},
		{
			args: args{ctx: ctx, key: "b", defaultValue: []int16{1}, setValue: func() (int16, bool) {
				return 0, true
			}}, want: 0, name: "GetInt16ValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "max", defaultValue: []int16{1}, setValue: func() (int16, bool) {
				return math.MaxInt16, true
			}}, want: math.MaxInt16, name: "GetMaxInt16ValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "min", defaultValue: []int16{1}, setValue: func() (int16, bool) {
				return math.MinInt16, true
			}}, want: math.MinInt16, name: "GetMinInt16ValueWithoutDefault",
		},
	}

	for _, item := range tests {
		value, set := item.args.setValue()
		if set {
			SetContextValue(item.args.ctx, item.args.key, value)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetContextValueInt16(tt.args.ctx, tt.args.key, tt.args.defaultValue...); got != tt.want {
				t.Errorf("GetContextValueInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetContextValueInt32(t *testing.T) {
	ctx := NewContext(context.Background())

	type args struct {
		ctx          context.Context
		key          string
		defaultValue []int32
		setValue     func() (int32, bool)
	}

	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			args: args{ctx: ctx, key: "a", defaultValue: []int32{1}, setValue: func() (int32, bool) {
				return 0, false
			}}, want: 1, name: "GetInt32ValueWithDefault",
		},
		{
			args: args{ctx: ctx, key: "b", defaultValue: []int32{1}, setValue: func() (int32, bool) {
				return 0, true
			}}, want: 0, name: "GetInt32ValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "max", defaultValue: []int32{1}, setValue: func() (int32, bool) {
				return math.MaxInt32, true
			}}, want: math.MaxInt32, name: "GetMaxInt32ValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "min", defaultValue: []int32{1}, setValue: func() (int32, bool) {
				return math.MinInt32, true
			}}, want: math.MinInt32, name: "GetMinInt32ValueWithoutDefault",
		},
	}

	for _, item := range tests {
		value, set := item.args.setValue()
		if set {
			SetContextValue(item.args.ctx, item.args.key, value)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetContextValueInt32(tt.args.ctx, tt.args.key, tt.args.defaultValue...); got != tt.want {
				t.Errorf("GetContextValueInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetContextValueInt64(t *testing.T) {
	ctx := NewContext(context.Background())

	type args struct {
		ctx          context.Context
		key          string
		defaultValue []int64
		setValue     func() (int64, bool)
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{ctx: ctx, key: "a", defaultValue: []int64{1}, setValue: func() (int64, bool) {
				return 0, false
			}}, want: 1, name: "GetInt64ValueWithDefault",
		},
		{
			args: args{ctx: ctx, key: "b", defaultValue: []int64{1}, setValue: func() (int64, bool) {
				return 0, true
			}}, want: 0, name: "GetInt64ValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "max", defaultValue: []int64{1}, setValue: func() (int64, bool) {
				return math.MaxInt64, true
			}}, want: math.MaxInt64, name: "GetMaxInt64ValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "min", defaultValue: []int64{1}, setValue: func() (int64, bool) {
				return math.MinInt64, true
			}}, want: math.MinInt64, name: "GetMinInt64ValueWithoutDefault",
		},
	}

	for _, item := range tests {
		value, set := item.args.setValue()
		if set {
			SetContextValue(item.args.ctx, item.args.key, value)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetContextValueInt64(tt.args.ctx, tt.args.key, tt.args.defaultValue...); got != tt.want {
				t.Errorf("GetContextValueInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetContextValueUInt(t *testing.T) {
	ctx := NewContext(context.Background())

	max, min := uint(0), uint(0)
	if util.IsSystem64Bit() {
		max = math.MaxUint64
		min = 0
	} else {
		max = math.MaxUint32
		min = 0
	}

	type args struct {
		ctx          context.Context
		key          string
		defaultValue []uint
		setValue     func() (uint, bool)
	}

	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			args: args{ctx: ctx, key: "a", defaultValue: []uint{1}, setValue: func() (uint, bool) {
				return 0, false
			}}, want: 1, name: "GetUIntValueWithDefault",
		},
		{
			args: args{ctx: ctx, key: "b", defaultValue: []uint{1}, setValue: func() (uint, bool) {
				return 0, true
			}}, want: 0, name: "GetUIntValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "max", defaultValue: []uint{1}, setValue: func() (uint, bool) {
				return max, true
			}}, want: max, name: "GetMaxUIntValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "min", defaultValue: []uint{1}, setValue: func() (uint, bool) {
				return min, true
			}}, want: min, name: "GetMinUIntValueWithoutDefault",
		},
	}

	for _, item := range tests {
		value, set := item.args.setValue()
		if set {
			SetContextValue(item.args.ctx, item.args.key, value)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetContextValueUInt(tt.args.ctx, tt.args.key, tt.args.defaultValue...); got != tt.want {
				t.Errorf("GetContextValueUInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetContextValueUInt8(t *testing.T) {
	ctx := NewContext(context.Background())

	type args struct {
		ctx          context.Context
		key          string
		defaultValue []uint8
		setValue     func() (uint8, bool)
	}

	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			args: args{ctx: ctx, key: "a", defaultValue: []uint8{1}, setValue: func() (uint8, bool) {
				return 0, false
			}}, want: 1, name: "GetInt8ValueWithDefault",
		},
		{
			args: args{ctx: ctx, key: "b", defaultValue: []uint8{1}, setValue: func() (uint8, bool) {
				return 0, true
			}}, want: 0, name: "GetInt8ValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "max", defaultValue: []uint8{1}, setValue: func() (uint8, bool) {
				return math.MaxUint8, true
			}}, want: math.MaxUint8, name: "GetMaxUInt8ValueWithoutDefault",
		},
	}

	for _, item := range tests {
		value, set := item.args.setValue()
		if set {
			SetContextValue(item.args.ctx, item.args.key, value)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetContextValueUInt8(tt.args.ctx, tt.args.key, tt.args.defaultValue...); got != tt.want {
				t.Errorf("GetContextValueUInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetContextValueUInt16(t *testing.T) {
	ctx := NewContext(context.Background())

	type args struct {
		ctx          context.Context
		key          string
		defaultValue []uint16
		setValue     func() (uint16, bool)
	}

	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			args: args{ctx: ctx, key: "a", defaultValue: []uint16{1}, setValue: func() (uint16, bool) {
				return 0, false
			}}, want: 1, name: "GetUInt16ValueWithDefault",
		},
		{
			args: args{ctx: ctx, key: "b", defaultValue: []uint16{1}, setValue: func() (uint16, bool) {
				return 0, true
			}}, want: 0, name: "GetUInt16ValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "max", defaultValue: []uint16{1}, setValue: func() (uint16, bool) {
				return math.MaxUint16, true
			}}, want: math.MaxUint16, name: "GetMaxUInt16ValueWithoutDefault",
		},
	}

	for _, item := range tests {
		value, set := item.args.setValue()
		if set {
			SetContextValue(item.args.ctx, item.args.key, value)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetContextValueUInt16(tt.args.ctx, tt.args.key, tt.args.defaultValue...); got != tt.want {
				t.Errorf("GetContextValueUInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetContextValueUInt32(t *testing.T) {
	ctx := NewContext(context.Background())

	type args struct {
		ctx          context.Context
		key          string
		defaultValue []uint32
		setValue     func() (uint32, bool)
	}

	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			args: args{ctx: ctx, key: "a", defaultValue: []uint32{1}, setValue: func() (uint32, bool) {
				return 0, false
			}}, want: 1, name: "GetUInt32ValueWithDefault",
		},
		{
			args: args{ctx: ctx, key: "b", defaultValue: []uint32{1}, setValue: func() (uint32, bool) {
				return 0, true
			}}, want: 0, name: "GetUInt32ValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "max", defaultValue: []uint32{1}, setValue: func() (uint32, bool) {
				return math.MaxUint32, true
			}}, want: math.MaxUint32, name: "GetMaxUInt32ValueWithoutDefault",
		},
	}

	for _, item := range tests {
		value, set := item.args.setValue()
		if set {
			SetContextValue(item.args.ctx, item.args.key, value)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetContextValueUInt32(tt.args.ctx, tt.args.key, tt.args.defaultValue...); got != tt.want {
				t.Errorf("GetContextValueUInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetContextValueUInt64(t *testing.T) {
	ctx := NewContext(context.Background())

	type args struct {
		ctx          context.Context
		key          string
		defaultValue []uint64
		setValue     func() (uint64, bool)
	}

	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			args: args{ctx: ctx, key: "a", defaultValue: []uint64{1}, setValue: func() (uint64, bool) {
				return 0, false
			}}, want: 1, name: "GetUInt64ValueWithDefault",
		},
		{
			args: args{ctx: ctx, key: "b", defaultValue: []uint64{1}, setValue: func() (uint64, bool) {
				return 0, true
			}}, want: 0, name: "GetUInt64ValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "max", defaultValue: []uint64{1}, setValue: func() (uint64, bool) {
				return math.MaxUint64, true
			}}, want: math.MaxUint64, name: "GetMaxUInt64ValueWithoutDefault",
		},
	}

	for _, item := range tests {
		value, set := item.args.setValue()
		if set {
			SetContextValue(item.args.ctx, item.args.key, value)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetContextValueUInt64(tt.args.ctx, tt.args.key, tt.args.defaultValue...); got != tt.want {
				t.Errorf("GetContextValueUInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetContextValueBool(t *testing.T) {
	ctx := NewContext(context.Background())

	type args struct {
		ctx          context.Context
		key          string
		defaultValue []bool
		setValue     func() (bool, bool)
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{ctx: ctx, key: "a", defaultValue: []bool{true}, setValue: func() (bool, bool) {
				return false, false
			}}, want: true, name: "GetBoolValueWithDefault",
		},
		{
			args: args{ctx: ctx, key: "b", defaultValue: []bool{true}, setValue: func() (bool, bool) {
				return false, true
			}}, want: false, name: "GetBoolValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "c", defaultValue: []bool{true}, setValue: func() (bool, bool) {
				return true, true
			}}, want: true, name: "GetBoolValueWithoutDefault",
		},
	}

	for _, item := range tests {
		value, set := item.args.setValue()
		if set {
			SetContextValue(item.args.ctx, item.args.key, value)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetContextValueBool(tt.args.ctx, tt.args.key, tt.args.defaultValue...); got != tt.want {
				t.Errorf("GetContextValueBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetContextValueByte(t *testing.T) {
	ctx := NewContext(context.Background())

	type args struct {
		ctx          context.Context
		key          string
		defaultValue []byte
		setValue     func() (byte, bool)
	}

	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			args: args{ctx: ctx, key: "a", defaultValue: []byte{1}, setValue: func() (byte, bool) {
				return 0, false
			}}, want: 1, name: "GetByteValueWithDefault",
		},
		{
			args: args{ctx: ctx, key: "b", defaultValue: []byte{1}, setValue: func() (byte, bool) {
				return 0, true
			}}, want: 0, name: "GetByteValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "c", defaultValue: []byte{1}, setValue: func() (byte, bool) {
				return 2, true
			}}, want: 2, name: "GetByteValueWithoutDefault",
		},
	}

	for _, item := range tests {
		value, set := item.args.setValue()
		if set {
			SetContextValue(item.args.ctx, item.args.key, value)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetContextValueByte(tt.args.ctx, tt.args.key, tt.args.defaultValue...); got != tt.want {
				t.Errorf("GetContextValueByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetContextValueBytes(t *testing.T) {
	ctx := NewContext(context.Background())

	type args struct {
		ctx          context.Context
		key          string
		defaultValue [][]byte
		setValue     func() ([]byte, bool)
	}

	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{ctx: ctx, key: "a", defaultValue: [][]byte{{1}}, setValue: func() ([]byte, bool) {
				return nil, false
			}}, want: []byte{1}, name: "GetBytesValueWithDefault",
		},
		{
			args: args{ctx: ctx, key: "b", defaultValue: [][]byte{{1}}, setValue: func() ([]byte, bool) {
				return nil, true
			}}, want: nil, name: "GetBytesValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "c", defaultValue: [][]byte{{1}}, setValue: func() ([]byte, bool) {
				return []byte{1}, true
			}}, want: []byte{1}, name: "GetBytesValueWithoutDefault",
		},
	}

	for _, item := range tests {
		value, set := item.args.setValue()
		if set {
			SetContextValue(item.args.ctx, item.args.key, value)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetContextValueBytes(tt.args.ctx, tt.args.key, tt.args.defaultValue...)
			if got == nil && tt.want != nil {
				t.Errorf("GetContextValueBytes() = %v, want %v", got, tt.want)
			} else if got != nil && tt.want == nil {
				t.Errorf("GetContextValueBytes() = %v, want %v", got, tt.want)
			} else {
				assert.ElementsMatch(t, got, tt.want)
			}
		})
	}
}

func TestGetContextValueTime(t *testing.T) {
	ctx := NewContext(context.Background())
	now := time.Now()
	zeroTime := time.Time{}

	type args struct {
		ctx          context.Context
		key          string
		defaultValue []time.Time
		setValue     func() (time.Time, bool)
	}

	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			args: args{ctx: ctx, key: "a", defaultValue: []time.Time{now}, setValue: func() (time.Time, bool) {
				return zeroTime, false
			}}, want: now, name: "GetTimeValueWithDefault",
		},
		{
			args: args{ctx: ctx, key: "b", defaultValue: []time.Time{now}, setValue: func() (time.Time, bool) {
				return zeroTime, true
			}}, want: zeroTime, name: "GetTimeValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "c", defaultValue: []time.Time{now}, setValue: func() (time.Time, bool) {
				return now, true
			}}, want: now, name: "GetTimeValueWithoutDefault",
		},
	}

	for _, item := range tests {
		value, set := item.args.setValue()
		if set {
			SetContextValue(item.args.ctx, item.args.key, value)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetContextValueTime(tt.args.ctx, tt.args.key, tt.args.defaultValue...); got != tt.want {
				t.Errorf("GetContextValueTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetContextValueFloat32(t *testing.T) {
	ctx := NewContext(context.Background())

	type args struct {
		ctx          context.Context
		key          string
		defaultValue []float32
		setValue     func() (float32, bool)
	}

	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			args: args{ctx: ctx, key: "a", defaultValue: []float32{1}, setValue: func() (float32, bool) {
				return 0, false
			}}, want: 1, name: "GetFloat32ValueWithDefault",
		},
		{
			args: args{ctx: ctx, key: "b", defaultValue: []float32{1}, setValue: func() (float32, bool) {
				return 0, true
			}}, want: 0, name: "GetFloat32ValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "max", defaultValue: []float32{1}, setValue: func() (float32, bool) {
				return math.MaxFloat32, true
			}}, want: math.MaxFloat32, name: "GetMaxFloat32ValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "min", defaultValue: []float32{1}, setValue: func() (float32, bool) {
				return math.SmallestNonzeroFloat32, true
			}}, want: math.SmallestNonzeroFloat32, name: "GetSmallestNonzeroFloat32ValueWithoutDefault",
		},
	}

	for _, item := range tests {
		value, set := item.args.setValue()
		if set {
			SetContextValue(item.args.ctx, item.args.key, value)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetContextValueFloat32(tt.args.ctx, tt.args.key, tt.args.defaultValue...); got != tt.want {
				t.Errorf("GetContextValueFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetContextValueFloat64(t *testing.T) {
	ctx := NewContext(context.Background())

	type args struct {
		ctx          context.Context
		key          string
		defaultValue []float64
		setValue     func() (float64, bool)
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{ctx: ctx, key: "a", defaultValue: []float64{1}, setValue: func() (float64, bool) {
				return 0, false
			}}, want: 1, name: "GetFloat64ValueWithDefault",
		},
		{
			args: args{ctx: ctx, key: "b", defaultValue: []float64{1}, setValue: func() (float64, bool) {
				return 0, true
			}}, want: 0, name: "GetFloat64ValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "max", defaultValue: []float64{1}, setValue: func() (float64, bool) {
				return math.MaxFloat64, true
			}}, want: math.MaxFloat64, name: "GetMaxFloat64ValueWithoutDefault",
		},
		{
			args: args{ctx: ctx, key: "min", defaultValue: []float64{1}, setValue: func() (float64, bool) {
				return math.SmallestNonzeroFloat64, true
			}}, want: math.SmallestNonzeroFloat64, name: "GetSmallestNonzeroFloat64ValueWithoutDefault",
		},
	}

	for _, item := range tests {
		value, set := item.args.setValue()
		if set {
			SetContextValue(item.args.ctx, item.args.key, value)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetContextValueFloat64(tt.args.ctx, tt.args.key, tt.args.defaultValue...); got != tt.want {
				t.Errorf("GetContextValueFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
