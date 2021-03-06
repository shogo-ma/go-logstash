# go-logstash
[![Build Status](https://travis-ci.org/shogo-ma/go-logstash.svg?branch=master)](https://travis-ci.org/shogo-ma/go-logstash)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](license)

A library for call HTTP Request to logstash

## Installation
Just execute the following commands

```
$ go get -u github.com/shogo-ma/go-logstash
```

## Example

```
import (
    "context"
    "fmt"

    "github.com/shogo-ma/go-logstash"
)                                                                                                                                                                                                                  
func main() {
    client, err := logstash.NewClient("http://localhost:9600")
    if err != nil { 
        panic(err) 
    }             

    ctx := context.Background()
    nos, err := client.NodePipelineStatsInfo().Do(ctx) 

    fmt.Printf("%+v\n", nos)
}                                 
```

## LICENSE
The content of this repository are licensed under the MIT License unless otherwise noted. Please read [LICENSE][license] for the detail.

[license]: LICENSE.txt
