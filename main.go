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
	ch := flag.Args()

	if ch == nil || len(ch) == 0 {
		Usage()
		return
	}

	if ch[0] == "help" || ch[0] == "h" {
		Usage()
		return
	}

	switch ch[0] {
	case "zip":
		if len(ch) != 3 {
			fmt.Println("USAGE: ys zip <filename>")
			return
		}
		tools.Zip(ch[1], ch[2])
	case "tar":
		if len(ch) != 3 {
			fmt.Println("USAGE: ys tar <filename>")
			return
		}
		tools.Tar(ch[1], ch[2])
	default:
		Usage()
	}

}
