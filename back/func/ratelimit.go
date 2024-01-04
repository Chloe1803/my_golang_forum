package forum

import "time"

type RateLimiter struct {
	Requests     chan time.Time
	Rate         int
	Burst        int
	RequestLimit int
}

func NewRateLimiter(rate, burst, requestLimit int) *RateLimiter {
	rl := &RateLimiter{
		Requests:     make(chan time.Time, burst),
		Rate:         rate,
		Burst:        burst,
		RequestLimit: requestLimit,
	}
	// Lancement d'une goroutine pour gérer le rate limiting
	go rl.manageRateLimit()
	return rl
}
func (rl *RateLimiter) manageRateLimit() {
	ticker := time.NewTicker(time.Second / time.Duration(rl.Rate))
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			// Réinitialisez le canal pour suivre les demandes dans la fenêtre de taux
			rl.resetRequests()
		case <-rl.Requests:
			// Traitement de la demande
		}
	}
}
func (rl *RateLimiter) resetRequests() {
	for len(rl.Requests) > 0 {
		<-rl.Requests
	}
}
func (rl *RateLimiter) IsAllowed() bool {
	if len(rl.Requests) < rl.RequestLimit {
		rl.Requests <- time.Now()
		return true
	}
	return false
}
