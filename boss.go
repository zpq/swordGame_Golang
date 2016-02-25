package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Boss struct {
	base   Base
	target []*Player
}

func (b *Boss) attack() {
	for {
		if len(b.target) == 0 {
			return
		}
		if b.base.health <= 0 {
			return
		}
		target := b.selectTarget()

		if target.base.health <= 0 {
			fmt.Println(target.base.Name, " dead ")
			b.removeTarget(target.base.Name)
			if len(b.target) == 0 {
				return
			}
			target = b.selectTarget()
		}

		randAtt := rand.Intn(b.base.maxAttack-b.base.minAttack) + b.base.minAttack
		damage := randAtt - target.base.defense
		if damage > 0 {
			target.base.health -= damage
		} else {
			damage = 0
		}

		fmt.Println(b.base.Name, " damaged ", damage, " on ",
			target.base.Name, " left ", target.base.health)
		time.Sleep(time.Millisecond * 1000)
	}
}

func (b *Boss) selectTarget() (p *Player) {
	r := rand.Intn(len(b.target))
	p = b.target[r]
	return
}

func (self *Boss) removeTarget(name string) {
	index := 0
	flag := false
	for k, v := range self.target {
		if v.base.Name == name {
			index = k
			flag = true
			break
		}
	}
	if flag {
		result := make([]*Player, 0)
		result = append(result, self.target[0:index]...)
		result = append(result, self.target[index+1:]...)
		self.target = result
	}
}
