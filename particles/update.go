package particles

import (
	"project-particles/config"
	"math/rand"
)
// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {
	for element := s.Content.Front() ; element != nil ; element = element.Next(){
		particule_individuelle := element.Value.(*Particle)
		particule_individuelle.PositionX += particule_individuelle.VitesseX
		particule_individuelle.PositionY += particule_individuelle.VitesseY
	}



	
	for i:=0; i< int(config.General.SpawnRate); i++{
		newParticule := Particle{
			PositionX: float64(config.General.SpawnX),
			PositionY: float64(config.General.SpawnY),
			ScaleX:    1, ScaleY: 1,
			ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
			Opacity: 1,
			VitesseX: float64((rand.Float64()*config.General.MaxVitesseX)-(config.General.MaxVitesseX/2)),
			VitesseY: float64((rand.Float64()*config.General.MaxVitesseY)-(config.General.MaxVitesseY/2)),
		}

		if config.General.RandomSpawn == true{
			newParticule.PositionX = float64(rand.Intn(config.General.WindowSizeX))
			newParticule.PositionY = float64(rand.Intn(config.General.WindowSizeY))
		} 
		s.Content.PushFront(&newParticule)
	}
}


