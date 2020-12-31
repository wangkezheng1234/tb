package main

import (
	"td/app/admin"
	"td/app/api"
	"td/routers"
)

func main(){
	routers.Include(api.Routers,admin.Routers)
	r := routers.Init()
	r.Run(":8083")
}