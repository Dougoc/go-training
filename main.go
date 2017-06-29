package main

import (
	"fmt"
	_ "bufio"
	"os"
)

func main() {
		//reader := bufio.NewReader(os.Stdin)

		//fmt.Print("The complete folder path to be restored: ")

		// path_origin, _ := reader.ReadString('\n')

		path_origin := "/tmp/"

		fmt.Println("Right, i go to restore this")

		d, err := os.Open(path_origin)
		if err != nil {
			fmt.Println("The informed directory does not exist %s", path_origin)
			os.Exit(1)
		}

		defer d.Close()

		files, err := d.Readdir(-1)
		fmt.Println("Reading "+ path_origin)

		for _, file := range files {
			if file.Mode().IsRegular() {
				fmt.Println(file.Name(), file.Size(), "bytes")
			}
		}
}
