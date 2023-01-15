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

	if config.General.Debug {
		ebitenutil.DebugPrintAt(screen,fmt.Sprint("Nombre de TPS  : ", ebiten.ActualTPS()), config.General.WindowSizeX-250, 0)
		ebitenutil.DebugPrintAt(screen,fmt.Sprint("Nombre de FPS  : ", ebiten.ActualFPS()),  config.General.WindowSizeX-250, 15)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Nombres de particules : ",g.system.Content.Len()),config.General.WindowSizeX-250,30)
		ebitenutil.DebugPrint(screen, fmt.Sprint("WindowSizeX: ",config.General.WindowSizeX))
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("WindowSizeY : ",config.General.WindowSizeY), 0, 15)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("InitNumParticles : ",config.General.InitNumParticles),0,30)
		

		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Type du générateur : ",config.General.TypeGenerateur),0,60)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("SpawnY des particules : ",config.General.SpawnX),0,75)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("SpawnY des particules: ",config.General.SpawnY),0,90)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("RayonSpawnX : ",config.General.RayonSpawnX),0,105)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("RayonSpawnY : ",config.General.RayonSpawnY),0,120)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("SpawnRate : ",config.General.SpawnRate),0,135)

		ebitenutil.DebugPrintAt(screen, fmt.Sprint("ClickMouseSpawn : ",config.General.ClickMouseSpawn),0,165)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("ClickSpawnRate : ",config.General.ClickSpawnRate),0,180)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("FollowMouseSpawn : ",config.General.FollowMouseSpawn),0,195)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Bounce : ",config.General.Bounce),0,210)

		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MaxVitesseX : ",config.General.MaxVitesseX),0,240)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MaxVitesseY : ",config.General.MaxVitesseY),0,255)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("GravityX : ",config.General.GravityX),0,270)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("GravityY : ",config.General.MaxVitesseY),0,285)


		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Kill_particule_WindowSizeX : ",config.General.Kill_particule_WindowSizeX),0,315)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Kill_particule_WindowSizeY : ",config.General.Kill_particule_WindowSizeY),0,330)

		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Lifetime : ",config.General.Lifetime),0,360)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("OpacityLifetime : ",config.General.OpacityLifetime),0,375)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("ChangeOpacity : ",config.General.ChangeOpacity),0,390)

		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MinColorRed : ",config.General.MinColorRed),0,420)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MaxColorRed : ",config.General.MaxColorRed),0,435)

		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MinColorGreen : ",config.General.MaxVitesseY),0,465)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MaxColorGreen : ",config.General.MaxVitesseY),0,480)

		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MinColorBlue : ",config.General.MaxVitesseY),0,510)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MaxColorBlue : ",config.General.MaxVitesseY),0,525)

		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MinScaleX : ",config.General.MaxVitesseY),0,555)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MaxScaleX : ",config.General.MaxVitesseY),0,570)

		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MinScaleY : ",config.General.MinScaleY),0,600)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MaxScaleY : ",config.General.MaxScaleY),0,615)

		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MaxScaleY : ",config.General.MaxScaleY),0,645)

	}

}
