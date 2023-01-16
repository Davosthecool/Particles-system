package main

import("github.com/hajimehoshi/ebiten/v2/inpututil"
		"github.com/hajimehoshi/ebiten/v2"
		"project-particles/config")
// Update se charge d'appeler la fonction Update du système de particules
// g.system. Elle est appelée automatiquement exactement 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction ne devrait pas être modifiée sauf
// pour les deux dernières extensions.
func (g *game) Update() error {
	g.system.Update()

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape){
		config.General.Debug = !config.General.Debug
		g.SelectMode = false
	}


	if inpututil.IsKeyJustPressed(ebiten.KeyT){
		g.ReadMode = !g.ReadMode
		print(g.Value)
	}
	if g.ReadMode{
		g.Value += NumPressed()
		if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) && len(g.Value)>0{
			g.Value = g.Value[:len(g.Value)-1]
		}
	}else{
		g.Value = ""
	}


	if inpututil.IsKeyJustPressed(ebiten.KeyS){
		g.ReadMode = !g.ReadMode
	}

	if g.SelectMode && config.General.Debug{
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft){
			// posx, posy := ebiten.CursorPosition()
			g.SelectMode = false
		}
		// if posx>0 && posx<100{

		// }
	}
	
	
	return nil
}



















func NumPressed() string{
	switch{
	case inpututil.IsKeyJustPressed(ebiten.KeyNumpad0):
		return "0"
	case inpututil.IsKeyJustPressed(ebiten.KeyNumpad1):
		return "1"
	case inpututil.IsKeyJustPressed(ebiten.KeyNumpad2):
		return "2"
	case inpututil.IsKeyJustPressed(ebiten.KeyNumpad3):
		return "3"
	case inpututil.IsKeyJustPressed(ebiten.KeyNumpad4):
		return "4"
	case inpututil.IsKeyJustPressed(ebiten.KeyNumpad5):
		return "5"
	case inpututil.IsKeyJustPressed(ebiten.KeyNumpad6):
		return "6"
	case inpututil.IsKeyJustPressed(ebiten.KeyNumpad7):
		return "7"
	case inpututil.IsKeyJustPressed(ebiten.KeyNumpad8):
		return "8"
	case inpututil.IsKeyJustPressed(ebiten.KeyNumpad9):
		return "9"
	case inpututil.IsKeyJustPressed(ebiten.KeyNumpadDecimal):
		return "."
	}
	
	return ""
}