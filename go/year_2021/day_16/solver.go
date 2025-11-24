package day_16

import (
	_ "embed"
	"fmt"
	"iter"
	"slices"
	"strconv"
	"strings"
)

//go:embed input
var Input string

type Packet struct {
	Version uint
	TypeID  uint
	Body    string
}

type LiteralPacket struct {
	Packet Packet
}

type OperatorPacket struct {
	Packet             Packet
	LengthTypeID       uint
	LengthOfSubPackets uint
	NumberOfSubPackets uint
}

const PacketHeaderLength = 6
const LiteralPacketChunkSize = 5

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
	return
}

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

func (p Packet) ParsePacketReader() PacketReader {
	switch p.TypeID {
	case 4:
		return LiteralPacket{Packet: p}
	default:
		return ParseOperatorPacket(p)
	}
}

func (p OperatorPacket) Read(subPackets *[]Packet) iter.Seq[string] {
	fmt.Printf("%+v\n", p)
	*subPackets = append(*subPackets, p.Packet)
	return func(yield func(string) bool) {
		switch p.LengthTypeID {
		case 0:
			body := p.Packet.Body[0:p.LengthOfSubPackets]
			var leftOffset, rightOffset int
			for {
				reader := ParsePacket(body[leftOffset:]).ParsePacketReader()
				rightOffset += len(strings.Join(slices.Collect(reader.Read(subPackets)), "")) + reader.GetHeaderOffset()
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
				rightOffset += len(strings.Join(slices.Collect(reader.Read(subPackets)), "")) + reader.GetHeaderOffset()
				if !yield(body[leftOffset:rightOffset]) {
					return
				}
				leftOffset = rightOffset
			}
		default:
			panic("type not implemented")
		}
	}
}

func (p LiteralPacket) Read(subPackets *[]Packet) iter.Seq[string] {
	fmt.Printf("%+v\n", p)
	*subPackets = append(*subPackets, p.Packet)
	return func(yield func(string) bool) {
		for chunk := range slices.Chunk([]rune(p.Packet.Body), LiteralPacketChunkSize) {
			isLastGroupBit := chunk[0]
			if !yield(string(chunk)) || isLastGroupBit == '0' {
				return
			}
		}
	}
}

func (p OperatorPacket) GetBodyLength() int {
	fmt.Printf("%+v\n", p)
	switch p.LengthTypeID {
	case 0:
		return int(p.LengthOfSubPackets)
	case 1:
		body := p.Packet.Body
		var leftOffset, rightOffset int
		for range p.NumberOfSubPackets {
			reader := ParsePacket(body[leftOffset:]).ParsePacketReader()
			rightOffset += reader.GetBodyLength() + reader.GetHeaderOffset()
			leftOffset = rightOffset
		}
		return rightOffset
	default:
		panic("type not implemented")
	}
}

func (p LiteralPacket) GetBodyLength() int {
	fmt.Printf("%+v\n", p)
	return len(p.Packet.Body) / LiteralPacketChunkSize
}

type PacketReader interface {
	Read(subPackets *[]Packet) iter.Seq[string]
	GetHeaderOffset() int
	GetBodyLength() int
}

func (p LiteralPacket) Evaluate() uint {
	var values []rune
	var i int
	for chunk := range p.Read(&[]Packet{}) {
		runes := []rune(chunk)
		isLastGroupBit, value := runes[0], runes[1:]
		values = append(values, value...)
		i += 1
		if isLastGroupBit == '0' {
			break
		}
	}
	result, err := strconv.ParseUint(string(values), 2, 0)
	if err != nil {
		panic(err)
	}
	return uint(result)
}

func (packet Packet) GetVersionSum() (sum uint) {
	// for l := range ParsePacket("110100101111111000101000").ParsePacketReader().Read() {
	// 	fmt.Println(l)
	// }

	// 110100010100101001000100100
	// 11010001010
	// 0101001000100100
	// ---

	reader := packet.ParsePacketReader()
	subPackets := []Packet{}
	for range reader.Read(&subPackets) {
	}
	for _, subPacket := range subPackets {
		sum += subPacket.Version
	}
	// for elem := range reader.Read(&subPackets) {
	// 	sum += ParsePacket(elem).Version
	// 	fmt.Println("p:", elem)
	// }

	// reader := packet.ParsePacketReader()
	// fmt.Println(reader.GetBodyLength())

	return sum

	// fmt.Printf("%+v\n", reader)
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
