package particles

import (
	"project-particles/config"
	"math/rand"
)
func (s *System) newParticle() {

	//newParticule comporte les caractéristique d'une particule tels que sa position, sa couleur, sa taille afin d'afficher correctement les caractéristiques demandées sur l'interface.
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

	newParticule.SetSpawn()


	
	s.Content.PushFront(&newParticule) // On insère cette nouvelle particule dans la liste du système afin que celui-ci puisse l'afficher.
}




func (p *Particle) UpdateOpacity(){

	if config.General.OpacityLifetime && config.General.Lifetime>0{
		p.Opacity-= p.Base_Opacity/config.General.Lifetime
	}else{
		p.Opacity+= config.General.ChangeOpacity
	}
	
}


//crée une valeur random de spawn , on crée les bords de notre zone de spawn. Si les bords de notre zone de spawn dépasse les bornes de la killzone alors on les rétrécit pour les rentrer dedans.
func (p *Particle) SetSpawn(){

	//Si le TypeGenerateur est -1 depuis config.json alors chaque particule aura une position différente sur les axes X et Y comprise entre 0 et la valeur défini dans WindowSizeX et WindowSizeY
	if config.General.TypeGenerateur == -1{
		p.PositionX = RandomBetweenFloat(-config.General.Kill_particule_WindowSizeX,float64(config.General.WindowSizeX)+config.General.Kill_particule_WindowSizeX)
		p.PositionY = RandomBetweenFloat(-config.General.Kill_particule_WindowSizeY,float64(config.General.WindowSizeY)+config.General.Kill_particule_WindowSizeY)
	}

	if config.General.TypeGenerateur == 1{
		haut := float64(config.General.SpawnY)-config.General.RayonSpawnY
		bas := float64(config.General.SpawnY)+config.General.RayonSpawnY
		gauche :=float64(config.General.SpawnX)-config.General.RayonSpawnX
		droite := float64(config.General.SpawnX)+config.General.RayonSpawnX
		
	
		if gauche < -config.General.Kill_particule_WindowSizeX{
			gauche = -config.General.Kill_particule_WindowSizeX
		}
		if droite> float64(config.General.WindowSizeX)+config.General.Kill_particule_WindowSizeX{
			droite = float64(config.General.WindowSizeX)+config.General.Kill_particule_WindowSizeX
		}
		if haut < -config.General.Kill_particule_WindowSizeY{
			haut = -config.General.Kill_particule_WindowSizeY
		}
		if bas > float64(config.General.WindowSizeY)+config.General.Kill_particule_WindowSizeY{
			bas = float64(config.General.WindowSizeY)+config.General.Kill_particule_WindowSizeY
		}
		p.PositionX =  RandomBetweenFloat(gauche,droite)
		p.PositionY =  RandomBetweenFloat(haut,bas)
	}

	if config.General.TypeGenerateur == 2{
		vecx,vecy:=RandomVecteur(RandomBetweenFloat(0,config.General.RayonSpawnX))
		p.PositionX+=vecx
		p.PositionY+=vecy
	}

	if config.General.TypeGenerateur == 3{
		vecx:=RandomBetweenFloat(-config.General.RayonSpawnX,config.General.RayonSpawnX)
		vecy:=config.General.RayonSpawnX-vecx

		p.PositionX+=vecx
		p.PositionY+=vecy

	}
}




