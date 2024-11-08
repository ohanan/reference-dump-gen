package main

import "flag"

func main() {
	var root string
	flag.StringVar(&root, "root", "", "root object")
	flag.Parse()
	if root == "" {
		panic("root object is empty")
	}

}
