package api

type ServerConfig struct {
	ListenAddr string `yaml:"LISTEN_ADDR"`
	DebugMode  bool   `yaml:"DEBUG_MODE"`
}
