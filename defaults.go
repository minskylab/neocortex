package neocortex

var defaultPerformance = func(dialog *Dialog) float64 {
	goods := 0
	for _, out := range dialog.Outs {
		valids := 0
		for _, intent := range out.Output.Intents {
			if intent.Confidence > 0.1 {
				valids++
			}
		}
		if valids > 0 {
			goods++
		}
	}

	if totalOuts := float64(len(dialog.Outs)); totalOuts > 0.0 {
		return float64(goods) / totalOuts
	}

	return 0.0
}
