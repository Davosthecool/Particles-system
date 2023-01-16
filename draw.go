package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particles"
)

// Draw se charge d'afficher à l'écran l'état actuel du système de particules
// g.system. Elle est appelée automatiquement environ 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction pourra être légèrement modifiée quand
// c'est précisé dans le sujet.
func (g *game) Draw(screen *ebiten.Image) {
	for e := g.system.Content.Front(); e != nil; e = e.Next() {
		p, ok := e.Value.(*particles.Particle)
		if ok{
			options := ebiten.DrawImageOptions{}
			options.GeoM.Rotate(p.Rotation)
			options.GeoM.Scale(p.ScaleX, p.ScaleY)
			options.GeoM.Translate(p.PositionX, p.PositionY)
			options.ColorM.Scale(p.ColorRed, p.ColorGreen, p.ColorBlue, p.Opacity)
			screen.DrawImage(assets.ParticleImage, &options)
		}
	}
	// var ListeLignes []LineDraw

	// ListeLignes = append(ListeLignes,LineDraw{1,"WindowSizeX :",&config.General.WindowSizeX})
	// ListeLignes = append(ListeLignes,LineDraw{2,"WindowSizeY :",&config.General.WindowSizeY})
	// var LineY int
	// if config.General.Debug {
	// ebitenutil.DebugPrintAt(screen,fmt.Sprint("Nombre de TPS  : ", ebiten.ActualTPS()), config.General.WindowSizeX-250, 0)
	// ebitenutil.DebugPrintAt(screen,fmt.Sprint("Nombre de FPS  : ", ebiten.ActualFPS()),  config.General.WindowSizeX-250, 15)
	// ebitenutil.DebugPrintAt(screen, fmt.Sprint("Nombres de particules : ",g.system.Content.Len()),config.General.WindowSizeX-250,30)
	
	// for _,e := range(ListeLignes){
	// 	ebitenutil.DebugPrintAt(screen, fmt.Sprint(e.NameValue,*e.Value),0,LineY)
	// 	LineY+=20
	// }
	if config.General.Debug {
	ebitenutil.DebugPrint(screen, fmt.Sprint("WindowSizeX: ",config.General.WindowSizeX))
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("WindowSizeY : ",config.General.WindowSizeY), 0, 15)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("InitNumParticles : ",config.General.InitNumParticles),0,30)
	
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("Type du générateur : ",config.General.TypeGenerateur),0,60)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("SpawnY des particules : ",config.General.SpawnX),0,75)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("SpawnY des particules: ",config.General.SpawnY),0,90)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("RayonSpawnX : ",config.General.RayonSpawnX),0,105)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("RayonSpawnY : ",config.General.RayonSpawnY),0,120)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("SpawnRate : ",config.General.SpawnRate),0,135)

	ebitenutil.DebugPrintAt(screen, fmt.Sprint("ClickMouseParticles : ",config.General.ClickMouseParticules),0,165)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("ClickSpawnRate : ",config.General.ClickSpawnRate),0,180)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("FollowMouseSpawn : ",config.General.FollowMouseSpawn),0,195)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("Bounce : ",config.General.Bounce),0,210)

	ebitenutil.DebugPrintAt(screen, fmt.Sprint("MaxVitesseX : ",int(config.General.MaxVitesseX)),0,240)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("MaxVitesseY : ",config.General.MaxVitesseY),0,255)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("GravityX : ",config.General.GravityX),0,270)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("GravityY : ",config.General.GravityY),0,285)


	ebitenutil.DebugPrintAt(screen, fmt.Sprint("Kill_particule_WindowSizeX : ",config.General.Kill_particule_WindowSizeX),0,315)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("Kill_particule_WindowSizeY : ",config.General.Kill_particule_WindowSizeY),0,330)

	ebitenutil.DebugPrintAt(screen, fmt.Sprint("Lifetime : ",config.General.Lifetime),0,360)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("OpacityLifetime : ",config.General.OpacityLifetime),0,375)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("ChangeOpacity : ",config.General.ChangeOpacity),0,390)

	ebitenutil.DebugPrintAt(screen, fmt.Sprint("MinColorRed : ",config.General.MinColorRed),0,420)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("MaxColorRed : ",config.General.MaxColorRed),0,435)

	ebitenutil.DebugPrintAt(screen, fmt.Sprint("MinColorGreen : ",config.General.MinColorGreen),0,465)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("MaxColorGreen : ",config.General.MaxColorGreen),0,480)

	ebitenutil.DebugPrintAt(screen, fmt.Sprint("MinColorBlue : ",config.General.MinColorBlue),0,510)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("MaxColorBlue : ",config.General.MaxColorBlue),0,525)

	ebitenutil.DebugPrintAt(screen, fmt.Sprint("MinScaleX : ",config.General.MinScaleX),0,555)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("MaxScaleX : ",config.General.MaxScaleX),0,570)

	ebitenutil.DebugPrintAt(screen, fmt.Sprint("MinScaleY : ",config.General.MinScaleY),0,600)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("MaxScaleY : ",config.General.MaxScaleY),0,615)

	}else if g.ReadHelp{

		ebitenutil.DebugPrint(screen,fmt.Sprint("Le Clic gauche de la souris vous permet de faire spawner des particules en fonction du ClickSpawnRate "))
		ebitenutil.DebugPrintAt(screen,fmt.Sprint("Actuellement, lorsque vous cliquez / maintenez le clic gauche, vous faites apparaitre ", config.General.ClickSpawnRate, " particules par clic"), 0, 20)
		ebitenutil.DebugPrintAt(screen,fmt.Sprint("Le Clic droit de la souris vous permet de faire changer le point de spawn des particules en fonction des coordonnées de la souris "), 0, 60)
		ebitenutil.DebugPrintAt(screen,fmt.Sprint("Actuellement, votre SpawnX de particules est situé à : ", config.General.SpawnX, " pixels"), 0, 80)
		ebitenutil.DebugPrintAt(screen,fmt.Sprint("Actuellement, votre SpawnY de particules est situé à : ", config.General.SpawnY, " pixels"), 0, 100)
	}else{
		ebitenutil.DebugPrint(screen,fmt.Sprint("Appuyez sur Echap pour voir/retirer le menu"))
		ebitenutil.DebugPrintAt(screen,fmt.Sprint("Appuyez sur H pour afficher le menu des commandes"), 0, 30)
	}

	if g.ReadMode{
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Valeur : ",g.Value),500,50)
	}


}
