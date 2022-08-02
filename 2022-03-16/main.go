package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

/*
   判断文件或文件夹是否存在
   如果返回的错误为nil,说明文件或文件夹存在
   如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
   如果返回的错误为其它类型,则不确定是否在存在
*/
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	var cmd *exec.Cmd
	var err error
	var exist bool
	paths := []string{
		"C:\\Program Files (x86)\\Mobatek\\MobaXterm\\MobaXterm.exe",
		"C:\\Program Files\\Typora\\Typora.exe",
		"D:\\Program Files\\JetBrains\\PhpStorm 2021.1.3\\bin\\phpstorm64.exe",
		"C:\\Users\\fenno\\AppData\\Local\\SourceTree\\SourceTree.exe",
		"D:\\Program Files\\JetBrains\\GoLand 2021.1.3\\bin\\goland64.exe",
		"D:\\Program Files\\PremiumSoft\\Navicat Premium 15\\navicat.exe",
		"C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe",
	}

	for _, path := range paths {

		if path == "" {
			continue
		}

		exist, err = PathExists(path)
		if err != nil {
			fmt.Printf("PathExists(%s),err(%v)\n", path, err)
		}

		if !exist {
			fmt.Println(path + "文件不存在")
			continue
		}

		cmd = exec.Command(path)
		err = cmd.Run()

		if err != nil {
			fmt.Println(err)
		}
		
		fmt.Println(path,"启动成功")

		time.Sleep(time.Microsecond * 1)
	}

}
