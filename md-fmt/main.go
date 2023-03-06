package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/88250/lute"
)

func main() {
	luteEngine := lute.New()
	FmtDir(luteEngine, "./")
	fmt.Println("ok")
}

func FmtDir(luteEngine *lute.Lute, dirname string) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}
	if len(files) == 0 {
		return
	}

	for _, f := range files {
		filePath := path.Join(dirname, f.Name())

		if f.IsDir() {
			FmtDir(luteEngine, filePath)
			continue
		}
		if !strings.HasSuffix(strings.ToLower(f.Name()), ".md") {
			continue
		}

		md, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err)
		}

		md = luteEngine.Format(f.Name(), md)
		if err := ioutil.WriteFile(filePath, md, 0644); err != nil {
			panic(err)
		}

		fmt.Println("format: " + filePath)
	}
}
