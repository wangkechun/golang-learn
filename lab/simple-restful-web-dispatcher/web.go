// from https://gist.github.com/shaunlee/8895120
package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var _ = fmt.Print
var _ = log.Print
var _ = http.DefaultMaxHeaderBytes
var _ = regexp.Compile
var _ = strings.Compare

const defaultMaxMemory = 32 << 20
