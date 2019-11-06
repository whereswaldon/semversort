package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/blang/semver"
)

func main() {
	var (
		tolerant bool
	)
	flag.BoolVar(&tolerant, "tolerant", false, "allow versions that do not perfectly conform to semver spec")
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)
	versions := make([]semver.Version, 0, 10)
	for {
		// read a semver from stdin
		line, err := reader.ReadString([]byte("\n")[0])
		if len(line) > 0 {
			// remove trailing newlines from input
			line = line[:len(line)-1]
			if len(line) > 0 {
				// remove leading "v" if present
				if line[0] == "v"[0] {
					line = line[1:]
				}
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("failed to read input line: %v", err)
		}

		// use tolerant parse function if requested
		parseFunc := semver.Parse
		if tolerant {
			parseFunc = semver.ParseTolerant
		}

		ver, err := parseFunc(line)
		if err != nil {
			log.Fatalf("failed to parse semver \"%s\": %v", line, err)
		}
		// collect this version
		versions = append(versions, ver)
	}
	// sort all versions
	semver.Sort(versions)
	// emit sorted on stdout
	for _, v := range versions {
		fmt.Printf("%s\n", v)
	}
}
