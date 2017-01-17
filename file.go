// Copyright 2017 Cwen. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package justso

import "net/http"

type (
	onlyfilesFS struct {
		fs http.FileSystem
	}
)

// Dir returns a http.Filesystem that can be used by http.FileServer(). It is used interally
// in router.Static().
// if listDirectory == true, then it works the same as http.Dir() otherwise it returns
// a filesystem that prevents http.FileServer() to list the directory files.
func Dir(root string, listDirectory bool) http.FileSystem {
	fs := http.Dir(root)
	if listDirectory {
		return fs
	}
	return &onlyfilesFS{fs}
}
