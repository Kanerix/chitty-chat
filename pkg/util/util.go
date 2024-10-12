package util

import (
	"os"
	"path"
)

var ChippyPath = path.Join(os.TempDir(), ".chitty")
