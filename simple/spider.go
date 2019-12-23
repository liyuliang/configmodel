package main

import (
	"github.com/BurntSushi/toml"
	"os"
	"github.con/liyuliang/configmodel"
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
