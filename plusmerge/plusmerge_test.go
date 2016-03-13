package plusmerge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestisExist(t *testing.T) {
	assert.True(t, isExist(`/`))
	assert.False(t, isExist(`testfailfailtest11111`))
}

func TestFolderCreateFolder(t *testing.T) {
	var path = ""
	assert.True(t, createFolder(path))
}

func TestMoveFile(t *testing.T) {
	var pathA = ""
	var pathB = ""
	var pathC = ""

	assert.True(t, moveFile(pathA, pathC))
	assert.True(t, moveFile(pathA, pathB, pathC))
}
