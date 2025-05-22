package middleware

import (
	"net/http"
	"time"
)

var CommentRateLimits = make(map[int]*RateLimitComments)

func CheckRateLimitComment(ratelimit *RateLimitComments, window time.Duration) bool {
	if time.Now().Before(ratelimit.BlockedUntil) {
		return false
	}
	if ratelimit.count >= 50 {
		ratelimit.BlockedUntil = time.Now().Add(window)
		return false 
	}
	if time.Since(ratelimit.FirstTime) < window { 
		ratelimit.count += 1
		ratelimit.FirstTime = time.Now()
		return true
	}

	ratelimit.count++
	return true
}

func UserInfosComments(r *http.Request) (*RateLimitComments, bool) {
	rateLimit := &RateLimitComments{
		count:        0,
		FirstTime:    time.Now(),
		BlockedUntil: time.Time{},
		UserId:       -1,
	}
	UserID := GetUserId(r)
	if UserID == -1 {
		return rateLimit, false
	}
	rateLimit.UserId = UserID
	return rateLimit, true
}

func RateLimitCommentsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRateLimit, ok := UserInfosComments(r)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ratelimit, exists := CommentRateLimits[userRateLimit.UserId]
		if !exists {
			AddUserToTheMap_comment(userRateLimit)
			ratelimit = userRateLimit
		}

		if !CheckRateLimitComment(ratelimit, 1 * time.Minute) {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func AddUserToTheMap_comment(ratelimit *RateLimitComments) {
	CommentRateLimits[ratelimit.UserId] = ratelimit
}
