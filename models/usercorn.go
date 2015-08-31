package models

import (
    uc "github.com/lunixbochs/unicorn"
)

// TODO upstream Unicorn interface into bindings
type Unicorn interface {
    MemMap(addr, size uint64) error
    MemMapProt(addr, size uint64, prot int) error
    MemRead(addr, size uint64) ([]byte, error)
    MemReadInto(dst []byte, addr uint64) error
    MemReadStr(addr uint64) (string, error)
    MemWrite(addr uint64, data []byte) error
    RegRead(reg int) (uint64, error)
    RegWrite(reg int, value uint64) error
    Start(begin, until uint64) error
    StartWithOptions(begin, until uint64, options *uc.UcOptions) error
    Stop() error
}

type Usercorn interface {
    Unicorn
    Mmap(addr, size uint64) (uint64, error)
    Pop() (uint64, error)
    Push(n uint64) error
    PackAddr(buf []byte, n uint64) error
    UnpackAddr(buf []byte) uint64
}
