# cregis-sdk-go

The sdk of cregis by golang is used to cregis server

## Getting started

### Prerequisites

- **[Go](https://go.dev/)**: go version >=1.20

### Getting SDK

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```
import "github.com/0xcregis/cregis-sdk-go"
```

to your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.

Otherwise, run the following Go command to install the `sdk` package:

```sh
$ go get -u github.com/0xcregis/cregis-sdk-go
```

### Using SDK

First you need to import sdk package for using, one simplest example likes the follow `example.go`:

```go
package main

import (
	"fmt"

	sdk "github.com/0xcregis/cregis-go-sdk"
)

func main() {
	c := sdk.NewClient("http://a0c1369e-12ec-467f-9989-7aba384a25e3.apple806.cc:81", "a4b0e563414a4e4dbeb407c89ce2f127", 1388205706190848)
	r, err := c.AddressLegal("195", "TXsmKpEuW7qWnXzJLGP9eDLvWPR2GRn1FS")
	if err != nil {
		fmt.Printf("err:%v", err.Error())
	} else {
		fmt.Printf("%+v", r)
	}
}

```

### Learn more examples

Learn and practice more examples, please read the [Test Case](https://github.com/0xcregis/cregis-sdk-go/blob/main/api_test.go) .

## Documentation

See [API documentation and descriptions](https://apifox.com/apidoc/shared-31f77b03-7fd6-4b95-b008-f709a1c84cdb/api-90991413)
for package.
