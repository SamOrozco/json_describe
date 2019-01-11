package main

import (
	js "github.com/go-restit/lzjson"
	"io/ioutil"
	"os"
	"strings"
)

// Choosing 7 because {"a":a}
const JSON_MIN_LENGTH = 7

type FieldDescription struct {

}

func main() {
	jsonFileLoc := os.Args[1]
	if len(jsonFileLoc) < 1 {
		panic("Invalid file location")
	}

	// read all contents from file into body
	body, err := readContentsFromFile(jsonFileLoc)
	if err != nil {
		panic(err)
	}
	if len(body) < 1 {
		panic("Invalid json body")
	}
	stringReader := strings.NewReader(body)
	root := js.Decode(stringReader)
	// we want to check if the root is anything other than an object
	if root.Type() == js.TypeObject {
	}

}

// we know that we are only going to get objects here
func describe(jsonNode js.Node) ([][][]string, error) {
}

func describeType(jsonNode js.Node) ([][]string, error) {

}

func printDescription(val [][][]string) {
	println(val[0][0][0]);
}

func readContentsFromFile(loc string) (string, error) {
	b, err := ioutil.ReadFile(loc)
	if err != nil {
		return "", err
	}
	return string(b), err
}
