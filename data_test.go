package mplus

import (
	assert "github.com/stretchr/testify/require"
	"testing"
)

func TestData_With(t *testing.T) {

	base := map[string]interface{}{"a": 1, "b": "b", "c": true}

	for key, value := range (Data{}.With(base)) {
		assert.Equal(t, base[key], value)
	}
}

func TestData_Convert(t *testing.T) {

	base := map[string]interface{}{"a": 1, "b": "b", "c": true}

	for key, value := range Data(base) {
		assert.Equal(t, base[key], value)
	}
}

func TestData_WithNotExists(t *testing.T) {
	d := Data{"a": 1, "b": "b", "c": true}

	for key, value := range d.WithNotExists(map[string]interface{}{"a": 2, "b": "b2", "c": false}) {
		assert.Equal(t, d[key], value)
	}
}

func TestData_Push(t *testing.T) {
	d := Data{}

	d.Push("a", "a")
	assert.Equal(t, d["a"], "a")

	d.Push("a", "b")
	assert.Equal(t, d["a"], "b")

}

func TestData_PushD(t *testing.T) {
	d := Data{}

	d.PushD("a", func() interface{} {
		return "a"
	})
	assert.Equal(t, d["a"], "a")

	d.PushD("a", func() interface{} {
		return "b"
	})
	assert.Equal(t, d["a"], "b")
}

func TestData_PushIf(t *testing.T) {
	d := Data{}

	d.PushIf(true, "a", "a")
	assert.Equal(t, d["a"], "a")

	d.PushIf(false, "b", "b")
	assert.Nil(t, d["b"])
}

func TestData_PushIfD(t *testing.T) {
	d := Data{}

	d.PushIfD(true, "a", func() interface{} {
		return "a"
	})
	assert.Equal(t, d["a"], "a")

	d.PushIfD(false, "a", func() interface{} {
		return "b"
	})
	assert.Equal(t, d["a"], "a")
}

func TestData_PushPairs(t *testing.T) {
	d := Data{}

	// 入参个数必须为 2 的整数倍
	d.PushPairs("a", "a", "b")

	assert.Len(t, d, 0)

	pairs := []interface{}{"b", "b", "c", "c"}
	d.PushPairs("a", "a", pairs...)

	list := append([]interface{}{"a", "a"}, pairs...)
	for i := 0; i < len(list); i += 2 {
		assert.Equal(t, list[i], list[i+1])
	}
}

func TestData_PushPairsIf(t *testing.T) {
	d := Data{}

	// 入参个数必须为 2 的整数倍
	d.PushPairsIf(true, "a", "a", "b")

	assert.Len(t, d, 0)

	pairs := []interface{}{"b", "b", "c", "c"}
	d.PushPairsIf(false, "a", "a", pairs...)

	assert.Len(t, d, 0)

	d.PushPairsIf(true, "a", "a", pairs...)

	list := append([]interface{}{"a", "a"}, pairs...)
	for i := 0; i < len(list); i += 2 {
		assert.Equal(t, list[i], list[i+1])
	}
}

func TestData_PushNotExists(t *testing.T) {
	d := Data{"a": 1, "b": "b", "c": true}

	d.PushNotExists("a", 2).PushNotExists("d", nil)

	assert.Equal(t, 1, d["a"])
	assert.Len(t, d, 4)
}

func TestData_PushNotExistsD(t *testing.T) {
	d := Data{"a": 1, "b": "b", "c": true}

	d.PushNotExistsD("a", func() interface{} {
		return 2
	}).PushNotExistsD("d", func() interface{} {
		return 3
	})

	assert.Equal(t, 1, d["a"])
	assert.Len(t, d, 4)

}

func TestData_PushNotExistsIf(t *testing.T) {
	d := Data{"a": 1}.
		PushNotExistsIf(false, "b", "b").
		PushNotExistsIf(true, "d", "d")

	assert.Nil(t, d["b"])
	assert.Equal(t, "d", d["d"])

}

func TestData_PushNotExistsIfD(t *testing.T) {
	d := Data{"a": 1}.
		PushNotExistsIfD(false, "b", func() interface{} {
			return "b"
		}).
		PushNotExistsIfD(true, "d", func() interface{} {
			return "d"
		})

	assert.Nil(t, d["b"])
	assert.Equal(t, "d", d["d"])

}

func TestData_PushPairsNotExists(t *testing.T) {
	d := Data{}

	// 入参个数必须为 2 的整数倍
	d.PushPairsNotExists("a", "a", "b")

	assert.Len(t, d, 0)

	d = Data{"a": "a"}
	pairs := []interface{}{"b", "b", "c", "c"}
	d.PushPairsNotExists("a", "a2", pairs...)

	list := append([]interface{}{"a", "a"}, pairs...)
	for i := 0; i < len(list); i += 2 {
		assert.Equal(t, list[i], list[i+1])
	}
}

func TestData_PushPairsNotExistsIf(t *testing.T) {
	d := Data{}

	// 入参个数必须为 2 的整数倍
	d.PushPairsNotExistsIf(true, "a", "a", "b")

	assert.Len(t, d, 0)

	d = Data{"a": "a"}

	pairs := []interface{}{"b", "b", "c", "c"}
	d.PushPairsNotExistsIf(false, "a", "a2", pairs...)

	assert.Len(t, d, 1)

	d.PushPairsNotExistsIf(true, "a", "a2", pairs...)

	list := append([]interface{}{"a", "a"}, pairs...)
	for i := 0; i < len(list); i += 2 {
		assert.Equal(t, list[i], list[i+1])
	}
}

func TestData_Del(t *testing.T) {
	d := Data{"a": 1, "b": "b", "c": true}.Del("a").Del("b")

	assert.Nil(t, d["a"])
	assert.Nil(t, d["b"])
	assert.Len(t, d, 1)

}

func TestData_DelAll(t *testing.T) {
	d := Data{"a": 1, "b": "b", "c": true}.DelAll("a", "b")
	assert.Nil(t, d["a"])
	assert.Nil(t, d["b"])
	assert.Len(t, d, 1)

}

func TestData_Exists(t *testing.T) {
	d := Data{"a": 1, "b": "b", "c": true}
	assert.True(t, d.Exists("a"))
	assert.False(t, d.Exists("d"))

}

func TestData_ForEach(t *testing.T) {
	d, mark := Data{"a": 1, "b": "b", "c": true}, Data{}

	d.ForEach(func(key string, value interface{}, data Data) {
		mark.Push(key, value)
	})

	assert.Equal(t, d.Len(), mark.Len())

	mark.ForEach(func(key string, value interface{}, data Data) {
		assert.Equal(t, value, d[key])
	})
}

func TestData_Keys(t *testing.T) {
	assert.ElementsMatch(t, Data{"a": 1, "b": "b", "c": true}.Keys(), []string{"a", "b", "c"})
}

func TestData_Len(t *testing.T) {
	assert.Len(t, Data{"a": 1, "b": "b", "c": true}, 3)
}
