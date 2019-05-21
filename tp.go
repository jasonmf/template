/* This is a trivial template processor for golang/pkg/text/template. It takes its data as a JSON file specified in the first argument and the template specified in the second argument. Output is to stdout.
 */
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

var (
	fLax = flag.Bool("lax", false, "allow missing keys")
)

func usage() {
	fmt.Println("usage: tp [-lax] <vars file> <template file>")
	os.Exit(-1)
}

func fatalIfError(err error, msg string) {
	if err != nil {
		log.Fatal("error ", msg, ": ", err)
	}
}

func main() {
	flag.Parse()
	varsFile := flag.Arg(0)
	tmplFile := flag.Arg(1)
	if varsFile == "" || tmplFile == "" {
		usage()
	}

	varsData, err := ioutil.ReadFile(varsFile)
	fatalIfError(err, "reading vars file")

	vars := map[string]interface{}{}
	fatalIfError(json.Unmarshal(varsData, &vars), "parsing vars JSON")

	t, err := template.ParseFiles(tmplFile)
	fatalIfError(err, "parsing template file")

	if !*fLax {
		t.Option("missingkey=error")
	}

	fatalIfError(t.Execute(os.Stdout, vars), "rendering template")
}
