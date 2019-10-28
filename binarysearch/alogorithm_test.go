package binarysearch_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andremissaglia/algorithms_study/binarysearch"
	"github.com/andremissaglia/algorithms_study/internal/utils"
)

func TestBinarySearch(t *testing.T) {
	list := []utils.Element{
		{Key: "A", Value: "123"},
		{Key: "B", Value: "456"},
		{Key: "C", Value: "789"},
		{Key: "D", Value: "012"},
		{Key: "E", Value: "345"},
		{Key: "F", Value: "678"},
	}
	result, err := binarysearch.BinarySearch(list, "D")
	assert.NoError(t, err)
	assert.Equal(t, utils.Element{Key: "D", Value: "012"}, result)
}

func TestBinarySearchNotFound(t *testing.T) {
	list := []utils.Element{
		{Key: "A", Value: "123"},
		{Key: "B", Value: "456"},
		{Key: "C", Value: "789"},
		{Key: "E", Value: "345"},
		{Key: "F", Value: "678"},
	}
	_, err := binarysearch.BinarySearch(list, "D")
	assert.Error(t, err)
}

func TestBinarySearchEmpty(t *testing.T) {
	list := []utils.Element{}
	_, err := binarysearch.BinarySearch(list, "D")
	assert.Error(t, err)
}
