package 限流器

import (
	"go-utils/限流器/limiter"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "logistics",
		FillInterval: time.Second,
		Capacity:     100,
		Quantum:      100,
	},
)

r := gin.Default()
r.Use(middleware.RateLimiter(methodLimiters))