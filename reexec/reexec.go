package reexec

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var registeredInitializers = make(map[string]func())

// Register adds an initialization func under the specified name
func Register(name string, initializer func()) {
	if _, exists := registeredInitializers[name]; exists {
		panic(fmt.Sprintf("reexec func already registred under name %q", name))
	}

	registeredInitializers[name] = initializer
}

// Init is called as the first part of the exec process and returns true if an
// initialization function was called.
func Init() bool {
	initializer, exists := registeredInitializers[os.Args[0]]
	if exists {
		initializer()

		return true
	}

	return false
}

// Self returns the path to the current processes binary
func Self() string {
	name := os.Args[0]

	if filepath.Base(name) == name {
		if lp, err := exec.LookPath(name); err == nil {
			name = lp
		}
	}

	return name
}
