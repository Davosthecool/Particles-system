package particles

import (
	// "container/list"
	"math/rand"
	"project-particles/config"
	"testing"
)

//Tests concernant le fichier new.go

func TestNewLongueurZero(t *testing.T) {
	config.General.InitNumParticles = 0
	if NewSystem().Content.Len()!=config.General.InitNumParticles  {
		t.Error("Votre système devrait contenir" , config.General.InitNumParticles , "particules. Pourtant, il ne contient que", NewSystem().Content.Len(), "particules")
	}
}

func TestNewLongueurRandom(t *testing.T) {
	for i:=0; i< 100;i++{
		config.General.InitNumParticles = rand.Intn(10000) //Nous avons mis une limite à 10000 sinon le programme prendrait beaucoup de temps à trouver le résultat.
		if NewSystem().Content.Len()!=config.General.InitNumParticles  {
			t.Error("Votre système devrait contenir" , config.General.InitNumParticles ,"particules. Pourtant, il ne contient que", NewSystem().Content.Len(), "particules")
		}
	}
}





//Tests concernant le fichier particles.go

func TestParticlePositionFixed(t *testing.T) {
	config.Get("../config.json")
	config.General.TypeGenerateur = 0
	config.General.InitNumParticles = rand.Intn(100)
	var posx,posy float64
	
	for element := NewSystem().Content.Front() ; element != nil ; element = element.Next(){
		if element.Value.(*Particle).PositionX != float64(config.General.SpawnX) || element.Value.(*Particle).PositionY != float64(config.General.SpawnY) {
			posx,posy = element.Value.(*Particle).PositionX , element.Value.(*Particle).PositionY
			t.Error("Votre particule est à la position (",posx,",",posy,") alors que la position du spawn d'une particule est (",config.General.SpawnX ,",", config.General.SpawnY,")")
			break
		}
	}
	

}



func TestParticlePositionRandom(t *testing.T) {
	config.Get("../config.json")
	config.General.TypeGenerateur = -1
	config.General.InitNumParticles = rand.Intn(100)
	var compt int
	for element := NewSystem().Content.Front() ; element != nil ; element = element.Next(){
		if element.Value.(*Particle).PositionX == float64(config.General.SpawnX) || element.Value.(*Particle).PositionY == float64(config.General.SpawnY) {
			compt+=1
		}
		if compt == NewSystem().Content.Len(){
			t.Error("Vos particules comportent une position fixe alors que votre RandomSpawn est à", config.General.TypeGenerateur)
		}
			
	}
}



func TestParticleVitesseMax(t *testing.T) {
	config.Get("../config.json")
	config.General.InitNumParticles = rand.Intn(100)
	for element := NewSystem().Content.Front() ; element != nil ; element = element.Next(){		
		if Abs(element.Value.(*Particle).VitesseX) > config.General.MaxVitesseX || Abs(element.Value.(*Particle).VitesseY) > config.General.MaxVitesseY{
			t.Error("La vitesse de la particule est plus rapide que la vitesse maximale autorisée")
			break
		}
	}
}




// Tests concernant la fonction update.go 

func TestUpdateParticleVitesse(t *testing.T){
	config.Get("../config.json")
	sys := NewSystem()
	for element := sys.Content.Front() ; element != nil ; element = element.Next(){
		posx:=element.Value.(*Particle).PositionX
		posy:=element.Value.(*Particle).PositionY
		sys.Update()
		if element.Value.(*Particle).PositionX != (posx + element.Value.(*Particle).VitesseX){
			t.Error("Votre particule ne se déplace pas comme vous le souhaitez sur l'axe X")
			break
		}
		
		if element.Value.(*Particle).PositionY != posy + element.Value.(*Particle).VitesseY{
			t.Error("Votre particule ne se déplace pas comme vous le souhaitez sur l'axe Y")
			break
		}
	}
}

func TestUpdateParticleSpawnZero(t *testing.T){
	config.Get("../config.json")
	sys := NewSystem()
	config.General.SpawnRate = 0

	for i:=0;i<100;i++{
		longueur := sys.Content.Len()
		sys.Update()
		if longueur != sys.Content.Len(){
			t.Error("La fonction Update ne fait pas apparaître" , config.General.SpawnRate , "particules par seconde !")
			break
		}
	}
}

func TestUpdateParticleSpawnFloat(t *testing.T){
	config.Get("../config.json")
	sys := NewSystem()
	config.General.SpawnRate = rand.Float64()
	var spawn float64
	for i:=0;i<100;i++{
		longueur := sys.Content.Len()
		spawn+=config.General.SpawnRate
		sys.Update()
		if longueur != sys.Content.Len()-int(spawn){
			t.Error("La fonction Update ne fait pas apparaître" , config.General.SpawnRate , "particules par seconde !")
			break
		}
		spawn-=float64(int(spawn))
	}
}

func TestUpdateParticleSpawnInt(t *testing.T){
	config.Get("../config.json")
	sys := NewSystem()
	config.General.SpawnRate = float64(rand.Intn(1000))

	for i:=0;i<100;i++{
		longueur := sys.Content.Len()
		sys.Update()
		if longueur != sys.Content.Len()-int(config.General.SpawnRate){
			t.Error("La fonction Update ne fait pas apparaître" , config.General.SpawnRate , "particules par seconde !")
			break
		}
	}
}


func TestOutParticule(t *testing.T){
	config.Get("../config.json")
	sys := NewSystem()
	config.General.WindowSizeX = 800 
	config.General.WindowSizeY = 600
	config.General.Kill_particule_WindowSizeX = float64(rand.Intn(100))*NegaOrPosa()
	for element := sys.Content.Front() ; element != nil ; element = element.Next(){
		posx:=element.Value.(*Particle).PositionX 
		posy:=element.Value.(*Particle).PositionY
		sys.Update()
		if posx > float64(config.General.WindowSizeX) + config.General.Kill_particule_WindowSizeX || posy > float64(config.General.WindowSizeY) +config.General.Kill_particule_WindowSizeY{
			t.Error("Une particule a réussi à se situer dans la kill_particule_zone !")
			break
		}
	}
}


func LifetimeParticule(t *testing.T){
	config.Get("../config.json")
	sys := NewSystem()
	for element := sys.Content.Front() ; element != nil ; element = element.Next(){
		LifeParticule := element.Value.(*Particle).Lifetime
		LifeParticule--
		sys.Update()
		if LifeParticule <= 0 && element != nil{
			t.Error("La particule ne comporte plus de vie. Pourtant, elle se situe encore dans votre liste !")
		} 	
	}
}

func LifetimeParticule(t *testing.T){}