package neocortex

type Match struct {
	Is         string
	Confidence float64
}

type Matcher struct {
	Entity Match
	Intent Match
	AND    *Matcher
	OR     *Matcher
}

func match(output *Output, matcher *Matcher) bool {
	ok := false
	for _, i := range output.Intents {
		if i.Intent == matcher.Intent.Is && i.Confidence > matcher.Intent.Confidence {
			ok = true
		}
	}

	for _, e := range output.Entities {
		if e.Entity == matcher.Entity.Is && e.Confidence > matcher.Entity.Confidence {
			ok = true
		}
	}

	if matcher.AND != nil {
		if match(output, matcher.AND) && ok {
			ok = true
		} else {
			ok = false
		}
	}

	if matcher.OR != nil {
		if match(output, matcher.OR) || ok {
			ok = true
		} else {
			ok = false
		}
	}

	return ok
}
