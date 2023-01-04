package particles

import (
	"project-particles/config"
)
// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {

	for element := s.Content.Front() ; element != nil ; element = element.Next(){
		particule_individuelle := element.Value.(*Particle)

		//ccccc
		particule_individuelle.PositionX += particule_individuelle.VitesseX
		particule_individuelle.PositionY += particule_individuelle.VitesseY


		particule_individuelle.VitesseX+= config.General.GravityX 
		particule_individuelle.VitesseY+= config.General.GravityY


		
		particule_individuelle.Lifetime--

		if config.General.Lifetime>0{
			particule_individuelle.Opacity-= (1.0/config.General.Lifetime)
		}


		if particule_individuelle.PositionX > float64(config.General.WindowSizeX)+config.General.Kill_particule_WindowSizeX-10 || particule_individuelle.PositionX < -config.General.Kill_particule_WindowSizeX|| particule_individuelle.PositionY > float64(config.General.WindowSizeY)+config.General.Kill_particule_WindowSizeY-10 || particule_individuelle.PositionY < -config.General.Kill_particule_WindowSizeY{
			s.Content.Remove(element)
		}

		if (particule_individuelle.Lifetime <= 0 && config.General.Lifetime>0) || particule_individuelle.Opacity<=0{
			s.Content.Remove(element)
		}


	}


	/*Partie réservé au SpawnRate */ 
	s.SpawnRate+=config.General.SpawnRate
	for i:=0; i< int(s.SpawnRate); i++{
		s.newParticle()
	}
	s.SpawnRate-=float64(int(s.SpawnRate))

	
}


