package main

import "fmt"

func hexDump(data string) string {
	const bytesPerLine = 16
	var dump string

	for i := 0; i < len(data); i += bytesPerLine {
		// Print the hex values
		for j := i; j < i+bytesPerLine && j < len(data); j++ {
			dump += fmt.Sprintf("%02x ", data[j])
		}

		// Pad the line if it's shorter than bytesPerLine
		if i+bytesPerLine > len(data) {
			for j := 0; j < bytesPerLine-(len(data)-i); j++ {
				dump += "   "
			}
		}

		dump += " "

		// Print ASCII characters (if printable)
		for j := i; j < i+bytesPerLine && j < len(data); j++ {
			if data[j] >= 32 && data[j] <= 126 {
				dump += fmt.Sprintf("%c", data[j])
			} else {
				dump += "."
			}
		}

		dump += "\n"
	}

	return dump
}
