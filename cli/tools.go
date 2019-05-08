package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/iancoleman/strcase"
	"io/ioutil"
	"os/exec"
	"strings"
)

type Intent struct {
	Name        string `json:"intent"`
	Description string `json:"description"`
}

type Entity struct {
	Name       string `json:"entity"`
	FuzzyMatch bool   `json:"fuzzy_match"`
}

type Skill struct {
	Intents  []*Intent `json:"intents"`
	Entities []*Entity `json:"entities"`
}

func ExamineWatsonSkill(skillFile string) error {
	data, err := ioutil.ReadFile(skillFile)
	if err != nil {
		return err
	}
	skill := new(Skill)
	err = json.Unmarshal(data, skill)
	if err != nil {
		return nil
	}

	f := strings.Split(skillFile, ".")
	filename := strcase.ToSnake(f[0]) + ".go"

	packageName := "skills"
	fileBlank := fmt.Sprintf("package %s\n\n", packageName)
	fileBody := bytes.NewBufferString(fileBlank)

	fileBody.WriteString("// Intents constants\n")
	fileBody.WriteString("const (\n\t")
	for _, i := range skill.Intents {
		constName := strcase.ToCamel(i.Name)
		line := fmt.Sprintf(`%s = "%s"
		`, constName, i.Name)
		_, err = fileBody.WriteString(line)
		if err != nil {
			return err
		}
	}

	fileBody.WriteString(")\n\n")

	fileBody.WriteString("// Entities constants\n")
	fileBody.WriteString("const (\n\t")
	for _, e := range skill.Entities {
		constName := strcase.ToCamel(e.Name)
		line := fmt.Sprintf(`%s = "%s"
		`, constName, e.Name)
		_, err = fileBody.WriteString(line)
		if err != nil {
			return err
		}
	}

	fileBody.WriteString(")\n\n")

	err = ioutil.WriteFile(filename, fileBody.Bytes(), 0644)
	if err != nil {
		return err
	}

	err = exec.Command("go", "fmt").Run()
	if err != nil {
		return err
	}

	return nil
}
