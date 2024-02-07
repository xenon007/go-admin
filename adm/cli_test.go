package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/xenon007/go-admin/modules/system"
)

func TestGetLatestVersion(t *testing.T) {
	assert.Equal(t, getLatestVersion(), system.Version())
}
