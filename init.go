// Copyright 2018 The Redix Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.
package redix

import (
	"fmt"
	"net/url"
	"os"
	"runtime"
	"sync"

	"github.com/alash3al/go-color"
	"github.com/alash3al/go-pubsub"
	"github.com/bwmarrin/snowflake"
)

func init() {
	runtime.GOMAXPROCS(Workers)

	if !supportedEngines[Engine] {
		fmt.Println(color.RedString("Invalid strorage engine specified"))
		os.Exit(0)
		return
	}

	databases = new(sync.Map)
	changelog = pubsub.NewBroker()
	webhooks = new(sync.Map)
	websockets = new(sync.Map)
	engineOptions = (func() url.Values {
		opts, _ := url.ParseQuery(EngineOptions)
		return opts
	})()

	snowflakenode, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(color.RedString(err.Error()))
		os.Exit(0)
		return
	}

	snowflakeGenerator = snowflakenode

	// initDBs()
}

// // initDBs - initialize databases from the disk for faster access
// func initDBs() {
// 	os.MkdirAll(StorageDir, 0755)

// 	dirs, _ := ioutil.ReadDir(get)

// 	for _, f := range dirs {
// 		if !f.IsDir() {
// 			continue
// 		}

// 		name := filepath.Base(f.Name())

// 		_, err := selectDB(name)
// 		if err != nil {
// 			log.Println(color.RedString(err.Error()))
// 			continue
// 		}
// 	}
// }
