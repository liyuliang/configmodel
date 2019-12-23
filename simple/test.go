package main

import (
	"github.com/BurntSushi/toml"
	"github.con/liyuliang/configmodel"
	"os"
)

func main() {

	as := new(configmodel.Actions)
	_,err := toml.DecodeFile("./test.toml",as)
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}


	for _, action := range as.Action {
		println(action.Target.Key)
	}
}
