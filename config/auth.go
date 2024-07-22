package config

import "github.com/jetnoli/go-router/utils"

var Auth = &utils.AuthConfig{
	Time:    3,
	Memory:  32 * 1024,
	Threads: 4,
	KeyLen:  256,
	SaltLen: 32,
}
