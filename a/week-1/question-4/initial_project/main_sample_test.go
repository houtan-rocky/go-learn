package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPage(t *testing.T) {
	history := NewBrowserHistory()
	history.VisitNewPage("page1.com")
	currentURL := history.GetCurrentURL()
	assert.Equal(t, "page1.com", currentURL, "Current URL should be 'page1.com'")
}

func TestStructFields(t *testing.T) {
	history := NewBrowserHistory()
	historyType := reflect.TypeOf(history).Elem()

	// Check 'BrowserHistory' struct fields
	currentField, found := historyType.FieldByName("current")
	assert.True(t, found, "BrowserHistory struct should have a 'current' field")

	// Check type of 'current' field
	expectedCurrentType := reflect.TypeOf(&Node{})
	assert.Equal(t, expectedCurrentType, currentField.Type, "Type of 'current' field should be '*Node'")

	// Check 'Node' struct fields
	nodeType := currentField.Type.Elem()

	urlField, found := nodeType.FieldByName("url")
	assert.True(t, found, "Node struct should have a 'url' field")
	assert.Equal(t, reflect.TypeOf(""), urlField.Type, "Type of 'url' field should be 'string'")

	prevField, found := nodeType.FieldByName("prev")
	assert.True(t, found, "Node struct should have a 'prev' field")
	assert.Equal(t, expectedCurrentType, prevField.Type, "Type of 'prev' field should be '*Node'")

	nextField, found := nodeType.FieldByName("next")
	assert.True(t, found, "Node struct should have a 'next' field")
	assert.Equal(t, expectedCurrentType, nextField.Type, "Type of 'next' field should be '*Node'")
}
