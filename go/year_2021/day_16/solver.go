package day_16

import (
	"cmp"
	_ "embed"
	"fmt"
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
	Children           []Packeter
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

func (p Packet) ParsePacketer() Packeter {
	switch p.TypeID {
	case TypeLiteralValue:
		return LiteralPacket{Packet: p}
	default:
		return ParseOperatorPacket(p)
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
			packeter := ParsePacket(body[leftOffset:]).ParsePacketer()
			rightOffset += packeter.GetBodyLength() + packeter.GetHeaderOffset()
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

func (p LiteralPacket) GetPacket() Packet {
	return p.Packet
}

func (p OperatorPacket) GetPacket() Packet {
	return p.Packet
}

type Packeter interface {
	GetHeaderOffset() int
	GetBodyLength() int
	GetPacket() Packet
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
		result = slices.MinFunc(p.Children, func(a, b Packeter) int {
			return cmp.Compare(a.EvalValue(), b.EvalValue())
		}).EvalValue()
	case TypeMaximum:
		result = slices.MaxFunc(p.Children, func(a, b Packeter) int {
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
	packeter := p.ParsePacketer()
	switch p := packeter.(type) {
	case OperatorPacket:
		p.PopulatePacketTree()
		queue := []Packeter{p}
		for len(queue) > 0 {
			first := queue[0]
			queue = queue[1:]
			sum += first.GetPacket().Version
			if p, ok := first.(OperatorPacket); ok {
				queue = append(queue, p.Children...)
			}
		}
	case LiteralPacket:
		sum = p.Packet.Version
	default:
		panic("packet type not implemented")
	}
	return
}

func (p *OperatorPacket) PopulatePacketTree() {
	switch p.LengthTypeID {
	case 0:
		body := p.Packet.Body[0:p.LengthOfSubPackets]
		var leftOffset, rightOffset int
		for {
			packeter := ParsePacket(body[leftOffset:]).ParsePacketer()
			child, ok := packeter.(OperatorPacket)
			if ok {
				child.PopulatePacketTree()
				p.Children = append(p.Children, child)
			} else {
				p.Children = append(p.Children, packeter)
			}
			rightOffset += packeter.GetBodyLength() + packeter.GetHeaderOffset()
			if rightOffset >= len(body) {
				break
			}
			leftOffset = rightOffset
		}
	case 1:
		body := p.Packet.Body
		var leftOffset, rightOffset int
		for range p.NumberOfSubPackets {
			packeter := ParsePacket(body[leftOffset:]).ParsePacketer()
			child, ok := packeter.(OperatorPacket)
			if ok {
				child.PopulatePacketTree()
				p.Children = append(p.Children, child)
			} else {
				p.Children = append(p.Children, packeter)
			}
			rightOffset += packeter.GetBodyLength() + packeter.GetHeaderOffset()
			leftOffset = rightOffset
		}
	default:
		panic("length type not implemented")
	}
}

func (p Packet) EvalPacketExpression() int {
	packeter := p.ParsePacketer()

	switch p := packeter.(type) {
	case OperatorPacket:
		p.PopulatePacketTree()
		return p.EvalValue()
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
