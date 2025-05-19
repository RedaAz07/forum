package middleware

import "time"

func CheckRateLimit(ratelimit *RateLimitPosts, window time.Duration) bool {
	//Posts' limit
	if time.Now().Before(ratelimit.BlockedUntil) {
		return false
	}
	if time.Since(ratelimit.FirstTime) < window { //check ila dazt sa3a 3la awl post. resetiw lhssab
		ratelimit.count = 0
		ratelimit.FirstTime = time.Now()
		return true
	}
	if ratelimit.count >= 10 {
		ratelimit.BlockedUntil = time.Now().Add(window)
		return false // block user from posting
	}
	ratelimit.count++
	return true
}


