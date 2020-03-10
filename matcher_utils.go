package neocortex

func (m *Matcher) And(and *Matcher) *Matcher {
	m.AND = and
	return m
}

func (m *Matcher) Or(or *Matcher) *Matcher {
	m.OR = or
	return m
}

func IntentIs(intent string, confidence ...float64) *Matcher {
	conf := 0.0
	if len(confidence) > 0 {
		conf = confidence[0]
	}
	return &Matcher{
		Intent: Match{
			Is:         intent,
			Confidence: conf,
		},
	}
}

func (m *Matcher) AndIntentIs(intent string, confidence ...float64) *Matcher {
	return m.And(IntentIs(intent, confidence...))
}

func (m *Matcher) OrIntentIs(intent string, confidence ...float64) *Matcher {
	return m.Or(IntentIs(intent, confidence...))
}

func IfEntityIs(entity string, confidence ...float64) *Matcher {
	conf := 0.0
	if len(confidence) > 0 {
		conf = confidence[0]
	}
	return &Matcher{
		Entity: Match{
			Is:         entity,
			Confidence: conf,
		},
	}
}

func (m *Matcher) AndEntityIs(entity string, confidence ...float64) *Matcher {
	return m.And(IfEntityIs(entity, confidence...))
}

func (m *Matcher) OrEntityIs(entity string, confidence ...float64) *Matcher {
	return m.Or(IfEntityIs(entity, confidence...))
}

func IfContextVariableIs(name string, value interface{}) *Matcher {
	return &Matcher{
		ContextVariable: CMatch{
			Name:  name,
			Value: value,
		},
	}
}

func (m *Matcher) AndIfContextVariableIs(name string, value interface{}) *Matcher {
	return m.And(IfContextVariableIs(name, value))
}

func (m *Matcher) OrIfContextVariableIs(name string, value interface{}) *Matcher {
	return m.Or(IfContextVariableIs(name, value))
}

func IfDialogNodeTitleIs(title string) *Matcher {
	return &Matcher{
		DialogNode: DialogNodeMatch{
			Title: title,
		},
	}
}

func (m *Matcher) AndIfDialogNodeTitleIs(title string) *Matcher {
	return m.And(IfDialogNodeTitleIs(title))
}

func (m *Matcher) OrIfDialogNodeTitleIs(title string) *Matcher {
	return m.Or(IfDialogNodeTitleIs(title))
}

func IfDialogNodeNameIs(name string) *Matcher {
	return &Matcher{
		DialogNode: DialogNodeMatch{
			Name: name,
		},
	}
}

func (m *Matcher) AndIfDialogNodeNameIs(name string) *Matcher {
	return m.And(IfDialogNodeNameIs(name))
}

func (m *Matcher) OrIfDialogNodeNameIs(name string) *Matcher {
	return m.Or(IfDialogNodeNameIs(name))
}

func IfDialogNodeIs(title, name string) *Matcher {
	return &Matcher{
		DialogNode: DialogNodeMatch{
			Title: title,
			Name:  name,
		},
	}
}

func (m *Matcher) AndIfDialogNodeIs(title, name string) *Matcher {
	return m.And(IfDialogNodeIs(title, name))
}

func (m *Matcher) OrIfDialogNodeIs(title, name string) *Matcher {
	return m.Or(IfDialogNodeIs(title, name))
}
