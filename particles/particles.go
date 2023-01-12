package particles

import (
	"project-particles/config"
	"math/rand"
)
func (s *System) newParticle() {
	newParticule := Particle{
		PositionX: float64(config.General.SpawnX),
		PositionY: float64(config.General.SpawnY),
		ScaleX: RandomBetweenFloat(config.General.MinScaleX, config.General.MaxScaleX), 
		ScaleY: RandomBetweenFloat(config.General.MinScaleY, config.General.MaxScaleY),
		ColorRed: RandomBetweenFloat(config.General.MinColorRed, config.General.MaxColorRed),
		ColorGreen: RandomBetweenFloat(config.General.MinColorGreen, config.General.MaxColorGreen),
		ColorBlue: RandomBetweenFloat(config.General.MinColorBlue, config.General.MaxColorBlue),
		Opacity: rand.Float64(),
		Base_Opacity : 0,
		VitesseX: float64(rand.Float64()*config.General.MaxVitesseX-config.General.MaxVitesseX/2),
		VitesseY: float64(rand.Float64()*config.General.MaxVitesseY-config.General.MaxVitesseY/2),
		Lifetime : config.General.Lifetime,
		Alive : true,
	}
	newParticule.Base_Opacity = newParticule.Opacity

	if config.General.RandomSpawn{
		newParticule.PositionX = float64(rand.Intn(config.General.WindowSizeX))
		newParticule.PositionY = float64(rand.Intn(config.General.WindowSizeY))
	}
	s.Content.PushFront(&newParticule)
}


func (p *Particle) UpdateOpacity(){

	if config.General.OpacityLifetime && config.General.Lifetime>0{
		p.Opacity-= p.Base_Opacity/config.General.Lifetime
	}else{
		p.Opacity+= config.General.ChangeOpacity
	}
	
}
