package main

import (
	"bytes"
	"fmt"
	"net/http"
	core "rediss/core"
	"strings"

	"github.com/rs/cors"
)

type InMemoryBufferConn struct {
	*bytes.Buffer
}

func (m *InMemoryBufferConn) Write(p []byte) (n int, err error) {
	return m.Buffer.Write(p)
}

func (m *InMemoryBufferConn) Read(p []byte) (n int, err error) {
	return m.Buffer.Read(p)
}

type SetRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	EX    int64  `json:"ex,omitempty"`
}

func executeCommand(cmdName string, args []string) (interface{}, error) {
	var encodedArgs []string
	encodedArgs = append(encodedArgs, cmdName)
	encodedArgs = append(encodedArgs, args...)

	var buf []byte
	argBuf := bytes.NewBuffer(buf)
	argBuf.Write([]byte(fmt.Sprintf("*%d\r\n", len(encodedArgs))))
	for _, arg := range encodedArgs {
		argBuf.Write([]byte(fmt.Sprintf("$%d\r\n%s\r\n", len(arg), arg)))
	}

	rawCmdBytes := argBuf.Bytes()

	decodedCmds, err := core.Decode(rawCmdBytes)
	if err != nil {
		return nil, err
	}

	if len(decodedCmds) == 0 {
		return nil, fmt.Errorf("failed to decode command")
	}

	var redisCmds core.RedisCmds
	for _, dCmd := range decodedCmds {
		arr, ok := dCmd.([]interface{})
		if !ok || len(arr) == 0 {
			continue
		}
		cmdStr, ok := arr[0].(string)
		if !ok {
			continue
		}
		var argsList []string
		for i := 1; i < len(arr); i++ {
			if s, ok := arr[i].(string); ok {
				argsList = append(argsList, s)
			}
		}
		redisCmds = append(redisCmds, &core.RedisCmd{
			Cmd:  strings.ToUpper(cmdStr),
			Args: argsList,
		})
	}

	responseBuffer := &InMemoryBufferConn{Buffer: bytes.NewBuffer(nil)}
	core.EvalAndRespond(redisCmds, responseBuffer)

	respBytes := responseBuffer.Bytes()
	if len(respBytes) == 0 {
		return nil, fmt.Errorf("empty response from evaluator")
	}

	decodedResponse, _, err := core.DecodeOne(respBytes)
	return decodedResponse, err
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/get/", handleGet)
	mux.HandleFunc("/set", handleSet)
	mux.HandleFunc("/del/", handleDel)
	mux.HandleFunc("/incr/", handleIncr)
	mux.HandleFunc("/ttl/", handleTtl)
	mux.HandleFunc("/stats", handleStats)
	mux.HandleFunc("/getall", handleGetAll)
	mux.HandleFunc("/seed", handleSeeding)
	mux.HandleFunc("/flushall", handleFlushAll)

	fmt.Println("Proxy server running on port 8001...")
	handler := cors.New(
		cors.Options{
			AllowedOrigins: []string{
				"http://localhost:3000",
				"http://localhost:300",
			},
			AllowedMethods: []string{
				"GET",
				"POST",
				"PUT",
				"DELETE",
				"OPTIONS",
				"PATCH",
				"HEAD",
			},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: true,
			Debug:            true,
		},
	).Handler(mux)
	http.ListenAndServe(":8001", handler)
}
