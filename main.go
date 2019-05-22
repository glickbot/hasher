// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code is a prototype and not engineered for production use.
// Error handling is incomplete or inappropriate for usage beyond
// a development sample.

package main

import (
    "bufio"
    "fmt"
    "os"
    "flag"
    "strings"
    "github.com/howeyc/crc16"
    "github.com/dgryski/go-farm"
    "crypto/sha1"
)

func main() {
    hashType := flag.String("type", "[farm|crc|sha1]", "hash type")
    flag.Parse()
    var result func(string) string
    switch option := strings.ToLower(*hashType); option {
        case "farm":
	    result = func(text string) string {
		    return fmt.Sprintf("%d",farm.Hash64([]byte(text)))
	    }
	case "crc":
	    result = func(text string) string {
		    return fmt.Sprintf("%d",crc16.Checksum([]byte(text),crc16.MakeTable(1024)))
	    }
        case "sha1":
	    result = func(text string) string {
		    return fmt.Sprintf("%x", sha1.Sum([]byte(text)))
	    }
	default:
	    fmt.Println("unknown options, please use --type [farm|crc|sha1|pass]")
	    os.Exit(1)
    }
    input, err := getStdin()
    if err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
    fmt.Println(result(input))
}

func getStdin() (string,error) {
    scanner := bufio.NewScanner(os.Stdin)
    text := ""
    for scanner.Scan() {
        text += scanner.Text()
    }
    if err := scanner.Err(); err != nil {
	return "", err
    }
    return text, nil
}
