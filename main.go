package main

import (
	"flag"
	"fmt"
	"yasuo/tools"
)

var Usage = func() {
	fmt.Println("USAGE: compress command [arguments] ...")
	fmt.Println("\nThe commands are:\n\ttype\tcompress type[zip, tar].")
	fmt.Println("\tcompress dir\tcompress dir or compress file.")
	fmt.Println("\tsaved filename\tsaved filename.")
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		Usage()
		return
	}

	if args[0] == "help" || args[0] == "h" {
		Usage()
		return
	}

	switch args[0] {
	case "zip":
		if len(args) != 3 {
			fmt.Println("USAGE: ys zip <filename>")
			return
		}
		tools.Zip(args[1], args[2])
	case "tar":
		if len(args) != 3 {
			fmt.Println("USAGE: ys tar <filename>")
			return
		}
		tools.Tar(args[1], args[2])
	default:
		Usage()
	}

}
