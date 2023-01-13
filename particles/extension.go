package particles

import("math/rand"
		"project-particles/config")

//Contiendra les extensions 

func Abs(x float64) float64{
	if x<0{
		return -x
	}
	return x
}

func RandomBetweenInt(x, y int) int{
	if x==y{
		return y
	}
	if x>y{
		x,y=y,x
	}
	return rand.Intn(y-x) + x
}

func RandomBetweenFloat(x, y float64) float64{
	if x==y{
		return y
	}
	if x>y{
		x,y=y,x
	}
	return rand.Float64()*(y-x)+x
}

func OutOfScreen(p *Particle) bool{
	Kill_particule_WindowSizeX := 0.1*float64(config.General.WindowSizeX)
	Kill_particule_WindowSizeY := 0.1*float64(config.General.WindowSizeY)

	if p.PositionX > float64(config.General.WindowSizeX)+Kill_particule_WindowSizeX|| 
	p.PositionX < -Kill_particule_WindowSizeX|| 
	p.PositionY > float64(config.General.WindowSizeY)+Kill_particule_WindowSizeY|| 
	p.PositionY < -Kill_particule_WindowSizeY{
		return true
	}
	return false
}

func OutOfKillScreen(x,y float64) bool{
	if x > float64(config.General.WindowSizeX)+config.General.Kill_particule_WindowSizeX|| 
	x < -config.General.Kill_particule_WindowSizeX|| 
	y > float64(config.General.WindowSizeY)+config.General.Kill_particule_WindowSizeY|| 
	y < -config.General.Kill_particule_WindowSizeY{
		return true
	}
	return false
}

func (p *Particle)WallBounce(){
	posx,posy := p.PositionX,p.PositionY 
	switch{
	case posx < -config.General.Kill_particule_WindowSizeX || posx> float64(config.General.WindowSizeX)+config.General.Kill_particule_WindowSizeX:
		p.VitesseX*=-1
	case posy < -config.General.Kill_particule_WindowSizeY || posy> float64(config.General.WindowSizeY)+config.General.Kill_particule_WindowSizeY:
		p.VitesseY*=-1	
	}
}

