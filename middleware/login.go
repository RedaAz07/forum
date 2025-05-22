package middleware

import (
	"net"
	"net/http"
	"strconv"
	"time"
)

var LoginRateLimits = make(map[int]*RateLimitLogin)

func CheckRateLimitLogin(ratelimit *RateLimitLogin, window time.Duration) bool {
	//Likess' limit
	if time.Now().Before(ratelimit.BlockedUntil) {
		return false
	}
	if ratelimit.count >= 3 {
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
		UserIP:       -1,
	}
	userIP := GetUserIP(r)
	if userIP == -1 {
		return rateLimit, false
	}
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

func GetUserIP(r *http.Request) int {
	var userIP int
	// Addr := r.RemoteAddr //l output kikoun haka : 127.0.0.1:12345 ghaliban
	temp, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return -1
	}
	userIP, err1 := strconv.Atoi(temp)
	if err1 != nil {
		return -1
	}
	return userIP
}

func AddUserToTheMap_Login(ratelimit *RateLimitLogin) {
	LoginRateLimits[ratelimit.UserIP] = ratelimit
}
