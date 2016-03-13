package plusmerge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestisExist(t *testing.T) {
	assert.True(t, IsExist(`/`))
	assert.False(t, IsExist(`testfailfailtest11111`))
}

func TestFolderCreateFolder(t *testing.T) {
	var path = ""
	assert.True(t, CreateFolder(path))
}

func TestMoveFile(t *testing.T) {
	var pathA = ""
	var pathB = ""
	var pathC = ""

	assert.True(t, MoveFile(pathA, pathC))
	assert.True(t, MoveFile(pathA, pathB, pathC))
}
