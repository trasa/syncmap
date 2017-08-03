package syncmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapList_Add(t *testing.T) {
	m := NewMapList()
	m.Add("foo", 1)
	m.Add("foo", 2)
	m.Add("foo", 3)
	m.Add("bar", 4)
	m.Add("bar", 5)
	m.Add("bar", 6)

	assert.Equal(t, 3, len(m.Get("foo")))
	assert.Equal(t, 3, len(m.Get("bar")))
	assert.Equal(t, 1, m.Get("foo")[0])
	assert.Equal(t, 1, m.Get("foo")[0])
	assert.Equal(t, 1, m.Get("foo")[0])
	assert.Equal(t, 1, m.Get("foo")[0])
}

func TestMapList_Remove(t *testing.T) {
	m := NewMapList()
	m.Add("foo", 1)
	m.Add("foo", 2)
	m.Add("foo", 3)
	m.Add("bar", 4)
	m.Add("bar", 5)
	m.Add("bar", 6)

	m.Remove("foo")
	assert.Equal(t, 0, len(m.Get("foo")))
	assert.Equal(t, 3, len(m.Get("bar")))
}

func TestMapList_RemoveItem(t *testing.T) {
	m := NewMapList()
	m.Add("foo", 1)
	m.Add("foo", 2)
	m.Add("foo", 3)

	m.RemoveItem("foo", 2)
	assert.Equal(t, 2, len(m.Get("foo")))
	assert.Equal(t, 1, m.Get("foo")[0])
	assert.Equal(t, 3, m.Get("foo")[1])
}
