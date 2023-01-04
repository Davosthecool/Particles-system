package particles

import (
	"project-particles/config"
	"math/rand"
)
func (s *System) newParticle() {
	newParticule := Particle{
		PositionX: float64(config.General.SpawnX),
		PositionY: float64(config.General.SpawnY),
		ScaleX:    1, ScaleY: 1,
		ColorRed: float64(rand.Float64()), ColorGreen: float64(rand.Float64()), ColorBlue: float64(rand.Float64()),
		Opacity: 1,
		VitesseX: float64(rand.Float64()*config.General.MaxVitesseX-config.General.MaxVitesseX/2),
		VitesseY: float64(rand.Float64()*config.General.MaxVitesseY-config.General.MaxVitesseY/2),
		Lifetime : config.General.Lifetime,
	}


	if config.General.RandomSpawn{
		newParticule.PositionX = float64(rand.Intn(config.General.WindowSizeX))
		newParticule.PositionY = float64(rand.Intn(config.General.WindowSizeY))
	}
	s.Content.PushFront(&newParticule)
}