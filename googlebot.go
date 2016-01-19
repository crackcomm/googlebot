package googlebot

import (
	"net"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
)

var cacheIPs *cache.Cache

func init() {
	cacheIPs = cache.New(time.Second*360, time.Second*30)
}

// IsGoogleBot - Checks if given IP address is really Google Bot IP address.
func IsGoogleBot(addr string) (yes bool, err error) {
	addrs, err := net.LookupAddr(addr)
	if err != nil || len(addrs) == 0 {
		return
	}
	yes = strings.HasSuffix(addrs[0], ".googlebot.com.")
	return
}

// IsGoogleBotWithCache - Checks if given IP address is really Google Bot IP.
// Uses in-memory cache.
func IsGoogleBotWithCache(addr string) (yes bool, err error) {
	if v, ok := cacheIPs.Get(addr); ok {
		yes, _ = v.(bool)
		return
	}
	yes, err = IsGoogleBot(addr)
	if err != nil {
		return
	}
	cacheIPs.Set(addr, yes, cache.DefaultExpiration)
	return
}
