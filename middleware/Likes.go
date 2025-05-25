package middleware

import (
	"net/http"
	"time"

	"forum/helpers"
	"forum/utils"
)

var LikesRateLimits = make(map[int]*RateLimitLikes)

func CheckRateLimitLikes(ratelimit *RateLimitLikes, window time.Duration) bool {
	if time.Now().Before(ratelimit.BlockedUntil) {
		return false
	}
	if time.Now().After(ratelimit.BlockedUntil) && ratelimit.count > 100 {
		ratelimit.FirstTime = time.Now()
		ratelimit.BlockedUntil = time.Time{}
		ratelimit.count = 0
	}

	ratelimit.count++
	if ratelimit.count > 100 {
		ratelimit.BlockedUntil = time.Now().Add(window)
		return false
	}
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
		userRateLimit, _ := UserInfosLikes(r)

		ratelimit, exists := LikesRateLimits[userRateLimit.UserId]
		if !exists {

			AddUserToTheMap_Likes(userRateLimit)
			ratelimit = userRateLimit
		}

		if !CheckRateLimitLikes(ratelimit, 1*time.Minute) {
			helpers.RanderTemplate(w, "statusPage.html", http.StatusTooManyRequests, utils.ErrorToManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func AddUserToTheMap_Likes(ratelimit *RateLimitLikes) {
	LikesRateLimits[ratelimit.UserId] = ratelimit
}
