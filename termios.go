/*
 * Go bindings for termios(3).
 */
package termios

/*
#include <stdio.h>
#include <termios.h>
#include <unistd.h>
*/
import "C"
import "errors"

/*
 * Special Control Characters
 * Index into Termios.CC[] character array.
 *
 * Name	          Enabled by
 */
const (
	VEOF     = iota // ICANON
	VEOL            // ICANON
	VEOL2           // ICANON together with IEXTEN
	VERASE          // ICANON
	VWERASE         // ICANON together with IEXTEN
	VKILL           // ICANON
	VREPRINT        // ICANON together with IEXTEN
	_               // spare 1
	VINTR           // ISIG
	VQUIT           // ISIG
	VSUSP           // ISIG
	VDSUSP          // ISIG together with IEXTEN
	VSTART          // IXON, IXOFF
	VSTOP           // IXON, IXOFF
	VLNEXT          // IEXTEN
	VDISCARD        // IEXTEN
	VMIN            // !ICANON
	VTIME           // !ICANON
	VSTATUS         // ICANON together with IEXTEN
	_               // spare 2
	NCCS            // size of c_cc[]
)

/*
 * Input flags - software input processing
 */
const (
	IGNBRK  = 0x00000001 // ignore BREAK condition
	BRKINT  = 0x00000002 // map BREAK to SIGINTR
	IGNPAR  = 0x00000004 // ignore (discard) parity errors
	PARMRK  = 0x00000008 // mark parity and framing errors
	INPCK   = 0x00000010 // enable checking of parity errors
	ISTRIP  = 0x00000020 // strip 8th bit off chars
	INLCR   = 0x00000040 // map NL into CR
	IGNCR   = 0x00000080 // ignore CR
	ICRNL   = 0x00000100 // map CR to NL (ala CRMOD)
	IXON    = 0x00000200 // enable output flow control
	IXOFF   = 0x00000400 // enable input flow control
	IXANY   = 0x00000800 // any char will restart after stop
	IMAXBEL = 0x00002000 // ring bell on input queue full
	IUTF8   = 0x00004000 // maintain state for UTF-8 VERASE
)

/*
 * Output flags - software output processing
 */
const (
	OPOST  = 0x00000001 // enable following output processing
	ONLCR  = 0x00000002 // map NL to CR-NL (ala CRMOD)
	OXTABS = 0x00000004 // expand tabs to spaces
	ONOEOT = 0x00000008 // discard EOT's (^D) on output)

	// unimplemented features
	OCRNL  = 0x00000010 // map CR to NL on output
	ONOCR  = 0x00000020 // no CR output at column 0
	ONLRET = 0x00000040 // NL performs CR function
	OFILL  = 0x00000080 // use fill characters for delay
	NLDLY  = 0x00000300 // \n delay
	TABDLY = 0x00000c04 // horizontal tab delay
	CRDLY  = 0x00003000 // \r delay
	FFDLY  = 0x00004000 // form feed delay
	BSDLY  = 0x00008000 // \b delay
	VTDLY  = 0x00010000 // vertical tab delay
	OFDEL  = 0x00020000 // fill is DEL, else NUL
	NL0    = 0x00000000
	NL1    = 0x00000100
	NL2    = 0x00000200
	NL3    = 0x00000300
	TAB0   = 0x00000000
	TAB1   = 0x00000400
	TAB2   = 0x00000800
	TAB3   = 0x00000004
	CR0    = 0x00000000
	CR1    = 0x00001000
	CR2    = 0x00002000
	CR3    = 0x00003000
	FF0    = 0x00000000
	FF1    = 0x00004000
	BS0    = 0x00000000
	BS1    = 0x00008000
	VT0    = 0x00000000
	VT1    = 0x00010000
)

/*
 * "Local" flags - dumping ground for other state
 *
 * Warning: some flags in this structure begin with
 * the letter "I" and look like they belong in the
 * input flag.
 */
const (
	ECHOKE     = 0x00000001 // visual erase for line kill
	ECHOE      = 0x00000002 // visually erase chars
	ECHOK      = 0x00000004 // echo NL after line kill
	ECHO       = 0x00000008 // enable echoing
	ECHONL     = 0x00000010 // echo NL even if ECHO is off
	ECHOPRT    = 0x00000020 // visual erase mode for hardcopy
	ECHOCTL    = 0x00000040 // echo control chars as ^(Char)
	ISIG       = 0x00000080 // enable signals INTR, QUIT, [D]SUSP
	ICANON     = 0x00000100 // canonicalize input lines
	ALTWERASE  = 0x00000200 // use alternate WERASE algorithm
	IEXTEN     = 0x00000400 // enable DISCARD and LNEXT
	EXTPROC    = 0x00000800 // external processing
	TOSTOP     = 0x00400000 // stop background jobs from output
	FLUSHO     = 0x00800000 // output being flushed (state)
	NOKERNINFO = 0x02000000 // no kernel output from VSTATUS
	PENDIN     = 0x20000000 // XXX retype pending input (state)
	NOFLSH     = 0x80000000 // don't flush after interrupt
)

/*
 * Commands passed to SetAttr() for setting the termios structure.
 */
const (
	TCSANOW   = 0    // make change immediate
	TCSADRAIN = 1    // drain output, then change
	TCSAFLUSH = 2    // drain output, flush input
	TCSASOFT  = 0x10 // flag - don't alter h.w. state
)

/*
 * Standard speeds
 */
const (
	B0      = 0
	B50     = 50
	B75     = 75
	B110    = 110
	B134    = 134
	B150    = 150
	B200    = 200
	B300    = 300
	B600    = 600
	B1200   = 1200
	B1800   = 1800
	B2400   = 2400
	B4800   = 4800
	B9600   = 9600
	B19200  = 19200
	B38400  = 38400
	B7200   = 7200
	B14400  = 14400
	B28800  = 28800
	B57600  = 57600
	B76800  = 76800
	B115200 = 115200
	B230400 = 230400
	EXTA    = 19200
	EXTB    = 38400
)

const (
	TCIFLUSH  = 1
	TCOFLUSH  = 2
	TCIOFLUSH = 3
	TCOOFF    = 1
	TCOON     = 2
	TCIOFF    = 3
	TCION     = 4
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

/*
 * Wrapper of tcgetattr(3).
 * The GetAttr() function copies the parameters associated with the terminal.
 */
func (t *Termios) GetAttr(fd int) error {
	var cTerm C.struct_termios

	if C.tcgetattr(C.int(fd), &cTerm) == -1 {
		return errors.New("tcgetattr failure")
	}
	*t = *goTermios(&cTerm)
	return nil
}

/*
 * Wrapper of tcsetattr(3).
 * The SetAttr() function sets the parameters associated with the terminal.
 */
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
