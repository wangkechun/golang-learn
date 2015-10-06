package main

import (
	"fmt"
	"strings"

	conf "github.com/qiniu/api.v6/conf"
	io "github.com/qiniu/api.v6/io"
	rs "github.com/qiniu/api.v6/rs"
)

func init() {
	conf.ACCESS_KEY = "4TDk-3CBZb9pxkxTiqKsFmatQ4wj7_H_cIHI15eR"
	conf.SECRET_KEY = "BzQEIIO5yA5LuIHcWjWpCV90lxcokwAi420oQsRE"

}

func main() {
	c := (&rs.PutPolicy{Scope: "test"}).Token(nil)
	fmt.Println(c)
	var ret io.PutRet
	s := "fdsfdsgdfsg"
	for i := 0; i < 100; i++ {
		s = s + s
		io.PutWithoutKey(nil, &ret, c, strings.NewReader(s), nil)
		fmt.Println(len(s), ret)
	}

}
