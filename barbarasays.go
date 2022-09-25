package barbarasays

import (
	"github.com/waltervargas/barbarasays/colombian"
	"github.com/waltervargas/barbarasays/english"
	"github.com/waltervargas/barbarasays/germany"
	"github.com/waltervargas/barbarasays/venezuelan"
)

// knows how to read the failing tests +15pts.
func Say(language string) string {
	if language == "ve" {
		return venezuelan.Dice()
	}
	if language == "co" {
		return colombian.Dice()
	}
	// +5pts
	if language == "de" {
		return germany.Sagt()
	}
	// by default it will return english
	// +5 pts
	return english.Says()
}
