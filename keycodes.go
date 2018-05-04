package main

var mapcode map[int]string

const (
	VK_LBUTTON int = 0 + iota
	VK_RBUTTON
	VK_CANCEL
	VK_MBUTTON
	VK_XBUTTON1
	VK_XBUTTON2
	VK_BACK
	VK_TAB
	VK_CLEAR
	VK_RETURN
	VK_SHIFT
	VK_CONTROL
	VK_MENU
	VK_PAUSE
	VK_CAPITAL
	VK_KANA
	VK_HANGUL
	VK_JUNJA
	VK_FINAL
	VK_HANJA
	VK_KANJI
	VK_ESCAPE
	VK_CONVERT
	VK_NONCONVERT
	VK_ACCEPT
	VK_MODECHANGE
	VK_SPACE
	VK_PRIOR
	VK_NEXT
	VK_END
	VK_HOME
	VK_LEFT
	VK_UP
	VK_RIGHT
	VK_DOWN
	VK_SELECT
	VK_PRINT
	VK_EXECUTE
	VK_SNAPSHOT
	VK_INSERT
	VK_DELETE
	VK_0
	VK_1
	VK_2
	VK_3
	VK_4
	VK_5
	VK_6
	VK_7
	VK_8
	VK_9
)
const (
	VK_A int = 65 + iota
	VK_B
	VK_C
	VK_D
	VK_E
	VK_F
	VK_G
	VK_H
	VK_I
	VK_J
	VK_K
	VK_L
	VK_M
	VK_N
	VK_O
	VK_P
	VK_Q
	VK_R
	VK_S
	VK_T
	VK_U
	VK_V
	VK_W
	VK_X
	VK_Y
	VK_Z
	VK_91
	VK_92
	VK_93
)
const (
	VK_95 int = 95 + iota
	VK_NUMPAD0
	VK_NUMPAD1
	VK_NUMPAD2
	VK_NUMPAD3
	VK_NUMPAD4
	VK_NUMPAD5
	VK_NUMPAD6
	VK_NUMPAD7
	VK_NUMPAD8
	VK_NUMPAD9
	VK_MULTIPLY
	VK_ADD
	VK_SEPARATOR
	VK_SUBTRACT
	VK_DECIMAL
	VK_DIVIDE
	VK_F1
	VK_F2
	VK_F3
	VK_F4
	VK_F5
	VK_F6
	VK_F7
	VK_F8
	VK_F9
	VK_F10
	VK_F11
	VK_F12
	VK_F13
	VK_F14
	VK_F15
	VK_F16
	VK_F17
	VK_F18
	VK_F19
	VK_F20
	VK_F21
	VK_F22
	VK_F23
	VK_F24
)
const (
	VK_NUMLOCK int = 144 + iota
	VK_SCROLL
)
const (
	VK_LSHIFT int = 160 + iota
	VK_RSHIFT
	VK_LCONTROL
	VK_RCONTROL
	VK_LMENU
	VK_RMENU
	VK_BROWSER_BACK
	VK_BROWSER_FORWARD
	VK_BROWSER_REFRESH
	VK_BROWSER_STOP
	VK_BROWSER_SEARCH
	VK_BROWSER_FAVORITES
	VK_BROWSER_HOME
	VK_VOLUME_MUTE
	VK_VOLUME_DOWN
	VK_VOLUME_UP
	VK_MEDIA_NEXT_TRACK
	VK_MEDIA_PREV_TRACK
	VK_MEDIA_STOP
	VK_MEDIA_PLAY_PAUSE
	VK_LAUNCH_MAIL
	VK_LAUNCH_MEDIA_SELECT
	VK_LAUNCH_APP1
	VK_LAUNCH_APP2
	VK_OEM_COLON
	VK_OEM_PLUS
	VK_OEM_COMMA
	VK_OEM_MINUS
)
const (
	VK_OEM_PERIOD int = 190 + iota
	VK_OEM_QUEST
	VK_OEM_BACKQUOTE
)
const (
	VK_OEM_LEFTBRACKET int = 219 + iota
	VK_OEM_BACKSLASH
	VK_OEM_RIGHTBRACKET
	VK_OEM_QUOTE
	VK_OEM_8
)
const (
	VK_OEM_102    int = 226 + iota
	VK_PACKET     int = 231 + iota
	VK_PROCESSKEY int = 229 + iota
)
const (
	VK_ATTN int = 246 + iota
	VK_CRSEL
	VK_EXSEL
	VK_EREOF
	VK_PLAY
	VK_ZOOM
	VK_NONAME
	VK_PA1
	VK_OEM_CLE
	VK_BOTTOM
)

func mapinit() map[int]string {
	return map[int]string{
		1:   "VK_LBUTTON", //鼠标左键
		2:   "VK_RBUTTON", //鼠标右键
		3:   "VK_CANCEL",  //Cancel
		4:   "VK_MBUTTON", //鼠标中键
		5:   "VK_XBUTTON1",
		6:   "VK_XBUTTON2",
		8:   "VK_BACK",    //Backspace
		9:   "VK_TAB",     //Tab
		12:  "VK_CLEAR",   //Clear
		13:  "VK_RETURN",  //Enter
		16:  "VK_SHIFT",   //Shift
		17:  "VK_CONTROL", //Ctrl
		18:  "VK_MENU",    //Alt
		19:  "VK_PAUSE",   //Pause
		20:  "VK_CAPITAL", //Caps Lock
		21:  "VK_KANA",
		22:  "VK_HANGUL",
		23:  "VK_JUNJA",
		24:  "VK_FINAL",
		25:  "VK_HANJA",
		26:  "VK_KANJI",  //*
		27:  "VK_ESCAPE", //Esc
		28:  "VK_CONVERT",
		29:  "VK_NONCONVERT",
		30:  "VK_ACCEPT",
		31:  "VK_MODECHANGE",
		32:  "VK_SPACE",    //Space
		33:  "VK_PRIOR",    //Page Up
		34:  "VK_NEXT",     //Page Down
		35:  "VK_END",      //End
		36:  "VK_HOME",     //Home
		37:  "VK_LEFT",     //Left Arrow
		38:  "VK_UP",       //Up Arrow
		39:  "VK_RIGHT",    //Right Arrow
		40:  "VK_DOWN",     //Down Arrow
		41:  "VK_SELECT",   //Select
		42:  "VK_PRINT",    //Print
		43:  "VK_EXECUTE",  //Execute
		44:  "VK_SNAPSHOT", //Snapshot
		45:  "VK_INSERT",   //Insert
		46:  "VK_DELETE",   //Delete
		47:  "VK_HELP",     //Help
		48:  "VK_0",
		49:  "VK_2",
		50:  "VK_3",
		51:  "VK_4",
		52:  "VK_5",
		53:  "VK_6",
		54:  "VK_7",
		55:  "VK_8",
		56:  "VK_9",
		65:  "VK_A",
		66:  "VK_B",
		67:  "VK_C",
		68:  "VK_D",
		69:  "VK_E",
		70:  "VK_F",
		71:  "VK_G",
		72:  "VK_H",
		73:  "VK_I",
		74:  "VK_J",
		75:  "VK_K",
		76:  "VK_L",
		77:  "VK_M",
		78:  "VK_N",
		79:  "VK_O",
		80:  "VK_P",
		81:  "VK_Q",
		82:  "VK_R",
		83:  "VK_S",
		84:  "VK_T",
		85:  "VK_U",
		86:  "VK_V",
		87:  "VK_W",
		88:  "VK_X",
		89:  "VK_Y",
		90:  "VK_Z",
		91:  "VK_LWIN", //left wins
		92:  "VK_RWIN", //rightwins
		93:  "VK_APPS",
		95:  "VK_SLEEP",
		96:  "VK_NUMPAD0",   //小键盘 0
		97:  "VK_NUMPAD1",   //小键盘 1
		98:  "VK_NUMPAD2",   //小键盘 2
		99:  "VK_NUMPAD3",   //小键盘 3
		100: "VK_NUMPAD4",   //小键盘 4
		101: "VK_NUMPAD5",   //小键盘 5
		102: "VK_NUMPAD6",   //小键盘 6
		103: "VK_NUMPAD7",   //小键盘 7
		104: "VK_NUMPAD8",   //小键盘 8
		105: "VK_NUMPAD9",   //小键盘 9
		106: "VK_MULTIPLY",  //小键盘 *
		107: "VK_ADD",       //小键盘 +
		108: "VK_SEPARATOR", //小键盘 Enter
		109: "VK_SUBTRACT",  //小键盘 -
		110: "VK_DECIMAL",   //小键盘 .
		111: "VK_DIVIDE",    //小键盘 /
		112: "VK_F1",        //F1
		113: "VK_F2",        //F2
		114: "VK_F3",        //F3
		115: "VK_F4",        //F4
		116: "VK_F5",        //F5
		117: "VK_F6",        //F6
		118: "VK_F7",        //F7
		119: "VK_F8",        //F8
		120: "VK_F9",        //F9
		121: "VK_F10",       //F10
		122: "VK_F11",       //F11
		123: "VK_F12",       //F12
		124: "VK_F13",
		125: "VK_F14",
		126: "VK_F15",
		127: "VK_F16",
		128: "VK_F17",
		129: "VK_F18",
		130: "VK_F19",
		131: "VK_F20",
		132: "VK_F21",
		133: "VK_F22",
		134: "VK_F23",
		135: "VK_F24",
		144: "VK_NUMLOCK", //Num Lock
		145: "VK_SCROLL",  //Scroll
		160: "VK_LSHIFT",
		161: "VK_RSHIFT",
		162: "VK_LCONTROL",
		163: "VK_RCONTROL",
		164: "VK_LMENU",
		165: "VK_RMENU",
		166: "VK_BROWSER_BACK",
		167: "VK_BROWSER_FORWARD",
		168: "VK_BROWSER_REFRESH",
		169: "VK_BROWSER_STOP",
		170: "VK_BROWSER_SEARCH",
		171: "VK_BROWSER_FAVORITES",
		172: "VK_BROWSER_HOME",
		173: "VK_VOLUME_MUTE", //VolumeMute
		174: "VK_VOLUME_DOWN", //VolumeDown
		175: "VK_VOLUME_UP",   //VolumeUp
		176: "VK_MEDIA_NEXT_TRACK",
		177: "VK_MEDIA_PREV_TRACK",
		178: "VK_MEDIA_STOP",
		179: "VK_MEDIA_PLAY_PAUSE",
		180: "VK_LAUNCH_MAIL",
		181: "VK_LAUNCH_MEDIA_SELECT",
		182: "VK_LAUNCH_APP1",
		183: "VK_LAUNCH_APP2",
		186: "VK_OEM_COLON",        //; :
		187: "VK_OEM_PLUS",         //= +
		188: "VK_OEM_COMMA",        //, <
		189: "VK_OEM_MINUS",        //- _
		190: "VK_OEM_PERIOD",       //.>
		191: "VK_OEM_QUEST",        /// ?
		192: "VK_OEM_BACKQUOTE",    //` ~
		219: "VK_OEM_LEFTBRACKET",  //[ {
		220: "VK_OEM_BACKSLASH",    //\ |
		221: "VK_OEM_RIGHTBRACKET", //] }
		222: "VK_OEM_QUOTE",        //' "
		223: "VK_OEM_8",
		226: "VK_OEM_102",
		231: "VK_PACKET",
		229: "VK_PROCESSKEY",
		246: "VK_ATTN",
		247: "VK_CRSEL",
		248: "VK_EXSEL",
		249: "VK_EREOF",
		250: "VK_PLAY",
		251: "VK_ZOOM",
		252: "VK_NONAME",
		253: "VK_PA1",
		254: "VK_OEM_CLE", //,,AR 	254
		255: "VK_BOTTOM",
	}
}
