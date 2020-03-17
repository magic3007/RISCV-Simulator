package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
	"path"
	"text/template"
	"utils"
)

var (
	filepath = flag.String("f", "action_tables.csv", "filepath of the action table CSV file")
	tmplpath = flag.String("t", "action.tmpl", "filepath of template file")
)

// ReadCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
func ReadCsv(filename string) ([][]string, error) {

    // Open CSV file
    f, err := os.Open(filename)
    if err != nil {
        return [][]string{}, err
    }
    defer f.Close()

    // Read File into a Variable
    lines, err := csv.NewReader(f).ReadAll()
    if err != nil {
        return [][]string{}, err
    }

    return lines, nil
}

type Action struct {
    Name string
    Type string
    Opcode string
    Funct3 string
    Funct7 string
    BitConstraint string
    Action1 string
    Action2 string
}

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Llongfile)
	var (
		PACKAGE = "simulator"
		IMPORT = []string{"bit_utils", "isa", "log", "math/bits"}
		CONST = map[string]string{"XLEN":"64"}
		dest = path.Join(utils.GetCurrentPath(), PACKAGE, "action.go")
	)

	lines, err := ReadCsv(*filepath)
    if err != nil {
        log.Fatal(err)
    }

    reformat := func(s string) string {
    	if s[len(s)-1] == ';' {
    		s = s[:len(s)-1]
		}
		return s
	}

    actionSet := make([]Action, 0, len(lines) - 1)
    for _, line := range lines[1:] {
        data := Action{
            Name: line[0],
            Type: line[1],
            Opcode: line[2],
    		Funct3: line[3],
    		Funct7: line[4],
    		BitConstraint: line[5],
    		Action1: reformat(line[6]),
    		Action2: reformat(line[7]),
        }
        actionSet = append(actionSet, data)
    }

	vals := struct{
		PACKAGE string
		IMPORT []string
		CONST map[string]string
		ActionSet []Action
	}{
		PACKAGE,
		IMPORT,
		CONST,
		actionSet,
	}

	tmpl, err := template.ParseFiles(*tmplpath)
	if err != nil {
		log.Fatal(err)
	}

	outFile, err := os.Create(dest)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	if err := tmpl.Execute(outFile, vals); err != nil {
		log.Fatalln("executing template:", err)
	}
}

