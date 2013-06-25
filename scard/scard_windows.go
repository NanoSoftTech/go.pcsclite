package scard

import (
	"bytes"
	"fmt"
	"syscall"
	"unsafe"
)

var (
	modwinscard          = syscall.NewLazyDLL("winscard.dll")
	procEstablishContext = modwinscard.NewProc("SCardEstablishContext")
	procRelease          = modwinscard.NewProc("SCardReleaseContext")
	procListReaders      = modwinscard.NewProc("SCardListReadersA")
	procConnect          = modwinscard.NewProc("SCardConnectA")
	procDisconnect	     = modwinscard.NewProc("SCardDisconnect")
	procStatus	     = modwinscard.NewProc("SCardStatusA")
	procGetAttrib	     = modwinscard.NewProc("SCardGetAttrib")
)

type Context struct {
	ctx uintptr
}

type Card struct {
	handle uintptr
	activeProtocol uintptr
}

func (e scardError) Error() string {
	//return "scard: " + C.GoString(C.pcsc_stringify_error(C.LONG(e)))
	return fmt.Sprintf("scard: error code %x", e)
}

// wraps SCardEstablishContext
func EstablishContext() (*Context, error) {
	var ctx Context

	r, _, _ := procEstablishContext.Call(2, uintptr(0), uintptr(0), uintptr(unsafe.Pointer(&ctx.ctx)))
	if scardError(r) != S_SUCCESS {
		return nil, scardError(r)
	}

	return &ctx, nil
}

// wraps SCardIsValidContext
func (ctx *Context) IsValid() (bool, error) {
	panic("scard: not implemented") // TODO
}

// wraps SCardCancel
func (ctx *Context) Cancel() error {
	panic("scard: not implemented") // TODO
}

// wraps SCardReleaseContext
func (ctx *Context) Release() error {
	r, _, _ := procRelease.Call(uintptr(ctx.ctx))
	if scardError(r) != S_SUCCESS {
		return scardError(r)
	}
	return nil
}

// wraps SCardListReaders
func (ctx *Context) ListReaders() ([]string, error) {
	var needed uint32

	r, _, _ := procListReaders.Call(
		ctx.ctx,
		0,
		0,
		uintptr(unsafe.Pointer(&needed)))
	if scardError(r) != S_SUCCESS {
		return nil, scardError(r)
	}

	data := make([]byte, needed)

	r, _, _ = procListReaders.Call(
		ctx.ctx, 
		0, 
		uintptr(unsafe.Pointer(&data[0])), 
		uintptr(unsafe.Pointer(&needed)))
	if scardError(r) != S_SUCCESS {
		return nil, scardError(r)
	}

	var readers []string
	for _, b := range bytes.Split(data, []byte{0}) {
		if len(b) > 0 {
			readers = append(readers, string(b))
		}
	}

	return readers, nil
}

// wraps SCardListReaderGroups
func (ctx *Context) ListReaderGroups() ([]string, error) {
	panic("scard: not implemented")
}

// wraps SCardGetStatusChange
func (ctx *Context) GetStatusChange(readerStates []ReaderState, timeout Timeout) error {
	panic("scard: not implemented")
}

// wraps SCardConnect
func (ctx *Context) Connect(reader string, mode ShareMode, proto Protocol) (*Card, error) {
	var card Card

	creader := []byte(reader)
	creader = append(creader, 0)

	r, _, _ := procConnect.Call(
		ctx.ctx, 
		uintptr(unsafe.Pointer(&creader[0])),
		uintptr(mode),
		uintptr(proto),
		uintptr(unsafe.Pointer(&card.handle)),
		uintptr(unsafe.Pointer(&card.activeProtocol)))

	if scardError(r) != S_SUCCESS {
		return nil, scardError(r)
	}

	return &card, nil
}

// wraps SCardDisconnect
func (card *Card) Disconnect(d Disposition) error {
	r, _, _ := procDisconnect.Call(card.handle, uintptr(d))
	if scardError(r) != S_SUCCESS {
		return scardError(r) 
	}
	return nil
}

// wraps SCardReconnect
func (card *Card) Reconnect(mode ShareMode, protocol Protocol, init Disposition) error {
	panic("scard: not implemented")
}

// wraps SCardBeginTransaction
func (card *Card) BeginTransaction() error {
	panic("scard: not implemented")
}

// wraps SCardEndTransaction
func (card *Card) EndTransaction(d Disposition) error {
	panic("scard: not implemented")
}

// wraps SCardStatus
func (card *Card) Status() (*CardStatus, error) {
	var reader [MAX_READERNAME + 1]byte
	var readerLen = uint32(len(reader))
	var state, proto uint32
	var atr [MAX_ATR_SIZE]byte
	var atrLen = uint32(len(atr))

	r, _, _ := procStatus.Call(
		card.handle,
		uintptr(unsafe.Pointer(&reader[0])),
		uintptr(unsafe.Pointer(&readerLen)),
		uintptr(unsafe.Pointer(&state)),
		uintptr(unsafe.Pointer(&proto)),
		uintptr(unsafe.Pointer(&atr[0])),
		uintptr(unsafe.Pointer(&atrLen)))
	if scardError(r) != S_SUCCESS {
		return nil, scardError(r)
	}

	status := &CardStatus {
		Reader:         string(reader[0:readerLen]),
		State:          State(state),
		ActiveProtocol: Protocol(proto),
		ATR:            atr[0:atrLen],
	}

	return status, nil
}

// wraps SCardTransmit
func (card *Card) Transmit(cmd []byte) ([]byte, error) {
	panic("scard: not implemented") // TODO
}

// wraps SCardControl
func (card *Card) Control(ctrl uint32, cmd []byte) ([]byte, error) {
	panic("scard: not implemented") // TODO
}

// wraps SCardGetAttrib
func (card *Card) GetAttrib(id uint32) ([]byte, error) {
	var needed uintptr

	r, _, _ := procGetAttrib.Call(
		card.handle, 
		uintptr(id),
		0,
		uintptr(unsafe.Pointer(&needed)))
	if scardError(r) != S_SUCCESS {
		return nil, scardError(r)
	}

	var attrib = make([]byte, needed)

	r, _, _ = procGetAttrib.Call(
		card.handle, 
		uintptr(id),
		uintptr(unsafe.Pointer(&attrib[0])),
		uintptr(unsafe.Pointer(&needed)))
	if scardError(r) != S_SUCCESS {
		return nil, scardError(r)
	}

	return attrib[0:needed], nil
}

// wraps SCardSetAttrib
func (card *Card) SetAttrib(id uint32, data []byte) error {
	panic("scard: not implemented") // TODO
}
