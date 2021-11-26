package middlewares

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var limiter = newIPRateLimiter(1, 2)

func ReteLimitter(c *gin.Context) {
	lim := limiter.getLimiter(c.Request.RemoteAddr)
	if t := lim.Allow(); !t {
		m := make(map[string]string)
		m["code"] = "429"
		m["message"] = "Too many requests"
		c.Status(429)
		c.AbortWithStatusJSON(http.StatusTooManyRequests, m)
	}
	c.Next()
}

//IPRateLimiter struct limitter
type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

// NewIPRateLimiter baroi limit
func newIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	i := &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}
	return i
}

// AddIP создает новый ограничитель скорости и добавляет его в карту ips,
// используя IP-адрес в качестве ключа
func (i *IPRateLimiter) addIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()
	limiter := rate.NewLimiter(i.r, i.b)
	i.ips[ip] = limiter
	return limiter
}

// GetLimiter возвращает ограничитель для переданного IP-адреса, если тот существует.
// В обратном случае вызывает AddIP, чтобы добавить IP-адрес в карту
func (i *IPRateLimiter) getLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.ips[ip]
	if !exists {
		i.mu.Unlock()
		return i.addIP(ip)
	}
	i.mu.Unlock()
	return limiter
}
