package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"

	"golang.org/x/mod/semver"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	versions := make([]string, 0, 10)
	for {
		// read a semver from stdin
		line, err := reader.ReadString([]byte("\n")[0])
		if len(line) > 0 {
			// remove trailing newlines from input
			line = line[:len(line)-1]
			if len(line) > 0 {
				// add leading "v" if not present
				if line[0] != "v"[0] {
					line = fmt.Sprintf("v%s", line)
				}
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("failed to read input line: %v", err)
		}

		if !semver.IsValid(line) {
			log.Fatalf("failed to parse semver \"%s\": %v", line, err)
		}
		// collect this version
		versions = append(versions, line)
	}
	// sort all versions
	sort.Slice(versions, func(i, j int) bool {
		return semver.Compare(versions[i], versions[j]) < 0
	})
	// emit sorted on stdout
	for _, v := range versions {
		fmt.Printf("%s\n", v)
	}
}
