// 0.0.1
// github.com/r00m3
/*******************************************************************************


		Program is called : Boggart.

			-- "a generic name for an apparition".
			-- "shape-shifters whose true form is unknown".


Purpose of program Boggart is to convert user input
from one form to another.
Some features are missing for now.

*******************************************************************************/

package main

import (
	"Boggart/lines"
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

func main() {
	print_greeting()
	print_user_selects()
}

/*******************************************************************************

		Stores user input.

*******************************************************************************/

type Input struct {
	tokens  []string
	hex     []uint64
	decimal []uint64
	bytes   []uint64
	runes   []rune
}

/*******************************************************************************

		Triggers user input read from stdin.
		User input gets instantly:

			- collected.
			- tokenized.
			- validated.
			- added to proper struct fields.

******************************************************************************/

func (i *Input) seed() {

	// Collects user input from stdin.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	i.tokens = strings.Fields(scanner.Text())
	i.runes = []rune(scanner.Text())

	// Seeds < hex []uint64 >.
	for _, token := range i.tokens {
		token_parsed, err := strconv.ParseUint(
			token, 16, 64)
		if err == nil {
			i.hex = append(i.hex, token_parsed)
		}
	}

	// Seeds < decimal []uint64 >.
	for _, token := range i.tokens {
		token_parsed, err := strconv.ParseUint(
			token, 10, 64)
		if err == nil {
			i.decimal = append(i.decimal, token_parsed)
		}
	}

	// Seeds < bytes []uint64 >.
	for _, token := range i.tokens {
		token_parsed, err := strconv.ParseUint(
			token, 2, 64)
		if err == nil {
			i.bytes = append(i.bytes, token_parsed)
		}
	}
}

/*******************************************************************************

		Prints out particular struct field
		as blocks of arrays in:

			- base-1	< green >.
			- base-10	< yellow >.
			- base-2	< uncolored >.
			- Unicode	< uncolored >.

*******************************************************************************/

func (i *Input) array_hex() {
	print_array(i.hex)
}

func (i *Input) array_decimal() {
	print_array(i.decimal)
}

func (i *Input) array_bytes() {
	print_array(i.bytes)
}

func (i *Input) array_unicode() {
	print_array(i.runes)
}

/*******************************************************************************

		Prints out particular struct field
		as tokens in:

			- base-16 -> base-10 -> base-2 -> Unicode.

*******************************************************************************/

func (i *Input) grid_hex() {

	for _, token := range i.hex {

		if token <= 0xffffffff {
			len_bits := bits.Len64(token)
			len_bytes := bits_to_bytes(len_bits)
			print_grid(token, len_bits, len_bytes)

		} else if token > 0xffffffff {
			len_bits := bits.Len64(token)
			len_bytes := bits_to_bytes(len_bits)
			print_grid_big(token, len_bits, len_bytes)
		}
	}
}

func (i *Input) grid_decimal() {

	for _, token := range i.decimal {

		if token <= 0xffffffff {
			len_bits := bits.Len64(token)
			len_bytes := bits_to_bytes(len_bits)
			print_grid(token, len_bits, len_bytes)

		} else if token > 0xffffffff {
			len_bits := bits.Len64(token)
			len_bytes := bits_to_bytes(len_bits)
			print_grid_big(token, len_bits, len_bytes)
		}
	}
}

func (i *Input) grid_bytes() {

	for _, token := range i.bytes {

		if token <= 0xffffffff {
			len_bits := bits.Len64(token)
			len_bytes := bits_to_bytes(len_bits)
			print_grid(token, len_bits, len_bytes)

		} else if token > 0xffffffff {
			len_bits := bits.Len64(token)
			len_bytes := bits_to_bytes(len_bits)
			print_grid_big(token, len_bits, len_bytes)
		}
	}
}

func (i *Input) grid_unicode() {

	for _, token := range i.runes {

		len_bits := bits.Len32(uint32(token))
		len_bytes := bits_to_bytes(len_bits)
		print_grid(token, len_bits, len_bytes)
	}
}

/*******************************************************************************


		Helper function used to format
		and print out argument.
		Argument should be one of arrays
		from Input struct instance fields.


tok_array
	Already properly sanitized and valid
	array containing tokens to print.

*******************************************************************************/

func print_array(tok_array any) {

	ln := lines.New{}
	ln.
		// Prints tok_array as [hexadecimal].
		ASCII_new_ln().
		ASCII_new_ln().
		ADD("Hexadecimal:").
		ASCII_new_ln().
		DECOR_margin().
		FG_green().
		ADD("%#x").
		RESET().
		ASCII_new_ln().
		ASCII_new_ln().

		// Prints tok_array as [decimal].
		ADD("Decimal:").
		ASCII_new_ln().
		DECOR_margin().
		FG_yellow().
		ADD("%d").
		RESET().
		ASCII_new_ln().
		ASCII_new_ln().

		// Prints tok_array as [bits].
		ADD("Bit:").
		ASCII_new_ln().
		DECOR_margin().
		ADD("%b").
		ASCII_new_ln().
		ASCII_new_ln().

		// Prints tok_array as [Unicode characters].
		ADD("Unicode character:").
		ASCII_new_ln().
		DECOR_margin().
		ADD("%-5q").
		ASCII_new_ln().
		ASCII_new_ln().
		PRINT(tok_array, tok_array, tok_array, tok_array)
}

/*******************************************************************************


		Helper function used to format
		and print out argument.


tok
	Already properly sanitized and valid
	standalone token. Created by iterating over
	one of arrays from Input struct instance fields.

len_bits
	Minimum number of bits required to represent token.
	len_bits := bits.Len64(token).

len_bytes
	Minimum number of bytes required to represent token.
	len_bytes := bits_to_bytes(len_bits).


if token <= 0xffffffff	{ print_grid() }
if token > 0xffffffff	{ print_grid_big() }

*******************************************************************************/

func print_grid(tok any, len_bits int, len_bytes int) {

	ln := lines.New{}
	ln.
		// Prints tok as hexadecimal.
		ASCII_new_ln().
		DECOR_margin().
		FG_green().
		ADD("%#-10x").
		FG_magenta().
		ADD(" -> ").

		// Prints tok as decimal.
		FG_yellow().
		ADD("%-10d").
		FG_magenta().
		ADD("-[ ").
		RESET().
		ADD("%-2d ").
		FG_magenta().
		ADD("| ").
		RESET().
		ADD("%d").
		FG_magenta().
		ADD(" ]-> ").
		RESET().

		// Prints tok as bits.
		ADD("%#-35b").
		FG_magenta().
		ADD(" -> ").
		RESET().

		// Prints tok as Unicode
		// U+ code and actual char.
		ADD("%#-5U").
		PRINT(tok, tok, len_bits, len_bytes, tok, tok)
}

/*******************************************************************************


		Helper function used to format
		and print out argument.


tok
	Already properly sanitized and valid
	standalone token. Created by iterating over
	one of arrays from Input struct instance fields.

len_bits
	Minimum number of bits required to represent token.
	len_bits := bits.Len64(token).

len_bytes
	Minimum number of bytes required to represent token.
	len_bytes := bits_to_bytes(len_bits).


if token <= 0xffffffff	{ print_grid() }
if token > 0xffffffff	{ print_grid_big() }

*******************************************************************************/

func print_grid_big(tok any, len_bits int, len_bytes int) {

	// In new separate line:
	ln := lines.New{}
	ln.
		// Prints tok as hexadecimal.
		ASCII_new_ln().
		ASCII_new_ln().
		FG_green().
		ADD("%#-18x").
		FG_magenta().
		ADD(" -> ").

		// Prints tok as decimal.
		ASCII_new_ln().
		DECOR_margin().
		DECOR_margin().
		ADD("   -> ").
		FG_yellow().
		ADD("%-10d").

		// Prints tok as bits.
		ASCII_new_ln().
		DECOR_margin().
		FG_magenta().
		ADD("-[ ").
		RESET().
		ADD("%-2d ").
		FG_magenta().
		ADD("| ").
		RESET().
		ADD("%d").
		FG_magenta().
		ADD(" ]-> ").
		RESET().
		ADD("%#b").

		// Prints tok as Unicode
		// U+ code and actual char.
		ASCII_new_ln().
		DECOR_margin().
		DECOR_margin().
		FG_magenta().
		ADD("   -> ").
		RESET().
		ADD("%#U").
		ASCII_new_ln().
		PRINT(tok, tok, len_bits, len_bytes, tok, tok)
}

/*******************************************************************************

		Prints out program name and purpose.

*******************************************************************************/

func print_greeting() {

	ln := lines.New{}
	ln.
		DEL_term().
		ASCII_new_ln().
		ASCII_new_ln().
		DECOR_margin().
		DECOR_bold().
		FG_BR_green().
		ADD("Boggart").
		RESET().
		ASCII_new_ln().
		DECOR_margin().
		DECOR_margin().
		ADD("-- convert input to different encoding representations.").
		ASCII_new_ln().
		ASCII_new_ln().
		ASCII_new_ln().
		PRINT()
}

/*******************************************************************************

		Prints out selection menu.
		Reads and validates selection.
		Calls handler function
		to navigate to proper window:

			- print_window_hex().
			- print_window_bytes().
			- print_window_decimal().
			- print_window_string().
			- print_window_info().
			- print_user_selects().

*******************************************************************************/

func print_user_selects() {

	// Shows user operations to pick from.
	ln := lines.New{}
	ln.
		ADD("Convert from").
		ASCII_new_ln().

		// Hex.
		DECOR_margin().
		FG_cyan().
		ADD("[ h ]").
		RESET().
		ADD("hex").
		ASCII_new_ln().

		// Bytes.
		DECOR_margin().
		FG_cyan().
		ADD("[ b ]").
		RESET().
		ADD("bytes").
		ASCII_new_ln().

		// Decimal.
		DECOR_margin().
		FG_cyan().
		ADD("[ d ]").
		RESET().
		ADD("decimal").
		ASCII_new_ln().

		// String literal.
		DECOR_margin().
		FG_cyan().
		ADD("[ s ]").
		RESET().
		ADD("string literal").
		ASCII_new_ln().

		// Information window.
		DECOR_margin().
		FG_cyan().
		ADD("[ i ]").
		RESET().
		ADD("information window").
		ASCII_new_ln().

		// Enter selection.
		FG_cyan().
		ADD("Enter selection: ").
		RESET().
		PRINT()

	// User picks operation.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	selected_operation := scanner.Text()

	// Handles picked operation.
	if selected_operation == "h" {
		print_window_hex()
	} else if selected_operation == "b" {
		print_window_bytes()
	} else if selected_operation == "d" {
		print_window_decimal()
	} else if selected_operation == "s" {
		print_window_string()
	} else if selected_operation == "i" {
		print_window_info()
	} else {
		print_user_selects()
	}
}

/*******************************************************************************

		All the handler functions:

			- print_window_hex().
			- print_window_bytes().
			- print_window_decimal().
			- print_window_string().
			- print_window_info().


		Each of those:

			- Prints out Boggart name and purpose.
			- Prints short intro about current window.
			- Creates new Input instance.
			- Seeds created instance reading stdin.
			- Converts input to every needed encoding.
			- Prints converted input as array blocks.
			- Prints converted input as token lines.
			- Prints out selection menu to keep going.

*******************************************************************************/

func print_window_hex() {

	print_greeting()
	ln := lines.New{}
	ln.
		ADD("1f9d9 1F52E 1F5DD -> ").
		ADD("\U0001f9d9 \U0001F52E \U0001F5DD").
		ASCII_new_ln().
		FG_cyan().
		ADD("Enter hexadecimal numbers separated by spaces:").
		RESET().
		ASCII_new_ln().
		PRINT()

	in := Input{}
	in.seed()
	in.array_hex()
	in.grid_hex()

	fmt.Printf("\n\n")
	print_user_selects()
}

func print_window_bytes() {

	print_greeting()
	ln := lines.New{}
	ln.
		FG_cyan().
		ADD("Enter byte sequences separated by spaces:").
		RESET().
		ASCII_new_ln().
		PRINT()

	in := Input{}
	in.seed()
	in.array_bytes()
	in.grid_bytes()

	fmt.Printf("\n\n")
	print_user_selects()
}

func print_window_decimal() {

	print_greeting()
	ln := lines.New{}
	ln.
		FG_cyan().
		ADD("Enter decimal numbers separated by spaces:").
		RESET().
		ASCII_new_ln().
		PRINT()

	in := Input{}
	in.seed()
	in.array_decimal()
	in.grid_decimal()

	fmt.Printf("\n\n")
	print_user_selects()
}

func print_window_string() {

	print_greeting()
	ln := lines.New{}
	ln.
		FG_cyan().
		ADD("Enter string literal:").
		RESET().
		ASCII_new_ln().
		PRINT()

	in := Input{}
	in.seed()
	in.array_unicode()
	in.grid_unicode()

	fmt.Printf("\n\n")
	print_user_selects()
}

func print_window_info() {

	print_greeting()
	ln := lines.New{}
	ln.
		FG_red().
		ADD("In future..").
		RESET().
		ASCII_new_ln().
		PRINT()

	// Implement in future.

	fmt.Printf("\n\n")
	print_user_selects()
}

func bits_to_bytes(bits int) int {
	if bits <= 0 {
		return 0
	} else {
		full_bytes := bits / 8
		if bits%8 != 0 {
			full_bytes += 1
		}
		return full_bytes
	}
}
