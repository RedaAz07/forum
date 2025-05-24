package middleware

import (
	"net/http"
	"strings"
	"time"

	"forum/helpers"
	"forum/utils"
)

var LoginRateLimits = make(map[string]*RateLimitLogin)

func CheckRateLimitLogin(ratelimit *RateLimitLogin, window time.Duration) bool {
	if time.Now().Before(ratelimit.BlockedUntil) {
		return false
	}
	if time.Now().After(ratelimit.BlockedUntil) && ratelimit.count > 10 {
		ratelimit.FirstTime = time.Now()
		ratelimit.BlockedUntil = time.Time{}
		ratelimit.count = 0
	}
	ratelimit.count++
	if ratelimit.count > 10 {
		ratelimit.BlockedUntil = time.Now().Add(window)
		return false
	}
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

		if !CheckRateLimitLogin(ratelimit, 1*time.Minute) {
			helpers.RanderTemplate(w, "statusPage.html", http.StatusTooManyRequests, utils.ErrorToManyRequests)
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
