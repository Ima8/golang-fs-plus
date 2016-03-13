package golang_fs_plus_merge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileisExist(t *testing.T) {
	var path = ""
	assert.True(t, isFileExist(path))
}

func TestFolderisExist(t *testing.T) {
	var path = ""
	assert.True(t, isFolderExist(path))
}

func TestFolderCreateFolder(t *testing.T) {
	var path = ""
	assert.True(t, createFolder(path))
}

func TestMoveFile(t *testing.T) {
	var pathA = ""
	var pathB = ""
	assert.True(t, moveFile(pathA), moveFile(pathB))
}
