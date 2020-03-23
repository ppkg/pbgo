package pbgo

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Generate() error {
	exe := func(name string) error {
		cmd := exec.Command("cmd.exe", "/C", "protoc --go_out=plugins=grpc:. "+name)
		w := bytes.NewBuffer(nil)
		cmd.Stderr = w
		log.Println(fmt.Sprintf("开始生成%s pb.go文件", name))
		if err := cmd.Run(); err != nil {
			fmt.Println(fmt.Sprintf("生成%s pb.go文件出错，%s %s", name, err.Error(), string(w.Bytes())))
			return errors.New(string(w.Bytes()))
		}
		log.Println(fmt.Sprintf("%s pb.go文件生成成功", name))
		return nil
	}

	var err error
	file, _ := ioutil.ReadDir(".")
	for _, v := range file {
		if strings.HasSuffix(v.Name(), ".proto") {
			if err = exe(v.Name()); err != nil {
				break
			}
		}
	}

	if err == nil {
		for _, v := range file {
			if strings.HasSuffix(v.Name(), ".pb.go") {
				if f, err := os.OpenFile(v.Name(), os.O_RDWR, v.Mode().Perm()); err != nil {
					log.Println(err.Error())
				} else {
					defer f.Close()
					if b, err := ioutil.ReadAll(f); err != nil {
						println("err", err.Error())
					} else {
						n := strings.ReplaceAll(string(b), ",omitempty", "")
						os.Truncate(v.Name(), 0)
						f.WriteAt([]byte(n), 0)
					}
				}
			}
		}
	}
	return err
}
