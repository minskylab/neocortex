package neocortex

type Match struct {
	Is         string
	Confidence float64
}

type CMatch struct {
	Name  string
	Value interface{}
}

type DialogNodeMatch struct {
	Title string
	Name  string
}

type Matcher struct {
	DialogNode      DialogNodeMatch
	Entity          Match
	Intent          Match
	ContextVariable CMatch
	AND             *Matcher
	OR              *Matcher
}

func (out *Output) Match(matcher *Matcher) bool {
	ok := false
	for _, i := range out.Intents {
		if i.Intent == matcher.Intent.Is && i.Confidence > matcher.Intent.Confidence {
			ok = true
		}
	}

	for _, e := range out.Entities {
		if e.Entity == matcher.Entity.Is && e.Confidence > matcher.Entity.Confidence {
			ok = true
		}
	}

	if out.Context.Variables != nil {
		for varName, varValue := range out.Context.Variables {
			if matcher.ContextVariable.Name == varName {
				if matcher.ContextVariable.Value == varValue {
					ok = true
				}
			}
		}
	}

	if matcher.DialogNode.Title != "" || matcher.DialogNode.Name != "" {
		for _, n := range out.VisitedNodes {
			if matcher.DialogNode.Name != "" {
				if matcher.DialogNode.Title != "" {
					if n.Name == matcher.DialogNode.Name && n.Title == matcher.DialogNode.Title {
						ok = true
					}
				}

				if n.Name == matcher.DialogNode.Name {
					ok = true
				}
			} else if matcher.DialogNode.Title != "" {
				if n.Title == matcher.DialogNode.Title {
					ok = true
				}
			}
		}
	}

	if matcher.AND != nil {
		if out.Match(matcher.AND) && ok {
			ok = true
		} else {
			ok = false
		}
	}

	if matcher.OR != nil {
		if out.Match(matcher.OR) || ok {
			ok = true
		} else {
			ok = false
		}
	}

	return ok
}

func (in *Input) Match(matcher *Matcher) bool {
	ok := false
	for _, i := range in.Intents {
		if i.Intent == matcher.Intent.Is && i.Confidence > matcher.Intent.Confidence {
			ok = true
		}
	}

	for _, e := range in.Entities {
		if e.Entity == matcher.Entity.Is && e.Confidence > matcher.Entity.Confidence {
			ok = true
		}
	}

	if in.Context.Variables != nil {
		for varName, varValue := range in.Context.Variables {
			if matcher.ContextVariable.Name == varName {
				if matcher.ContextVariable.Value == varValue {
					ok = true
				}
			}
		}
	}

	if matcher.AND != nil {
		if in.Match(matcher.AND) && ok {
			ok = true
		} else {
			ok = false
		}
	}

	if matcher.OR != nil {
		if in.Match(matcher.OR) || ok {
			ok = true
		} else {
			ok = false
		}
	}

	return ok

}
