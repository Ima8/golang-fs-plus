package plusmerge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestisExist(t *testing.T) {
	assert.True(t, IsExist(`/`), "Test isExist path '/'")
	assert.False(t, IsExist(`wgefwherw`), "Test isExist path 'wgefwherw'")
}

func TestFolderCreateFolder(t *testing.T) {
	var path = "./max"
	assert.True(t, CreateFolder(path), "Test CreateFolder")
	assert.True(t, IsExist(path), "Test isExist after CreateFolder")
	assert.True(t, DeleteFolder(path), "Test Delete Folder after CreateFolder")
}

func TestGetChild(t *testing.T) {
	var path = "tmp"
	assert.True(t, IsExist(path), "Test isExist in TestGet	Child")
	assert.NotEqual(t, GetChild(path), false, "Test GetChild after isExist")
}

func TestMoveFile(t *testing.T) {
	var pathA = ""
	var pathB = ""
	var pathC = ""

	assert.True(t, MoveFile(pathA, pathC))
	assert.True(t, MoveFile(pathA, pathB, pathC))
}
