package main

import (
	"flag"
	"fmt"
	"github.com/hojabri/http-call-md5/app"
	"os"
)

func main() {
	parallel := flag.Int("parallel", 10, "number of parallel requests")
	flag.Parse()
	flag.Usage =usage
	commandLineURLs := flag.Args()

	if len(commandLineURLs) == 0 {
		flag.Usage()
		os.Exit(0)
	}


	app.StartApplication(commandLineURLs, *parallel)

}


// command customized usage function
func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "%s [-parallel <size> ] <URL1> <URL2> ...\n", os.Args[0])
	flag.PrintDefaults()
}
