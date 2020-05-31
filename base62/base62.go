package base62

import (
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const encodingChunkSize = 2

// no of bytes required in base62 to represent hex encoded string value of length encodingChunkSize
// given by formula :: int(math.Ceil(math.Log(math.Pow(16, 2*encodingChunkSize)-1) / math.Log(62)))
const decodingChunkSize = 3

func Encode(str string) string {
	var encoded strings.Builder

	inBytes := []byte(str)
	byteLength := len(inBytes)

	for i := 0; i < byteLength; i += encodingChunkSize {
		chunk := inBytes[i:minOf(i+encodingChunkSize, byteLength)]
		s := hex.EncodeToString(chunk)
		val, _ := strconv.ParseUint(s, 16, 64)
		w := padLeft(toBase62(val), "0", decodingChunkSize)
		encoded.WriteString(w)
	}
	return encoded.String()
}

func Decode(encoded string) (string, error) {
	decodedBytes := []byte{}
	for i := 0; i < len(encoded); i += decodingChunkSize {
		chunk := encoded[i:minOf(i+decodingChunkSize, len(encoded))]
		val, err := fromBase62(chunk)
		if err != nil {
			return "", err
		}
		chunkHex := strconv.FormatUint(val, 16)
		dst := make([]byte, hex.DecodedLen(len([]byte(chunkHex))))
		_, err = hex.Decode(dst, []byte(chunkHex))
		if err != nil {
			return "", errors.Wrap(err, "malformed input")
		}
		decodedBytes = append(decodedBytes, dst...)
	}
	s := string(decodedBytes)
	return s, nil
}

func minOf(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func padLeft(str, pad string, length int) string {
	for len(str) < length {
		str = pad + str
	}
	return str
}
