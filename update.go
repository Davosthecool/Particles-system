package main

import("github.com/hajimehoshi/ebiten/v2/inpututil")
// Update se charge d'appeler la fonction Update du système de particules
// g.system. Elle est appelée automatiquement exactement 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction ne devrait pas être modifiée sauf
// pour les deux dernières extensions.
func (g *game) Update() error {

	g.system.Update()
	g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])

	for _, key := range g.pressedKeys {
		switch key.String() {
		case "ArrowUp":
			print("lol")
		case "ArrowDown":
			print("lal")
		}
	}


	return nil
}
