package middleware

import (
	"net/http"
	"strings"
	"time"
)

var LoginRateLimits = make(map[string]*RateLimitLogin)

func CheckRateLimitLogin(ratelimit *RateLimitLogin, window time.Duration) bool {
	if time.Now().Before(ratelimit.BlockedUntil) {
		return false
	}
	if ratelimit.count >= 10 {
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

func UserInfosLogin(r *http.Request) (*RateLimitLogin, bool) {
	rateLimit := &RateLimitLogin{
		count:        0,
		FirstTime:    time.Now(),
		BlockedUntil: time.Time{},
		UserIP:       "",
	}
	userIP := GetUserIP(r)
	rateLimit.UserIP = userIP
	return rateLimit, true
}

func RateLimitLoginMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRateLimit, ok := UserInfosLogin(r)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ratelimit, exists := LoginRateLimits[userRateLimit.UserIP]
		if !exists {
			AddUserToTheMap_Login(userRateLimit)
			ratelimit = userRateLimit
		}

		if !CheckRateLimitLogin(ratelimit, 1*time.Hour) {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func GetUserIP(r *http.Request) string {
	temp := r.RemoteAddr
	userIP := strings.Split(temp, ":")[0]
	return userIP
}

func AddUserToTheMap_Login(ratelimit *RateLimitLogin) {
	LoginRateLimits[ratelimit.UserIP] = ratelimit
}
