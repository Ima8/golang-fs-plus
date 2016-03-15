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
	assert.True(t, DeleteFolder(path), "Test Delete Folder after CreateFolder")
}

func TestGetChild(t *testing.T) {
	var path = "tmp"
	TestisExist := assert.True(t, IsExist(path), "Test isExist in TestGet	Child")
	if TestisExist {
		assert.NotEmpty(t, GetChild(path), "Test GetChild is NotEmpty when Exist")
	}
}

func TestMoveFile(t *testing.T) {
	var pathA = ""
	var pathB = ""
	var pathC = ""

	assert.NotEqual(t, MoveFile(pathA, pathC), "FAIL")
	assert.NotEqual(t, MoveFile(pathA, pathB, pathC), "FAIL")
	assert.NotEqual(t, MoveFile(pathA, pathB, pathC, "barbarbaer"), "FAIL")
	assert.Equal(t, MoveFile(pathA), "FAIL")
}
