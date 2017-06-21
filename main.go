package main

import (
	"yasuo/tools"
	"os"
	"fmt"
)

var Usage = func() {
	fmt.Println("USAGE: compress command [arguments] ...")
	fmt.Println("\nThe commands are:\n\ttype\tcompress type[zip, tar].")
	fmt.Println("\tcompress dir\tcompress dir or compress file.")
	fmt.Println("\tsaved filename\tsaved filename.")
}

func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		Usage()
		return
	}

	if args[1] == "help" || args[1] == "h" {
		Usage()
		return
	}

	switch args[1] {
	case "zip":
		if len(args) != 4 {
			fmt.Println("USAGE: ys zip <filename>")
			return
		}
		tools.Zip(args[2], args[3]);
	case "tar":
		if len(args) != 4 {
			fmt.Println("USAGE: ys tar <filename>")
			return
		}
		tools.Tar(args[2], args[3]);
	default:
		Usage()
	}

}
