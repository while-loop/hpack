hpack
=====

<p align="center">
  <img src="https://github.com/while-loop/hpack/blob/master/doc/box.png">
  <br>
  <a href="https://godoc.org/github.com/while-loop/hpack"><img src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square"></a>
  <a href="https://travis-ci.org/while-loop/hpack"><img src="https://img.shields.io/travis/while-loop/hpack.svg?style=flat-square"></a>
  <a href="https://github.com/while-loop/hpack/releases"><img src="https://img.shields.io/github/release/while-loop/hpack.svg?style=flat-square"></a>
  <a href="https://coveralls.io/github/while-loop/hpack"><img src="https://img.shields.io/coveralls/while-loop/hpack.svg?style=flat-square"></a>
  <a href="https://github.com/while-loop/hpack/blob/master/LICENSE"><img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=flat-square"></a>
</p>

Go implementation of HTTP/2's compression format [RFC 7541](https://tools.ietf.org/html/rfc7541) for efficiently representing HTTP header fields.

Installation
------------

```
$ go get github.com/while-loop/hpack
```

Usage
-----

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
}
```

Changelog
---------

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

[CHANGELOG.md](https://github.com/while-loop/hpack/blob/master/CHANGELOG.md)

License
-------
hpack is licensed under the Apache 2.0 license. See
[LICENSE](https://github.com/while-loop/hpack/blob/master/LICENSE)
for details.

Author
------

Anthony Alves
