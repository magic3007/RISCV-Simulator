package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
	"utils"
)

var (
	filepath = flag.String("f", "action_table.csv", "filepath of the action table CSV file")
	tmplpath = flag.String("t", "microaction.tmpl", "filepath of template file")
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

type MicroAction struct {
	Name                 string
	Type                 string
	Opcode               string
	Funct3               string
	Funct7               string
	BitConstraint        string
	DisplayFormat        string
	IsBranch             string
	IsIndirectJump       string
	MemoryAccessFunction string
	ALUFunction          string
	ValCFunction         string
	PositiveOptionPC     string
	NegativeOptionPC     string
	EStagePeriod         string
	DstE                 string
	DstM                 string
	M_velE_Source        string
}

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Llongfile)
	var (
		PACKAGE = "pipeline"
		IMPORT  = []string{"bit_utils", "isa", "math/bits", "memory"}
		CONST   = map[string]string{"XLEN": "64"}
		dest    = path.Join(utils.GetCurrentPath(), PACKAGE, "microaction.go")
	)

	lines, err := ReadCsv(*filepath)
	if err != nil {
		log.Fatal(err)
	}

	actionSet := make([]MicroAction, 0, len(lines)-1)
	for _, line := range lines[1:] {
		data := MicroAction{
			Name:                 line[0],
			Type:                 line[1],
			Opcode:               line[2],
			Funct3:               line[3],
			Funct7:               line[4],
			BitConstraint:        line[5],
			DisplayFormat:        line[8],
			IsBranch:             strings.ToLower(line[9]),
			IsIndirectJump:       strings.ToLower(line[10]),
			MemoryAccessFunction: line[11],
			ALUFunction:          line[12],
			ValCFunction:         line[13],
			PositiveOptionPC:     line[14],
			NegativeOptionPC:     line[15],
			EStagePeriod:         line[16],
			DstE:                 line[17],
			DstM:                 line[18],
			M_velE_Source:        line[19],
		}
		actionSet = append(actionSet, data)
	}

	vals := struct {
		PACKAGE   string
		IMPORT    []string
		CONST     map[string]string
		ActionSet []MicroAction
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
