package main 

import ("testing"
	"project-particles/config"
	"project-particles/particles"
	"github.com/hajimehoshi/ebiten/v2"
)


func TestClickMouseParticules(t *testing.T){
	config.Get("config.json")
	sys := NewSystem()
	oldnbrparticule := float64(sys.Content.Len())
	config.General.ClickMouseParticules = true
	config.General.ClickSpawnRate = 15

	oldspawnx,oldspawny:=config.General.SpawnX,config.General.SpawnY
	if config.General.ClickMouseParticules{
		for i:=0.0;i<config.General.ClickSpawnRate;i++{
			config.General.SpawnX,config.General.SpawnY= ebiten.CursorPosition()
			sys.newParticle()
		}
		config.General.SpawnX,config.General.SpawnY=oldspawnx,oldspawny
	}

	
	if oldnbrparticule == float64(sys.Content.Len()){
		t.Error("Test")
	}
}
