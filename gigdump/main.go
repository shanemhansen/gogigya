// gigdump is an example of using the accounts.search api
// with strong typing. json results are marshaled into structs
// it's intended as an example. Reports are generated using
// text/template. Pagination is handled automatically assuming
// that appending limit %d start %d makes sense.
// example: gigdump -report '{{.UID}}' -apiKey $APIKEY -secretKey $SECRETKEY
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/shanemhansen/gogigya"
	"os"
	"text/template"
)

var query string
var report string
var count int

func main() {
	flag.StringVar(&query, "query", `SELECT * from accounts order by UID`, "Query to send to gigya")
	flag.StringVar(&report,
		"report", "{{.UID}}\t{{.LoginIds.Emails}}",
		"a text/template representation of the desired data")
	flag.IntVar(&count, "count", 0, "How many records to retrieve 0 for all")
	flag.Parse()
	if len(*gogigya.ApiKey) == 0 || len(*gogigya.SecretKey) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	tmpl := template.Must(template.New("report").Parse(report))
	if count == 0 {
		var err error
		count, err = GetCount()
		if err != nil {
			panic(err)
		}
	}
	limit := 1000
	if limit > count {
		limit = count
	}
	for start := 0; start < count; start += limit {
		results, err := GetResults(limit, start)
		if err != nil {
			panic(err)
		}
		for _, account := range results.Results {
			tmpl.Execute(os.Stdout, account)
			os.Stdout.Write([]byte{'\n'})
		}
	}
}
func GetResults(limit, start int) (gogigya.QueryResult, error) {
	query := fmt.Sprintf(query+" limit %d start %d", limit, start)
	req := gogigya.New(&gogigya.Request{Method: "accounts.search",
		ApiKey:    *gogigya.ApiKey,
		SecretKey: *gogigya.SecretKey,
		Format:    "json"})
	req.Params.Add("query", query)
	response, err := req.Send()
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(response.Body)
	results := gogigya.QueryResult{}
	err = decoder.Decode(&results)
	if err != nil {
		panic(err)
	}
	if results.ErrorCode != 0 {
		return results, errors.New(results.ErrorMessage)
	}
	return results, nil
}
func GetCount() (int, error) {
	results, err := GetResults(1, 0)
	return results.TotalCount, err
}
