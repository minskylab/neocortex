package neocortex

// SetPerformanceMetric links a custom performance metric function
func (engine *Engine) SetPerformanceMetric(perf func(*Dialog) float64) {
	engine.dialogPerformanceFunc = perf
}
