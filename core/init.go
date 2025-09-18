package core

import (
	"os"
)

func init() {
	err := os.MkdirAll(dataPath, 0755)
	if err != nil {
		panic(err)
	}
}
