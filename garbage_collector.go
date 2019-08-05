package neocortex

import "time"

type garbageCollector struct {
	tickTime        time.Duration
	maxLastResponse time.Duration
}

func defaultGarbageCollector(maxSessiontime time.Duration) *garbageCollector {
	return &garbageCollector{
		tickTime:        1 * time.Second,
		maxLastResponse: maxSessiontime,
	}
}

func (e *Engine) runGarbageCollector(g *garbageCollector) {
	ticker := time.NewTicker(g.tickTime)
	go func() {
		for t := range ticker.C {
			for c, diag := range e.ActiveDialogs {
				if t.Sub(diag.LastActivity) > g.maxLastResponse {
					e.onContextIsDone(c)
				}
			}
		}
	}()
}
