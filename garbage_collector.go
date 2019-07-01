package neocortex

import "time"

type garbageCollector struct {
	tickTime        time.Duration
	maxLastResponse time.Duration
}

func defaultGarbageCollector() *garbageCollector {
	return &garbageCollector{
		tickTime:        1 * time.Second,
		maxLastResponse: 10 * time.Minute,
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
