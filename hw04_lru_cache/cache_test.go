package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache_EvictByCapacity(t *testing.T) {
	c := NewCache(3)
	c.Set("a", 1)
	c.Set("b", 2)
	c.Set("c", 3)
	c.Set("d", 4) // вытесняет "a"

	_, ok := c.Get("a")
	require.False(t, ok)
	v, ok := c.Get("b")
	require.True(t, ok)
	require.Equal(t, 2, v)
	v, ok = c.Get("c")
	require.True(t, ok)
	require.Equal(t, 3, v)
	v, ok = c.Get("d")
	require.True(t, ok)
	require.Equal(t, 4, v)
}

func TestCache_EvictLeastRecentlyUsed(t *testing.T) {
	c := NewCache(3)
	c.Set("a", 1)  // a
	c.Set("b", 2)  // a b
	c.Set("c", 3)  // a b c
	c.Get("a")     // b c a
	c.Set("b", 20) // c a b
	c.Set("d", 4)  // a b d  -> вытеснился c

	_, ok := c.Get("c")
	require.False(t, ok)
	v, ok := c.Get("a")
	require.True(t, ok)
	require.Equal(t, 1, v)
	v, ok = c.Get("b")
	require.True(t, ok)
	require.Equal(t, 20, v)
	v, ok = c.Get("d")
	require.True(t, ok)
	require.Equal(t, 4, v)
}

func TestCache_SetFlag(t *testing.T) {
	c := NewCache(2)
	require.False(t, c.Set("x", 1))
	require.True(t, c.Set("x", 2))
	v, ok := c.Get("x")
	require.True(t, ok)
	require.Equal(t, 2, v)
}

func TestCache_Clear(t *testing.T) {
	c := NewCache(2)
	c.Set("a", 1)
	c.Set("b", 2)
	c.Clear()
	_, okA := c.Get("a")
	_, okB := c.Get("b")
	require.False(t, okA)
	require.False(t, okB)
	c.Set("c", 3)
	v, ok := c.Get("c")
	require.True(t, ok)
	require.Equal(t, 3, v)
}
