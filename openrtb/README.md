# OpenRTB v2.4

[Build Status]
<!--
[![Build Status](https://travis-ci.org/ilovelili/Andromeda/openrtb.svg?branch=master)](https://travis-ci.org/ilovelili/Andromeda/openrtb)
-->
OpenRTB implementation for Go

## Installation

To install, use `go get`:

```shell
go get github.com/ilovelili/Andromeda/openrtb
```

## Usage

Import the package:

```go
package main

import (
  "log"
  "github.com/ilovelili/Andromeda/openrtb"
)

func main() {
  file, err := os.Open("stored.json")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  var req *openrtb.BidRequest
  err = json.NewDecoder(file).Decode(&req)
  if err != nil {
    log.Fatal(err)
  }

  log.Printf("%+v\n", req)
}
```

## Examples
    Some test examples were taken from:
    https://code.google.com/p/openrtb/wiki/OpenRTB_Examples
