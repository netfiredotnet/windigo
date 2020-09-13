/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package co

// WM_GETDLGCODE return value.
type DLGC uint32

const (
	DLGC_WANTARROWS      DLGC = 0x0001 // Control wants arrow keys
	DLGC_WANTTAB         DLGC = 0x0002 // Control wants tab keys
	DLGC_WANTALLKEYS     DLGC = 0x0004 // Control wants all keys
	DLGC_WANTMESSAGE     DLGC = 0x0004 // Pass message to control
	DLGC_HASSETSEL       DLGC = 0x0008 // Understands EM_SETSEL message
	DLGC_DEFPUSHBUTTON   DLGC = 0x0010 // Default pushbutton
	DLGC_UNDEFPUSHBUTTON DLGC = 0x0020 // Non-default pushbutton
	DLGC_RADIOBUTTON     DLGC = 0x0040 // Radio button
	DLGC_WANTCHARS       DLGC = 0x0080 // Want WM_CHAR messages
	DLGC_STATIC          DLGC = 0x0100 // Static item: don't include
	DLGC_BUTTON          DLGC = 0x2000 // Button item: can be checked
)

// SetProcessDpiAwarenessContext() value.
type DPI_AWARE_CTX int32

const (
	DPI_AWARE_CTX_UNAWARE           DPI_AWARE_CTX = -1
	DPI_AWARE_CTX_SYSTEM_AWARE      DPI_AWARE_CTX = -2
	DPI_AWARE_CTX_PER_MON_AWARE     DPI_AWARE_CTX = -3
	DPI_AWARE_CTX_PER_MON_AWARE_V2  DPI_AWARE_CTX = -4
	DPI_AWARE_CTX_UNAWARE_GDISCALED DPI_AWARE_CTX = -5
)

// Composes SFGAO.
type DROPEFFECT uint32

const (
	DROPEFFECT_NONE   DROPEFFECT = 0
	DROPEFFECT_COPY   DROPEFFECT = 1
	DROPEFFECT_MOVE   DROPEFFECT = 2
	DROPEFFECT_LINK   DROPEFFECT = 4
	DROPEFFECT_SCROLL DROPEFFECT = 0x80000000
)

// DateTimePicker messages.
type DTM WM

const (
	_DTM_FIRST DTM = 0x1000

	DTM_GETSYSTEMTIME         DTM = _DTM_FIRST + 1
	DTM_SETSYSTEMTIME         DTM = _DTM_FIRST + 2
	DTM_GETRANGE              DTM = _DTM_FIRST + 3
	DTM_SETRANGE              DTM = _DTM_FIRST + 4
	DTM_SETFORMAT             DTM = _DTM_FIRST + 50
	DTM_SETMCCOLOR            DTM = _DTM_FIRST + 6
	DTM_GETMCCOLOR            DTM = _DTM_FIRST + 7
	DTM_GETMONTHCAL           DTM = _DTM_FIRST + 8
	DTM_SETMCFONT             DTM = _DTM_FIRST + 9
	DTM_GETMCFONT             DTM = _DTM_FIRST + 10
	DTM_SETMCSTYLE            DTM = _DTM_FIRST + 11
	DTM_GETMCSTYLE            DTM = _DTM_FIRST + 12
	DTM_CLOSEMONTHCAL         DTM = _DTM_FIRST + 13
	DTM_GETDATETIMEPICKERINFO DTM = _DTM_FIRST + 14
	DTM_GETIDEALSIZE          DTM = _DTM_FIRST + 15
)

// DateTimePicker notifications, sent via WM_NOTIFY.
type DTN NM

const (
	_DTN_FIRST  DTN = -740
	_DTN_FIRST2 DTN = -753

	DTN_CLOSEUP        DTN = _DTN_FIRST2 - 0
	DTN_DROPDOWN       DTN = _DTN_FIRST2 - 1
	DTN_DATETIMECHANGE DTN = _DTN_FIRST2 - 6
	DTN_FORMATQUERY    DTN = _DTN_FIRST - 2
	DTN_FORMAT         DTN = _DTN_FIRST - 3
	DTN_WMKEYDOWN      DTN = _DTN_FIRST - 4
	DTN_USERSTRING     DTN = _DTN_FIRST - 5
)

// DateTimePicker styles.
type DTS WS

const (
	DTS_UPDOWN                 DTS = 0x0001
	DTS_SHOWNONE               DTS = 0x0002
	DTS_SHORTDATEFORMAT        DTS = 0x0000
	DTS_LONGDATEFORMAT         DTS = 0x0004
	DTS_SHORTDATECENTURYFORMAT DTS = 0x000C
	DTS_TIMEFORMAT             DTS = 0x0009
	DTS_APPCANPARSE            DTS = 0x0010
	DTS_RIGHTALIGN             DTS = 0x0020
)

// Edit control messages.
type EM WM

const (
	EM_GETSEL              EM = 0x00B0
	EM_SETSEL              EM = 0x00B1
	EM_GETRECT             EM = 0x00B2
	EM_SETRECT             EM = 0x00B3
	EM_SETRECTNP           EM = 0x00B4
	EM_SCROLL              EM = 0x00B5
	EM_LINESCROLL          EM = 0x00B6
	EM_SCROLLCARET         EM = 0x00B7
	EM_GETMODIFY           EM = 0x00B8
	EM_SETMODIFY           EM = 0x00B9
	EM_GETLINECOUNT        EM = 0x00BA
	EM_LINEINDEX           EM = 0x00BB
	EM_SETHANDLE           EM = 0x00BC
	EM_GETHANDLE           EM = 0x00BD
	EM_GETTHUMB            EM = 0x00BE
	EM_LINELENGTH          EM = 0x00C1
	EM_REPLACESEL          EM = 0x00C2
	EM_GETLINE             EM = 0x00C4
	EM_LIMITTEXT           EM = 0x00C5
	EM_CANUNDO             EM = 0x00C6
	EM_UNDO                EM = 0x00C7
	EM_FMTLINES            EM = 0x00C8
	EM_LINEFROMCHAR        EM = 0x00C9
	EM_SETTABSTOPS         EM = 0x00CB
	EM_SETPASSWORDCHAR     EM = 0x00CC
	EM_EMPTYUNDOBUFFER     EM = 0x00CD
	EM_GETFIRSTVISIBLELINE EM = 0x00CE
	EM_SETREADONLY         EM = 0x00CF
	EM_SETWORDBREAKPROC    EM = 0x00D0
	EM_GETWORDBREAKPROC    EM = 0x00D1
	EM_GETPASSWORDCHAR     EM = 0x00D2
	EM_SETMARGINS          EM = 0x00D3
	EM_GETMARGINS          EM = 0x00D4
	EM_SETLIMITTEXT        EM = EM_LIMITTEXT
	EM_GETLIMITTEXT        EM = 0x00D5
	EM_POSFROMCHAR         EM = 0x00D6
	EM_CHARFROMPOS         EM = 0x00D7
	EM_SETIMESTATUS        EM = 0x00D8
	EM_GETIMESTATUS        EM = 0x00D9
)

// NMLVEMPTYMARKUP dwFlags.
type EMF uint32

const (
	EMF_NULL     EMF = 0x00000000
	EMF_CENTERED EMF = 0x00000001
)

// WM_ENDSESSION event.
type ENDSESSION uint32

const (
	ENDSESSION_RESTARTORSHUTDOWN ENDSESSION = 0
	ENDSESSION_CLOSEAPP          ENDSESSION = 0x00000001
	ENDSESSION_CRITICAL          ENDSESSION = 0x40000000
	ENDSESSION_LOGOFF            ENDSESSION = 0x80000000
)

// Edit control styles.
type ES WS

const (
	ES_LEFT        ES = 0x0000
	ES_CENTER      ES = 0x0001
	ES_RIGHT       ES = 0x0002
	ES_MULTILINE   ES = 0x0004
	ES_UPPERCASE   ES = 0x0008
	ES_LOWERCASE   ES = 0x0010
	ES_PASSWORD    ES = 0x0020
	ES_AUTOVSCROLL ES = 0x0040
	ES_AUTOHSCROLL ES = 0x0080
	ES_NOHIDESEL   ES = 0x0100
	ES_OEMCONVERT  ES = 0x0400
	ES_READONLY    ES = 0x0800
	ES_WANTRETURN  ES = 0x1000
	ES_NUMBER      ES = 0x2000
)

// WM_APPCOMMAND input event.
type FAPPCOMMAND uint32

const (
	FAPPCOMMAND_MOUSE FAPPCOMMAND = 0x8000
	FAPPCOMMAND_KEY   FAPPCOMMAND = 0
	FAPPCOMMAND_OEM   FAPPCOMMAND = 0x1000
)

// CreateFile() dwFlagsAndAttributes.
type FILE_ATTRIBUTE uint32

const (
	FILE_ATTRIBUTE_INVALID               FILE_ATTRIBUTE = 0xFFFFFFFF // -1
	FILE_ATTRIBUTE_READONLY              FILE_ATTRIBUTE = 0x00000001
	FILE_ATTRIBUTE_HIDDEN                FILE_ATTRIBUTE = 0x00000002
	FILE_ATTRIBUTE_SYSTEM                FILE_ATTRIBUTE = 0x00000004
	FILE_ATTRIBUTE_DIRECTORY             FILE_ATTRIBUTE = 0x00000010
	FILE_ATTRIBUTE_ARCHIVE               FILE_ATTRIBUTE = 0x00000020
	FILE_ATTRIBUTE_DEVICE                FILE_ATTRIBUTE = 0x00000040
	FILE_ATTRIBUTE_NORMAL                FILE_ATTRIBUTE = 0x00000080
	FILE_ATTRIBUTE_TEMPORARY             FILE_ATTRIBUTE = 0x00000100
	FILE_ATTRIBUTE_SPARSE_FILE           FILE_ATTRIBUTE = 0x00000200
	FILE_ATTRIBUTE_REPARSE_POINT         FILE_ATTRIBUTE = 0x00000400
	FILE_ATTRIBUTE_COMPRESSED            FILE_ATTRIBUTE = 0x00000800
	FILE_ATTRIBUTE_OFFLINE               FILE_ATTRIBUTE = 0x00001000
	FILE_ATTRIBUTE_NOT_CONTENT_INDEXED   FILE_ATTRIBUTE = 0x00002000
	FILE_ATTRIBUTE_ENCRYPTED             FILE_ATTRIBUTE = 0x00004000
	FILE_ATTRIBUTE_INTEGRITY_STREAM      FILE_ATTRIBUTE = 0x00008000
	FILE_ATTRIBUTE_VIRTUAL               FILE_ATTRIBUTE = 0x00010000
	FILE_ATTRIBUTE_NO_SCRUB_DATA         FILE_ATTRIBUTE = 0x00020000
	FILE_ATTRIBUTE_EA                    FILE_ATTRIBUTE = 0x00040000
	FILE_ATTRIBUTE_PINNED                FILE_ATTRIBUTE = 0x00080000
	FILE_ATTRIBUTE_UNPINNED              FILE_ATTRIBUTE = 0x00100000
	FILE_ATTRIBUTE_RECALL_ON_OPEN        FILE_ATTRIBUTE = 0x00040000
	FILE_ATTRIBUTE_RECALL_ON_DATA_ACCESS FILE_ATTRIBUTE = 0x00400000
)

// CreateFile() dwCreationDisposition.
type FILE_DISPO uint32

const (
	FILE_DISPO_CREATE_ALWAYS     FILE_DISPO = 2
	FILE_DISPO_CREATE_NEW        FILE_DISPO = 1
	FILE_DISPO_OPEN_ALWAYS       FILE_DISPO = 4
	FILE_DISPO_OPEN_EXISTING     FILE_DISPO = 3
	FILE_DISPO_TRUNCATE_EXISTING FILE_DISPO = 5
)

// CreateFile() dwFlagsAndAttributes.
type FILE_FLAG uint32

const (
	FILE_FLAG_NONE                  FILE_FLAG = 0
	FILE_FLAG_WRITE_THROUGH         FILE_FLAG = 0x80000000
	FILE_FLAG_OVERLAPPED            FILE_FLAG = 0x40000000
	FILE_FLAG_NO_BUFFERING          FILE_FLAG = 0x20000000
	FILE_FLAG_RANDOM_ACCESS         FILE_FLAG = 0x10000000
	FILE_FLAG_SEQUENTIAL_SCAN       FILE_FLAG = 0x08000000
	FILE_FLAG_DELETE_ON_CLOSE       FILE_FLAG = 0x04000000
	FILE_FLAG_BACKUP_SEMANTICS      FILE_FLAG = 0x02000000
	FILE_FLAG_POSIX_SEMANTICS       FILE_FLAG = 0x01000000
	FILE_FLAG_SESSION_AWARE         FILE_FLAG = 0x00800000
	FILE_FLAG_OPEN_REPARSE_POINT    FILE_FLAG = 0x00200000
	FILE_FLAG_OPEN_NO_RECALL        FILE_FLAG = 0x00100000
	FILE_FLAG_FIRST_PIPE_INSTANCE   FILE_FLAG = 0x00080000
	FILE_FLAG_OPEN_REQUIRING_OPLOCK FILE_FLAG = 0x00040000
)

// MapViewOfFile() dwDesiredAccess.
type FILE_MAP uint32

const (
	FILE_MAP_WRITE           FILE_MAP = FILE_MAP(SECTION_MAP_WRITE)
	FILE_MAP_READ            FILE_MAP = FILE_MAP(SECTION_MAP_READ)
	FILE_MAP_ALL_ACCESS      FILE_MAP = FILE_MAP(SECTION_ALL_ACCESS)
	FILE_MAP_EXECUTE         FILE_MAP = FILE_MAP(SECTION_MAP_EXECUTE_EXPLICIT)
	FILE_MAP_COPY            FILE_MAP = 0x00000001
	FILE_MAP_RESERVE         FILE_MAP = 0x80000000
	FILE_MAP_TARGETS_INVALID FILE_MAP = 0x40000000
	FILE_MAP_LARGE_PAGES     FILE_MAP = 0x20000000
)

// SetFilePointerEx() dwMoveMethod.
type FILE_SETPTR uint32

const (
	FILE_SETPTR_BEGIN   FILE_SETPTR = 0
	FILE_SETPTR_CURRENT FILE_SETPTR = 1
	FILE_SETPTR_END     FILE_SETPTR = 2
)

// CreateFile() dwShareMode.
type FILE_SHARE uint32

const (
	FILE_SHARE_NONE   FILE_SHARE = 0
	FILE_SHARE_READ   FILE_SHARE = 0x00000001
	FILE_SHARE_WRITE  FILE_SHARE = 0x00000002
	FILE_SHARE_DELETE FILE_SHARE = 0x00000004
)

// LOGFONT lfWeight.
type FW uint32

const (
	FW_DONTCARE   FW = 0
	FW_THIN       FW = 100
	FW_EXTRALIGHT FW = 200
	FW_ULTRALIGHT FW = FW_EXTRALIGHT
	FW_LIGHT      FW = 300
	FW_NORMAL     FW = 400
	FW_REGULAR    FW = 400
	FW_MEDIUM     FW = 500
	FW_SEMIBOLD   FW = 600
	FW_DEMIBOLD   FW = FW_SEMIBOLD
	FW_BOLD       FW = 700
	FW_EXTRABOLD  FW = 800
	FW_ULTRABOLD  FW = FW_EXTRABOLD
	FW_HEAVY      FW = 900
	FW_BLACK      FW = FW_HEAVY
)
