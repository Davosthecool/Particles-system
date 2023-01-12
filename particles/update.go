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

	for element := s.Content.Front() ; element != nil; {

		p, ok := element.Value.(*Particle)

		if !ok {
			continue
		}

		if p.Alive==false{
			element = element.Next()
			continue
		}
		particule_individuelle := p

		//ccccc
		particule_individuelle.PositionX += particule_individuelle.VitesseX
		particule_individuelle.PositionY += particule_individuelle.VitesseY


		particule_individuelle.VitesseX+= config.General.GravityX 
		particule_individuelle.VitesseY+= config.General.GravityY


		
		particule_individuelle.Lifetime--

		particule_individuelle.UpdateOpacity()

		next := element.Next()

		if OutOfScreen(element.Value.(*Particle)) || particule_individuelle.Lifetime <= 0 && config.General.Lifetime>0 || particule_individuelle.Opacity<=0{
			s.Content.PushBack(element)
			element.Value.(*Particle).Alive=false
		}
		element = next


	}


	/*Partie réservé au SpawnRate */ 
	s.SpawnRate+=config.General.SpawnRate
	for i:=0; i< int(s.SpawnRate); i++{
		s.newParticle()
	}
	s.SpawnRate-=float64(int(s.SpawnRate))

	
}


