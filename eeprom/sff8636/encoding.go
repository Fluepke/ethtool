package SFF8636

import (
	"fmt"
)

type Encoding byte

const (
	EncodingUnspecified    Encoding = 0x00
	Encoding8b10b          Encoding = 0x01
	Encoding4b5b           Encoding = 0x02
	EncodingNrz            Encoding = 0x03
	EncodingSonetScrambled Encoding = 0x04
	Encoding64b66b         Encoding = 0x05
	EncodingManchester     Encoding = 0x06
	Encoding256b           Encoding = 0x07
	EncodingPam4           Encoding = 0x08
)

func (e Encoding) String() string {
	mapping := map[Encoding]string{
		EncodingUnspecified:    "Unspecified",
		Encoding8b10b:          "8B/10B",
		Encoding4b5b:           "4B/5B",
		EncodingNrz:            "NRZ",
		EncodingManchester:     "Manchester",
		EncodingSonetScrambled: "SONET Scrambled",
		Encoding64b66b:         "64B/66B",
		Encoding256b:           "256B/257B (transcoded FEC-enabled data)",
		EncodingPam4:           "PAM4",
	}

	str, found := mapping[e]
	if found {
		return str
	}
	return "Invalid or unknown"
}

func (e Encoding) MarshalJson() map[string]interface{} {
	return map[string]interface{}{
		"ascii": e.String(),
		"hex":   fmt.Sprintf("%#02x", byte(e)),
	}
}