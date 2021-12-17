package main

import (
	"AdventOfCode2021/utils"
	"flag"
	"fmt"
	"math"
	"strconv"
)

var (
	useTestInputs = flag.Bool("test", false, "Should we use the test inputs")
)

func init() {
	flag.Parse()
}

func main() {
	var hexString string
	if *useTestInputs {
		hexString = "9C0141080250320F1802104A08"
	} else {
		lines, _ := utils.ReadInputFile(16)
		hexString = lines[0]
	}

	binaryString := ""

	for i := 0; i < len(hexString); i += 1 {
		chunk := string(hexString[i])
		number, _ := strconv.ParseUint(chunk, 16, 64)
		binaryString += fmt.Sprintf("%04b", number)
	}

	utils.PrintDayResultsWithDuration(16, part1(binaryString), part2(binaryString))
}

func countVersions(b string, subPackets map[string][]string) (int, string, int) {
	version, _ := strconv.ParseInt(b[0:3], 2, 32)
	typeId, _ := strconv.ParseInt(b[3:6], 2, 32)

	index := 6

	if typeId == 4 {
		litervalValueString := ""
		for {
			nextBit := string(b[index])

			litervalValueString += string(b[index+1 : index+5])

			index += 5
			if nextBit == "0" {
				break
			}
		}

		subPackets[b] = []string{}

		return int(version), b[index:], index
	} else {
		lengthTypeId := string(b[index])

		if lengthTypeId == "0" {
			bitsLengthBin := string(b[index+1 : index+16])
			bitLength, _ := strconv.ParseInt(bitsLengthBin, 2, 32)

			totalVersions := int(version)
			remaining := b[index+16:]

			totalBitsRead := 0

			var sPackets []string

			for {
				sPackets = append(sPackets, remaining)
				subPacketVersion, newRemaining, bitsRead := countVersions(remaining, subPackets)
				totalVersions += subPacketVersion
				totalBitsRead += bitsRead
				remaining = newRemaining

				if totalBitsRead == int(bitLength) {
					break
				}
			}
			subPackets[b] = sPackets

			return int(totalVersions), b[index+totalBitsRead+16:], index + totalBitsRead + 16
		}

		if lengthTypeId == "1" {
			numberOfPacketsBin := string(b[index+1 : index+12])
			numberOfPackets, _ := strconv.ParseInt(numberOfPacketsBin, 2, 32)

			totalVersions := int(version)
			remaining := b[index+12:]

			numberOfPacketsRead := 0
			totalBitsRead := 0

			var sPackets []string
			for {
				sPackets = append(sPackets, remaining)
				subPacketVersion, newRemaining, bitsRead := countVersions(remaining, subPackets)
				totalVersions += subPacketVersion
				remaining = newRemaining
				totalBitsRead += bitsRead
				numberOfPacketsRead += 1

				if numberOfPacketsRead == int(numberOfPackets) {
					break
				}
			}

			subPackets[b] = sPackets
			return int(totalVersions), b[index+totalBitsRead+12:], index + totalBitsRead + 12
		}
	}

	return 0, "", 0
}

func part1(binaryString string) utils.ResultWithTime {
	subPackets := make(map[string][]string)
	totalVersions, _, _ := countVersions(binaryString, subPackets)

	return utils.ResultWithTime{
		Value: totalVersions,
	}
}

func part2(binaryString string) utils.ResultWithTime {
	subPackets := make(map[string][]string)
	countVersions(binaryString, subPackets)

	value := packetValue(binaryString, subPackets)

	fmt.Printf("%v\n", value)

	return utils.ResultWithTime{
		Value: 0,
	}
}

func decodeLiteral(packet string) int64 {
	index := 6
	litervalValueString := ""

	for {
		nextBit := string(packet[index])

		litervalValueString += string(packet[index+1 : index+5])

		if nextBit == "0" {
			break
		} else {
			index += 5
		}
	}

	litervalValue, _ := strconv.ParseInt(litervalValueString, 2, 64)

	return litervalValue
}

func packetValue(packet string, subPacketsMap map[string][]string) int64 {
	typeId, _ := strconv.ParseInt(packet[3:6], 2, 32)

	if typeId == 0 {
		sum := int64(0)

		subPackets := subPacketsMap[packet]

		for _, subPacket := range subPackets {
			sum += packetValue(subPacket, subPacketsMap)
		}

		return sum
	} else if typeId == 1 {
		product := int64(1)

		subPackets := subPacketsMap[packet]

		for _, subPacket := range subPackets {
			product *= packetValue(subPacket, subPacketsMap)
		}

		return product
	} else if typeId == 2 {
		minValue := int64(math.MaxInt64)

		subPackets := subPacketsMap[packet]

		for _, subPacket := range subPackets {
			value := packetValue(subPacket, subPacketsMap)

			if value < minValue {
				minValue = value
			}
		}

		return minValue
	} else if typeId == 3 {
		maxValue := int64(math.MinInt64)

		subPackets := subPacketsMap[packet]

		for _, subPacket := range subPackets {
			value := packetValue(subPacket, subPacketsMap)

			if value > maxValue {
				maxValue = value
			}
		}

		return maxValue
	} else if typeId == 4 {
		return decodeLiteral(packet)
	} else if typeId == 5 {
		subPackets := subPacketsMap[packet]

		firstPacketValue := packetValue(subPackets[0], subPacketsMap)
		secondPacketValue := packetValue(subPackets[1], subPacketsMap)

		if firstPacketValue > secondPacketValue {
			return 1
		} else {
			return 0
		}
	} else if typeId == 6 {
		subPackets := subPacketsMap[packet]

		firstPacketValue := packetValue(subPackets[0], subPacketsMap)
		secondPacketValue := packetValue(subPackets[1], subPacketsMap)

		if firstPacketValue < secondPacketValue {
			return 1
		} else {
			return 0
		}
	} else if typeId == 7 {
		subPackets := subPacketsMap[packet]

		firstPacketValue := packetValue(subPackets[0], subPacketsMap)
		secondPacketValue := packetValue(subPackets[1], subPacketsMap)

		if firstPacketValue == secondPacketValue {
			return 1
		} else {
			return 0
		}
	} else {
		fmt.Printf("New type ID found! %v\n", typeId)
	}

	return 0
}
