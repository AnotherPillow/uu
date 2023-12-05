package main

import (
	"os"
	"strings"

	"github.com/gookit/color"
)

func main() {
	var full_args_length = len(os.Args)
	if full_args_length <= 1 {
		color.Red.Println("Please specify a command!")
		return
	}
	var cmd = os.Args[1]
	var args = os.Args[2:]

	if len(args) == 0 {
		color.Red.Println("No input/type specified.")
	}

	switch cmd {
	case "hash":
		if len(args) == 1 {
			color.Red.Println("No algorithm specified.")
		}
		if len(args) == 2 {
			color.Red.Println("No input specified.")
		}
		var hasher = args[0]
		var hashtype = args[1]
		var tohash = args[2]

		if !contains(validHashes, hasher) {
			color.Red.Printf("Unknown algorihtm %s.\n", color.BgDarkGray.Sprintf(hasher))
			color.Yellow.Printf("Valid commands are %s\n", strings.Join(validHashes, ", "))
			return
		}

		var hash = ""
		if hashtype == "file" {
			hash = filehash(tohash, hasher)
		} else if hashtype == "string" {
			hash = stringhash(tohash, hasher)
		} else {
			var output = ""
			output += color.Red.Sprintf("Unknown type ")
			output += RED_CODEBLOCK.Sprintf(hashtype)
			output += color.Red.Sprintf(". Options are ")
			output += RED_CODEBLOCK.Sprintf("file")
			output += color.Red.Sprintf(", ")
			output += RED_CODEBLOCK.Sprintf("string")
			output += "\n"

			print(output)
			return
		}

		color.Magenta.Printf("%s hash of %s %s\n", hasher, hashtype, color.BgDarkGray.Sprintf(tohash))
		color.Green.Printf("    %s", hash)
	default:

		color.Red.Printf("Unknown command %s.\n", color.BgDarkGray.Sprintf(cmd))
		var validCommands = []string{
			"hash",
		}
		color.Yellow.Printf("Valid commands are %s\n", strings.Join(validCommands, ", "))
	}
}
