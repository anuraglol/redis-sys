package core

import "strconv"

func deduceTypeEncoding(v string) (uint8, uint8) {
	oType := OBJ_TYPE_STRING
	if _, err := strconv.ParseInt(v, 10, 64); err == nil { // if we're able to conv to string
		return oType, OBJ_ENCODING_INT
	}
	if len(v) <= 44 { // embedded string
		return oType, OBJ_ENCODING_EMBSTR
	}
	return oType, OBJ_ENCODING_RAW // else raw
}
