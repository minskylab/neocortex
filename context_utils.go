package neocortex

func (c *Context) SetContextVariable(name string, value interface{}) {
	if c.Variables == nil {
		c.Variables = map[string]interface{}{}
	}

	c.Variables[name] = value
}

func (c *Context) DeleteContextVariable(name string, value interface{}) {
	if c.Variables == nil {
		c.Variables = map[string]interface{}{}

	}

	delete(c.Variables, name)

}

func (c *Context) GetStringContextVariable(name string, _default string) string {
	if c.Variables == nil {
		c.Variables = map[string]interface{}{}
		return _default
	}

	value, isOk := c.Variables[name].(string)
	if !isOk {
		return _default
	}

	return value
}

func (c *Context) GetIntContextVariable(name string, _default int) int {
	if c.Variables == nil {
		c.Variables = map[string]interface{}{}
		return _default
	}

	value, isOk := c.Variables[name].(int)
	if !isOk {
		return _default
	}

	return value
}

func (c *Context) GetFloat64ContextVariable(name string, _default float64) float64 {
	if c.Variables == nil {
		c.Variables = map[string]interface{}{}
		return _default
	}

	value, isOk := c.Variables[name].(float64)
	if !isOk {
		return _default
	}

	return value
}
