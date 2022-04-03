package model

import "time"

// ApplicationConfig ...
type ApplicationConfig struct {
	AppName            string
	ElasticSearchNodes string
	GRPCTimeout        time.Duration
	CacheExpiry        time.Duration
	CacheCleanup       time.Duration
	MineDBHost         string
	MineDBPort         string
	MineDBUser         string
	MineDBName         string
}
