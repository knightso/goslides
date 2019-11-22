package hello

import (
	"os"
	"testing"
)

// START OMIT
func TestMain(m *testing.M) {

	// パッケージ前処理

	cd := m.Run()

	// パッケージ後処理

	os.Exit(cd)
}

// END OMIT
