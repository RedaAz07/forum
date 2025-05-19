package middleware

import "time"

type RateLimitComments struct {
	count        int
	FirstTime    time.Time
	BlockedUntil time.Time
	UserId       int
}
type RateLimitLikes struct {
	count        int
	FirstTime    time.Time
	BlockedUntil time.Time
	UserId       int
}
type RateLimitPosts struct {
	count        int
	FirstTime    time.Time
	BlockedUntil time.Time
	UserId       int
}


// var PostRateLimits = make(map[int]*RateLimitPosts)