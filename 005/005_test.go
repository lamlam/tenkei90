package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestAnswer(t *testing.T) {
    inbuf := readFile("./sample_input.txt")
    outbuf := stubIO(inbuf, func() {
        main()
    })

    sampleResult := readFile("./sample_result.txt")
    if strings.TrimSuffix(outbuf, "\n") != strings.TrimSuffix(sampleResult, "\n") {
        t.Fatalf("failed with \n %s\n", outbuf)
    }
}

func readFile(fileName string) string {
    bytes, err := ioutil.ReadFile(fileName)
    if err != nil {
        panic(err)
    }

    return string(bytes)
}

func stubIO(inbuf string, fn func()) string {
    inr, inw, _ := os.Pipe()
    outr, outw, _ := os.Pipe()
    _, _ = inw.Write([]byte(inbuf))
    inw.Close()
    origStdin := os.Stdin
    origStdout := os.Stdout
	os.Stdin = inr
    os.Stdout = outw
    fn()
    os.Stdin = origStdin
    os.Stdout = origStdout
    outw.Close()
    outbuf, _ := ioutil.ReadAll(outr)
    return string(outbuf)
}