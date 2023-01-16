package particles

import (
	"project-particles/config"
	"github.com/hajimehoshi/ebiten/v2"
)
// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.


func (s *System) Update() {
	if config.General.FollowMouseSpawn{
		config.General.SpawnX,config.General.SpawnY= ebiten.CursorPosition()
	}


	posx,posy :=ebiten.CursorPosition()
	oldspawnx,oldspawny:=config.General.SpawnX,config.General.SpawnY
	if config.General.ClickMouseParticules && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && !OutOfKillScreen(float64(posx),float64(posy)){
		for i:=0.0;i<config.General.ClickSpawnRate;i++{
			config.General.SpawnX,config.General.SpawnY= ebiten.CursorPosition()
			//s.newParticle()
		}
		config.General.SpawnX,config.General.SpawnY=oldspawnx,oldspawny
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) && !OutOfKillScreen(float64(posx),float64(posy)){
		config.General.SpawnX,config.General.SpawnY= ebiten.CursorPosition()
	}

	// if config.General.Explosion_Spawn && ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight){
		
		
	// }



	//Chaque particule du système est transférée dans une variable "particule_individuelle".
	//Dans celle-ci, nous ajoutons une vitesse aléatoire (voir particle.go) à la position de la particule située dans "particule_individuelle".
	//Sachant que la vitesse est aléatoire, chaque particule comportera une position différente en fonction de la valeur de leur vitesseX et VitesseY.
	for element := s.Content.Front() ; element != nil; {

		p, _ := element.Value.(*Particle)

		//ccccc
		if config.General.Bounce{
			p.WallBounce()
		}
		p.PositionX += p.VitesseX
		p.PositionY += p.VitesseY


		p.VitesseX+= config.General.GravityX 
		p.VitesseY+= config.General.GravityY

		p.Lifetime--

		p.UpdateOpacity()
		p.Rotation += 0.01*config.General.RotationParticule

		next := element.Next()

		if OutOfScreen(element.Value.(*Particle)) || p.Lifetime <= 0 && config.General.Lifetime>0 || p.Opacity<=0{
			s.Content.Remove(element)
		}

		// if config.General.EffectExplosion && ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight){
		// 	posx,posy :=ebiten.CursorPosition()
		// 	ecart_x := float64(posx) - p.PositionX
		// 	ecart_y := float64(posy) - p.PositionY

		// 	if Abs(ecart_x + ecart_y) < 100 {
		// 		// p.VitesseX = ecart_x*(p.VitesseX/ecart_x)
		// 		// p.VitesseY = ecart_y*(p.VitesseY/ecart_y)
		// 		switch{
		// 		case Abs(ecart_x) <= 100:
		// 			p.VitesseX*=-1
		// 		case Abs(ecart_y) <= 100:
		// 			p.VitesseY*=-1	
		// 		}
		// 	}

		// }
	
		element = next
	}



	
	//Ce bout de code permet de générer un nombre de particules, défini par SpawRate dans config.json, par seconde.
	//La boucle permet ainsi de créer ce nombre de particule grâce à la méthode newParticle (présente dans particle.go)
	// **Mettre ligne pour OutofKillScreen
	s.SpawnRate+=config.General.SpawnRate
	if !OutOfKillScreen(float64(config.General.SpawnX),float64(config.General.SpawnY)){
		for i:=0; i< int(s.SpawnRate); i++{
			s.newParticle()
		}
	}
	s.SpawnRate-=float64(int(s.SpawnRate))

	
}


