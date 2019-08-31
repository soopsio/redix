// Copyright 2018 The Redix Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.
package redix

import (
	"fmt"
	"strconv"

	"github.com/alash3al/go-color"
)

func StartRedix() error {
	//fmt.Println(color.MagentaString(redixBrand))
	fmt.Printf("⇨ redix server version: %s \n", color.GreenString(redixVersion))
	fmt.Printf("⇨ redix selected engine: %s \n", color.GreenString(Engine))
	fmt.Printf("⇨ redix workers count: %s \n", color.GreenString(strconv.Itoa(Workers)))
	fmt.Printf("⇨ redix resp server available at: %s \n", color.GreenString(RESPListenAddr))
	fmt.Printf("⇨ redix http server available at: %s \n", color.GreenString(HTTPListenAddr))

	err := make(chan error)

	go (func() {
		err <- initRespServer()
	})()

	go (func() {
		err <- initHTTPServer()
	})()

	e := <-err
	color.Red(e.Error())
	return e
}
