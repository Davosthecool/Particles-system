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
		particule_individuelle.PositionX += particule_individuelle.VitesseX
		particule_individuelle.PositionY += particule_individuelle.VitesseY


		//Concerne la Gravité
		if config.General.Gravity == true{
			particule_individuelle.VitesseY+= config.General.Gravity_Value
		}
	}
	if config.General.SpawnRateOnOFF{
		for i:=0; i< int(config.General.SpawnRate); i++{
			s.newParticle()
		}
	}
	
}


