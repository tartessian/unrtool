package crypt

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"strconv"
	"unicode/utf16"

	"github.com/tartessian/unrtool/internal/crypt/xor"
)

const NO_CRYPT = -1
const HEADER_SIZE = 28
const HEADER_PREFIX = "Lineage2Ver"

func readHeader(input io.Reader) (int, error) {
	header := make([]byte, HEADER_SIZE)
	if _, err := io.ReadFull(input, header); err != nil {
		return NO_CRYPT, err
	}

	utf16Header := make([]uint16, HEADER_SIZE/2)
	binary.Read(bytes.NewReader(header), binary.LittleEndian, &utf16Header)
	headerStr := string(utf16.Decode(utf16Header))
	if !bytes.HasPrefix([]byte(headerStr), []byte(HEADER_PREFIX)) {
		err := errors.New("missing " + HEADER_PREFIX + " prefix")
		return NO_CRYPT, err
	}

	versionStr := headerStr[len(HEADER_PREFIX):]
	version, err := strconv.Atoi(versionStr)
	if err != nil {
		return NO_CRYPT, err
	}

	return version, nil
}

func Decrypt(input io.Reader, filename string) (io.Reader, error) {
	version, err := readHeader(input)
	if err != nil {
		return nil, err
	}

	switch version {
	case 111, 121:
		var xorKey int
		if version == 111 {
			xorKey = xor.XOR_KEY_111
		} else {
			xorKey = xor.GetXORKey121(filename)
		}
		return xor.NewL2Ver1x1InputStream(input, xorKey), nil
	default:
		return nil, errors.New("unsupported crypt version: " + strconv.Itoa(version))
	}
}
