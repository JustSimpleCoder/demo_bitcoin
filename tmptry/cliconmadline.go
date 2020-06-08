package tmptry

import (
	"fmt"
	"os"
)

func CliDemo() {

	fmt.Println(len(os.Args))
	for i, cmd := range os.Args {
		fmt.Printf("num:%d %s \n", i, cmd)
	}
}
