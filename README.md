# auto-mouse-keyboard

## Usage

1. Download exe file into an abtrary directory
2. Create a config file named `amk.conf` with config you need
3. Input `powershell` in Explorer to start `pwsh.exe`
4. Start amk with input `.\amk.exe` in the new window

Warning:

- All config with a suffix of `.conf` will be executed in the directory

## Description of `amk.conf`

All valid names are allowed, as long as the config end with `.conf` extension

Every line is a command, with three types: comment, setting and process

### Comment

Comment start with `#`, e.g.:

```bash
# This is a comment
```

Comments line, an empty line and a line with all whitespace are all empty lines, these lines will be ignored

### Setting

- `SHIM`: A very short time to suspend between two commands, in Millisecond
- `SCALE`: Scale for your monitor, keep the same with your system settings

In the settings below, the second command will delay 500ms after the first command finished, and the monitor scale 150%

```conf
SHIM=500
SCALE=1.5
```

### Process

There are six types:

- Mouse move
- Mouse click
- Keyboard Input
- Keyboard Press
- Suspend
- Loop

**Mouse move**

Say, we have a monitor with display 1920\*1080, left up point is (0,0), right bottom point is (1920,1080)

We want to move cursor to the point (300,200)

```bash
M=300,200
```

**Mouse click**

Click in current position

1. Left single click
2. Left double click
3. Right single click
4. Right double click

```conf
C=left
C=left,double
C=right
C=right,double
```

**Keyboard Input**

Input any content in current position (Must be an input box)

```conf
I=23333
I=Content can be English, also 中文
I=%^*(&())(  special chars are support
I=And Even whitespace
```

**Keyboard press**

Press can be single press and compound press

Single press: `a-z` / `0-9` and normal function keys

Compound press can be any types, say `Ctrl+C`/ `Ctrl+V`

All supported keys can be found at: [https://github.com/go-vgo/robotgo/blob/master/docs/keys.md](https://github.com/go-vgo/robotgo/blob/master/docs/keys.md)

```conf
T=1
T=cmd
T=esc
T=c,ctrl
T=v,ctrl
T=d,cmd
T=e,cmd
```

Some examples:

```bash
	"cmd"		is the "win" key for windows
	"alt"
	"ctrl"
	"shift"
	"capslock"
	"space"
	"backspace"
	"delete"
	"enter"
	"tab"
	"esc"
	"escape"
	"up"		Up arrow key
	"down"		Down arrow key
	"right"		Right arrow key
	"left"		Left arrow key
	"home"
	"end"
	"pageup"
	"pagedown"

	"f1"
	"f2"
	"f3"
	"f4"
	"f5"
	"f6"
	"f7"
	"f8"
	"f9"
	"f10"
	"f11"
	"f12"
```

**Suspend**

Increase the delay, the actual delay will be (SHIM+S)ms

If settings contains `SHIM=400` with config:

```conf
T=enter
S=300
I=wow
```

The delay between enter and input `wow` will be `SHIM+S+SHIM=400+300+400=1100 ms`

**Loop**

Loop a slice of comamnds for N times:

```conf
S=2000
C=left
I= begin

L=3
C=left
I=wow
T=enter*3
L

I= done
```

Suspend 2000ms, then click, then input text `begin`, and the loop begins.

In every loop, click left, then input text `wow`, then three ENTER will be pressed

After three loops done, input `done`

`L` must be in pairs, or the commands between the last `L` to the end of config will be looped

The number in the first of the paired `L` will be the times to be executed, empty will be zero
