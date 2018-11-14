package conf

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type AppConf struct {
	Port int
	Url  string
}

func InitConf() {
	app_path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file, err := os.Open(app_path + "/.env")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(line))
		if strings.Index(s, "#") == 0 {
			continue
		}
		n1 := strings.Index(s, "[")
		n2 := strings.LastIndex(s, "]")
		if n1 > -1 && n2 > -1 && n2 > n1+1 {
			key := strings.TrimSpace(s[n1+1 : n2])

		}
	}
}

func mapTo(s string, v interface{}) {

}
