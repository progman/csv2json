//--------------------------------------------------------------------------------------------------------------------------------------------//
// 1.0.0
// Alexey Potehin <gnuplanet@gmail.com>, http://www.gnuplanet.ru/doc/cv
//--------------------------------------------------------------------------------------------------------------------------------------------//
package main
//--------------------------------------------------------------------------------------------------------------------------------------------//
import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
	"fmt"
)
//--------------------------------------------------------------------------------------------------------------------------------------------//
func csv2json(name string) (out []byte, err error) {

	var csvMapLine map[string]string = map[string]string{}
	var csvMapList []map[string]string = []map[string]string{}


	var f *os.File
	f, err = os.Open(name)
	if err != nil {
		return
	}
	r := csv.NewReader(f)


	var header []string
	header, err = r.Read()
	if err == io.EOF {
		return
	}


	var index int
	for {
		index++;

		var record []string
		record, err = r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return
		}

		if len(header) != len(record) {
			err = fmt.Errorf("invalid length of line %d", index)
			return
		}

		for i := 0; i < len(record); i++ {
			csvMapLine[header[i]] = record[i]
		}

		csvMapList = append(csvMapList, csvMapLine)
		csvMapLine = map[string]string{}
	}


	out, err = json.Marshal(csvMapList)
	if err != nil {
		return
	}


	return
}
//--------------------------------------------------------------------------------------------------------------------------------------------//
func help() {
	fmt.Printf("example: %s FILE.CSV\n", os.Args[0])
}
//--------------------------------------------------------------------------------------------------------------------------------------------//
func do_it() (err error) {
	var out []byte

	if len(os.Args) != 2 {
		help()
		return
	}


	out, err = csv2json(os.Args[1])
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}
	fmt.Printf("%s\n", string(out))


	return
}
//--------------------------------------------------------------------------------------------------------------------------------------------//
func main() {

	err := do_it()
	if err != nil {
		os.Exit(1)
	}


	os.Exit(0)
}
//--------------------------------------------------------------------------------------------------------------------------------------------//
