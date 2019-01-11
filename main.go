package main

import (
	"encoding/json"
	"fmt"
	js "github.com/go-restit/lzjson"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// Choosing 7 because {"a":a}
const JSON_MIN_LENGTH = 7

type JsonDescription struct {
	Key      string
	DataType string
	Children []JsonDescription
}

func NewJsonDescription(key, dataType string) JsonDescription {
	return JsonDescription{Key: key, DataType: dataType}
}

func (desc JsonDescription) printDescription() {
	println(fmt.Sprintf("Key: %s", desc.Key))
	println(fmt.Sprintf("Data Type: %s", desc.DataType))
	if len(desc.Children) > 0 {
		for _, v := range desc.Children {
			v.printDescription()
		}
	}
}

func (desc JsonDescription) printDescriptionJson() {
	jso, _ := json.Marshal(desc)
	println(string(jso))
}

func main() {
	jsonFileLoc := os.Args[1]
	if len(jsonFileLoc) < 1 {
		panic("Invalid file location")
	}
	DescribeJson(jsonFileLoc)
}

func DescribeJson(loc string) {
	// read all contents from file into body
	body, err := readContentsFromFile(loc)
	if err != nil {
		panic(err)
	}
	if len(body) < JSON_MIN_LENGTH {
		panic("Invalid json body")
	}
	stringReader := strings.NewReader(body)
	root := js.Decode(stringReader)
	// we want to check if the root is anything other than an object
	if root.Type() == js.TypeObject {
		desc, _ := describe(root, "root")
		desc.printDescriptionJson()
	}
}

// we know that we are only going to get objects here
func describe(jsonNode js.Node, key string) (JsonDescription, error) {
	if jsonNode.Type() == js.TypeBool {
		return NewJsonDescription(key, "bool"), nil
	}
	if jsonNode.Type() == js.TypeNumber {
		return NewJsonDescription(key, "number"), nil
	}
	if jsonNode.Type() == js.TypeString {
		val := jsonNode.String()
		specificType := getSpecificStringType(val)
		return NewJsonDescription(key, specificType), nil
	}
	if jsonNode.Type() == js.TypeNull {
		return NewJsonDescription(key, "null"), nil
	}
	if jsonNode.Type() == js.TypeObject {
		childKeys := jsonNode.GetKeys()
		if len(childKeys) < 1 {
			return JsonDescription{}, nil
		}
		root := JsonDescription{DataType: "item", Key: key}
		root.Children = make([]JsonDescription, 1)
		for _, key := range childKeys {
			childNode := jsonNode.Get(key)
			node, err := describe(childNode, key)
			if err != nil {
				panic(err)
			}
			root.Children = append(root.Children, node)
		}

		return root, nil
	}

	if jsonNode.Type() == js.TypeArray {
		len := jsonNode.Len()
		if len < 1 {
			return JsonDescription{}, nil
		}
		root := JsonDescription{DataType: "relationship", Key: key}
		root.Children = make([]JsonDescription, 0)
		// if the type is array we are only going to
		// parse the first child and assume all the rest are the same
		node, err := describe(jsonNode.GetN(0), "obj"+string(0))
		if err != nil {
			panic(err)
		}
		root.Children = append(root.Children, node)
		return root, nil
	}

	return JsonDescription{}, nil
}

// this method rips our all white space and iterates to
// see if the string is made up of
//2018-11-01T13:30:37 - patters
//2018-11-01
var DateTimePatterns = [...]string{
	"^[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}",
	// you an add more patterns if you want it will affect performance
	// every single string pattern is checked
}

var DatePatterns = [...]string{
	"^[0-9]{4}-[0-9]{2}-[0-9]{2}",
	// you an add more patterns if you want it will affect performance
	// every single string pattern is checked
}

var cache = make(map[string]string, 100)
var dateTime = "date-time"
var date = "date-time"
var stringType = "string"

func getSpecificStringType(val string) string {
	if val, ok := cache[val]; ok {
		return val
	}
	for _, pattern := range DateTimePatterns {
		if match, _ := regexp.Match(pattern, []byte(val)); match {
			cache[val] = dateTime
			return dateTime
		}
	}

	for _, pattern := range DatePatterns {
		if match, _ := regexp.Match(pattern, []byte(val)); match {
			cache[val] = date
			return date
		}
	}

	cache[val] = stringType
	return stringType
}

//func describeType(jsonNode js.Node) (FieldDescription, error) {
//	return NewFieldDescription("unknown"), nil
//}

func readContentsFromFile(loc string) (string, error) {
	b, err := ioutil.ReadFile(loc)
	if err != nil {
		return "", err
	}
	return string(b), err
}
