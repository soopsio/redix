// Copyright 2018 The Redix Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.
package redix

import (
	"net/url"
	"runtime"
	"sync"

	"github.com/alash3al/go-pubsub"
	"github.com/bwmarrin/snowflake"
)

var (
	RESPListenAddr = ":6380"
	HTTPListenAddr = ":7090"
	StorageDir     = "./redix-data"
	Engine         = "leveldb"
	EngineOptions  = ""
	Workers        = runtime.NumCPU()
	Verbose        = true
	ACK            = true
)

var (
	databases          *sync.Map
	changelog          *pubsub.Broker
	webhooks           *sync.Map
	websockets         *sync.Map
	snowflakeGenerator *snowflake.Node
	kvjobs             chan func()
)

var (
	commands = map[string]CommandHandler{
		// strings
		"set":    setCommand,
		"mset":   msetCommand,
		"get":    getCommand,
		"mget":   mgetCommand,
		"del":    delCommand,
		"exists": existsCommand,
		"incr":   incrCommand,
		"ttl":    ttlCommand,
		"keys":   keysCommand,

		// lists
		"lpush":      lpushCommand,
		"lpushu":     lpushuCommand,
		"lrange":     lrangeCommand,
		"lrem":       lremCommand,
		"lcount":     lcountCommand,
		"lcard":      lcountCommand,
		"lsum":       lsumCommand,
		"lavg":       lavgCommand,
		"lmin":       lminCommand,
		"lmax":       lmaxCommand,
		"lsrch":      lsearchCommand,
		"lsrchcount": lsearchcountCommand,

		// sets (list alias)
		"sadd":     lpushuCommand,
		"smembers": lrangeCommand,
		"srem":     lremCommand,
		"scard":    lcountCommand,
		"sscan":    lrangeCommand,

		// hashes
		"hset":    hsetCommand,
		"hget":    hgetCommand,
		"hdel":    hdelCommand,
		"hgetall": hgetallCommand,
		"hkeys":   hkeysCommand,
		"hmset":   hmsetCommand,
		"hexists": hexistsCommand,
		"hincr":   hincrCommand,
		"httl":    httlCommand,
		"hlen":    hlenCommand,

		// pubsub
		"publish":        publishCommand,
		"subscribe":      subscribeCommand,
		"webhookset":     webhooksetCommand,
		"webhookdel":     webhookdelCommand,
		"websocketopen":  websocketopenCommand,
		"websocketclose": websocketcloseCommand,

		// utils
		"encode":   encodeCommand,
		"uuidv4":   uuid4Command,
		"uniqid":   uniqidCommand,
		"randstr":  randstrCommand,
		"randint":  randintCommand,
		"time":     timeCommand,
		"dbsize":   dbsizeCommand,
		"gc":       gcCommand,
		"info":     infoCommand,
		"echo":     echoCommand,
		"flushdb":  flushdbCommand,
		"flushall": flushallCommand,

		// ratelimit
		"ratelimitset":  ratelimitsetCommand,
		"ratelimittake": ratelimittakeCommand,
		"ratelimitget":  ratelimitgetCommand,
	}
)

var (
	supportedEngines = map[string]bool{
		"badgerdb": true,
		"boltdb":   true,
		"leveldb":  true,
		"null":     true,
		"sqlite":   true,
	}
	engineOptions         = url.Values{}
	defaultPubSubAllTopic = "*"
)

const (
	redixVersion = "1.10"
	redixBrand   = `

		 _______  _______  ______  _________         
		(  ____ )(  ____ \(  __  \ \__   __/|\     /|
		| (    )|| (    \/| (  \  )   ) (   ( \   / )
		| (____)|| (__    | |   ) |   | |    \ (_) / 
		|     __)|  __)   | |   | |   | |     ) _ (  
		| (\ (   | (      | |   ) |   | |    / ( ) \ 
		| ) \ \__| (____/\| (__/  )___) (___( /   \ )
		|/   \__/(_______/(______/ \_______/|/     \|

A high-concurrency standalone NoSQL datastore with the support for redis protocol 
and multiple backends/engines, also there is a native support for
real-time apps via webhook & websockets besides the basic redis channels.

	`
)
