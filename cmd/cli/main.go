package main

import (
	"flag"
	"fmt"

	"github.com/shoaibcldcvr/uuidgen/pkg/uuid"
)

func main() {
	version := flag.String("v", "4", "UUID version (1, 4, or 7)")
	count := flag.Int("n", 1, "Number of UUIDs to generate")
	flag.Parse()

	for i := 0; i < *count; i++ {
		switch *version {
		case "1":
			u, _ := uuid.NewV1()
			fmt.Println(u.String())
		case "4":
			u, _ := uuid.NewV4()
			fmt.Println(u.String())
		default:
			fmt.Println("Unsupported version. Use -v=1 or -v=4")
		}
	}
}
