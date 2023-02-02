package dear

import(
	"flag"
)

const (
	defaultConfigPath  = "./dear_go.yaml"
)

var path = defaultConfigPath

func ServerConfigPath() (path string){
	if path == defaultConfigPath {
		flag.StringVar(&path, "conf", defaultConfigPath, "server config path")
		flag.Parse()
	}
	return path
}