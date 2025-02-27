package co

// ACCELL fVirt.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-accel
type ACCELF uint8

const (
	ACCELF_NONE    ACCELF = 0
	ACCELF_VIRTKEY ACCELF = 1
	ACCELF_SHIFT   ACCELF = 0x04
	ACCELF_CONTROL ACCELF = 0x08
	ACCELF_ALT     ACCELF = 0x10
)

// NMTVASYNCDRAW dwRetFlags, don't seem to be defined anywhere, values are unconfirmed.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvasyncdraw
type ADRF uint32

const (
	ADRF_DRAWSYNC     ADRF = 0
	ADRF_DRAWNOTHING  ADRF = 1
	ADRF_DRAWFALLBACK ADRF = 2
	ADRF_DRAWIMAGE    ADRF = 3
)

// WM_APPCOMMAND command.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-appcommand
type APPCOMMAND int16

const (
	APPCOMMAND_BROWSER_BACKWARD                  APPCOMMAND = 1
	APPCOMMAND_BROWSER_FORWARD                   APPCOMMAND = 2
	APPCOMMAND_BROWSER_REFRESH                   APPCOMMAND = 3
	APPCOMMAND_BROWSER_STOP                      APPCOMMAND = 4
	APPCOMMAND_BROWSER_SEARCH                    APPCOMMAND = 5
	APPCOMMAND_BROWSER_FAVORITES                 APPCOMMAND = 6
	APPCOMMAND_BROWSER_HOME                      APPCOMMAND = 7
	APPCOMMAND_VOLUME_MUTE                       APPCOMMAND = 8
	APPCOMMAND_VOLUME_DOWN                       APPCOMMAND = 9
	APPCOMMAND_VOLUME_UP                         APPCOMMAND = 10
	APPCOMMAND_MEDIA_NEXTTRACK                   APPCOMMAND = 11
	APPCOMMAND_MEDIA_PREVIOUSTRACK               APPCOMMAND = 12
	APPCOMMAND_MEDIA_STOP                        APPCOMMAND = 13
	APPCOMMAND_MEDIA_PLAY_PAUSE                  APPCOMMAND = 14
	APPCOMMAND_LAUNCH_MAIL                       APPCOMMAND = 15
	APPCOMMAND_LAUNCH_MEDIA_SELECT               APPCOMMAND = 16
	APPCOMMAND_LAUNCH_APP1                       APPCOMMAND = 17
	APPCOMMAND_LAUNCH_APP2                       APPCOMMAND = 18
	APPCOMMAND_BASS_DOWN                         APPCOMMAND = 19
	APPCOMMAND_BASS_BOOST                        APPCOMMAND = 20
	APPCOMMAND_BASS_UP                           APPCOMMAND = 21
	APPCOMMAND_TREBLE_DOWN                       APPCOMMAND = 22
	APPCOMMAND_TREBLE_UP                         APPCOMMAND = 23
	APPCOMMAND_MICROPHONE_VOLUME_MUTE            APPCOMMAND = 24
	APPCOMMAND_MICROPHONE_VOLUME_DOWN            APPCOMMAND = 25
	APPCOMMAND_MICROPHONE_VOLUME_UP              APPCOMMAND = 26
	APPCOMMAND_HELP                              APPCOMMAND = 27
	APPCOMMAND_FIND                              APPCOMMAND = 28
	APPCOMMAND_NEW                               APPCOMMAND = 29
	APPCOMMAND_OPEN                              APPCOMMAND = 30
	APPCOMMAND_CLOSE                             APPCOMMAND = 31
	APPCOMMAND_SAVE                              APPCOMMAND = 32
	APPCOMMAND_PRINT                             APPCOMMAND = 33
	APPCOMMAND_UNDO                              APPCOMMAND = 34
	APPCOMMAND_REDO                              APPCOMMAND = 35
	APPCOMMAND_COPY                              APPCOMMAND = 36
	APPCOMMAND_CUT                               APPCOMMAND = 37
	APPCOMMAND_PASTE                             APPCOMMAND = 38
	APPCOMMAND_REPLY_TO_MAIL                     APPCOMMAND = 39
	APPCOMMAND_FORWARD_MAIL                      APPCOMMAND = 40
	APPCOMMAND_SEND_MAIL                         APPCOMMAND = 41
	APPCOMMAND_SPELL_CHECK                       APPCOMMAND = 42
	APPCOMMAND_DICTATE_OR_COMMAND_CONTROL_TOGGLE APPCOMMAND = 43
	APPCOMMAND_MIC_ON_OFF_TOGGLE                 APPCOMMAND = 44
	APPCOMMAND_CORRECTION_LIST                   APPCOMMAND = 45
	APPCOMMAND_MEDIA_PLAY                        APPCOMMAND = 46
	APPCOMMAND_MEDIA_PAUSE                       APPCOMMAND = 47
	APPCOMMAND_MEDIA_RECORD                      APPCOMMAND = 48
	APPCOMMAND_MEDIA_FAST_FORWARD                APPCOMMAND = 49
	APPCOMMAND_MEDIA_REWIND                      APPCOMMAND = 50
	APPCOMMAND_MEDIA_CHANNEL_UP                  APPCOMMAND = 51
	APPCOMMAND_MEDIA_CHANNEL_DOWN                APPCOMMAND = 52
	APPCOMMAND_DELETE                            APPCOMMAND = 53
	APPCOMMAND_DWM_FLIP3D                        APPCOMMAND = 54
)

// BITMAPINFOHEADER biCompression.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmapinfoheader
type BI uint32

const (
	BI_RGB       BI = 0
	BI_RLE8      BI = 1
	BI_RLE4      BI = 2
	BI_BITFIELDS BI = 3
	BI_JPEG      BI = 4
	BI_PNG       BI = 5
)

// SetBkMode() mode. Originally has no prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setbkmode
type BKMODE int32

const (
	BKMODE_TRANSPARENT BKMODE = 1
	BKMODE_OPAQUE      BKMODE = 2
)

// LOGBRUSH lbStyle. Originally with BS prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logbrush
type BRS uint32

const (
	BRS_SOLID         BRS = 0
	BRS_NULL          BRS = 1
	BRS_HOLLOW        BRS = BRS_NULL
	BRS_HATCHED       BRS = 2
	BRS_PATTERN       BRS = 3
	BRS_INDEXED       BRS = 4
	BRS_DIBPATTERN    BRS = 5
	BRS_DIBPATTERNPT  BRS = 6
	BRS_PATTERN8X8    BRS = 7
	BRS_DIBPATTERN8X8 BRS = 8
	BRS_MONOPATTERN   BRS = 9
)

// Button control styles.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/button-styles
type BS WS

const (
	BS_PUSHBUTTON      BS = 0x0000_0000
	BS_DEFPUSHBUTTON   BS = 0x0000_0001
	BS_CHECKBOX        BS = 0x0000_0002
	BS_AUTOCHECKBOX    BS = 0x0000_0003
	BS_RADIOBUTTON     BS = 0x0000_0004
	BS_3STATE          BS = 0x0000_0005
	BS_AUTO3STATE      BS = 0x0000_0006
	BS_GROUPBOX        BS = 0x0000_0007
	BS_USERBUTTON      BS = 0x0000_0008
	BS_AUTORADIOBUTTON BS = 0x0000_0009
	BS_PUSHBOX         BS = 0x0000_000a
	BS_OWNERDRAW       BS = 0x0000_000b
	BS_TYPEMASK        BS = 0x0000_000f
	BS_LEFTTEXT        BS = 0x0000_0020
	BS_TEXT            BS = 0x0000_0000
	BS_ICON            BS = 0x0000_0040
	BS_BITMAP          BS = 0x0000_0080
	BS_LEFT            BS = 0x0000_0100
	BS_RIGHT           BS = 0x0000_0200
	BS_CENTER          BS = 0x0000_0300
	BS_TOP             BS = 0x0000_0400
	BS_BOTTOM          BS = 0x0000_0800
	BS_VCENTER         BS = 0x0000_0c00
	BS_PUSHLIKE        BS = 0x0000_1000
	BS_MULTILINE       BS = 0x0000_2000
	BS_NOTIFY          BS = 0x0000_4000 // Button will send BN_DISABLE, BN_PUSHED, BN_KILLFOCUS, BN_PAINT, BN_SETFOCUS, and BN_UNPUSHED notifications.
	BS_FLAT            BS = 0x0000_8000
	BS_RIGHTBUTTON     BS = BS_LEFTTEXT
)

// BroadcastSystemMessage() lpInfo.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-broadcastsystemmessagew
type BSM uint32

const (
	BSM_ALLCOMPONENTS BSM = 0x00000000
	BSM_ALLDESKTOPS   BSM = 0x00000010
	BSM_APPLICATIONS  BSM = 0x00000008
)

// BroadcastSystemMessage() flags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-broadcastsystemmessagew
type BSF uint32

const (
	BSF_ALLOWSFW           BSF = 0x0000_0080
	BSF_FLUSHDISK          BSF = 0x0000_0004
	BSF_FORCEIFHUNG        BSF = 0x0000_0020
	BSF_IGNORECURRENTTASK  BSF = 0x0000_0002
	BSF_NOHANG             BSF = 0x0000_0008
	BSF_NOTIMEOUTIFNOTHUNG BSF = 0x0000_0040
	BSF_POSTMESSAGE        BSF = 0x0000_0010
	BSF_QUERY              BSF = 0x0000_0001
	BSF_SENDNOTIFYMESSAGE  BSF = 0x0000_0100
)

// IsDlgButtonChecked() return value, among others.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isdlgbuttonchecked
type BST uint32

const (
	BST_UNCHECKED     BST = 0x0000
	BST_CHECKED       BST = 0x0001
	BST_INDETERMINATE BST = 0x0002
	BST_PUSHED        BST = 0x0004
	BST_FOCUS         BST = 0x0008
)

// Toolbar button styles.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/toolbar-control-and-button-styles
type BTNS uint8

const (
	BTNS_BUTTON        BTNS = BTNS(TBSTYLE_BUTTON)
	BTNS_SEP           BTNS = BTNS(TBSTYLE_SEP)
	BTNS_CHECK         BTNS = BTNS(TBSTYLE_CHECK)
	BTNS_GROUP         BTNS = BTNS(TBSTYLE_GROUP)
	BTNS_CHECKGROUP    BTNS = BTNS(TBSTYLE_CHECKGROUP)
	BTNS_DROPDOWN      BTNS = BTNS(TBSTYLE_DROPDOWN)
	BTNS_AUTOSIZE      BTNS = BTNS(TBSTYLE_AUTOSIZE)
	BTNS_NOPREFIX      BTNS = BTNS(TBSTYLE_NOPREFIX)
	BTNS_SHOWTEXT      BTNS = 0x0040
	BTNS_WHOLEDROPDOWN BTNS = 0x0080
)
