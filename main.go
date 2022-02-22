package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"gopkg.in/yaml.v2"
)

func main() {

	pathFile := os.Getenv("INPUT_FILE")

	fmt.Println("Ready to compile ...")

	filename, _ := filepath.Abs(pathFile)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var mapResult map[interface{}]interface{}
	err = yaml.Unmarshal(yamlFile, &mapResult)
	if err != nil {
		panic(err)
	}

	fmt.Println("Env variables will be replaced")

	err_two := checkIsPointer(&mapResult)
	if err_two != nil {
		panic(err_two)
	}
	valueOf := reflect.ValueOf(mapResult)
	val := reflect.Indirect(valueOf)
	switch val.Type().Kind() {
	case reflect.Map:
		envMap := mapResult
		for in, iv := range envMap {
			envName := in.(string)
			envVal := iv.(string)
			env := strings.Replace(strings.TrimSpace(envVal), "$", "", -1)
			envMap[envName] = os.Getenv(env)
		}
	default:
		panic("This is not supposed to happen, but if it does, good luck")
	}
	fmt.Println("Successfully compiled env variables")

	out, err := yaml.Marshal(mapResult)
	if err != nil {
		panic(err)
	}
	// write the whole body at once
	err = ioutil.WriteFile(pathFile, out, 0644)
	if err != nil {
		panic(err)
	}
}

func checkIsPointer(any interface{}) error {
	if reflect.ValueOf(any).Kind() != reflect.Ptr {
		return fmt.Errorf("you passed something that was not a pointer: %s", any)
	}
	return nil
}
