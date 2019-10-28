package avltree_test

import (
	"testing"

	"github.com/andremissaglia/algorithms_study/avltree"
	"github.com/andremissaglia/algorithms_study/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestAvlTreeState(t *testing.T) {
	tree := avltree.New()
	tree.Insert(utils.Element{Key: "A", Value: "123"})
	tree.Insert(utils.Element{Key: "B", Value: "456"})
	tree.Insert(utils.Element{Key: "C", Value: "789"})
	tree.Insert(utils.Element{Key: "D", Value: "012"})
	tree.Insert(utils.Element{Key: "BA", Value: "345"})
	tree.Insert(utils.Element{Key: "BAA", Value: "678"})

	preorder := []utils.Element{
		{Key: "BA", Value: "345"},
		{Key: "B", Value: "456"},
		{Key: "A", Value: "123"},
		{Key: "C", Value: "789"},
		{Key: "BAA", Value: "678"},
		{Key: "D", Value: "012"},
	}

	traversal := tree.PreOrderTraversal()

	assert.Equal(t, preorder, traversal, "insert test")

	tree.Remove("D")
	tree.Remove("B")
	tree.Remove("A")
	tree.Remove("J")

	preorder = []utils.Element{
		{Key: "BAA", Value: "678"},
		{Key: "BA", Value: "345"},
		{Key: "C", Value: "789"},
	}
}

func TestAvlTree(t *testing.T) {
	tree := avltree.New()
	tree.Insert(utils.Element{Key: "C", Value: "789"})
	tree.Insert(utils.Element{Key: "A", Value: "123"})
	tree.Insert(utils.Element{Key: "B", Value: "456"})
	tree.Insert(utils.Element{Key: "D", Value: "012"})
	result, err := tree.Get("B")
	assert.Equal(t, "456", result.Value)
	assert.NoError(t, err)

	_, err = tree.Get("J")
	assert.Error(t, err)
}
