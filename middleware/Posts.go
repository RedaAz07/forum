package middleware

import (
	"forum/utils"
	"net/http"
	"time"
)

var PostRateLimits = make(map[int]*RateLimitPosts)

func CheckRateLimitPost(ratelimit *RateLimitPosts, window time.Duration) bool {
	if time.Now().Before(ratelimit.BlockedUntil) {
		return false
	}
	if ratelimit.count >= 10 {
		ratelimit.BlockedUntil = time.Now().Add(window)
		return false // block l user bach maypostich
	}
	if time.Since(ratelimit.FirstTime) < window { //check ila dazt sa3a 3la awl post. resetiw lhssab
		ratelimit.count += 1
		ratelimit.FirstTime = time.Now()
		return true
	}

	ratelimit.count++
	return true
}

func UserInfosPosts(r *http.Request) (*RateLimitPosts, bool) {
	rateLimit := &RateLimitPosts{
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

func RateLimitPostsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRateLimit, ok := UserInfosPosts(r)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ratelimit, exists := PostRateLimits[userRateLimit.UserId]
		if !exists {
			AddUserToTheMap_Post(userRateLimit)
			ratelimit = userRateLimit
		}

		if !CheckRateLimitPost(ratelimit,1 * time.Hour) {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func GetUserId(r *http.Request) int {
	var userID int
	cookie, err := r.Cookie("session")
	if err != nil || cookie.Value == "" {
		return -1
	} else {
		// Check if the session is valid
		stmt := "SELECT id FROM users WHERE session = ?"
		err = utils.Db.QueryRow(stmt, cookie.Value).Scan(&userID)
		if err != nil {
			return -1
		}
	}
	return userID
}

func AddUserToTheMap_Post(ratelimit *RateLimitPosts) {
	PostRateLimits[ratelimit.UserId] = ratelimit
}
