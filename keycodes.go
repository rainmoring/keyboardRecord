package main

const (
	VK_LBUTTON    uint16 = itoa //1	鼠标左键
	VK_RBUTTON                  //2	鼠标右键
	VK_CANCEL                   //3	Cancel
	VK_MBUTTON                  //	4 	鼠标中键
	VK_XBUTTON1                 //	5
	VK_XBUTTON2                 //6
	VK_BACK                     //8	Backspace
	VK_TAB                      //9 	Tab
	VK_CLEAR                    //12 	Clear
	VK_RETURN                   //13 	Enter
	VK_SHIFT                    //16 	Shift
	VK_CONTROL                  //17 	Ctrl
	VK_MENU                     //18 	Alt
	_PAUSE                      //19 	Pause
	VK_CAPITAL                  //20 	Caps Lock
	VK_KANA                     //21
	VK_HANGUL                   //21
	VK_JUNJA                    //23
	VK_FINAL                    //24
	VK_HANJA                    //25
	VK_KANJI                    //25*
	VK_ESCAPE                   //27 	Esc
	VK_CONVERT                  //28
	VK_NONCONVERT               //29
	VK_ACCEPT                   //30
	VK_MODECHANGE               //31
	VK_SPACE                    //32 	Space
	VK_PRIOR                    //33 	Page Up
	VK_NEXT                     //34 	Page Down
	VK_END                      //35 	End
	VK_HOME                     //36 	Home
	VK_LEFT                     //37 	Left Arrow
	VK_UP                       //38 	Up Arrow
	VK_RIGHT                    //39 	Right Arrow
	VK_DOWN                     //40 	Down Arrow
	VK_SELECT                   //41 	Select
	VK_PRINT                    //42 	Print
	VK_EXECUTE                  //43 	Execute
	VK_SNAPSHOT                 //44 	Snapshot
	VK_INSERT                   //45 	Insert
	VK_DELETE                   //46 	Delete
	VK_HELP                     //47 	Help
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
	VK_A
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
	VK_95
	VK_NUMPAD0             //96 	小键盘 0
	VK_NUMPAD1             //97 	小键盘 1
	VK_NUMPAD2             //98 	小键盘 2
	VK_NUMPAD3             //99 	小键盘 3
	VK_NUMPAD4             //100 	小键盘 4
	VK_NUMPAD5             //101 	小键盘 5
	VK_NUMPAD6             //102 	小键盘 6
	VK_NUMPAD7             //103 	小键盘 7
	VK_NUMPAD8             //104 	小键盘 8
	VK_NUMPAD9             //105 	小键盘 9
	VK_MULTIPLY            //106 	小键盘 *
	VK_ADD                 //107 	小键盘 +
	VK_SEPARATOR           //108 	小键盘 Enter
	VK_SUBTRACT            //109 	小键盘 -
	VK_DECIMAL             //110 	小键盘 .
	VK_DIVIDE              //111 	小键盘 /
	VK_F1                  //112 	F1
	VK_F2                  //113 	F2
	VK_F3                  //114 	F3
	VK_F4                  //115 	F4
	VK_F5                  //116 	F5
	VK_F6                  //117 	F6
	VK_F7                  //118 	F7
	VK_F8                  //119 	F8
	VK_F9                  //120 	F9
	VK_F10                 //121 	F10
	VK_F11                 //122 	F11
	VK_F12                 //123 	F12
	VK_F13                 //124
	VK_F14                 //125
	VK_F15                 //126
	VK_F16                 //127
	VK_F17                 //128
	VK_F18                 //129
	VK_F19                 //130
	VK_F20                 //131
	VK_F21                 //132
	VK_F22                 //133
	VK_F23                 //134
	VK_F24                 //135
	VK_NUMLOCK             //144 	Num Lock
	VK_SCROLL              //145 	Scroll
	VK_LSHIFT              //160
	VK_RSHIFT              //161
	VK_LCONTROL            //162
	VK_RCONTROL            //163
	VK_LMENU               //164
	VK_RMENU               //165
	VK_BROWSER_BACK        //166
	VK_BROWSER_FORWARD     //167
	VK_BROWSER_REFRESH     //168
	VK_BROWSER_STOP        //169
	VK_BROWSER_SEARCH      //170
	VK_BROWSER_FAVORITES   //171
	VK_BROWSER_HOME        //172
	VK_VOLUME_MUTE         //173 	VolumeMute
	VK_VOLUME_DOWN         //174 	VolumeDown
	VK_VOLUME_UP           //175 	VolumeUp
	VK_MEDIA_NEXT_TRACK    //176
	VK_MEDIA_PREV_TRACK    //177
	VK_MEDIA_STOP          //178
	VK_MEDIA_PLAY_PAUSE    //179
	VK_LAUNCH_MAIL         //180
	VK_LAUNCH_MEDIA_SELECT //181
	VK_LAUNCH_APP1         //182
	VK_LAUNCH_APP2         //183
	VK_OEM_COLON           //186 	; :
	VK_OEM_PLUS            //187 	= +
	VK_OEM_COMMA           //188 , <
	VK_OEM_MINUS           //189 	- _
	VK_OEM_PERIOD          //190 .>
	VK_OEM_QUEST           //191 	/ ?
	VK_OEM_BACKQUOTE       //192 	` ~
	VK_OEM_LEFTBRACKET     //219 	[ {
	VK_OEM_BACKSLASH       //220 	\ |
	VK_OEM_RIGHTBRACKET    //221 	] }
	VK_OEM_QUOTE           //222 	' "
	VK_OEM_8               //223
	VK_OEM_102             //226
	VK_PACKET              //231
	VK_PROCESSKEY          //229
	VK_ATTN                //246
	VK_CRSEL               //247
	VK_EXSEL               //248
	VK_EREOF               //249
	VK_PLAY                //250
	VK_ZOOM                //251
	VK_NONAME              //252
	VK_PA1                 //253
	VK_OEM_CLE             //,,AR 	254
	VK_BOTTOM
)
