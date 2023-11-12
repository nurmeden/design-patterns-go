package _interface

type iGun interface {
	setName(name string)
	serPower(power int)
	getName() string
	getPower() int
}
