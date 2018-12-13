package main

import(
	"backend/conf"
	"backend/api/routers"
)

// Env dev/pro
const Env = "dev"

func main()  {
	conf.InitConfig(Env,"api")
	routers.InitRun()
}

