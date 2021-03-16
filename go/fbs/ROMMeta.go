// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ROMMetaT struct {
	BuildInfo *BuildInfoT
	LinkerInfo *LinkerMetaT
	Symbols []*SymbolT
	State MetaState
}

func (t *ROMMetaT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	buildInfoOffset := t.BuildInfo.Pack(builder)
	linkerInfoOffset := t.LinkerInfo.Pack(builder)
	symbolsOffset := flatbuffers.UOffsetT(0)
	if t.Symbols != nil {
		symbolsLength := len(t.Symbols)
		symbolsOffsets := make([]flatbuffers.UOffsetT, symbolsLength)
		for j := 0; j < symbolsLength; j++ {
			symbolsOffsets[j] = t.Symbols[j].Pack(builder)
		}
		ROMMetaStartSymbolsVector(builder, symbolsLength)
		for j := symbolsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(symbolsOffsets[j])
		}
		symbolsOffset = builder.EndVector(symbolsLength)
	}
	ROMMetaStart(builder)
	ROMMetaAddBuildInfo(builder, buildInfoOffset)
	ROMMetaAddLinkerInfo(builder, linkerInfoOffset)
	ROMMetaAddSymbols(builder, symbolsOffset)
	ROMMetaAddState(builder, t.State)
	return ROMMetaEnd(builder)
}

func (rcv *ROMMeta) UnPackTo(t *ROMMetaT) {
	t.BuildInfo = rcv.BuildInfo(nil).UnPack()
	t.LinkerInfo = rcv.LinkerInfo(nil).UnPack()
	symbolsLength := rcv.SymbolsLength()
	t.Symbols = make([]*SymbolT, symbolsLength)
	for j := 0; j < symbolsLength; j++ {
		x := Symbol{}
		rcv.Symbols(&x, j)
		t.Symbols[j] = x.UnPack()
	}
	t.State = rcv.State()
}

func (rcv *ROMMeta) UnPack() *ROMMetaT {
	if rcv == nil { return nil }
	t := &ROMMetaT{}
	rcv.UnPackTo(t)
	return t
}

type ROMMeta struct {
	_tab flatbuffers.Table
}

func GetRootAsROMMeta(buf []byte, offset flatbuffers.UOffsetT) *ROMMeta {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ROMMeta{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ROMMeta) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ROMMeta) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ROMMeta) BuildInfo(obj *BuildInfo) *BuildInfo {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(BuildInfo)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ROMMeta) LinkerInfo(obj *LinkerMeta) *LinkerMeta {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(LinkerMeta)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ROMMeta) Symbols(obj *Symbol, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ROMMeta) SymbolsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ROMMeta) State() MetaState {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return MetaState(rcv._tab.GetInt32(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *ROMMeta) MutateState(n MetaState) bool {
	return rcv._tab.MutateInt32Slot(10, int32(n))
}

func ROMMetaStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func ROMMetaAddBuildInfo(builder *flatbuffers.Builder, buildInfo flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(buildInfo), 0)
}
func ROMMetaAddLinkerInfo(builder *flatbuffers.Builder, linkerInfo flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(linkerInfo), 0)
}
func ROMMetaAddSymbols(builder *flatbuffers.Builder, symbols flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(symbols), 0)
}
func ROMMetaStartSymbolsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ROMMetaAddState(builder *flatbuffers.Builder, state MetaState) {
	builder.PrependInt32Slot(3, int32(state), 0)
}
func ROMMetaEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}