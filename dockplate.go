package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type Context struct {
	Env map[string]string
}

func main() {
	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("Dockerfile").Parse(string(stdin))
	if err != nil {
		log.Fatal(err)
	}

	env := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		env[pair[0]] = pair[1]
	}

	writer := bufio.NewWriter(os.Stdout)
	context := Context{env}
	err = tmpl.Execute(writer, context)
	if err != nil {
		log.Fatal(err)
	}
	writer.Flush()
}
