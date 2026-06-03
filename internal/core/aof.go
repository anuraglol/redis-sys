package core

import (
	"fmt"
	"log"
	"os"
	"rediss/config"
	"strings"
)

func dumpKey(fp *os.File, k string, obj *Obj) {
	cmd := fmt.Sprintf("SET %s %S", k, obj.Value)
	tokens := strings.Split(cmd, " ")
	fp.Write(Encode(tokens, false))
}

func DumpAllAOF() {
	fp, err := os.OpenFile(config.AOFFile, os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Print("error", err)
		return
	}
	log.Println("rewriting aof file")
	for k, obj := range store {
		dumpKey(fp, k, obj)
	}
	log.Println("completed writing aof file at: ", config.AOFFile)
}
