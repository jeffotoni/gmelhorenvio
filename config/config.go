package config

import (
	"time"

	"github.com/patrickmn/go-cache"
)

const DAY = time.Hour * 24

var Cache = cache.New(24*27*time.Hour, 24*28*time.Hour)
