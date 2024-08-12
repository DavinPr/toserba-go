package common_test

import (
	"testing"
	"time"

	"github.com/DavinPr/toserba-go/common"
	"github.com/stretchr/testify/assert"
)

type anyStruct struct {
	name string
	age  int
}

func TestToPtr(t *testing.T) {
	testCases := []struct {
		name string
		val  interface{}
	}{
		{
			name: "when val is string",
			val:  "test",
		},
		{
			name: "when val is int",
			val:  45,
		},
		{
			name: "when val is float",
			val:  50.5,
		},
		{
			name: "when val is bool",
			val:  true,
		},
		{
			name: "when val is struct",
			val:  time.Now(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ptr := common.ToPtr(tc.val)
			assert.NotNil(t, ptr)
			assert.Equal(t, &tc.val, ptr)
		})
	}
}

func TestFromPtr(t *testing.T) {
	testCases := []struct {
		name string
		val  interface{}
	}{
		{
			name: "when val is string",
			val:  "test",
		},
		{
			name: "when val is int",
			val:  45,
		},
		{
			name: "when val is float",
			val:  50.5,
		},
		{
			name: "when val is bool",
			val:  true,
		},
		{
			name: "when val is time",
			val:  time.Now(),
		},
		{
			name: "when val is any struct",
			val: anyStruct{
				name: "test",
				age:  20,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ptr := common.FromPtr(&tc.val)
			assert.Equal(t, tc.val, ptr)
		})
	}
}

func TestFromPtr_Nil(t *testing.T) {
	t.Run("when type is string", func(t *testing.T) {
		val := common.FromPtr((*string)(nil))
		assert.Equal(t, "", val)
	})
	t.Run("when type is int", func(t *testing.T) {
		val := common.FromPtr((*int)(nil))
		assert.Equal(t, 0, val)
	})
	t.Run("when type is uint64", func(t *testing.T) {
		val := common.FromPtr((*uint64)(nil))
		assert.Equal(t, uint64(0), val)
	})
	t.Run("when type is float64", func(t *testing.T) {
		val := common.FromPtr((*float64)(nil))
		assert.Equal(t, float64(0), val)
	})
	t.Run("when type is bool", func(t *testing.T) {
		val := common.FromPtr((*bool)(nil))
		assert.Equal(t, false, val)
	})
	t.Run("when type is time", func(t *testing.T) {
		val := common.FromPtr((*time.Time)(nil))
		assert.Equal(t, time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), val)
	})
	t.Run("when type is struct", func(t *testing.T) {
		val := common.FromPtr((*anyStruct)(nil))
		assert.Equal(t, anyStruct{}, val)
	})
}
