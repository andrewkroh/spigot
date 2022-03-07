package asa

import "fmt"

type config struct {
	Type             string `config:"type" validate:"required"`
	IncludeTimestamp bool   `config:"include_timestamp"`
}

func defaultConfig() config {
	return config{
		Type: Name,
	}
}

func (c *config) Validate() error {
	if c.Type != Name {
		return fmt.Errorf("'%s' is not a valid value for 'type' expected '%s'", c.Type, Name)
	}
	return nil
}