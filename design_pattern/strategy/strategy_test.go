package strategy

import "testing"

func TestNewLoggerManager(t *testing.T) {
	f := &FileLogger{}
	manager := NewLoggerManager(f)
	manager.Info("gfy")
	manager.Error("gfy")

	db := &DbLogger{}
	manager = NewLoggerManager(db)
	manager.Info("gfy")
	manager.Error("gfy")
}
