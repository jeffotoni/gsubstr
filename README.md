# gsubstr

>Just a simple lib, to gsubstr: returns a part of a string

#### Use

Example gsubstr
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

#### Install Using go mod in your project
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