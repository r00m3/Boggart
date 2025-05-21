/*******************************************************************************


		w := "World!"
		ln := lines.New{}
		ln.
			FG_red().
			DECOR_cross().
			DECOR_blink().
			ADD("Hello, %s").
			RESET().
			PRINT(w)


All method names are slightly longer,
yet slightly more informative - faster to catch up.
This is done intentionally due to consideration,
that method chaining is better done vertically.
So not taking space from other elements anyway.

ADD text via .ADD(w) instead of:
	ln.
		FG_red(w).
		PRINT()

Limitations contribute to predictability.
Each attribute supposed to be in new line.


.ADD()
	Adds text to < ln string > struct field.

.RESET()
	Resets all text attributes to defaults.
	"Terminal prints" in "normal text" after this.

.PRINT()
	Prints out string from < ln string > struct field.
	You can supply any number of arguments
	to satisfy those %s %d... specifiers.

*******************************************************************************/

package lines

import "fmt"

type New struct {
	ln string
}

func (n *New) ADD(s string) *New {
	n.ln += s
	return n
}

func (n *New) RESET() *New {
	n.ln += DECOR_RES
	return n
}

func (n *New) PRINT(a ...any) {
	fmt.Printf(n.ln, a...)
}

/*******************************************************************************

		Erase commands.

*******************************************************************************/

const (
	ERASE_DISPLAY            string = ESC + "J"
	ERASE_UNTIL_SCREEN_END   string = ESC + "0J"
	ERASE_UNTIL_SCREEN_START string = ESC + "1J"
	ERASE_SCREEN             string = ESC + "2J"
	ERASE_SAVED_LINES        string = ESC + "3J"
	ERASE_IN_LINE            string = ESC + "K"
	ERASE_UNTIL_LINE_END     string = ESC + "0K"
	ERASE_UNTIL_LINE_START   string = ESC + "1K"
	ERASE_LINE               string = ESC + "2K"
	CURSOR_HOME              string = ESC + "H" // Find a place for this.
	ERASE_TERM               string = CURSOR_HOME +
		ERASE_SAVED_LINES +
		ERASE_DISPLAY
)

func (n *New) DEL_display() *New {
	n.ln += ERASE_DISPLAY
	return n
}
func (n *New) DEL_scrn_end() *New {
	n.ln += ERASE_UNTIL_SCREEN_END
	return n
}
func (n *New) DEL_scrn_start() *New {
	n.ln += ERASE_UNTIL_SCREEN_START
	return n
}
func (n *New) DEL_scrn() *New {
	n.ln += ERASE_SCREEN
	return n
}
func (n *New) DEL_saved_ln() *New {
	n.ln += ERASE_SAVED_LINES
	return n
}
func (n *New) DEL_in_ln() *New {
	n.ln += ERASE_IN_LINE
	return n
}
func (n *New) DEL_ln_end() *New {
	n.ln += ERASE_UNTIL_LINE_END
	return n
}
func (n *New) DEL_ln_start() *New {
	n.ln += ERASE_UNTIL_LINE_START
	return n
}
func (n *New) DEL_ln() *New {
	n.ln += ERASE_LINE
	return n
}
func (n *New) DEL_term() *New {
	n.ln += ERASE_TERM
	return n
}

/*******************************************************************************

		ASCII special characters.

*******************************************************************************/

const (
	ESC                   string = "\x1B["
	ASCII_TERMINAL_BELL   string = "\a"
	ASCII_BACKSPACE       string = "\b"
	ASCII_HORIZONTAL_TAB  string = "\t"
	ASCII_NEWLINE         string = "\n"
	ASCII_VERTICAL_TAB    string = "\v"
	ASCII_NEW_PAGE        string = "\f"
	ASCII_CARRIAGE_RETURN string = "\r"
)

func (n *New) ASCII_bel() *New {
	n.ln += ASCII_TERMINAL_BELL
	return n
}
func (n *New) ASCII_backsp() *New {
	n.ln += ASCII_BACKSPACE
	return n
}
func (n *New) ASCII_tab_hor() *New {
	n.ln += ASCII_HORIZONTAL_TAB
	return n
}
func (n *New) ASCII_new_ln() *New {
	n.ln += ASCII_NEWLINE
	return n
}
func (n *New) ASCII_tab_vert() *New {
	n.ln += ASCII_VERTICAL_TAB
	return n
}
func (n *New) ASCII_new_page() *New {
	n.ln += ASCII_NEW_PAGE
	return n
}
func (n *New) ASCII_crg_return() *New {
	n.ln += ASCII_CARRIAGE_RETURN
	return n
}

/*******************************************************************************

		Decorations.

*******************************************************************************/

const (
	DECOR_RES              string = ESC + "0m"
	DECOR_BOLD             string = ESC + "1m"
	DECOR_DIM              string = ESC + "2m"
	DECOR_ITALIC           string = ESC + "3m"
	DECOR_UNDERLINE        string = ESC + "4m"
	DECOR_BLINK            string = ESC + "5m"
	DECOR_BLINK_RAPID      string = ESC + "6m"
	DECOR_INVERSE          string = ESC + "7m"
	DECOR_INVISIBLE_TEXT   string = ESC + "8m"
	DECOR_VISIBLE_TEXT     string = ESC + "28m"
	DECOR_STRIKETHROUGH    string = ESC + "9m"
	DECOR_INVISIBLE_CURSOR string = ESC + "?25l"
	DECOR_VISIBLE_CURSOR   string = ESC + "?25h"
	DECOR_MARGIN           string = "        "
)

func (n *New) DECOR_bold() *New {
	n.ln += DECOR_BOLD
	return n
}
func (n *New) DECOR_dim() *New {
	n.ln += DECOR_DIM
	return n
}
func (n *New) DECOR_ita() *New {
	n.ln += DECOR_ITALIC
	return n
}
func (n *New) DECOR_underln() *New {
	n.ln += DECOR_UNDERLINE
	return n
}
func (n *New) DECOR_blink() *New {
	n.ln += DECOR_BLINK
	return n
}
func (n *New) DECOR_blink_fast() *New {
	n.ln += DECOR_BLINK_RAPID
	return n
}
func (n *New) DECOR_inverse() *New {
	n.ln += DECOR_INVERSE
	return n
}
func (n *New) DECOR_hide_text() *New {
	n.ln += DECOR_INVISIBLE_TEXT
	return n
}
func (n *New) DECOR_unhide_text() *New {
	n.ln += DECOR_VISIBLE_TEXT
	return n
}
func (n *New) DECOR_cross() *New {
	n.ln += DECOR_STRIKETHROUGH
	return n
}
func (n *New) DECOR_hide_cursor() *New {
	n.ln += DECOR_INVISIBLE_CURSOR
	return n
}
func (n *New) DECOR_unhide_cursor() *New {
	n.ln += DECOR_VISIBLE_CURSOR
	return n
}
func (n *New) DECOR_margin() *New {
	n.ln += DECOR_MARGIN
	return n
}

/*******************************************************************************

		4-bit colors.

*******************************************************************************/

const (
	FG_BLACK   string = ESC + "30m"
	FG_RED     string = ESC + "31m"
	FG_GREEN   string = ESC + "32m"
	FG_YELLOW  string = ESC + "33m"
	FG_BLUE    string = ESC + "34m"
	FG_MAGENTA string = ESC + "35m"
	FG_CYAN    string = ESC + "36m"
	FG_WHITE   string = ESC + "37m"

	FG_BRIGHT_BLACK   string = ESC + "90m"
	FG_BRIGHT_RED     string = ESC + "91m"
	FG_BRIGHT_GREEN   string = ESC + "92m"
	FG_BRIGHT_YELLOW  string = ESC + "93m"
	FG_BRIGHT_BLUE    string = ESC + "94m"
	FG_BRIGHT_MAGENTA string = ESC + "95m"
	FG_BRIGHT_CYAN    string = ESC + "96m"
	FG_BRIGHT_WHITE   string = ESC + "97m"

	BG_BLACK   string = ESC + "40m"
	BG_RED     string = ESC + "41m"
	BG_GREEN   string = ESC + "42m"
	BG_YELLOW  string = ESC + "43m"
	BG_BLUE    string = ESC + "44m"
	BG_MAGENTA string = ESC + "45m"
	BG_CYAN    string = ESC + "46m"
	BG_WHITE   string = ESC + "47m"

	BG_BRIGHT_BLACK   string = ESC + "100m"
	BG_BRIGHT_RED     string = ESC + "101m"
	BG_BRIGHT_GREEN   string = ESC + "102m"
	BG_BRIGHT_YELLOW  string = ESC + "103m"
	BG_BRIGHT_BLUE    string = ESC + "104m"
	BG_BRIGHT_MAGENTA string = ESC + "105m"
	BG_BRIGHT_CYAN    string = ESC + "106m"
	BG_BRIGHT_WHITE   string = ESC + "107m"
)

// Foreground.
func (n *New) FG_black() *New {
	n.ln += FG_BLACK
	return n
}
func (n *New) FG_red() *New {
	n.ln += FG_RED
	return n
}
func (n *New) FG_green() *New {
	n.ln += FG_GREEN
	return n
}
func (n *New) FG_yellow() *New {
	n.ln += FG_YELLOW
	return n
}
func (n *New) FG_blue() *New {
	n.ln += FG_BLUE
	return n
}
func (n *New) FG_magenta() *New {
	n.ln += FG_MAGENTA
	return n
}
func (n *New) FG_cyan() *New {
	n.ln += FG_CYAN
	return n
}
func (n *New) FG_white() *New {
	n.ln += FG_WHITE
	return n
}

// Foreground bright.
func (n *New) FG_BR_black() *New {
	n.ln += FG_BRIGHT_BLACK
	return n
}
func (n *New) FG_BR_red() *New {
	n.ln += FG_BRIGHT_RED
	return n
}
func (n *New) FG_BR_green() *New {
	n.ln += FG_BRIGHT_GREEN
	return n
}
func (n *New) FG_BR_yellow() *New {
	n.ln += FG_BRIGHT_YELLOW
	return n
}
func (n *New) FG_BR_blue() *New {
	n.ln += FG_BRIGHT_BLUE
	return n
}
func (n *New) FG_BR_magenta() *New {
	n.ln += FG_BRIGHT_MAGENTA
	return n
}
func (n *New) FG_BR_cyan() *New {
	n.ln += FG_BRIGHT_CYAN
	return n
}
func (n *New) FG_BR_white() *New {
	n.ln += FG_BRIGHT_WHITE
	return n
}

// Background.
func (n *New) BG_black() *New {
	n.ln += BG_BLACK
	return n
}
func (n *New) BG_red() *New {
	n.ln += BG_RED
	return n
}
func (n *New) BG_green() *New {
	n.ln += BG_GREEN
	return n
}
func (n *New) BG_yellow() *New {
	n.ln += BG_YELLOW
	return n
}
func (n *New) BG_blue() *New {
	n.ln += BG_BLUE
	return n
}
func (n *New) BG_magenta() *New {
	n.ln += BG_MAGENTA
	return n
}
func (n *New) BG_cyan() *New {
	n.ln += BG_CYAN
	return n
}
func (n *New) BG_white() *New {
	n.ln += BG_WHITE
	return n
}

// Background bright.
func (n *New) BG_BR_black() *New {
	n.ln += BG_BRIGHT_BLACK
	return n
}
func (n *New) BG_BR_red() *New {
	n.ln += BG_BRIGHT_RED
	return n
}
func (n *New) BG_BR_green() *New {
	n.ln += BG_BRIGHT_GREEN
	return n
}
func (n *New) BG_BR_yellow() *New {
	n.ln += BG_BRIGHT_YELLOW
	return n
}
func (n *New) BG_BR_blue() *New {
	n.ln += BG_BRIGHT_BLUE
	return n
}
func (n *New) BG_BR_magenta() *New {
	n.ln += BG_BRIGHT_MAGENTA
	return n
}
func (n *New) BG_BR_cyan() *New {
	n.ln += BG_BRIGHT_CYAN
	return n
}
func (n *New) BG_BR_white() *New {
	n.ln += BG_BRIGHT_WHITE
	return n
}
