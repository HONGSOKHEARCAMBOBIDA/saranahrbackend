package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"

	ginlimiter "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memory "github.com/ulule/limiter/v3/drivers/store/memory"
)

func RateLimiterMiddleware() gin.HandlerFunc {
	rate, err := limiter.NewRateFromFormatted("5-M")
	if err != nil {
		panic(err)
	}

	store := memory.NewStore()

	middleware := ginlimiter.NewMiddleware(limiter.New(store, rate), ginlimiter.WithLimitReachedHandler(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message": "too many requests",
		})
	}))

	return middleware
}
