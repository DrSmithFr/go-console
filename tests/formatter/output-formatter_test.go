package main

import (
	"github.com/MrSmith777/go-console/pkg/formatter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyTag(t *testing.T) {
	format := formatter.NewOutputFormatter()

	assert.Equal(t, "foo<>bar", format.Format("foo<>bar"))
}

func TestLGCharEscaping(t *testing.T) {
	format := formatter.NewOutputFormatter()
	format.SetDecorated(true)

	assert.Equal(t, "foo<>bar", format.Format("foo<>bar"))
}