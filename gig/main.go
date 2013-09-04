// gig is a generic command line
// utility for interacting with the
// gigya REST api http://developers.gigya.com/037_API_reference
// The format is: gig $method key value key value
// for example:
// gig socialize.exportUsers limit 10
// the output is json
package main

import (
	"flag"
	"github.com/shanemhansen/gogigya"
	"io"
	"os"
)

func main() {
	flag.Parse()
	if len(*gogigya.ApiKey) == 0 || len(*gogigya.SecretKey) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	args := flag.Args()
	method := args[0]
	args = args[1:]
	req := gogigya.New(&gogigya.Request{Method: method,
		ApiKey:    *gogigya.ApiKey,
		SecretKey: *gogigya.SecretKey,
		Format:    "json"})
	if len(args) > 0 && len(args)%2 == 0 {
		for i := 0; i < len(args); i += 2 {
			req.Params.Add(args[i], args[i+1])
		}
	}
	response, err := req.Send()
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	io.Copy(os.Stdout, response.Body)
}
