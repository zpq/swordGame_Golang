package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	base   Base
	target []*Boss
}

func (p *Player) attack() {
	for {
		if p.base.health <= 0 {
			return
		}
		target := p.selectTarget()

		if target.base.health <= 0 {
			fmt.Println(target.base.Name, " dead ")
			return
		}
		randAtt := rand.Intn(p.base.maxAttack-p.base.minAttack) + p.base.minAttack
		damage := randAtt - target.base.defense
		sm.Lock()
		if damage > 0 {
			target.base.health -= damage
		} else {
			damage = 0
		}
		fmt.Println(p.base.Name, " damaged ", damage, " on ",
			target.base.Name, " left ", target.base.health)
		sm.Unlock()
		//	c <- true
		//		runtime.Gosched()
		time.Sleep(time.Millisecond * 500)
	}
}

func (p *Player) selectTarget() (b *Boss) {
	r := rand.Intn(len(p.target))
	b = p.target[r]
	return
}
