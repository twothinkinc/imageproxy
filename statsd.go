package imageproxy

import (
	"time"
)

// TODO: add tags to statsd metrics
func (p *Proxy) StatsdIncrement(name string) {
	if p.StatsdClient != nil {
		p.StatsdClient.Incr(name, p.StatsdTags, 1)
	}
}

func (p *Proxy) StatsdTiming(name string, start time.Time, rate float64) {
	duration := time.Since(start)

	if p.StatsdClient != nil {
		p.StatsdClient.Timing(name, duration, p.StatsdTags, rate)
	}
}
