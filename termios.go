// Go bindings for termios(3).

package termios

/*
#include <stdio.h>
#include <termios.h>
#include <unistd.h>
*/
import "C"
import "errors"

// Special Control Characters
// Index into Termios.CC[] character array.
// c_cc characters
const (
	// Name            Enabled by
	VINTR    =  0   // ISIG
	VQUIT    =  1   // ISIG
	VERASE   =  2   // ICANON
	VKILL    =  3   // ICANON
	VEOF     =  4   // ICANON
	VTIME    =  5   // !ICANON
	VMIN     =  6   // !ICANON
	VSWTC    =  7
	VSTART   =  8   // IXON, IXOFF
	VSTOP    =  9   // IXON, IXOFF
	VSUSP    = 10   // ISIG
	VEOL     = 11   // ICANON
	VREPRINT = 12   // ICANON together with IEXTEN
	VDISCARD = 13   // IEXTEN
	VWERASE  = 14   // ICANON together with IEXTEN
	VLNEXT   = 15   // IEXTEN
	VEOL2    = 16   // ICANON together with IEXTEN
	NCCS     = 32   // size of c_cc[]
)

// Input flags - software input processing
// c_iflag bits
const (
	IGNBRK  = 0000001 // ignore BREAK condition
	BRKINT  = 0000002 // map BREAK to SIGINTR
	IGNPAR  = 0000004 // ignore (discard) parity errors
	PARMRK  = 0000010 // mark parity and framing errors
	INPCK   = 0000020 // enable checking of parity errors
	ISTRIP  = 0000040 // strip 8th bit off chars
	INLCR   = 0000100 // map NL into CR
	IGNCR   = 0000200 // ignore CR
	ICRNL   = 0000400 // map CR to NL (ala CRMOD)
	IUCLC   = 0001000
	IXON    = 0002000 // enable output flow control
	IXANY   = 0004000 // any char will restart after stop
	IXOFF   = 0010000 // enable input flow control
	IMAXBEL = 0020000 // ring bell on input queue full
	IUTF8   = 0040000 // maintain state for UTF-8 VERASE
)

// Output flags - software output processing
// c_oflag bits
const (
	OPOST  = 0000001 // enable following output processing
	OLCUC  = 0000002
	ONLCR  = 0000004 // map NL to CR-NL (ala CRMOD)
	OCRNL  = 0000010 // map CR to NL on output
	ONOCR  = 0000020 // no CR output at column 0
	ONLRET = 0000040 // NL performs CR function
	OFILL  = 0000100 // use fill characters for delay
	OFDEL  = 0000200 // fill is DEL, else NUL
	NLDLY  = 0000400 // \n delay
	NL0    = 0000000
	NL1    = 0000400
	CRDLY  = 0003000 // \r delay
	CR0    = 0000000
	CR1    = 0001000
	CR2    = 0002000
	CR3    = 0003000
	TABDLY = 0014000 // horizontal tab delay
	TAB0   = 0000000
	TAB1   = 0004000
	TAB2   = 0010000
	TAB3   = 0014000
	BSDLY  = 0020000 // \b delay
	BS0    = 0000000
	BS1    = 0020000
	FFDLY  = 0100000 // form feed delay
	FF0    = 0000000
	FF1    = 0100000
	VTDLY  = 0040000 // vertical tab delay
	VT0    = 0000000
	VT1    = 0040000
	XTABS  = 0014000
)

// "Local" flags - dumping ground for other state
//
// Warning: some flags in this structure begin with
// the letter "I" and look like they belong in the
// input flag.
// c_lflag bits
const (
	ISIG       = 0000001 // enable signals INTR, QUIT, [D]SUSP
	ICANON     = 0000002 // canonicalize input lines
	XCASE      = 0000004
	ECHO       = 0000010 // enable echoing
	ECHOE      = 0000020 // visually erase chars
	ECHOK      = 0000040 // echo NL after line kill
	ECHONL     = 0000100 // echo NL even if ECHO is off
	NOFLSH     = 0000200 // don't flush after interrupt
	TOSTOP     = 0000400 // stop background jobs from output
	ECHOCTL    = 0001000 // echo control chars as ^(Char)
	ECHOPRT    = 0002000 // visual erase mode for hardcopy
	ECHOKE     = 0004000 // visual erase for line kill
	FLUSHO     = 0010000 // output being flushed (state)
	PENDIN     = 0040000 // XXX retype pending input (state)
	IEXTEN     = 0100000 // enable DISCARD and LNEXT
	EXTPROC    = 0200000 // external processing
)

// Commands passed to SetAttr() for setting the termios structure.
// tcsetattr use these
// tcsetattr uses these
const (
	TCSANOW   = 0    // make change immediate
	TCSADRAIN = 1    // drain output, then change
	TCSAFLUSH = 2    // drain output, flush input
)

// Standard speeds
// c_cflag bit
const (
	CBAUD     = 0010017
	B0        = 0000000 // hang up
	B50       = 0000001
	B75       = 0000002
	B110      = 0000003
	B134      = 0000004
	B150      = 0000005
	B200      = 0000006
	B300      = 0000007
	B600      = 0000010
	B1200     = 0000011
	B1800     = 0000012
	B2400     = 0000013
	B4800     = 0000014
	B9600     = 0000015
	B19200    = 0000016
	B38400    = 0000017
	EXTA      = 0x0B19200
	EXTB      = 0x0B38400
	CSIZE     = 0000060
	CS5       = 0000000
	CS6       = 0000020
	CS7       = 0000040
	CS8       = 0000060
	CSTOPB    = 0000100
	CREAD     = 0000200
	PARENB    = 0000400
	PARODD    = 0001000
	HUPCL     = 0002000
	CLOCAL    = 0004000
	CBAUDEX   = 0010000
	B57600    = 0010001
	B115200   = 0010002
	B230400   = 0010003
	B460800   = 0010004
	B500000   = 0010005
	B576000   = 0010006
	B921600   = 0010007
	B1000000  = 0010010
	B1152000  = 0010011
	B1500000  = 0010012
	B2000000  = 0010013
	B2500000  = 0010014
	B3000000  = 0010015
	B3500000  = 0010016
	B4000000  = 0010017
	_MAX_BAUD = 0xB4000000
	CIBAUD    = 002003600000 // input baud rate (not used)
	CMSPAR    = 010000000000   // mark or space (stick) parity
	CRTSCTS   = 020000000000   // flow control
)

// tcflow() and TCXONC use these
const (
	TCOOFF    = 0
	TCOON     = 1
	TCIOFF    = 2
	TCION     = 3
// tcflush() and TCFLSH use these
	TCIFLUSH  = 0
	TCOFLUSH  = 1
	TCIOFLUSH = 2
)

var (
	Stdin  = int(C.fileno(C.stdin))
	Stdout = int(C.fileno(C.stdout))
	Stderr = int(C.fileno(C.stderr))
)

type Flag uint
type CC byte
type Speed uint

type Termios struct {
	IFlag  Flag     // input flags
	OFlag  Flag     // output flags
	CFlag  Flag     // control flags
	LFlag  Flag     // local flags
	CC     [NCCS]CC // control chars
	ISpeed Speed    // input speed
	OSpeed Speed    // output speed
}

// Wrapper of tcgetattr(3).
// The GetAttr() function copies the parameters associated with the terminal.
func (t *Termios) GetAttr(fd int) error {
	var cTerm C.struct_termios

	if C.tcgetattr(C.int(fd), &cTerm) == -1 {
		return errors.New("tcgetattr failure")
	}
	*t = *goTermios(&cTerm)
	return nil
}

// Wrapper of tcsetattr(3).
// The SetAttr() function sets the parameters associated with the terminal.
func (t *Termios) SetAttr(fd int, opt int) error {
	if C.tcsetattr(C.int(fd), C.int(opt), cTermios(t)) == -1 {
		return errors.New("tcsetattr failure")
	}
	return nil
}

func goTermios(cTerm *C.struct_termios) *Termios {
	cc := [NCCS]CC{}
	for idx, ch := range cTerm.c_cc {
		cc[idx] = CC(ch)
	}

	return &Termios{
		IFlag:  Flag(cTerm.c_iflag),
		OFlag:  Flag(cTerm.c_oflag),
		CFlag:  Flag(cTerm.c_cflag),
		LFlag:  Flag(cTerm.c_lflag),
		CC:     cc,
		ISpeed: Speed(cTerm.c_ispeed),
		OSpeed: Speed(cTerm.c_ospeed),
	}
}

func cTermios(goTerm *Termios) *C.struct_termios {
	var cTerm C.struct_termios

	cTerm.c_iflag = C.tcflag_t(goTerm.IFlag)
	cTerm.c_oflag = C.tcflag_t(goTerm.OFlag)
	cTerm.c_cflag = C.tcflag_t(goTerm.CFlag)
	cTerm.c_lflag = C.tcflag_t(goTerm.LFlag)
	for idx, ch := range goTerm.CC {
		cTerm.c_cc[idx] = C.cc_t(ch)
	}
	cTerm.c_ispeed = C.speed_t(goTerm.ISpeed)
	cTerm.c_ospeed = C.speed_t(goTerm.OSpeed)

	return &cTerm
}
