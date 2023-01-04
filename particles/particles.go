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
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity: 1,
		VitesseX: float64(rand.Float64()*config.General.MaxVitesseX-config.General.MaxVitesseX/2),
		VitesseY: float64((rand.Float64()*config.General.MaxVitesseY)-(config.General.MaxVitesseY/2)),
	}


	if config.General.RandomSpawn == true{
		newParticule.PositionX = float64(rand.Intn(config.General.WindowSizeX))
		newParticule.PositionY = float64(rand.Intn(config.General.WindowSizeY))
	}
	s.Content.PushFront(&newParticule)
}