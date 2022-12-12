package moonphase

import (
	"math"
	"time"
)

// Phase is a moon phase.
type Phase uint8

const (
	New Phase = iota
	WaxingCrescent
	FirstQuarter
	WaxingGibbous
	Full
	WaningGibbous
	LastQuarter
	WaningCrescent
)

const lunarCycleSeconds = 2551443 // how long it takes for the moon to go through a full cycle
const firstNewMoon = 947175240    // the unix timestamp of the first new moon in 2000

// FindPhase returns the moon phase for the given time.
func FindPhase(t time.Time) Phase {
	totalSecs := t.Unix() - firstNewMoon
	currentSecs := math.Mod(float64(totalSecs), lunarCycleSeconds)
	if currentSecs < 0 {
		currentSecs += lunarCycleSeconds
	}
	currentFrac := currentSecs / float64(lunarCycleSeconds)
	r := math.Round(currentFrac * 8)
	if r >= 8 {
		r = 0
	}
	return Phase(r)
}
