package day_16

import (
	_ "embed"
	"fmt"
	"iter"
	"slices"
	"strconv"
	"strings"
)

//go:embed example_operator
var Input string

type Packet struct {
	Version uint
	TypeID  uint
	Body    string
}

type LiteralPacket struct {
	Packet Packet
	Value  uint
}

type OperatorPacket struct {
	Packet             Packet
	LengthTypeID       uint
	LengthOfSubPackets uint
	NumberOfSubPackets uint
}

const PacketHeaderLength = 6

func ParsePacket(binary string) (packet Packet) {
	version, err := strconv.ParseUint(binary[:3], 2, 0)
	if err != nil {
		panic(err)
	}
	typeID, err := strconv.ParseUint(binary[3:PacketHeaderLength], 2, 0)
	if err != nil {
		panic(err)
	}
	packet = Packet{Version: uint(version), TypeID: uint(typeID), Body: binary[PacketHeaderLength:]}
	// if packet.checkIsOperatorPacket() {
	// 	packet.tryParseOperatorPacket()
	// }
	// fmt.Printf("%+v\n", packet)
	return
}

// func parseLiteralValue(binary string) (uint, uint) {
// 	var values []rune
// 	var i int
// 	var chunkSize = 5
// 	for chunk := range slices.Chunk([]rune(binary), chunkSize) {
// 		isLastGroupBit, value := chunk[0], chunk[1:]
// 		values = append(values, value...)
// 		i += 1
// 		if isLastGroupBit == '0' {
// 			break
// 		}
// 	}

// 	result, err := strconv.ParseUint(string(values), 2, 0)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return uint(result), uint(i * chunkSize)
// }

func ParseOperatorPacket(packet Packet) OperatorPacket {
	operatorPacket := OperatorPacket{Packet: packet}
	lengthTypeId, err := strconv.ParseUint(string(packet.Body[0]), 2, 0)
	if err != nil {
		panic(err)
	}
	operatorPacket.LengthTypeID = uint(lengthTypeId)
	switch lengthTypeId {
	case 0:
		lengthOfSubpackets, err := strconv.ParseUint(packet.Body[1:16], 2, 0)
		if err != nil {
			panic(err)
		}
		// operatorPacket.Packet.Body = packet.Body[16 : lengthOfSubpackets+16]
		operatorPacket.Packet.Body = packet.Body[16:]
		operatorPacket.LengthOfSubPackets = uint(lengthOfSubpackets)
	case 1:
		numberOfSubpackets, err := strconv.ParseUint(packet.Body[1:12], 2, 0)
		if err != nil {
			panic(err)
		}
		operatorPacket.Packet.Body = packet.Body[12:]
		operatorPacket.NumberOfSubPackets = uint(numberOfSubpackets)
	default:
		panic("type not implemented")
	}
	return operatorPacket
}

func (p LiteralPacket) GetHeaderOffset() int {
	return PacketHeaderLength
}

func (p OperatorPacket) GetHeaderOffset() int {
	lengthTypeIdBit := 1
	var lengthTypeId int
	switch p.LengthTypeID {
	case 0:
		lengthTypeId = 15
	case 1:
		lengthTypeId = 11
	default:
		panic("type not implemented")
	}
	return PacketHeaderLength + lengthTypeIdBit + lengthTypeId
}

type PacketReader interface {
	Read() iter.Seq[string]
	GetHeaderOffset() int
}

func (p Packet) ParsePacketReader() PacketReader {
	switch p.TypeID {
	case 4:
		return LiteralPacket{Packet: p}
	default:
		return ParseOperatorPacket(p)
	}
}

func (p OperatorPacket) Read() iter.Seq[string] {
	fmt.Printf("%+v\n", p)
	return func(yield func(string) bool) {
		switch p.LengthTypeID {
		case 0:
			body := p.Packet.Body[0:p.LengthOfSubPackets]
			var leftOffset, rightOffset int
			for {
				reader := ParsePacket(body[leftOffset:]).ParsePacketReader()
				rightOffset += len(strings.Join(slices.Collect(reader.Read()), "")) + reader.GetHeaderOffset()
				if !yield(body[leftOffset:rightOffset]) || rightOffset >= len(body) {
					return
				}
				leftOffset = rightOffset
			}
		case 1:
			body := p.Packet.Body
			var leftOffset, rightOffset int
			for range p.NumberOfSubPackets {
				reader := ParsePacket(body[leftOffset:]).ParsePacketReader()
				rightOffset += len(strings.Join(slices.Collect(reader.Read()), "")) + reader.GetHeaderOffset()
				if !yield(body[leftOffset:rightOffset]) || rightOffset >= len(body) {
					return
				}
				leftOffset = rightOffset
			}
		default:
			panic("type not implemented")
		}
	}
}

func (p LiteralPacket) Read() iter.Seq[string] {
	fmt.Printf("%+v\n", p)
	chunkSize := 5
	return func(yield func(string) bool) {
		for chunk := range slices.Chunk([]rune(p.Packet.Body), chunkSize) {
			isLastGroupBit := chunk[0]
			if !yield(string(chunk)) || isLastGroupBit == '0' {
				return
			}
		}
	}
}

func (p LiteralPacket) Evaluate() {}

func (packet Packet) GetVersionSum() (sum uint) {
	// for l := range ParsePacket("110100101111111000101000").ParsePacketReader().Read() {
	// 	fmt.Println(l)
	// }

	// 110100010100101001000100100
	// 11010001010
	// 0101001000100100
	// ---

	sum += packet.Version
	reader := packet.ParsePacketReader()
	// fmt.Printf("%+v\n", reader)
	for elem := range reader.Read() {
		sum += ParsePacket(elem).Version
		fmt.Println("p:", elem)
	}

	return sum
}

func parseInput(hex string) Packet {
	var binary string
	for _, char := range hex { // strings.TrimRight(hex, "0") {
		elem, err := strconv.ParseUint(string(char), 16, 0)
		if err != nil {
			panic(err)
		}
		binary += fmt.Sprintf("%0.4b", elem)
	}
	return ParsePacket(binary)
}

func SolvePartOne(input string) (result int) {
	packet := parseInput(input)
	result = int(packet.GetVersionSum())
	return
}

func SolvePartTwo(input string) (result int) {
	return
}
