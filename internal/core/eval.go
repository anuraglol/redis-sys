package core

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"
)

var RESP_NIL []byte = []byte("$-1\r\n")
var RESP_OK []byte = []byte("+OK\r\n")
var RESP_ZERO []byte = []byte(":0\r\n")
var RESP_ONE []byte = []byte(":1\r\n")
var RESP_MINUS_1 []byte = []byte(":-1\r\n")
var RESP_MINUS_2 []byte = []byte(":-2\r\n")

func evalPING(args []string) []byte {
	var b []byte

	if len(args) >= 2 {
		return Encode(errors.New("ERR wrong number of arguments for 'ping' command"), false)
	}

	if len(args) == 0 {
		b = Encode("PONG", true)
	} else {
		b = Encode(args[0], false)
	}

	return b
}

func evalSET(args []string) []byte {
	if len(args) <= 1 {
		return Encode(errors.New("ERR wrong number of arguments for 'set' command"), false)
	}

	var key, value string
	var exDurationMs int64 = -1

	key, value = args[0], args[1]
	oType, oEnc := deduceTypeEncoding(value)

	for i := 2; i < len(args); i++ {
		switch args[i] {
		case "EX", "ex":
			i++
			if i == len(args) {
				return Encode(errors.New("ERR syntax error"), false)
			}

			exDurationSec, err := strconv.ParseInt(args[3], 10, 64)
			if err != nil {
				return Encode(errors.New("ERR value is not an integer or out of range"), false)
			}
			exDurationMs = exDurationSec * 1000
		default:
			return Encode(errors.New("ERR syntax error"), false)
		}
	}

	Put(key, NewObj(value, exDurationMs, oType, oEnc))
	return RESP_OK
}

func evalGET(args []string) []byte {
	if len(args) != 1 {
		return Encode(errors.New("ERR wrong number of arguments for 'get' command"), false)
	}

	var key string = args[0]

	obj := Get(key)

	if obj == nil {
		return RESP_NIL
	}

	if hasExpired(obj) {
		return RESP_NIL
	}

	return Encode(obj.Value, false)
}

func evalTTL(args []string) []byte {
	if len(args) != 1 {
		return Encode(errors.New("ERR wrong number of arguments for 'ttl' command"), false)
	}

	var key string = args[0]

	obj := Get(key)

	if obj == nil {
		return RESP_MINUS_2
	}

	exp, isExpirySet := getExpiry(obj)
	if !isExpirySet {
		return RESP_MINUS_1
	}

	if exp < uint64(time.Now().UnixMilli()) {
		return RESP_MINUS_2
	}

	durationMs := exp - uint64(time.Now().UnixMilli())

	return Encode(int64(durationMs/1000), false)
}

func evalDEL(args []string) []byte {
	var countDeleted int = 0

	for _, key := range args {
		if ok := Del(key); ok {
			countDeleted++
		}
	}

	return Encode(countDeleted, false)
}

func evalEXPIRE(args []string) []byte {
	if len(args) <= 1 {
		return Encode(errors.New("ERR wrong number of arguments for 'expire' command"), false)
	}

	var key string = args[0]
	exDurationSec, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return Encode(errors.New("ERR value is not an integer or out of range"), false)
	}

	obj := Get(key)

	if obj == nil {
		return RESP_ZERO
	}

	setExpiry(obj, exDurationSec*1000)

	return RESP_ONE
}

func evalBGREWRITEAOF(args []string) []byte {
	DumpAllAOF()
	return RESP_OK
}

func evalINCR(args []string) []byte {
	if len(args) != 1 {
		return Encode(errors.New("ERR wrong number of arguments for 'incr' command"), false)
	}

	var key string = args[0]
	obj := Get(key)
	if obj == nil {
		obj = NewObj("0", -1, OBJ_TYPE_STRING, OBJ_ENCODING_INT)
		Put(key, obj)
	}

	if err := assertType(obj.TypeEncoding, OBJ_TYPE_STRING); err != nil {
		return Encode(err, false)
	}

	if err := assertEncoding(obj.TypeEncoding, OBJ_ENCODING_INT); err != nil {
		return Encode(err, false)
	}

	i, _ := strconv.ParseInt(obj.Value.(string), 10, 64)
	i++
	obj.Value = strconv.FormatInt(i, 10)

	return Encode(i, false)
}

func evalCLIENT(args []string) []byte {
	return RESP_OK
}

func evalLATENCY(args []string) []byte {
	return Encode([]string{}, false)
}

func evalLRU(args []string) []byte {
	evictAllkeysLRU()
	return RESP_OK
}

func evalSTATS(args []string) []byte {
	s := GetStats()

	arr := []string{
		"total_commands", strconv.FormatInt(s.TotalCommands, 10),
		"set_cmds", strconv.FormatInt(s.SetCmds, 10),
		"get_cmds", strconv.FormatInt(s.GetCmds, 10),
		"ping_cmds", strconv.FormatInt(s.PingCmds, 10),
		"expire_cmds", strconv.FormatInt(s.ExpireCmds, 10),
	}

	return Encode(arr, false)
}

func evalGETALL(args []string) []byte {
	if len(args) != 0 {
		return Encode(errors.New("ERR wrong number of arguments for 'get all' command"), false)
	}

	arr := []string{}
	for key, obj := range store {
		valStr := obj.String()

		var ttlStr string
		if exp, exists := expires[obj]; exists {
			ttlStr = fmt.Sprintf("%d", exp)
		} else {
			ttlStr = "-1"
		}
		arr = append(arr, key, valStr, ttlStr)
	}

	return Encode(arr, false)
}

func evalFLUSHALL(args []string) []byte {
	if len(args) != 0 {
		return Encode(errors.New("ERR wrong number of arguments for 'flush all' command"), false)
	}

	count := 0
	for key := range store {
		if ok := Del(key); ok {
			count++
		}
	}

	return RESP_OK
}

func EvalAndRespond(cmds RedisCmds, c io.ReadWriter) {
	var response []byte
	buf := bytes.NewBuffer(response)

	for _, cmd := range cmds {
		IncrTotalCommands()
		switch cmd.Cmd {
		case "PING":
			IncrPingCmds()
			buf.Write(evalPING(cmd.Args))
		case "SET":
			IncrSetCmds()
			buf.Write(evalSET(cmd.Args))
		case "GET":
			IncrGetCmds()
			buf.Write(evalGET(cmd.Args))
		case "TTL":
			buf.Write(evalTTL(cmd.Args))
		case "DEL":
			buf.Write(evalDEL(cmd.Args))
		case "EXPIRE":
			IncrExpireCmds()
			buf.Write(evalEXPIRE(cmd.Args))
		case "BGREWRITEAOF":
			buf.Write(evalBGREWRITEAOF(cmd.Args))
		case "INCR":
			buf.Write(evalINCR(cmd.Args))
		case "CLIENT":
			buf.Write(evalCLIENT(cmd.Args))
		case "LATENCY":
			buf.Write(evalLATENCY(cmd.Args))
		case "LRU":
			buf.Write(evalLRU(cmd.Args))
		case "STATS":
			buf.Write(evalSTATS(cmd.Args))
		case "GETALL":
			buf.Write(evalGETALL(cmd.Args))
		case "FLUSHALL":
			buf.Write(evalFLUSHALL(cmd.Args))
		default:
			buf.Write(evalPING(cmd.Args))
		}
	}
	c.Write(buf.Bytes())
}
