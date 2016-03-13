package plusmerge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestisExist(t *testing.T) {
	assert.True(t, IsExist(`/`), "Test isExist path '/'")
	assert.False(t, IsExist(`wgefwherw`), "Test isExist path 'wgefwherw'")
}

func TestFolderCreateFolder(t *testing.T) {
	var path = "./max"
	assert.True(t, CreateFolder(path), "Test CreateFolder")
	assert.True(t, IsExist(path), "Test isExist after CreateFolder")
}

func TestMoveFile(t *testing.T) {
	var pathA = ""
	var pathB = ""
	var pathC = ""

	assert.True(t, MoveFile(pathA, pathC))
	assert.True(t, MoveFile(pathA, pathB, pathC))
}
