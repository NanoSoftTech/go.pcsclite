// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs -- -I/usr/include/PCSC const.go

package scard

const (
	ATTR_VENDOR_NAME              uint32 = 0x10100
	ATTR_VENDOR_IFD_TYPE          uint32 = 0x10101
	ATTR_VENDOR_IFD_VERSION       uint32 = 0x10102
	ATTR_VENDOR_IFD_SERIAL_NO     uint32 = 0x10103
	ATTR_CHANNEL_ID               uint32 = 0x20110
	ATTR_ASYNC_PROTOCOL_TYPES     uint32 = 0x30120
	ATTR_DEFAULT_CLK              uint32 = 0x30121
	ATTR_MAX_CLK                  uint32 = 0x30122
	ATTR_DEFAULT_DATA_RATE        uint32 = 0x30123
	ATTR_MAX_DATA_RATE            uint32 = 0x30124
	ATTR_MAX_IFSD                 uint32 = 0x30125
	ATTR_SYNC_PROTOCOL_TYPES      uint32 = 0x30126
	ATTR_POWER_MGMT_SUPPORT       uint32 = 0x40131
	ATTR_USER_TO_CARD_AUTH_DEVICE uint32 = 0x50140
	ATTR_USER_AUTH_INPUT_DEVICE   uint32 = 0x50142
	ATTR_CHARACTERISTICS          uint32 = 0x60150
	ATTR_CURRENT_PROTOCOL_TYPE    uint32 = 0x80201
	ATTR_CURRENT_CLK              uint32 = 0x80202
	ATTR_CURRENT_F                uint32 = 0x80203
	ATTR_CURRENT_D                uint32 = 0x80204
	ATTR_CURRENT_N                uint32 = 0x80205
	ATTR_CURRENT_W                uint32 = 0x80206
	ATTR_CURRENT_IFSC             uint32 = 0x80207
	ATTR_CURRENT_IFSD             uint32 = 0x80208
	ATTR_CURRENT_BWT              uint32 = 0x80209
	ATTR_CURRENT_CWT              uint32 = 0x8020a
	ATTR_CURRENT_EBC_ENCODING     uint32 = 0x8020b
	ATTR_EXTENDED_BWT             uint32 = 0x8020c
	ATTR_ICC_PRESENCE             uint32 = 0x90300
	ATTR_ICC_INTERFACE_STATUS     uint32 = 0x90301
	ATTR_CURRENT_IO_STATE         uint32 = 0x90302
	ATTR_ATR_STRING               uint32 = 0x90303
	ATTR_ICC_TYPE_PER_ATR         uint32 = 0x90304
	ATTR_ESC_RESET                uint32 = 0x7a000
	ATTR_ESC_CANCEL               uint32 = 0x7a003
	ATTR_ESC_AUTHREQUEST          uint32 = 0x7a005
	ATTR_MAXINPUT                 uint32 = 0x7a007
	ATTR_DEVICE_UNIT              uint32 = 0x7fff0001
	ATTR_DEVICE_IN_USE            uint32 = 0x7fff0002
	ATTR_DEVICE_FRIENDLY_NAME     uint32 = 0x7fff0003
	ATTR_DEVICE_SYSTEM_NAME       uint32 = 0x7fff0004
	ATTR_SUPRESS_T1_IFS_REQUEST   uint32 = 0x7fff0007
)

const (
	S_SUCCESS                 = scardError(0x0)
	F_INTERNAL_ERROR          = scardError(0x80100001)
	E_CANCELLED               = scardError(0x80100002)
	E_INVALID_HANDLE          = scardError(0x80100003)
	E_INVALID_PARAMETER       = scardError(0x80100004)
	E_INVALID_TARGET          = scardError(0x80100005)
	E_NO_MEMORY               = scardError(0x80100006)
	F_WAITED_TOO_LONG         = scardError(0x80100007)
	E_INSUFFICIENT_BUFFER     = scardError(0x80100008)
	E_UNKNOWN_READER          = scardError(0x80100009)
	E_TIMEOUT                 = scardError(0x8010000a)
	E_SHARING_VIOLATION       = scardError(0x8010000b)
	E_NO_SMARTCARD            = scardError(0x8010000c)
	E_UNKNOWN_CARD            = scardError(0x8010000d)
	E_CANT_DISPOSE            = scardError(0x8010000e)
	E_PROTO_MISMATCH          = scardError(0x8010000f)
	E_NOT_READY               = scardError(0x80100010)
	E_INVALID_VALUE           = scardError(0x80100011)
	E_SYSTEM_CANCELLED        = scardError(0x80100012)
	F_COMM_ERROR              = scardError(0x80100013)
	F_UNKNOWN_ERROR           = scardError(0x80100014)
	E_INVALID_ATR             = scardError(0x80100015)
	E_NOT_TRANSACTED          = scardError(0x80100016)
	E_READER_UNAVAILABLE      = scardError(0x80100017)
	P_SHUTDOWN                = scardError(0x80100018)
	E_PCI_TOO_SMALL           = scardError(0x80100019)
	E_READER_UNSUPPORTED      = scardError(0x8010001a)
	E_DUPLICATE_READER        = scardError(0x8010001b)
	E_CARD_UNSUPPORTED        = scardError(0x8010001c)
	E_NO_SERVICE              = scardError(0x8010001d)
	E_SERVICE_STOPPED         = scardError(0x8010001e)
	E_UNEXPECTED              = scardError(0x8010001f)
	E_UNSUPPORTED_FEATURE     = scardError(0x8010001f)
	E_ICC_INSTALLATION        = scardError(0x80100020)
	E_ICC_CREATEORDER         = scardError(0x80100021)
	E_FILE_NOT_FOUND          = scardError(0x80100024)
	E_NO_DIR                  = scardError(0x80100025)
	E_NO_FILE                 = scardError(0x80100026)
	E_NO_ACCESS               = scardError(0x80100027)
	E_WRITE_TOO_MANY          = scardError(0x80100028)
	E_BAD_SEEK                = scardError(0x80100029)
	E_INVALID_CHV             = scardError(0x8010002a)
	E_UNKNOWN_RES_MNG         = scardError(0x8010002b)
	E_NO_SUCH_CERTIFICATE     = scardError(0x8010002c)
	E_CERTIFICATE_UNAVAILABLE = scardError(0x8010002d)
	E_NO_READERS_AVAILABLE    = scardError(0x8010002e)
	E_COMM_DATA_LOST          = scardError(0x8010002f)
	E_NO_KEY_CONTAINER        = scardError(0x80100030)
	E_SERVER_TOO_BUSY         = scardError(0x80100031)
	W_UNSUPPORTED_CARD        = scardError(0x80100065)
	W_UNRESPONSIVE_CARD       = scardError(0x80100066)
	W_UNPOWERED_CARD          = scardError(0x80100067)
	W_RESET_CARD              = scardError(0x80100068)
	W_REMOVED_CARD            = scardError(0x80100069)
	W_SECURITY_VIOLATION      = scardError(0x8010006a)
	W_WRONG_CHV               = scardError(0x8010006b)
	W_CHV_BLOCKED             = scardError(0x8010006c)
	W_EOF                     = scardError(0x8010006d)
	W_CANCELLED_BY_USER       = scardError(0x8010006e)
	W_CARD_NOT_AUTHENTICATED  = scardError(0x8010006f)
)

type Protocol uint32

const (
	PROTOCOL_UNDEFINED Protocol = 0x0
	PROTOCOL_T0        Protocol = 0x1
	PROTOCOL_T1        Protocol = 0x2
	PROTOCOL_ANY       Protocol = PROTOCOL_T0 | PROTOCOL_T1
)

type ShareMode uint32

const (
	SHARE_EXCLUSIVE ShareMode = 0x1
	SHARE_SHARED    ShareMode = 0x2
	SHARE_DIRECT    ShareMode = 0x3
)

type Disposition uint32

const (
	LEAVE_CARD   Disposition = 0x0
	RESET_CARD   Disposition = 0x1
	UNPOWER_CARD Disposition = 0x2
	EJECT_CARD   Disposition = 0x3
)

type State uint32

const (
	UNKNOWN    State = 0x1
	ABSENT     State = 0x2
	PRESENT    State = 0x4
	SWALLOWED  State = 0x8
	POWERED    State = 0x10
	NEGOTIABLE State = 0x20
	SPECIFIC   State = 0x40
)

type StateFlag uint32

const (
	STATE_UNAWARE     StateFlag = 0x0
	STATE_IGNORE      StateFlag = 0x1
	STATE_CHANGED     StateFlag = 0x2
	STATE_UNKNOWN     StateFlag = 0x4
	STATE_UNAVAILABLE StateFlag = 0x8
	STATE_EMPTY       StateFlag = 0x10
	STATE_PRESENT     StateFlag = 0x20
	STATE_ATRMATCH    StateFlag = 0x40
	STATE_EXCLUSIVE   StateFlag = 0x80
	STATE_INUSE       StateFlag = 0x100
	STATE_MUTE        StateFlag = 0x200
	STATE_UNPOWERED   StateFlag = 0x400
)

type Timeout uint32

const (
	INFINITE Timeout = 0xffffffff
)

const (
	MAX_BUFFER_SIZE          = 0x108
	MAX_BUFFER_SIZE_EXTENDED = 0x1000c
	MAX_READERNAME           = 0x80
	MAX_ATR_SIZE             = 0x21
)
