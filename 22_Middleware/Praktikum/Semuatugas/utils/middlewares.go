package tulis

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstak/echo/v4"
)

type (
Stats struct {
	Uptime		time.Time		`json:"uptime"`
	RequestCount uint64     	`json: "requestCount"`
	Statuses	map[string]int	`json:"statuses"`
	mutex		sync.RWMutex
	}
)

func NewStats() *Stats {
	return &Stats{
		Uptime: time.Now(),
		Statuses: map[string]int{},
	}
}

// Process is the middleware function.
func (s *Stats) Process(next echo.HandlerFunc) echo.HandlerFunc
	return func(c echo.Context) error{
		if err := next(c); err != nil {
			c.Error(err)
		}
		s.mutex.Lock()
		defer s.mutex.Unlock()
		c.RequestCount()