package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var (
	// TODO: human readable format to snmp format
	stringToBeConverted = flag.String("convert", "", "Hex Representation SNMP DateAndTime. Example: 07 E9 07 0A 0E 1E 2D 00 2B 02 00")
)

func main() {
	flag.Parse()
	if stringToBeConverted == nil {
		panic("--convent cannot be empty")
	}
	s := strings.ReplaceAll(*stringToBeConverted, " ", "")
	if len(s) < 22 || len(s) > 22 {
		panic("Input too short or long for SNMP DateAndTime")
	}

	var hexChunks []string
	for i := 0; i < len(s); i += 2 {
		hexChunks = append(hexChunks, s[i:i+2])
	}

	year := (convertHexToDec("0x"+hexChunks[0]) << 8) + convertHexToDec("0x"+hexChunks[1])
	month := convertHexToDec("0x" + hexChunks[2])
	day := convertHexToDec("0x" + hexChunks[3])
	hour := convertHexToDec("0x" + hexChunks[4])
	minute := convertHexToDec("0x" + hexChunks[5])
	second := convertHexToDec("0x" + hexChunks[6])
	deciSecond := convertHexToDec("0x" + hexChunks[7])

	directionByte, err := hex.DecodeString(hexChunks[8])
	if err != nil {
		panic("Invalid direction byte: " + err.Error())
	}
	if string(directionByte) != "-" && string(directionByte) != "+" {
		panic("Invalid direction byte: " + string(directionByte))
	}
	direction := rune(directionByte[0])

	hourOffset := convertHexToDec("0x" + hexChunks[9])
	minOffset := convertHexToDec("0x" + hexChunks[10])

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%04d-%02d-%02d ", year, month, day))
	sb.WriteString(fmt.Sprintf("%02d:%02d:%02d.%d ", hour, minute, second, deciSecond))
	sb.WriteString(fmt.Sprintf("%c%02d:%02d", direction, hourOffset, minOffset))

	fmt.Println("Parsed SNMP DateTime:")
	fmt.Println(sb.String())
}

func convertHexToDec(hexStr string) int64 {
	val, err := strconv.ParseInt(hexStr, 0, 64)
	if err != nil {
		panic("Failed to parse hex: " + hexStr + " - " + err.Error())
	}
	return val
}
