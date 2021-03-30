# automatically generated by the FlatBuffers compiler, do not modify

# namespace: fbs

import flatbuffers
from flatbuffers.compat import import_numpy
np = import_numpy()

class VirtualSegment(object):
    __slots__ = ['_tab']

    # VirtualSegment
    def Init(self, buf, pos):
        self._tab = flatbuffers.table.Table(buf, pos)

    # VirtualSegment
    def Start(self): return self._tab.Get(flatbuffers.number_types.Uint64Flags, self._tab.Pos + flatbuffers.number_types.UOffsetTFlags.py_type(0))
    # VirtualSegment
    def Size(self): return self._tab.Get(flatbuffers.number_types.Uint64Flags, self._tab.Pos + flatbuffers.number_types.UOffsetTFlags.py_type(8))

def CreateVirtualSegment(builder, start, size):
    builder.Prep(8, 16)
    builder.PrependUint64(size)
    builder.PrependUint64(start)
    return builder.Offset()


class VirtualSegmentT(object):

    # VirtualSegmentT
    def __init__(self):
        self.start = 0  # type: int
        self.size = 0  # type: int

    @classmethod
    def InitFromBuf(cls, buf, pos):
        virtualSegment = VirtualSegment()
        virtualSegment.Init(buf, pos)
        return cls.InitFromObj(virtualSegment)

    @classmethod
    def InitFromObj(cls, virtualSegment):
        x = VirtualSegmentT()
        x._UnPack(virtualSegment)
        return x

    # VirtualSegmentT
    def _UnPack(self, virtualSegment):
        if virtualSegment is None:
            return
        self.start = virtualSegment.Start()
        self.size = virtualSegment.Size()

    # VirtualSegmentT
    def Pack(self, builder):
        return CreateVirtualSegment(builder, self.start, self.size)