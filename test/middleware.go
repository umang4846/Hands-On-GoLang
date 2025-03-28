package main

type RateLimit struct {
	rateLimiter   map[string]int
	ratePerSecond int
	burstTime     int
}

func NewRateLimiter(rate int, burstTime int) *RateLimit {
	return &RateLimit{
		rateLimiter:   make(map[string]int),
		ratePerSecond: rate,
		burstTime:     burstTime,
	}
}

func (r *RateLimit) CheckRateLimit(userid string) bool {
	//check weather should we allow the user or not
	return r.rateLimiter[userid] > r.ratePerSecond
}
