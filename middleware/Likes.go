package middleware

import (
	"net/http"
	"time"
)

var LikesRateLimits = make(map[int]*RateLimitLikes)

func CheckRateLimitLikes(ratelimit *RateLimitLikes, window time.Duration) bool {
	if time.Now().Before(ratelimit.BlockedUntil) {
		return false
	}
	if ratelimit.count >= 2 {
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

func UserInfosLikes(r *http.Request) (*RateLimitLikes, bool) {
	rateLimit := &RateLimitLikes{
		count:        0,
		FirstTime:    time.Now(),
		BlockedUntil: time.Time{},
		UserId:       -1,
	}
	userID := GetUserId(r)
	if userID == -1 {
		return rateLimit, false
	}
	rateLimit.UserId = userID
	return rateLimit, true
}

func RateLimitLikesMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRateLimit, ok := UserInfosLikes(r)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ratelimit, exists := LikesRateLimits[userRateLimit.UserId]
		if !exists {

			AddUserToTheMap_Likes(userRateLimit)
			ratelimit = userRateLimit
		}

		if !CheckRateLimitLikes(ratelimit, 1*time.Hour) {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func AddUserToTheMap_Likes(ratelimit *RateLimitLikes) {
	LikesRateLimits[ratelimit.UserId] = ratelimit
}
