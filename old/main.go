package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Base struct {
	Name           string
	health         int
	minAttack      int
	maxAttack      int
	defense        int
	attackInterval int
}

type Boss struct {
	base   Base
	target []*Player
}

type Player struct {
	base   Base
	target []*Boss
}

func (b *Boss) attack() {
	for {
		if len(b.target) == 0 {
			fmt.Println("Boss win")
			break
		}
		if b.base.health <= 0 {
			// fmt.Println(b.base.Name , " dead");
			break
		}
		target := b.selectTarget()

		if target.base.health <= 0 {
			fmt.Println(target.base.Name, " dead ")
			b.removeTarget(target.base.Name)
			continue
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

func (p *Player) attack() {
	for {
		if p.base.health <= 0 {
			// fmt.Println(p.base.Name , " dead");
			break
		}
		target := p.selectTarget()

		if target.base.health <= 0 {
			fmt.Println(target.base.Name, " dead ")
			fmt.Println(" player win ")
			break
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

//var waitgroup sync.WaitGroup
var c = make(chan bool)
var sm sync.Mutex
var bm sync.Mutex

func main() {
	b := Boss{
		base: Base{
			Name:           "boss",
			health:         20000,
			minAttack:      250,
			maxAttack:      400,
			defense:        200,
			attackInterval: 500,
		},
	}
	p := Player{
		base: Base{
			Name:           "jack",
			health:         2000,
			minAttack:      250,
			maxAttack:      300,
			defense:        130,
			attackInterval: 200,
		},
	}
	p2 := Player{
		base: Base{
			Name:           "roye",
			health:         2500,
			minAttack:      250,
			maxAttack:      400,
			defense:        100,
			attackInterval: 300,
		},
	}
	p3 := Player{
		base: Base{
			Name:           "bobe",
			health:         1800,
			minAttack:      400,
			maxAttack:      500,
			defense:        100,
			attackInterval: 300,
		},
	}
	b.target = append(b.target, &p)
	p.target = append(p.target, &b)
	b.target = append(b.target, &p2)
	p2.target = append(p2.target, &b)
	b.target = append(b.target, &p3)
	p3.target = append(p3.target, &b)

	// c1, c2 := make(chan bool), make(chan bool)

	//	fmt.Println(b.target[0], b.target[1])
	//	b.removeTarget(b.target[0].base.Name)
	//	fmt.Println(len(b.target))
	go b.attack()

	go p.attack()

	go p2.attack()

	go p3.attack()

	//		if !<-c {
	//			break
	//		}

	time.Sleep(60 * time.Second)

	// select {
	// case <-c1:
	// 	fmt.Println("ch1 pop one element")
	// case <-c2:
	// 	fmt.Println("ch2 pop one element")
	// }

	// fmt.Println(b)
	// fmt.Println(p)
	// t := b.target

	// fmt.Println(*t[0])

}
