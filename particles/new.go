package particles

import (
	"container/list"
	"project-particles/config"
	"math/rand"
	"time"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {

	l := list.New()
	rand.Seed(time.Now().UnixNano())

	for i:=0; i < config.General.InitNumParticles; i++{
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
		l.PushFront(&newParticule)
	}
	return System{Content: l}
}

