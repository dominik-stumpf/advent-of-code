package day_16

import (
	"cmp"
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
	TypeID  TypeID
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
	Children           []PacketReader
}

type TypeID uint

const (
	TypeSum TypeID = iota
	TypeProduct
	TypeMinimum
	TypeMaximum
	TypeLiteralValue
	TypeGreaterThan
	TypeLessThan
	TypeEqualTo
)

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
	packet = Packet{Version: uint(version), TypeID: TypeID(typeID), Body: binary[PacketHeaderLength:]}
	// fmt.Printf("%+v\n", packet)
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
	case TypeLiteralValue:
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
	fmt.Printf("literalValue: %d\n", p.EvalValue())
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
		panic("length type not implemented")
	}
}

func (p LiteralPacket) GetBodyLength() int {
	var length int
	for chunk := range slices.Chunk([]rune(p.Packet.Body), LiteralPacketChunkSize) {
		isLastGroupBit := chunk[0]
		length += 1
		if isLastGroupBit == '0' {
			break
		}
	}
	return length * LiteralPacketChunkSize
}

type PacketReader interface {
	Read(subPackets *[]Packet) iter.Seq[string]
	GetHeaderOffset() int
	GetBodyLength() int
	EvalValue() int
}

func (p OperatorPacket) EvalValue() (result int) {
	switch p.Packet.TypeID {
	case TypeSum:
		for _, child := range p.Children {
			result += child.EvalValue()
		}
	case TypeProduct:
		if len(p.Children) > 0 {
			result = 1
		}
		for _, child := range p.Children {
			result *= child.EvalValue()
		}
	case TypeMinimum:
		result = slices.MinFunc(p.Children, func(a, b PacketReader) int {
			return cmp.Compare(a.EvalValue(), b.EvalValue())
		}).EvalValue()
	case TypeMaximum:
		result = slices.MaxFunc(p.Children, func(a, b PacketReader) int {
			return cmp.Compare(a.EvalValue(), b.EvalValue())
		}).EvalValue()
	case TypeGreaterThan:
		if p.Children[0].EvalValue() > p.Children[1].EvalValue() {
			result = 1
		}
	case TypeLessThan:
		if p.Children[0].EvalValue() < p.Children[1].EvalValue() {
			result = 1
		}
	case TypeEqualTo:
		if p.Children[0].EvalValue() == p.Children[1].EvalValue() {
			result = 1
		}
	default:
		panic("packet type not implemented")
	}
	return result
}

func (p LiteralPacket) EvalValue() int {
	var values []rune
	var i int
	// for chunk := range p.Read(&[]Packet{}) {
	for chunk := range slices.Chunk([]rune(p.Packet.Body), LiteralPacketChunkSize) {
		isLastGroupBit, value := chunk[0], chunk[1:]
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
	return int(result)
}

func (p Packet) GetVersionSum() (sum uint) {
	reader := p.ParsePacketReader()
	subPackets := []Packet{}
	for range reader.Read(&subPackets) {
	}
	for _, subPacket := range subPackets {
		sum += subPacket.Version
	}
	return sum
}

func (p *OperatorPacket) PopulatePacketTree() {
	switch p.LengthTypeID {
	case 0:
		body := p.Packet.Body[0:p.LengthOfSubPackets]
		var leftOffset, rightOffset int
		for {
			reader := ParsePacket(body[leftOffset:]).ParsePacketReader()
			child, ok := reader.(OperatorPacket)
			// fmt.Printf("length: %v, %p %p\n", reader, &reader, &child)
			if ok {
				child.PopulatePacketTree()
				p.Children = append(p.Children, child)
			} else {
				p.Children = append(p.Children, reader)
			}
			rightOffset += reader.GetBodyLength() + reader.GetHeaderOffset()
			if rightOffset >= len(body) {
				break
			}
			leftOffset = rightOffset
		}
	case 1:
		body := p.Packet.Body
		var leftOffset, rightOffset int
		for range p.NumberOfSubPackets {
			reader := ParsePacket(body[leftOffset:]).ParsePacketReader()
			child, ok := reader.(OperatorPacket)
			// fmt.Println("number:", reader)
			if ok {
				child.PopulatePacketTree()
				p.Children = append(p.Children, child)
			} else {
				p.Children = append(p.Children, reader)
			}
			rightOffset += reader.GetBodyLength() + reader.GetHeaderOffset()
			leftOffset = rightOffset
		}
	default:
		panic("length type not implemented")
	}
}

// {Packet:{Version:0 TypeID:1 Body:1011000011001110001001} LengthTypeID:0 LengthOfSubPackets:22 NumberOfSubPackets:0 Children:[]}
// {Packet:{Version:5 TypeID:4 Body:0011001110001001}}
// literalValue: 6
// {Packet:{Version:3 TypeID:4 Body:01001}}
// literalValue: 9
func (p Packet) EvalPacketExpression() int {
	reader := p.ParsePacketReader()

	// var evalByTraverse func(OperatorPacket) int
	// evalByTraverse = func(parent OperatorPacket) (result int) {
	// 	for _, child := range parent.Children {
	// 		// fmt.Printf("%+v\n", child)
	// 		switch p := child.(type) {
	// 		case OperatorPacket:
	// 			result += evalByTraverse(p)
	// 		case LiteralPacket:
	// 			// fmt.Println(p.EvalLiteralValue())
	// 			result += p.EvalValue()
	// 		default:
	// 			panic("packet type not implemented")
	// 		}
	// 	}

	// 	return result
	// }

	switch p := reader.(type) {
	case OperatorPacket:
		p.PopulatePacketTree()
		return p.EvalValue()
		// j, _ := json.Marshal(p)
		// os.WriteFile("sample.json", j, 0644)
		// return evalByTraverse(p)
	case LiteralPacket:
		return int(p.EvalValue())
	default:
		panic("packet type not implemented")
	}
}

func parseInput(hex string) Packet {
	var binary string
	for _, char := range strings.TrimRight(hex, "0") {
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
	packet := parseInput(input)
	result = packet.EvalPacketExpression()
	return
}
