package animations

import (
	"context"
	"time"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/whgentry/gomidi-led/leds"
)

var VelocityBar = &Animation{
	Name:        "Velocity Bars",
	Key:         "velocity-bars",
	Description: "Bars go up corresponding to the velocity of the note",
	Run: func(ctx context.Context) {
		defer wg.Done()
		frameTicker := time.NewTicker(frameDuration)
		for {
			select {
			case <-frameTicker.C:
				for row := range pixels {
					for col, ps := range pixels[row] {
						// Decay Intensity Exponentially
						if kboard.Keys[col].IsNotePressed {
							ps.Intensity = kboard.Keys[col].GetAdjustedVelocityRatio()
						} else {
							ps.Intensity *= 0.95
						}
						// Determine led color on intensity
						if row >= int(ps.Intensity*float64(numRows)) {
							ps.Color = leds.ColorOff()
						} else if kboard.Keys[col].IsNotePressed {
							ps.Color = colorful.Hsv(360*float64(row)/float64(numRows), 1, 1)
						}
					}
				}
			case <-ctx.Done():
				return
			}
		}
	},
}
