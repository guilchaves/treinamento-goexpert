package main

import "github.com/guilchaves/treinamento-goexpert/6-APIs/1/configs"

func main(){
    config, _ := configs.LoadConfig(".")
    println(config.DBDriver)
}
