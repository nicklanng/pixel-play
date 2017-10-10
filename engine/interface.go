package engine

import (
	"time"
)

type Component interface {
	Update(time.Duration)
	Draw()
}
