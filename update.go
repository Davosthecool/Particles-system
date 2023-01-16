package main

import("github.com/hajimehoshi/ebiten/v2/inpututil"
		"github.com/hajimehoshi/ebiten/v2"
		"project-particles/config"
		)
// Update se charge d'appeler la fonction Update du système de particules
// g.system. Elle est appelée automatiquement exactement 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction ne devrait pas être modifiée sauf
// pour les deux dernières extensions.
func (g *game) Update() error {
	g.system.Update()

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape){
		g.ReadMode = true
		config.General.Debug = !config.General.Debug
		g.ReadHelp = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyH){
		g.ReadHelp = !g.ReadHelp
		g.ReadMode = false
		config.General.Debug = false
	}



	if g.SelectMode && config.General.Debug{
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft){
			// posx, posy := ebiten.CursorPosition()
			g.SelectMode = false
		}
		// if posx>0 && posx<100{

		// }
	}
	// if inpututil.IsKeyJustPressed(ebiten.KeyE){
	// 	g.ReadMode = !g.ReadMode
	// 	if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
	// 		config.General.MaxVitesseX += 5.0
	// 	}else{
	// 		config.General.MaxVitesseX -= 5.0
	// 	}
	// }
	g.lol()
	
	return nil
}

func (g *game) test()string{
	if g.ReadMode{
		g.Value += NumPressed()
		if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) && len(g.Value)>0{
			g.Value = g.Value[:len(g.Value)-1]
		}
	}
	return g.Value
}

func (g*game) lol(){
	g.test()
	switch{

	case g.Value == "B":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.InitNumParticles += 1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.InitNumParticles -= 1
		}

	case g.Value == "C":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			if config.General.TypeGenerateur <= 3{
				config.General.TypeGenerateur += 1
			}
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			if config.General.TypeGenerateur >= -1{
				config.General.TypeGenerateur -= 1
			}
		}

	case g.Value == "D1":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.SpawnX+= 1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.SpawnX -= 1
		}

	case g.Value == "D2":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.SpawnY+= 1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.SpawnY -= 1
		}

	case g.Value == "E":
		if ebiten.IsKeyPressed(ebiten.KeyKPAdd){
			config.General.SpawnRate += 0.1
		}else if ebiten.IsKeyPressed(ebiten.KeyKPSubtract){
			config.General.SpawnRate -= 0.1
		}

	case g.Value == "F":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.ClickMouseParticules = true
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.ClickMouseParticules = false
		}

	case g.Value == "G":
		if ebiten.IsKeyPressed(ebiten.KeyKPAdd){
			config.General.ClickSpawnRate += 0.1
		}else if ebiten.IsKeyPressed(ebiten.KeyKPSubtract){
			config.General.ClickSpawnRate -= 0.1
		}

	case g.Value == "I":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.FollowMouseSpawn = true
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.FollowMouseSpawn = false
		}

	case g.Value == "J":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.Bounce = true
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.Bounce = false
		}

	case g.Value == "K1":
		if ebiten.IsKeyPressed(ebiten.KeyKPAdd){
			config.General.MaxVitesseX += 0.1
		}else if ebiten.IsKeyPressed(ebiten.KeyKPSubtract){
			config.General.MaxVitesseX -= 0.1
		}

	case g.Value == "K2":
		if ebiten.IsKeyPressed(ebiten.KeyKPAdd){
			config.General.MaxVitesseY += 0.1
		}else if ebiten.IsKeyPressed(ebiten.KeyKPSubtract){
			config.General.MaxVitesseY -= 0.1
		}
	case g.Value == "L1":
		if ebiten.IsKeyPressed(ebiten.KeyKPAdd){
			config.General.GravityX += 0.1
		}else if ebiten.IsKeyPressed(ebiten.KeyKPSubtract){
			config.General.GravityX -= 0.1
		}

	case g.Value == "L2":
		if ebiten.IsKeyPressed(ebiten.KeyKPAdd){
			config.General.GravityY += 0.1
		}else if ebiten.IsKeyPressed(ebiten.KeyKPSubtract){
			config.General.GravityY -= 0.1
		}

	case g.Value == "M1":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.Kill_particule_WindowSizeX += 1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.Kill_particule_WindowSizeX += 1
		}

	case g.Value == "M2":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.Kill_particule_WindowSizeY += 1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.Kill_particule_WindowSizeY += 1
		}

	case g.Value == "N":
		if ebiten.IsKeyPressed(ebiten.KeyKPAdd){
			config.General.Lifetime += 0.1
		}else if ebiten.IsKeyPressed(ebiten.KeyKPSubtract){
			config.General.Lifetime -= 0.1
		}

	case g.Value == "O":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.OpacityLifetime = true
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.OpacityLifetime = false
		}

	case g.Value == "P":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.ChangeOpacity += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.ChangeOpacity -= 0.1
		}

	case g.Value == "Q1":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MinColorRed += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MinColorRed -= 0.1
		}

	case g.Value == "Q2":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MaxColorRed += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MaxColorRed -= 0.1
		}

	case g.Value == "R1":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MinColorGreen += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MinColorGreen -= 0.1
		}

	case g.Value == "R2":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MaxColorGreen += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MaxColorGreen -= 0.1
		}

	case g.Value == "S1":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MinColorBlue += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MinColorBlue -= 0.1
		}

	case g.Value == "S2":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MaxColorBlue += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MaxColorBlue -= 0.1
		}

	case g.Value == "U1":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MinScaleX += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MinScaleX -= 0.1
		}

	case g.Value == "U2":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MaxScaleX += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MaxScaleX -= 0.1
		}

	case g.Value == "V1":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MinScaleY += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MinScaleY -= 0.1
		}

	case g.Value == "V2":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MaxScaleY += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MaxScaleY -= 0.1
		}
	}
}
	







func NumPressed() string{
	switch{
	case inpututil.IsKeyJustPressed(ebiten.KeyA):
		return "A"
	case inpututil.IsKeyJustPressed(ebiten.KeyB):
		return "B"
	case inpututil.IsKeyJustPressed(ebiten.KeyC):
		return "C"
	case inpututil.IsKeyJustPressed(ebiten.KeyD):
		return "D"
	case inpututil.IsKeyJustPressed(ebiten.KeyE):
		return "E"
	case inpututil.IsKeyJustPressed(ebiten.KeyF):
		return "F"
	case inpututil.IsKeyJustPressed(ebiten.KeyG):
		return "G"
	case inpututil.IsKeyJustPressed(ebiten.KeyH):
		return "H"
	case inpututil.IsKeyJustPressed(ebiten.KeyI):
		return "I"
	case inpututil.IsKeyJustPressed(ebiten.KeyJ):
		return "J"
	case inpututil.IsKeyJustPressed(ebiten.KeyK):
		return "K"
	case inpututil.IsKeyJustPressed(ebiten.KeyM):
		return "M"
	case inpututil.IsKeyJustPressed(ebiten.KeyN):
		return "N"
	case inpututil.IsKeyJustPressed(ebiten.KeyO):
		return "O"
	case inpututil.IsKeyJustPressed(ebiten.KeyP):
		return "P"
	case inpututil.IsKeyJustPressed(ebiten.KeyQ):
		return "Q"
	case inpututil.IsKeyJustPressed(ebiten.KeyR):
		return "R"
	case inpututil.IsKeyJustPressed(ebiten.KeyS):
		return "S"
	case inpututil.IsKeyJustPressed(ebiten.KeyT):
		return "T"
	case inpututil.IsKeyJustPressed(ebiten.KeyU):
		return "U"
	case inpututil.IsKeyJustPressed(ebiten.KeyV):
		return "V"
	case inpututil.IsKeyJustPressed(ebiten.Key1):
		return "1"
	
	case inpututil.IsKeyJustPressed(ebiten.Key2):
		return "2"

	}
	return ""
}

