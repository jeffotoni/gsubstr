# gsubstr

[![GoDoc][1]][2]
[![GitHub release][3]][4]
[![Coverage Status][5]][6]
[![Build Status][7]][8]
[![Go Report Card][9]][10]
[![License][11]][11]

[1]: https://godoc.org/github.com/jeffotoni/gsubstr?status.svg
[2]: https://godoc.org/github.com/jeffotoni/gsubstr
[3]: https://img.shields.io/github/v/release/jeffotoni/gsubstr?include_prereleases
[4]: https://github.com/jeffotoni/gsubstr/releases
[5]: https://coveralls.io/repos/github/jeffotoni/gsubstr/badge.svg?branch=master
[6]: https://coveralls.io/github/jeffotoni/gsubstr?branch=master
[7]: https://travis-ci.com/jeffotoni/gsubstr.svg?branch=master
[8]: https://travis-ci.com/jeffotoni/gsubstr
[9]: https://goreportcard.com/badge/github.com/jeffotoni/gsubstr
[10]: https://goreportcard.com/report/github.com/jeffotoni/gsubstr
[11]: https://img.shields.io/github/license/jeffotoni/gsubstr

>A simple library to search for parts of a string, this function is known for slicing a string as needed. The shape that has been implemented makes it possible to use it with both positive and negative sizes. A simple, lightweight lib that we use in our day-to-day projects.

## Example 

#### Example gsubstr
```go
package main

import (
	gsub "github.com/jeffotoni/gsubstr"
)

func main() {
    println(gsub.Substr("DD-MM-YYYY", 0, 2))
    // out -> DD
    
    println(gsub.Substr("DD-MM-YYYY", 3, 2))
    // out -> MM
    
    println(gsub.Substr("DD-MM-YYYY", 6, 4))
    // out -> YYYY
    
    println(gsub.Substr("YYYYMMDD", 0, 4))
    // out -> YYYY

    println(gsub.Substr("ABCD", -1, -5))
    // out -> D
    
    println(gsub.Substr("ABCDEFGH", 0))
    // out -> ABCDEFGH

    println(gsub.Substr("ABCDEFGH", 3))
    // out -> DEFGH
    
    println(gsub.Substr("ABCDEFGH", 6))
    // out -> GH

    println(gsub.Substr("ABCDEFGH", 7))
    // out -> H

    println(gsub.Substr("ABCDEFGH", -2, 1))
    // out -> G

    println(gsub.Substr("ABCDEFGH", -4, 2))
    // out -> EF

    println(gsub.Substr("ABCDEFGH", -4, 3))
    // out -> EFG

    println(gsub.Substr("ABCDEFGH", 2, -3))
    // out -> CDE
   
    println(gsub.Substr("ABCDEFGH", 3, -3))
    // out -> DE

    println(gsub.Substr("ABCDEFGH", 2, -1))
    // out -> CDEFG

    println(gsub.Substr("ABCDEFGH", -6, -3))
    // out -> CDE

    println(gsub.Substr("ABCDEFGH", -6, -1))
    // out -> CDEFG
}
```

### Some types allowed
> - Substr("DD-MM-YYYY", 0, 2) == DD
> - Substr("DD-MM-YYYY", 3, 2) == MM
> - Substr("ABCDEFGH", -6, -1) == CDEFG
> - Substr("ABCDEFGH", -4, 2)  == EF
> - Substr("ABCDEFGH", 0)      == ABCDEFGH
> - Substr("ABCDEFGH", 2, -1)  == CDEFG


## Install Using go mod in your project
```bash
$ go mod init <your-dir>
$ go mod tidy
$ go run main.go
``````

#### Another possibility would be
```bash
$ go get -u github.com/jeffotoni/gsubstr
```

#### Creator

Jefferson Otoni, [@jeffotoni](https://twitter.com/jeffotoni), [github.com/jeffotoni](https://github.com/jeffotoni), [linkedin.com/in/jeffotoni](https://www.linkedin.com/in/jeffotoni)   


Distributed under the MIT license. See ``LICENSE`` for more information.

## Contributing

To contribuit is simples, just clone or fork the repository and send to us the pull request