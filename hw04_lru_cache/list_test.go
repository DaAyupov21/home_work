package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList_BasicOps(t *testing.T) {
	l := NewList()
	require.Equal(t, 0, l.Len())
	require.Nil(t, l.Front())
	require.Nil(t, l.Back())

	a := l.PushBack(10)
	b := l.PushBack(20)
	c := l.PushBack(30)
	require.Equal(t, 3, l.Len())
	require.Equal(t, a, l.Front())
	require.Equal(t, c, l.Back())

	l.Remove(b) // [10,30]
	require.Equal(t, 2, l.Len())
	require.Equal(t, 10, l.Front().Value)
	require.Equal(t, 30, l.Back().Value)

	l.MoveToFront(c) // [30,10]
	require.Equal(t, c, l.Front())
	require.Equal(t, a, l.Back())

	l.Remove(c) // [10]
	require.Equal(t, 1, l.Len())
	require.Equal(t, a, l.Front())
	require.Equal(t, a, l.Back())
}

func TestList_MoveToFront_Idempotent(t *testing.T) {
	l := NewList()
	x := l.PushFront("x")
	l.MoveToFront(x)
	require.Equal(t, x, l.Front())
	require.Equal(t, x, l.Back())
}
