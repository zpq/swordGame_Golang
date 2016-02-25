package main

import (
	"sync"
	"time"
)

//var waitgroup sync.WaitGroup

var c = make(chan bool)
var sm sync.Mutex
var bm sync.Mutex

func main() {
	b := Boss{
		base: Base{
			Name:           "boss",
			health:         20000,
			minAttack:      200,
			maxAttack:      300,
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
			attackInterval: 300,
		},
	}
	p2 := Player{
		base: Base{
			Name:           "roye",
			health:         2500,
			minAttack:      250,
			maxAttack:      300,
			defense:        100,
			attackInterval: 300,
		},
	}
	p3 := Player{
		base: Base{
			Name:           "bobe",
			health:         1500,
			minAttack:      350,
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

	go b.attack()
	go p.attack()
	go p2.attack()
	go p3.attack()

	time.Sleep(60 * time.Second)
}
