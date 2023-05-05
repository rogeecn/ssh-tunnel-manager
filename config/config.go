package config

type Config struct {
	Tunnel []Tunnel
}

type Tunnel struct {
	Name       string
	LocalPort  uint
	RemotePort uint
	RemoteHost string
	JumpServer string
}
