package co

// WM_HSCROLL, WM_HSCROLL, WM_HSCROLLCLIPBOARD and WM_VSCROLLCLIPBOARD request.
// Originally with SB prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/wm-hscroll
type SB_REQ uint16

const (
	SB_REQ_LINEUP        SB_REQ = 0
	SB_REQ_LINELEFT      SB_REQ = 0
	SB_REQ_LINEDOWN      SB_REQ = 1
	SB_REQ_LINERIGHT     SB_REQ = 1
	SB_REQ_PAGEUP        SB_REQ = 2
	SB_REQ_PAGELEFT      SB_REQ = 2
	SB_REQ_PAGEDOWN      SB_REQ = 3
	SB_REQ_PAGERIGHT     SB_REQ = 3
	SB_REQ_THUMBPOSITION SB_REQ = 4
	SB_REQ_THUMBTRACK    SB_REQ = 5
	SB_REQ_TOP           SB_REQ = 6
	SB_REQ_LEFT          SB_REQ = 6
	SB_REQ_BOTTOM        SB_REQ = 7
	SB_REQ_RIGHT         SB_REQ = 7
	SB_REQ_ENDSCROLL     SB_REQ = 8
)

// GetScrollInfo() nBar, among others. Originally has SB prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getscrollinfo
type SB_TYPE int32

const (
	SB_TYPE_HORZ SB_TYPE = 0
	SB_TYPE_VERT SB_TYPE = 1
	SB_TYPE_CTL  SB_TYPE = 2
)

// StatusBar styles.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/status-bar-styles
type SBARS WS

const (
	SBARS_SIZEGRIP SBARS = 0x0100 // The status bar control will include a sizing grip at the right end of the status bar. A sizing grip is similar to a sizing border; it is a rectangular area that the user can click and drag to resize the parent window.
	SBARS_TOOLTIPS SBARS = 0x0800 // Use this style to enable tooltips.
)

// WM_SYSCOMMAND type of requested command.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-syscommand
type SC uint32

const (
	SC_SIZE         SC = 0xf000
	SC_MOVE         SC = 0xf010
	SC_MINIMIZE     SC = 0xf020
	SC_MAXIMIZE     SC = 0xf030
	SC_NEXTWINDOW   SC = 0xf040
	SC_PREVWINDOW   SC = 0xf050
	SC_CLOSE        SC = 0xf060
	SC_VSCROLL      SC = 0xf070
	SC_HSCROLL      SC = 0xf080
	SC_MOUSEMENU    SC = 0xf090
	SC_KEYMENU      SC = 0xf100
	SC_ARRANGE      SC = 0xf110
	SC_RESTORE      SC = 0xf120
	SC_TASKLIST     SC = 0xf130
	SC_SCREENSAVE   SC = 0xf140
	SC_HOTKEY       SC = 0xf150
	SC_DEFAULT      SC = 0xf160
	SC_MONITORPOWER SC = 0xf170
	SC_CONTEXTHELP  SC = 0xf180
	SC_SEPARATOR    SC = 0xf00f
)

// CreateFileMapping() flProtect.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-createfilemappingw
type SEC uint32

const (
	SEC_NONE                   SEC = 0
	SEC_PARTITION_OWNER_HANDLE SEC = 0x0004_0000
	SEC_64K_PAGES              SEC = 0x0008_0000
	SEC_FILE                   SEC = 0x0080_0000
	SEC_IMAGE                  SEC = 0x0100_0000
	SEC_PROTECTED_IMAGE        SEC = 0x0200_0000
	SEC_RESERVE                SEC = 0x0400_0000
	SEC_COMMIT                 SEC = 0x0800_0000
	SEC_NOCACHE                SEC = 0x1000_0000
	SEC_WRITECOMBINE           SEC = 0x4000_0000
	SEC_LARGE_PAGES            SEC = 0x8000_0000
	SEC_IMAGE_NO_EXECUTE       SEC = SEC_IMAGE | SEC_NOCACHE
)

// CreateFile() dwFlagsAndAttributes.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createfilew
type SECURITY uint32

const (
	SECURITY_NONE             SECURITY = 0
	SECURITY_ANONYMOUS        SECURITY = 0 << 16
	SECURITY_IDENTIFICATION   SECURITY = 1 << 16
	SECURITY_IMPERSONATION    SECURITY = 2 << 16
	SECURITY_DELEGATION       SECURITY = 3 << 16
	SECURITY_CONTEXT_TRACKING SECURITY = 0x0004_0000
	SECURITY_EFFECTIVE_ONLY   SECURITY = 0x0008_0000
)

// SHFILEINFO dwAttributes.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-shfileinfow
type SFGAO uint32

const (
	_DROPEFFECT_NONE   SFGAO = 0
	_DROPEFFECT_COPY   SFGAO = 1
	_DROPEFFECT_MOVE   SFGAO = 2
	_DROPEFFECT_LINK   SFGAO = 4
	_DROPEFFECT_SCROLL SFGAO = 0x8000_0000

	SFGAO_CANCOPY         SFGAO = _DROPEFFECT_COPY
	SFGAO_CANMOVE         SFGAO = _DROPEFFECT_MOVE
	SFGAO_CANLINK         SFGAO = _DROPEFFECT_LINK
	SFGAO_STORAGE         SFGAO = 0x0000_0008
	SFGAO_CANRENAME       SFGAO = 0x0000_0010
	SFGAO_CANDELETE       SFGAO = 0x0000_0020
	SFGAO_HASPROPSHEET    SFGAO = 0x0000_0040
	SFGAO_DROPTARGET      SFGAO = 0x0000_0100
	SFGAO_CAPABILITYMASK  SFGAO = 0x0000_0177
	SFGAO_PLACEHOLDER     SFGAO = 0x0000_0800
	SFGAO_SYSTEM          SFGAO = 0x0000_1000
	SFGAO_ENCRYPTED       SFGAO = 0x0000_2000
	SFGAO_ISSLOW          SFGAO = 0x0000_4000
	SFGAO_GHOSTED         SFGAO = 0x0000_8000
	SFGAO_LINK            SFGAO = 0x0001_0000
	SFGAO_SHARE           SFGAO = 0x0002_0000
	SFGAO_READONLY        SFGAO = 0x0004_0000
	SFGAO_HIDDEN          SFGAO = 0x0008_0000
	SFGAO_DISPLAYATTRMASK SFGAO = 0x000f_c000
	SFGAO_FILESYSANCESTOR SFGAO = 0x1000_0000
	SFGAO_FOLDER          SFGAO = 0x2000_0000
	SFGAO_FILESYSTEM      SFGAO = 0x4000_0000
	SFGAO_HASSUBFOLDER    SFGAO = 0x8000_0000
	SFGAO_CONTENTSMASK    SFGAO = 0x8000_0000
	SFGAO_VALIDATE        SFGAO = 0x0100_0000
	SFGAO_REMOVABLE       SFGAO = 0x0200_0000
	SFGAO_COMPRESSED      SFGAO = 0x0400_0000
	SFGAO_BROWSABLE       SFGAO = 0x0800_0000
	SFGAO_NONENUMERATED   SFGAO = 0x0010_0000
	SFGAO_NEWCONTENT      SFGAO = 0x0020_0000
	SFGAO_CANMONIKER      SFGAO = 0x0040_0000
	SFGAO_HASSTORAGE      SFGAO = 0x0040_0000
	SFGAO_STREAM          SFGAO = 0x0040_0000
	SFGAO_STORAGEANCESTOR SFGAO = 0x0080_0000
	SFGAO_STORAGECAPMASK  SFGAO = 0x70c5_0008
	SFGAO_PKEYSFGAOMASK   SFGAO = 0x8104_4000
)

// SHGetFileInfo() uFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shgetfileinfow
type SHGFI uint32

const (
	SHGFI_NONE              SHGFI = 0
	SHGFI_ICON              SHGFI = 0x0000_0100
	SHGFI_DISPLAYNAME       SHGFI = 0x0000_0200
	SHGFI_TYPENAME          SHGFI = 0x0000_0400
	SHGFI_ATTRIBUTES        SHGFI = 0x0000_0800
	SHGFI_ICONLOCATION      SHGFI = 0x0000_1000
	SHGFI_EXETYPE           SHGFI = 0x0000_2000
	SHGFI_SYSICONINDEX      SHGFI = 0x0000_4000
	SHGFI_LINKOVERLAY       SHGFI = 0x0000_8000
	SHGFI_SELECTED          SHGFI = 0x0001_0000
	SHGFI_ATTR_SPECIFIED    SHGFI = 0x0002_0000
	SHGFI_LARGEICON         SHGFI = 0x0000_0000
	SHGFI_SMALLICON         SHGFI = 0x0000_0001
	SHGFI_OPENICON          SHGFI = 0x0000_0002
	SHGFI_SHELLICONSIZE     SHGFI = 0x0000_0004
	SHGFI_PIDL              SHGFI = 0x0000_0008
	SHGFI_USEFILEATTRIBUTES SHGFI = 0x0000_0010
	SHGFI_ADDOVERLAYS       SHGFI = 0x0000_0020
	SHGFI_OVERLAYINDEX      SHGFI = 0x0000_0040
)

// SCROLLINFO fMask.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-scrollinfo
type SIF uint32

const (
	SIF_RANGE           SIF = 0x0001
	SIF_PAGE            SIF = 0x0002
	SIF_POS             SIF = 0x0004
	SIF_DISABLENOSCROLL SIF = 0x0008
	SIF_TRACKPOS        SIF = 0x0010
	SIF_ALL             SIF = SIF_RANGE | SIF_PAGE | SIF_POS | SIF_TRACKPOS
)

// WM_SIZE request.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-size
type SIZE_REQ int32

const (
	SIZE_REQ_RESTORED  SIZE_REQ = 0 // The window has been resized, but neither the SIZE_REQ_MINIMIZED nor SIZE_REQ_MAXIMIZED value applies.
	SIZE_REQ_MINIMIZED SIZE_REQ = 1 // The window has been minimized.
	SIZE_REQ_MAXIMIZED SIZE_REQ = 2 // The window has been maximized.
	SIZE_REQ_MAXSHOW   SIZE_REQ = 3 // Message is sent to all pop-up windows when some other window has been restored to its former size.
	SIZE_REQ_MAXHIDE   SIZE_REQ = 4 // Message is sent to all pop-up windows when some other window is maximized.
)

// GetSystemMetrics() nIndex.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsystemmetrics
type SM int32

const (
	SM_CXSCREEN                    SM = 0
	SM_CYSCREEN                    SM = 1
	SM_CXVSCROLL                   SM = 2
	SM_CYHSCROLL                   SM = 3
	SM_CYCAPTION                   SM = 4
	SM_CXBORDER                    SM = 5
	SM_CYBORDER                    SM = 6
	SM_CXDLGFRAME                  SM = 7
	SM_CYDLGFRAME                  SM = 8
	SM_CYVTHUMB                    SM = 9
	SM_CXHTHUMB                    SM = 10
	SM_CXICON                      SM = 11
	SM_CYICON                      SM = 12
	SM_CXCURSOR                    SM = 13
	SM_CYCURSOR                    SM = 14
	SM_CYMENU                      SM = 15
	SM_CXFULLSCREEN                SM = 16
	SM_CYFULLSCREEN                SM = 17
	SM_CYKANJIWINDOW               SM = 18
	SM_MOUSEPRESENT                SM = 19
	SM_CYVSCROLL                   SM = 20
	SM_CXHSCROLL                   SM = 21
	SM_DEBUG                       SM = 22
	SM_SWAPBUTTON                  SM = 23
	SM_RESERVED1                   SM = 24
	SM_RESERVED2                   SM = 25
	SM_RESERVED3                   SM = 26
	SM_RESERVED4                   SM = 27
	SM_CXMIN                       SM = 28
	SM_CYMIN                       SM = 29
	SM_CXSIZE                      SM = 30
	SM_CYSIZE                      SM = 31
	SM_CXFRAME                     SM = 32
	SM_CYFRAME                     SM = 33
	SM_CXMINTRACK                  SM = 34
	SM_CYMINTRACK                  SM = 35
	SM_CXDOUBLECLK                 SM = 36
	SM_CYDOUBLECLK                 SM = 37
	SM_CXICONSPACING               SM = 38
	SM_CYICONSPACING               SM = 39
	SM_MENUDROPALIGNMENT           SM = 40
	SM_PENWINDOWS                  SM = 41
	SM_DBCSENABLED                 SM = 42
	SM_CMOUSEBUTTONS               SM = 43
	SM_CXFIXEDFRAME                SM = SM_CXDLGFRAME
	SM_CYFIXEDFRAME                SM = SM_CYDLGFRAME
	SM_CXSIZEFRAME                 SM = SM_CXFRAME
	SM_CYSIZEFRAME                 SM = SM_CYFRAME
	SM_SECURE                      SM = 44
	SM_CXEDGE                      SM = 45
	SM_CYEDGE                      SM = 46
	SM_CXMINSPACING                SM = 47
	SM_CYMINSPACING                SM = 48
	SM_CXSMICON                    SM = 49
	SM_CYSMICON                    SM = 50
	SM_CYSMCAPTION                 SM = 51
	SM_CXSMSIZE                    SM = 52
	SM_CYSMSIZE                    SM = 53
	SM_CXMENUSIZE                  SM = 54
	SM_CYMENUSIZE                  SM = 55
	SM_ARRANGE                     SM = 56
	SM_CXMINIMIZED                 SM = 57
	SM_CYMINIMIZED                 SM = 58
	SM_CXMAXTRACK                  SM = 59
	SM_CYMAXTRACK                  SM = 60
	SM_CXMAXIMIZED                 SM = 61
	SM_CYMAXIMIZED                 SM = 62
	SM_NETWORK                     SM = 63
	SM_CLEANBOOT                   SM = 67
	SM_CXDRAG                      SM = 68
	SM_CYDRAG                      SM = 69
	SM_SHOWSOUNDS                  SM = 70
	SM_CXMENUCHECK                 SM = 71
	SM_CYMENUCHECK                 SM = 72
	SM_SLOWMACHINE                 SM = 73
	SM_MIDEASTENABLED              SM = 74
	SM_MOUSEWHEELPRESENT           SM = 75
	SM_XVIRTUALSCREEN              SM = 76
	SM_YVIRTUALSCREEN              SM = 77
	SM_CXVIRTUALSCREEN             SM = 78
	SM_CYVIRTUALSCREEN             SM = 79
	SM_CMONITORS                   SM = 80
	SM_SAMEDISPLAYFORMAT           SM = 81
	SM_IMMENABLED                  SM = 82
	SM_CXFOCUSBORDER               SM = 83
	SM_CYFOCUSBORDER               SM = 84
	SM_TABLETPC                    SM = 86
	SM_MEDIACENTER                 SM = 87
	SM_STARTER                     SM = 88
	SM_SERVERR2                    SM = 89
	SM_MOUSEHORIZONTALWHEELPRESENT SM = 91
	SM_CXPADDEDBORDER              SM = 92
	SM_DIGITIZER                   SM = 94
	SM_MAXIMUMTOUCHES              SM = 95
	SM_CMETRICS                    SM = 97
	SM_REMOTESESSION               SM = 0x1000
	SM_SHUTTINGDOWN                SM = 0x2000
	SM_REMOTECONTROL               SM = 0x2001
	SM_CARETBLINKINGENABLED        SM = 0x2002
	SM_CONVERTIBLESLATEMODE        SM = 0x2003
	SM_SYSTEMDOCKED                SM = 0x2004
)

// Sort order identifier for locales.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/intl/sort-order-identifiers
type SORT uint16

const (
	SORT_DEFAULT                SORT = 0x0
	SORT_INVARIANT_MATH         SORT = 0x1
	SORT_JAPANESE_XJIS          SORT = 0x0
	SORT_JAPANESE_UNICODE       SORT = 0x1
	SORT_JAPANESE_RADICALSTROKE SORT = 0x4
	SORT_CHINESE_BIG5           SORT = 0x0
	SORT_CHINESE_PRCP           SORT = 0x0
	SORT_CHINESE_UNICODE        SORT = 0x1
	SORT_CHINESE_PRC            SORT = 0x2
	SORT_CHINESE_BOPOMOFO       SORT = 0x3
	SORT_CHINESE_RADICALSTROKE  SORT = 0x4
	SORT_KOREAN_KSC             SORT = 0x0
	SORT_KOREAN_UNICODE         SORT = 0x1
	SORT_GERMAN_PHONE_BOOK      SORT = 0x1
	SORT_HUNGARIAN_DEFAULT      SORT = 0x0
	SORT_HUNGARIAN_TECHNICAL    SORT = 0x1
	SORT_GEORGIAN_TRADITIONAL   SORT = 0x0
	SORT_GEORGIAN_MODERN        SORT = 0x1
)

// SystemParametersInfo() uiAction.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-systemparametersinfow
type SPI uint32

const (
	SPI_GETBEEP                     SPI = 0x0001
	SPI_SETBEEP                     SPI = 0x0002
	SPI_GETMOUSE                    SPI = 0x0003
	SPI_SETMOUSE                    SPI = 0x0004
	SPI_GETBORDER                   SPI = 0x0005
	SPI_SETBORDER                   SPI = 0x0006
	SPI_GETKEYBOARDSPEED            SPI = 0x000a
	SPI_SETKEYBOARDSPEED            SPI = 0x000b
	SPI_LANGDRIVER                  SPI = 0x000c
	SPI_ICONHORIZONTALSPACING       SPI = 0x000d
	SPI_GETSCREENSAVETIMEOUT        SPI = 0x000e
	SPI_SETSCREENSAVETIMEOUT        SPI = 0x000f
	SPI_GETSCREENSAVEACTIVE         SPI = 0x0010
	SPI_SETSCREENSAVEACTIVE         SPI = 0x0011
	SPI_GETGRIDGRANULARITY          SPI = 0x0012
	SPI_SETGRIDGRANULARITY          SPI = 0x0013
	SPI_SETDESKWALLPAPER            SPI = 0x0014
	SPI_SETDESKPATTERN              SPI = 0x0015
	SPI_GETKEYBOARDDELAY            SPI = 0x0016
	SPI_SETKEYBOARDDELAY            SPI = 0x0017
	SPI_ICONVERTICALSPACING         SPI = 0x0018
	SPI_GETICONTITLEWRAP            SPI = 0x0019
	SPI_SETICONTITLEWRAP            SPI = 0x001a
	SPI_GETMENUDROPALIGNMENT        SPI = 0x001b
	SPI_SETMENUDROPALIGNMENT        SPI = 0x001c
	SPI_SETDOUBLECLKWIDTH           SPI = 0x001d
	SPI_SETDOUBLECLKHEIGHT          SPI = 0x001e
	SPI_GETICONTITLELOGFONT         SPI = 0x001f
	SPI_SETDOUBLECLICKTIME          SPI = 0x0020
	SPI_SETMOUSEBUTTONSWAP          SPI = 0x0021
	SPI_SETICONTITLELOGFONT         SPI = 0x0022
	SPI_GETFASTTASKSWITCH           SPI = 0x0023
	SPI_SETFASTTASKSWITCH           SPI = 0x0024
	SPI_SETDRAGFULLWINDOWS          SPI = 0x0025
	SPI_GETDRAGFULLWINDOWS          SPI = 0x0026
	SPI_GETNONCLIENTMETRICS         SPI = 0x0029
	SPI_SETNONCLIENTMETRICS         SPI = 0x002a
	SPI_GETMINIMIZEDMETRICS         SPI = 0x002b
	SPI_SETMINIMIZEDMETRICS         SPI = 0x002c
	SPI_GETICONMETRICS              SPI = 0x002d
	SPI_SETICONMETRICS              SPI = 0x002e
	SPI_SETWORKAREA                 SPI = 0x002f
	SPI_GETWORKAREA                 SPI = 0x0030
	SPI_SETPENWINDOWS               SPI = 0x0031
	SPI_GETHIGHCONTRAST             SPI = 0x0042
	SPI_SETHIGHCONTRAST             SPI = 0x0043
	SPI_GETKEYBOARDPREF             SPI = 0x0044
	SPI_SETKEYBOARDPREF             SPI = 0x0045
	SPI_GETSCREENREADER             SPI = 0x0046
	SPI_SETSCREENREADER             SPI = 0x0047
	SPI_GETANIMATION                SPI = 0x0048
	SPI_SETANIMATION                SPI = 0x0049
	SPI_GETFONTSMOOTHING            SPI = 0x004a
	SPI_SETFONTSMOOTHING            SPI = 0x004b
	SPI_SETDRAGWIDTH                SPI = 0x004c
	SPI_SETDRAGHEIGHT               SPI = 0x004d
	SPI_SETHANDHELD                 SPI = 0x004e
	SPI_GETLOWPOWERTIMEOUT          SPI = 0x004f
	SPI_GETPOWEROFFTIMEOUT          SPI = 0x0050
	SPI_SETLOWPOWERTIMEOUT          SPI = 0x0051
	SPI_SETPOWEROFFTIMEOUT          SPI = 0x0052
	SPI_GETLOWPOWERACTIVE           SPI = 0x0053
	SPI_GETPOWEROFFACTIVE           SPI = 0x0054
	SPI_SETLOWPOWERACTIVE           SPI = 0x0055
	SPI_SETPOWEROFFACTIVE           SPI = 0x0056
	SPI_SETCURSORS                  SPI = 0x0057
	SPI_SETICONS                    SPI = 0x0058
	SPI_GETDEFAULTINPUTLANG         SPI = 0x0059
	SPI_SETDEFAULTINPUTLANG         SPI = 0x005a
	SPI_SETLANGTOGGLE               SPI = 0x005b
	SPI_GETWINDOWSEXTENSION         SPI = 0x005c
	SPI_SETMOUSETRAILS              SPI = 0x005d
	SPI_GETMOUSETRAILS              SPI = 0x005e
	SPI_SETSCREENSAVERRUNNING       SPI = 0x0061
	SPI_SCREENSAVERRUNNING          SPI = SPI_SETSCREENSAVERRUNNING
	SPI_GETFILTERKEYS               SPI = 0x0032
	SPI_SETFILTERKEYS               SPI = 0x0033
	SPI_GETTOGGLEKEYS               SPI = 0x0034
	SPI_SETTOGGLEKEYS               SPI = 0x0035
	SPI_GETMOUSEKEYS                SPI = 0x0036
	SPI_SETMOUSEKEYS                SPI = 0x0037
	SPI_GETSHOWSOUNDS               SPI = 0x0038
	SPI_SETSHOWSOUNDS               SPI = 0x0039
	SPI_GETSTICKYKEYS               SPI = 0x003a
	SPI_SETSTICKYKEYS               SPI = 0x003b
	SPI_GETACCESSTIMEOUT            SPI = 0x003c
	SPI_SETACCESSTIMEOUT            SPI = 0x003d
	SPI_GETSERIALKEYS               SPI = 0x003e
	SPI_SETSERIALKEYS               SPI = 0x003f
	SPI_GETSOUNDSENTRY              SPI = 0x0040
	SPI_SETSOUNDSENTRY              SPI = 0x0041
	SPI_GETSNAPTODEFBUTTON          SPI = 0x005f
	SPI_SETSNAPTODEFBUTTON          SPI = 0x0060
	SPI_GETMOUSEHOVERWIDTH          SPI = 0x0062
	SPI_SETMOUSEHOVERWIDTH          SPI = 0x0063
	SPI_GETMOUSEHOVERHEIGHT         SPI = 0x0064
	SPI_SETMOUSEHOVERHEIGHT         SPI = 0x0065
	SPI_GETMOUSEHOVERTIME           SPI = 0x0066
	SPI_SETMOUSEHOVERTIME           SPI = 0x0067
	SPI_GETWHEELSCROLLLINES         SPI = 0x0068
	SPI_SETWHEELSCROLLLINES         SPI = 0x0069
	SPI_GETMENUSHOWDELAY            SPI = 0x006a
	SPI_SETMENUSHOWDELAY            SPI = 0x006b
	SPI_GETWHEELSCROLLCHARS         SPI = 0x006c
	SPI_SETWHEELSCROLLCHARS         SPI = 0x006d
	SPI_GETSHOWIMEUI                SPI = 0x006e
	SPI_SETSHOWIMEUI                SPI = 0x006f
	SPI_GETMOUSESPEED               SPI = 0x0070
	SPI_SETMOUSESPEED               SPI = 0x0071
	SPI_GETSCREENSAVERRUNNING       SPI = 0x0072
	SPI_GETDESKWALLPAPER            SPI = 0x0073
	SPI_GETAUDIODESCRIPTION         SPI = 0x0074
	SPI_SETAUDIODESCRIPTION         SPI = 0x0075
	SPI_GETSCREENSAVESECURE         SPI = 0x0076
	SPI_SETSCREENSAVESECURE         SPI = 0x0077
	SPI_GETHUNGAPPTIMEOUT           SPI = 0x0078
	SPI_SETHUNGAPPTIMEOUT           SPI = 0x0079
	SPI_GETWAITTOKILLTIMEOUT        SPI = 0x007a
	SPI_SETWAITTOKILLTIMEOUT        SPI = 0x007b
	SPI_GETWAITTOKILLSERVICETIMEOUT SPI = 0x007c
	SPI_SETWAITTOKILLSERVICETIMEOUT SPI = 0x007d
	SPI_GETMOUSEDOCKTHRESHOLD       SPI = 0x007e
	SPI_SETMOUSEDOCKTHRESHOLD       SPI = 0x007f
	SPI_GETPENDOCKTHRESHOLD         SPI = 0x0080
	SPI_SETPENDOCKTHRESHOLD         SPI = 0x0081
	SPI_GETWINARRANGING             SPI = 0x0082
	SPI_SETWINARRANGING             SPI = 0x0083
	SPI_GETMOUSEDRAGOUTTHRESHOLD    SPI = 0x0084
	SPI_SETMOUSEDRAGOUTTHRESHOLD    SPI = 0x0085
	SPI_GETPENDRAGOUTTHRESHOLD      SPI = 0x0086
	SPI_SETPENDRAGOUTTHRESHOLD      SPI = 0x0087
	SPI_GETMOUSESIDEMOVETHRESHOLD   SPI = 0x0088
	SPI_SETMOUSESIDEMOVETHRESHOLD   SPI = 0x0089
	SPI_GETPENSIDEMOVETHRESHOLD     SPI = 0x008a
	SPI_SETPENSIDEMOVETHRESHOLD     SPI = 0x008b
	SPI_GETDRAGFROMMAXIMIZE         SPI = 0x008c
	SPI_SETDRAGFROMMAXIMIZE         SPI = 0x008d
	SPI_GETSNAPSIZING               SPI = 0x008e
	SPI_SETSNAPSIZING               SPI = 0x008f
	SPI_GETDOCKMOVING               SPI = 0x0090
	SPI_SETDOCKMOVING               SPI = 0x0091
)

// SystemParametersInfo() fWinIni.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-systemparametersinfow
type SPIF uint32

const (
	SPIF_UPDATEINIFILE    SPIF = 1
	SPIF_SENDWININICHANGE SPIF = 2
	SPIF_SENDCHANGE       SPIF = SPIF_SENDWININICHANGE
)

// Static control styles.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/static-control-styles
type SS WS

const (
	SS_LEFT            SS = 0x0000_0000
	SS_CENTER          SS = 0x0000_0001
	SS_RIGHT           SS = 0x0000_0002
	SS_ICON            SS = 0x0000_0003
	SS_BLACKRECT       SS = 0x0000_0004
	SS_GRAYRECT        SS = 0x0000_0005
	SS_WHITERECT       SS = 0x0000_0006
	SS_BLACKFRAME      SS = 0x0000_0007
	SS_GRAYFRAME       SS = 0x0000_0008
	SS_WHITEFRAME      SS = 0x0000_0009
	SS_USERITEM        SS = 0x0000_000a
	SS_SIMPLE          SS = 0x0000_000b
	SS_LEFTNOWORDWRAP  SS = 0x0000_000c
	SS_OWNERDRAW       SS = 0x0000_000d
	SS_BITMAP          SS = 0x0000_000e
	SS_ENHMETAFILE     SS = 0x0000_000f
	SS_ETCHEDHORZ      SS = 0x0000_0010
	SS_ETCHEDVERT      SS = 0x0000_0011
	SS_ETCHEDFRAME     SS = 0x0000_0012
	SS_TYPEMASK        SS = 0x0000_001f
	SS_REALSIZECONTROL SS = 0x0000_0040
	SS_NOPREFIX        SS = 0x0000_0080
	SS_NOTIFY          SS = 0x0000_0100
	SS_CENTERIMAGE     SS = 0x0000_0200
	SS_RIGHTJUST       SS = 0x0000_0400
	SS_REALSIZEIMAGE   SS = 0x0000_0800
	SS_SUNKEN          SS = 0x0000_1000
	SS_EDITCONTROL     SS = 0x0000_2000
	SS_ENDELLIPSIS     SS = 0x0000_4000
	SS_PATHELLIPSIS    SS = 0x0000_8000
	SS_WORDELLIPSIS    SS = 0x0000_c000
	SS_ELLIPSISMASK    SS = 0x0000_c000
)

// Standard access rights. These are generic and compose other access right
// types. Also includes unprefixed and SPECIFIC_RIGHT prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/secauthz/standard-access-rights
type STANDARD_RIGHTS uint32

const (
	STANDARD_RIGHTS_NONE STANDARD_RIGHTS = 0

	STANDARD_RIGHTS_DELETE       STANDARD_RIGHTS = 0x0001_0000
	STANDARD_RIGHTS_READ_CONTROL STANDARD_RIGHTS = 0x0002_0000
	STANDARD_RIGHTS_SYNCHRONIZE  STANDARD_RIGHTS = 0x0010_0000
	STANDARD_RIGHTS_WRITE_DAC    STANDARD_RIGHTS = 0x0004_0000
	STANDARD_RIGHTS_WRITE_OWNER  STANDARD_RIGHTS = 0x0008_0000

	STANDARD_RIGHTS_ALL      STANDARD_RIGHTS = 0x001f_0000
	STANDARD_RIGHTS_EXECUTE  STANDARD_RIGHTS = STANDARD_RIGHTS_READ_CONTROL
	STANDARD_RIGHTS_READ     STANDARD_RIGHTS = STANDARD_RIGHTS_READ_CONTROL
	STANDARD_RIGHTS_REQUIRED STANDARD_RIGHTS = 0x000f_0000
	STANDARD_RIGHTS_WRITE    STANDARD_RIGHTS = STANDARD_RIGHTS_READ_CONTROL
)

// STARTUPINFO dwFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/ns-processthreadsapi-startupinfow
type STARTF uint32

const (
	STARTF_FORCEONFEEDBACK  STARTF = 0x0000_0040
	STARTF_FORCEOFFFEEDBACK STARTF = 0x0000_0080
	STARTF_PREVENTPINNING   STARTF = 0x0000_2000
	STARTF_RUNFULLSCREEN    STARTF = 0x0000_0020
	STARTF_TITLEISAPPID     STARTF = 0x0000_1000
	STARTF_TITLEISLINKNAME  STARTF = 0x0000_0800
	STARTF_UNTRUSTEDSOURCE  STARTF = 0x0000_8000
	STARTF_USECOUNTCHARS    STARTF = 0x0000_0008
	STARTF_USEFILLATTRIBUTE STARTF = 0x0000_0010
	STARTF_USEHOTKEY        STARTF = 0x0000_0200
	STARTF_USEPOSITION      STARTF = 0x0000_0004
	STARTF_USESHOWWINDOW    STARTF = 0x0000_0001
	STARTF_USESIZE          STARTF = 0x0000_0002
	STARTF_USESTDHANDLES    STARTF = 0x0000_0100
)

// SetStretchBltMode() mode.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setstretchbltmode
type STRETCH int32

const (
	STRETCH_BLACKONWHITE STRETCH = 1
	STRETCH_WHITEONBLACK STRETCH = 2
	STRETCH_COLORONCOLOR STRETCH = 3
	STRETCH_HALFTONE     STRETCH = 4
	STRETCH_ANDSCANS     STRETCH = STRETCH_BLACKONWHITE
	STRETCH_ORSCANS      STRETCH = STRETCH_WHITEONBLACK
	STRETCH_DELETESCANS  STRETCH = STRETCH_COLORONCOLOR
)

// Sub-language identifier.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/intl/language-identifier-constants-and-strings
type SUBLANG uint16

const (
	SUBLANG_NEUTRAL                             SUBLANG = 0x00
	SUBLANG_DEFAULT                             SUBLANG = 0x01
	SUBLANG_SYS_DEFAULT                         SUBLANG = 0x02
	SUBLANG_CUSTOM_DEFAULT                      SUBLANG = 0x03
	SUBLANG_CUSTOM_UNSPECIFIED                  SUBLANG = 0x04
	SUBLANG_UI_CUSTOM_DEFAULT                   SUBLANG = 0x05
	SUBLANG_AFRIKAANS_SOUTH_AFRICA              SUBLANG = 0x01
	SUBLANG_ALBANIAN_ALBANIA                    SUBLANG = 0x01
	SUBLANG_ALSATIAN_FRANCE                     SUBLANG = 0x01
	SUBLANG_AMHARIC_ETHIOPIA                    SUBLANG = 0x01
	SUBLANG_ARABIC_SAUDI_ARABIA                 SUBLANG = 0x01
	SUBLANG_ARABIC_IRAQ                         SUBLANG = 0x02
	SUBLANG_ARABIC_EGYPT                        SUBLANG = 0x03
	SUBLANG_ARABIC_LIBYA                        SUBLANG = 0x04
	SUBLANG_ARABIC_ALGERIA                      SUBLANG = 0x05
	SUBLANG_ARABIC_MOROCCO                      SUBLANG = 0x06
	SUBLANG_ARABIC_TUNISIA                      SUBLANG = 0x07
	SUBLANG_ARABIC_OMAN                         SUBLANG = 0x08
	SUBLANG_ARABIC_YEMEN                        SUBLANG = 0x09
	SUBLANG_ARABIC_SYRIA                        SUBLANG = 0x0a
	SUBLANG_ARABIC_JORDAN                       SUBLANG = 0x0b
	SUBLANG_ARABIC_LEBANON                      SUBLANG = 0x0c
	SUBLANG_ARABIC_KUWAIT                       SUBLANG = 0x0d
	SUBLANG_ARABIC_UAE                          SUBLANG = 0x0e
	SUBLANG_ARABIC_BAHRAIN                      SUBLANG = 0x0f
	SUBLANG_ARABIC_QATAR                        SUBLANG = 0x10
	SUBLANG_ARMENIAN_ARMENIA                    SUBLANG = 0x01
	SUBLANG_ASSAMESE_INDIA                      SUBLANG = 0x01
	SUBLANG_AZERI_LATIN                         SUBLANG = 0x01
	SUBLANG_AZERI_CYRILLIC                      SUBLANG = 0x02
	SUBLANG_AZERBAIJANI_AZERBAIJAN_LATIN        SUBLANG = 0x01
	SUBLANG_AZERBAIJANI_AZERBAIJAN_CYRILLIC     SUBLANG = 0x02
	SUBLANG_BANGLA_INDIA                        SUBLANG = 0x01
	SUBLANG_BANGLA_BANGLADESH                   SUBLANG = 0x02
	SUBLANG_BASHKIR_RUSSIA                      SUBLANG = 0x01
	SUBLANG_BASQUE_BASQUE                       SUBLANG = 0x01
	SUBLANG_BELARUSIAN_BELARUS                  SUBLANG = 0x01
	SUBLANG_BENGALI_INDIA                       SUBLANG = 0x01
	SUBLANG_BENGALI_BANGLADESH                  SUBLANG = 0x02
	SUBLANG_BOSNIAN_BOSNIA_HERZEGOVINA_LATIN    SUBLANG = 0x05
	SUBLANG_BOSNIAN_BOSNIA_HERZEGOVINA_CYRILLIC SUBLANG = 0x08
	SUBLANG_BRETON_FRANCE                       SUBLANG = 0x01
	SUBLANG_BULGARIAN_BULGARIA                  SUBLANG = 0x01
	SUBLANG_CATALAN_CATALAN                     SUBLANG = 0x01
	SUBLANG_CENTRAL_KURDISH_IRAQ                SUBLANG = 0x01
	SUBLANG_CHEROKEE_CHEROKEE                   SUBLANG = 0x01
	SUBLANG_CHINESE_TRADITIONAL                 SUBLANG = 0x01
	SUBLANG_CHINESE_SIMPLIFIED                  SUBLANG = 0x02
	SUBLANG_CHINESE_HONGKONG                    SUBLANG = 0x03
	SUBLANG_CHINESE_SINGAPORE                   SUBLANG = 0x04
	SUBLANG_CHINESE_MACAU                       SUBLANG = 0x05
	SUBLANG_CORSICAN_FRANCE                     SUBLANG = 0x01
	SUBLANG_CZECH_CZECH_REPUBLIC                SUBLANG = 0x01
	SUBLANG_CROATIAN_CROATIA                    SUBLANG = 0x01
	SUBLANG_CROATIAN_BOSNIA_HERZEGOVINA_LATIN   SUBLANG = 0x04
	SUBLANG_DANISH_DENMARK                      SUBLANG = 0x01
	SUBLANG_DARI_AFGHANISTAN                    SUBLANG = 0x01
	SUBLANG_DIVEHI_MALDIVES                     SUBLANG = 0x01
	SUBLANG_DUTCH                               SUBLANG = 0x01
	SUBLANG_DUTCH_BELGIAN                       SUBLANG = 0x02
	SUBLANG_ENGLISH_US                          SUBLANG = 0x01
	SUBLANG_ENGLISH_UK                          SUBLANG = 0x02
	SUBLANG_ENGLISH_AUS                         SUBLANG = 0x03
	SUBLANG_ENGLISH_CAN                         SUBLANG = 0x04
	SUBLANG_ENGLISH_NZ                          SUBLANG = 0x05
	SUBLANG_ENGLISH_EIRE                        SUBLANG = 0x06
	SUBLANG_ENGLISH_SOUTH_AFRICA                SUBLANG = 0x07
	SUBLANG_ENGLISH_JAMAICA                     SUBLANG = 0x08
	SUBLANG_ENGLISH_CARIBBEAN                   SUBLANG = 0x09
	SUBLANG_ENGLISH_BELIZE                      SUBLANG = 0x0a
	SUBLANG_ENGLISH_TRINIDAD                    SUBLANG = 0x0b
	SUBLANG_ENGLISH_ZIMBABWE                    SUBLANG = 0x0c
	SUBLANG_ENGLISH_PHILIPPINES                 SUBLANG = 0x0d
	SUBLANG_ENGLISH_INDIA                       SUBLANG = 0x10
	SUBLANG_ENGLISH_MALAYSIA                    SUBLANG = 0x11
	SUBLANG_ENGLISH_SINGAPORE                   SUBLANG = 0x12
	SUBLANG_ESTONIAN_ESTONIA                    SUBLANG = 0x01
	SUBLANG_FAEROESE_FAROE_ISLANDS              SUBLANG = 0x01
	SUBLANG_FILIPINO_PHILIPPINES                SUBLANG = 0x01
	SUBLANG_FINNISH_FINLAND                     SUBLANG = 0x01
	SUBLANG_FRENCH                              SUBLANG = 0x01
	SUBLANG_FRENCH_BELGIAN                      SUBLANG = 0x02
	SUBLANG_FRENCH_CANADIAN                     SUBLANG = 0x03
	SUBLANG_FRENCH_SWISS                        SUBLANG = 0x04
	SUBLANG_FRENCH_LUXEMBOURG                   SUBLANG = 0x05
	SUBLANG_FRENCH_MONACO                       SUBLANG = 0x06
	SUBLANG_FRISIAN_NETHERLANDS                 SUBLANG = 0x01
	SUBLANG_FULAH_SENEGAL                       SUBLANG = 0x02
	SUBLANG_GALICIAN_GALICIAN                   SUBLANG = 0x01
	SUBLANG_GEORGIAN_GEORGIA                    SUBLANG = 0x01
	SUBLANG_GERMAN                              SUBLANG = 0x01
	SUBLANG_GERMAN_SWISS                        SUBLANG = 0x02
	SUBLANG_GERMAN_AUSTRIAN                     SUBLANG = 0x03
	SUBLANG_GERMAN_LUXEMBOURG                   SUBLANG = 0x04
	SUBLANG_GERMAN_LIECHTENSTEIN                SUBLANG = 0x05
	SUBLANG_GREEK_GREECE                        SUBLANG = 0x01
	SUBLANG_GREENLANDIC_GREENLAND               SUBLANG = 0x01
	SUBLANG_GUJARATI_INDIA                      SUBLANG = 0x01
	SUBLANG_HAUSA_NIGERIA_LATIN                 SUBLANG = 0x01
	SUBLANG_HAWAIIAN_US                         SUBLANG = 0x01
	SUBLANG_HEBREW_ISRAEL                       SUBLANG = 0x01
	SUBLANG_HINDI_INDIA                         SUBLANG = 0x01
	SUBLANG_HUNGARIAN_HUNGARY                   SUBLANG = 0x01
	SUBLANG_ICELANDIC_ICELAND                   SUBLANG = 0x01
	SUBLANG_IGBO_NIGERIA                        SUBLANG = 0x01
	SUBLANG_INDONESIAN_INDONESIA                SUBLANG = 0x01
	SUBLANG_INUKTITUT_CANADA                    SUBLANG = 0x01
	SUBLANG_INUKTITUT_CANADA_LATIN              SUBLANG = 0x02
	SUBLANG_IRISH_IRELAND                       SUBLANG = 0x02
	SUBLANG_ITALIAN                             SUBLANG = 0x01
	SUBLANG_ITALIAN_SWISS                       SUBLANG = 0x02
	SUBLANG_JAPANESE_JAPAN                      SUBLANG = 0x01
	SUBLANG_KANNADA_INDIA                       SUBLANG = 0x01
	SUBLANG_KASHMIRI_SASIA                      SUBLANG = 0x02
	SUBLANG_KASHMIRI_INDIA                      SUBLANG = 0x02
	SUBLANG_KAZAK_KAZAKHSTAN                    SUBLANG = 0x01
	SUBLANG_KHMER_CAMBODIA                      SUBLANG = 0x01
	SUBLANG_KICHE_GUATEMALA                     SUBLANG = 0x01
	SUBLANG_KINYARWANDA_RWANDA                  SUBLANG = 0x01
	SUBLANG_KONKANI_INDIA                       SUBLANG = 0x01
	SUBLANG_KOREAN                              SUBLANG = 0x01
	SUBLANG_KYRGYZ_KYRGYZSTAN                   SUBLANG = 0x01
	SUBLANG_LAO_LAO                             SUBLANG = 0x01
	SUBLANG_LATVIAN_LATVIA                      SUBLANG = 0x01
	SUBLANG_LITHUANIAN                          SUBLANG = 0x01
	SUBLANG_LOWER_SORBIAN_GERMANY               SUBLANG = 0x02
	SUBLANG_LUXEMBOURGISH_LUXEMBOURG            SUBLANG = 0x01
	SUBLANG_MACEDONIAN_MACEDONIA                SUBLANG = 0x01
	SUBLANG_MALAY_MALAYSIA                      SUBLANG = 0x01
	SUBLANG_MALAY_BRUNEI_DARUSSALAM             SUBLANG = 0x02
	SUBLANG_MALAYALAM_INDIA                     SUBLANG = 0x01
	SUBLANG_MALTESE_MALTA                       SUBLANG = 0x01
	SUBLANG_MAORI_NEW_ZEALAND                   SUBLANG = 0x01
	SUBLANG_MAPUDUNGUN_CHILE                    SUBLANG = 0x01
	SUBLANG_MARATHI_INDIA                       SUBLANG = 0x01
	SUBLANG_MOHAWK_MOHAWK                       SUBLANG = 0x01
	SUBLANG_MONGOLIAN_CYRILLIC_MONGOLIA         SUBLANG = 0x01
	SUBLANG_MONGOLIAN_PRC                       SUBLANG = 0x02
	SUBLANG_NEPALI_INDIA                        SUBLANG = 0x02
	SUBLANG_NEPALI_NEPAL                        SUBLANG = 0x01
	SUBLANG_NORWEGIAN_BOKMAL                    SUBLANG = 0x01
	SUBLANG_NORWEGIAN_NYNORSK                   SUBLANG = 0x02
	SUBLANG_OCCITAN_FRANCE                      SUBLANG = 0x01
	SUBLANG_ODIA_INDIA                          SUBLANG = 0x01
	SUBLANG_ORIYA_INDIA                         SUBLANG = 0x01
	SUBLANG_PASHTO_AFGHANISTAN                  SUBLANG = 0x01
	SUBLANG_PERSIAN_IRAN                        SUBLANG = 0x01
	SUBLANG_POLISH_POLAND                       SUBLANG = 0x01
	SUBLANG_PORTUGUESE                          SUBLANG = 0x02
	SUBLANG_PORTUGUESE_BRAZILIAN                SUBLANG = 0x01
	SUBLANG_PULAR_SENEGAL                       SUBLANG = 0x02
	SUBLANG_PUNJABI_INDIA                       SUBLANG = 0x01
	SUBLANG_PUNJABI_PAKISTAN                    SUBLANG = 0x02
	SUBLANG_QUECHUA_BOLIVIA                     SUBLANG = 0x01
	SUBLANG_QUECHUA_ECUADOR                     SUBLANG = 0x02
	SUBLANG_QUECHUA_PERU                        SUBLANG = 0x03
	SUBLANG_ROMANIAN_ROMANIA                    SUBLANG = 0x01
	SUBLANG_ROMANSH_SWITZERLAND                 SUBLANG = 0x01
	SUBLANG_RUSSIAN_RUSSIA                      SUBLANG = 0x01
	SUBLANG_SAKHA_RUSSIA                        SUBLANG = 0x01
	SUBLANG_SAMI_NORTHERN_NORWAY                SUBLANG = 0x01
	SUBLANG_SAMI_NORTHERN_SWEDEN                SUBLANG = 0x02
	SUBLANG_SAMI_NORTHERN_FINLAND               SUBLANG = 0x03
	SUBLANG_SAMI_LULE_NORWAY                    SUBLANG = 0x04
	SUBLANG_SAMI_LULE_SWEDEN                    SUBLANG = 0x05
	SUBLANG_SAMI_SOUTHERN_NORWAY                SUBLANG = 0x06
	SUBLANG_SAMI_SOUTHERN_SWEDEN                SUBLANG = 0x07
	SUBLANG_SAMI_SKOLT_FINLAND                  SUBLANG = 0x08
	SUBLANG_SAMI_INARI_FINLAND                  SUBLANG = 0x09
	SUBLANG_SANSKRIT_INDIA                      SUBLANG = 0x01
	SUBLANG_SCOTTISH_GAELIC                     SUBLANG = 0x01
	SUBLANG_SERBIAN_BOSNIA_HERZEGOVINA_LATIN    SUBLANG = 0x06
	SUBLANG_SERBIAN_BOSNIA_HERZEGOVINA_CYRILLIC SUBLANG = 0x07
	SUBLANG_SERBIAN_MONTENEGRO_LATIN            SUBLANG = 0x0b
	SUBLANG_SERBIAN_MONTENEGRO_CYRILLIC         SUBLANG = 0x0c
	SUBLANG_SERBIAN_SERBIA_LATIN                SUBLANG = 0x09
	SUBLANG_SERBIAN_SERBIA_CYRILLIC             SUBLANG = 0x0a
	SUBLANG_SERBIAN_CROATIA                     SUBLANG = 0x01
	SUBLANG_SERBIAN_LATIN                       SUBLANG = 0x02
	SUBLANG_SERBIAN_CYRILLIC                    SUBLANG = 0x03
	SUBLANG_SINDHI_INDIA                        SUBLANG = 0x01
	SUBLANG_SINDHI_PAKISTAN                     SUBLANG = 0x02
	SUBLANG_SINDHI_AFGHANISTAN                  SUBLANG = 0x02
	SUBLANG_SINHALESE_SRI_LANKA                 SUBLANG = 0x01
	SUBLANG_SOTHO_NORTHERN_SOUTH_AFRICA         SUBLANG = 0x01
	SUBLANG_SLOVAK_SLOVAKIA                     SUBLANG = 0x01
	SUBLANG_SLOVENIAN_SLOVENIA                  SUBLANG = 0x01
	SUBLANG_SPANISH                             SUBLANG = 0x01
	SUBLANG_SPANISH_MEXICAN                     SUBLANG = 0x02
	SUBLANG_SPANISH_MODERN                      SUBLANG = 0x03
	SUBLANG_SPANISH_GUATEMALA                   SUBLANG = 0x04
	SUBLANG_SPANISH_COSTA_RICA                  SUBLANG = 0x05
	SUBLANG_SPANISH_PANAMA                      SUBLANG = 0x06
	SUBLANG_SPANISH_DOMINICAN_REPUBLIC          SUBLANG = 0x07
	SUBLANG_SPANISH_VENEZUELA                   SUBLANG = 0x08
	SUBLANG_SPANISH_COLOMBIA                    SUBLANG = 0x09
	SUBLANG_SPANISH_PERU                        SUBLANG = 0x0a
	SUBLANG_SPANISH_ARGENTINA                   SUBLANG = 0x0b
	SUBLANG_SPANISH_ECUADOR                     SUBLANG = 0x0c
	SUBLANG_SPANISH_CHILE                       SUBLANG = 0x0d
	SUBLANG_SPANISH_URUGUAY                     SUBLANG = 0x0e
	SUBLANG_SPANISH_PARAGUAY                    SUBLANG = 0x0f
	SUBLANG_SPANISH_BOLIVIA                     SUBLANG = 0x10
	SUBLANG_SPANISH_EL_SALVADOR                 SUBLANG = 0x11
	SUBLANG_SPANISH_HONDURAS                    SUBLANG = 0x12
	SUBLANG_SPANISH_NICARAGUA                   SUBLANG = 0x13
	SUBLANG_SPANISH_PUERTO_RICO                 SUBLANG = 0x14
	SUBLANG_SPANISH_US                          SUBLANG = 0x15
	SUBLANG_SWAHILI_KENYA                       SUBLANG = 0x01
	SUBLANG_SWEDISH                             SUBLANG = 0x01
	SUBLANG_SWEDISH_FINLAND                     SUBLANG = 0x02
	SUBLANG_SYRIAC_SYRIA                        SUBLANG = 0x01
	SUBLANG_TAJIK_TAJIKISTAN                    SUBLANG = 0x01
	SUBLANG_TAMAZIGHT_ALGERIA_LATIN             SUBLANG = 0x02
	SUBLANG_TAMAZIGHT_MOROCCO_TIFINAGH          SUBLANG = 0x04
	SUBLANG_TAMIL_INDIA                         SUBLANG = 0x01
	SUBLANG_TAMIL_SRI_LANKA                     SUBLANG = 0x02
	SUBLANG_TATAR_RUSSIA                        SUBLANG = 0x01
	SUBLANG_TELUGU_INDIA                        SUBLANG = 0x01
	SUBLANG_THAI_THAILAND                       SUBLANG = 0x01
	SUBLANG_TIBETAN_PRC                         SUBLANG = 0x01
	SUBLANG_TIGRIGNA_ERITREA                    SUBLANG = 0x02
	SUBLANG_TIGRINYA_ERITREA                    SUBLANG = 0x02
	SUBLANG_TIGRINYA_ETHIOPIA                   SUBLANG = 0x01
	SUBLANG_TSWANA_BOTSWANA                     SUBLANG = 0x02
	SUBLANG_TSWANA_SOUTH_AFRICA                 SUBLANG = 0x01
	SUBLANG_TURKISH_TURKEY                      SUBLANG = 0x01
	SUBLANG_TURKMEN_TURKMENISTAN                SUBLANG = 0x01
	SUBLANG_UIGHUR_PRC                          SUBLANG = 0x01
	SUBLANG_UKRAINIAN_UKRAINE                   SUBLANG = 0x01
	SUBLANG_UPPER_SORBIAN_GERMANY               SUBLANG = 0x01
	SUBLANG_URDU_PAKISTAN                       SUBLANG = 0x01
	SUBLANG_URDU_INDIA                          SUBLANG = 0x02
	SUBLANG_UZBEK_LATIN                         SUBLANG = 0x01
	SUBLANG_UZBEK_CYRILLIC                      SUBLANG = 0x02
	SUBLANG_VALENCIAN_VALENCIA                  SUBLANG = 0x02
	SUBLANG_VIETNAMESE_VIETNAM                  SUBLANG = 0x01
	SUBLANG_WELSH_UNITED_KINGDOM                SUBLANG = 0x01
	SUBLANG_WOLOF_SENEGAL                       SUBLANG = 0x01
	SUBLANG_XHOSA_SOUTH_AFRICA                  SUBLANG = 0x01
	SUBLANG_YAKUT_RUSSIA                        SUBLANG = 0x01
	SUBLANG_YI_PRC                              SUBLANG = 0x01
	SUBLANG_YORUBA_NIGERIA                      SUBLANG = 0x01
	SUBLANG_ZULU_SOUTH_AFRICA                   SUBLANG = 0x01
)

// ShowWindow() nCmdShow.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showwindow
type SW int32

const (
	SW_HIDE            SW = 0
	SW_SHOWNORMAL      SW = 1
	SW_SHOWMINIMIZED   SW = 2
	SW_SHOWMAXIMIZED   SW = 3
	SW_MAXIMIZE        SW = 3
	SW_SHOWNOACTIVATE  SW = 4
	SW_SHOW            SW = 5
	SW_MINIMIZE        SW = 6
	SW_SHOWMINNOACTIVE SW = 7
	SW_SHOWNA          SW = 8
	SW_RESTORE         SW = 9
	SW_SHOWDEFAULT     SW = 10
	SW_FORCEMINIMIZE   SW = 11
)

// SetWindowPos(), DeferWindowPos() uFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwindowpos
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-deferwindowpos
type SWP uint32

const (
	SWP_NOSIZE         SWP = 0x0001
	SWP_NOMOVE         SWP = 0x0002
	SWP_NOZORDER       SWP = 0x0004
	SWP_NOREDRAW       SWP = 0x0008
	SWP_NOACTIVATE     SWP = 0x0010
	SWP_FRAMECHANGED   SWP = 0x0020
	SWP_SHOWWINDOW     SWP = 0x0040
	SWP_HIDEWINDOW     SWP = 0x0080
	SWP_NOCOPYBITS     SWP = 0x0100
	SWP_NOOWNERZORDER  SWP = 0x0200
	SWP_NOSENDCHANGING SWP = 0x0400
	SWP_DRAWFRAME      SWP = SWP_FRAMECHANGED
	SWP_NOREPOSITION   SWP = SWP_NOOWNERZORDER
	SWP_DEFERERASE     SWP = 0x2000
	SWP_ASYNCWINDOWPOS SWP = 0x4000
)

// WM_SHOWWINDOW return value. Originally has SW prefix.
type SWS uint8

const (
	SWS_OTHERUNZOOM   SWS = 4 // The window is being uncovered because a maximize window was restored or minimized.
	SWS_OTHERZOOM     SWS = 2 // The window is being covered by another window that has been maximized.
	SWS_PARENTCLOSING SWS = 1 // The window's owner window is being minimized.
	SWS_PARENTOPENING SWS = 3 // The window's owner window is being restored.
)
