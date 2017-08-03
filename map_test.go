package syncmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap_AddGet(t *testing.T) {
	m := New()
	m.Add("foo", 1)
	m.Get("foo")

	assert.Equal(t, 1, m.Get("foo"))
	assert.Nil(t, m.Get("bar"))
}

func TestMap_Remove(t *testing.T) {
	m := New()
	m.Add("foo", 5)
	m.Remove("foo")

	assert.Nil(t, m.Get("foo"))
}

func TestMap_Iter(t *testing.T) {
	m := New()
	m.Add("foo", 123)
	m.Iter(func(v interface{}) {
		i := v.(int)
		assert.Equal(t, 123, i)
	})
}
