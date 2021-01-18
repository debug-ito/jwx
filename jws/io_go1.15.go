// +build !go1.16

// Automatically generated by internal/cmd/genreadfile/main.go. DO NOT EDIT

package jws

import (
	"os"

	"github.com/pkg/errors"
)

// ReadFile reads a JWK set at the specified location.
//

// for go >= 1.16 where io/fs is supported, you may pass `WithFS()` option
// to provide an alternate location to load the files from to provide an
// alternate location to load the files from (if you are reading
// this message, your go (or your go doc) is probably running go < 1.16)
func ReadFile(path string) (*Message, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrapf(err, `failed to open %s`, path)
	}

	defer f.Close()
	return ParseReader(f)
}
