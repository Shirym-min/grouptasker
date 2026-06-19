package config

import "fmt"

func (c *Config) DeleteTask(taskName string) error {
	if _, exists := c.Tasks[taskName]; !exists {
		return fmt.Errorf("task \"%s\" not found", taskName)
	}
	delete(c.Tasks, taskName)
	return Save(c)
}
