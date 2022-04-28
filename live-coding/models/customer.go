package models

import "time"

type Customer struct {
	Country   string      `json:"-"`
	Createdat time.Time   `json:"CreatedAt"`
	Email     interface{} `json:"Email"`
	ID        string      `json:"ID"`
	Isactive  bool        `json:"IsActive"`
	Name      string      `json:"-"`
	Updatedat time.Time   `json:"-"`
}

func (c *Customer) RefactorEmail() {
	if emailStr, isString := c.Email.(string); isString {
		c.Email = []string{emailStr}
	}
}
