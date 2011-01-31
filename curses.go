package curses

// #include <stdlib.h>
// #include <ncurses.h>
// void Printw (const char* str) { printw (str); } 
// void Wprintw (WINDOW* win, const char* str) { wprintw (win, str); } 
// void Mvprintw (int y, int x, const char* str) { mvprintw (y, x, str); } 
// void Mvwprintw (WINDOW* win, int y, int x, const char* str) { mvwprintw (win, y, x, str); } 
import "C"
import "unsafe"
import "fmt"

// Curses window type.
type Window C.WINDOW

// Standard window.
var Stdscr *Window = (*Window) (C.stdscr)

// Initializes curses.
// This function should be called before using the package.
func Initscr() (*Window) {
    Stdscr = (*Window) (C.initscr())
    return Stdscr
}

// Raw input. No buffering.
// CTRL+Z and CTRL+C passed to the application.
func Raw() {
    C.raw()
}

// No buffering.
func Cbreak() {
    C.cbreak()
}

// Enable character echoing while reading.
func Echo() {
    C.echo()
}

// Disables character echoing while reading.
func Noecho() {
    C.noecho()
}

// Enable reading of function keys.
func (win *Window) Keypad (on bool) {
    C.keypad((*C.WINDOW) (win), C.bool(on))
}

// Get char from the standard in.
func (win *Window) Getch() int {
    return int(C.wgetch((*C.WINDOW)(win)))
}

// Get char from the standard in.
func Getch() int {
    return int(C.getch())
}

// Print in current location.
func Printw (str ... interface{}) {
    res := (*C.char) (C.CString(fmt.Sprint(str...)))
    defer C.free(unsafe.Pointer(res))
    C.Printw(res)
}

// Print in current window.
func (win *Window) Printw (str ... interface{}) {
    res := (*C.char) (C.CString(fmt.Sprint(str...)))
    defer C.free(unsafe.Pointer(res))
    C.Wprintw((*C.WINDOW) (win), res)
}

// Move cursor than print.
func Mvprintw (y, x int, str ... interface{}) {
    res := (*C.char) (C.CString(fmt.Sprint(str...)))
    defer C.free(unsafe.Pointer(res))
    C.Mvprintw(C.int(y), C.int(x), res)
}

// Move cursor than print inside window.
func (win *Window) Mvprintw (y, x int, str ... interface{}) {
    res := (*C.char) (C.CString(fmt.Sprint(str...)))
    defer C.free(unsafe.Pointer(res))
    C.Mvwprintw((*C.WINDOW) (win), C.int(y), C.int(x), res)
}

// Enable attribute
func Attron (attr int) {
    C.attron(C.int(attr))
}

// Disable attribute
func Attroff (attr int) {
    C.attroff(C.int(attr))
}

// Refresh screen.
func Refresh() {
    C.refresh()
}

// Refresh given window.
func (win *Window) Refresh() {
    C.wrefresh((*C.WINDOW) (win))
}

// Finalizes curses.
func End() {
    C.endwin()
}

// Create new window.
func NewWindow (height, width, starty, startx int) *Window {
    return (*Window) (C.newwin(C.int(height), C.int(width),
    C.int(starty), C.int(startx)))
}

// Set box lines.
func (win *Window) Box (v, h int) {
    C.box((*C.WINDOW) (win), C.chtype(v), C.chtype(h))
}

// Set border characters.
// 1. ls: character to be used for the left side of the window 
// 2. rs: character to be used for the right side of the window 
// 3. ts: character to be used for the top side of the window 
// 4. bs: character to be used for the bottom side of the window 
// 5. tl: character to be used for the top left corner of the window 
// 6. tr: character to be used for the top right corner of the window 
// 7. bl: character to be used for the bottom left corner of the window 
// 8. br: character to be used for the bottom right corner of the window
func (win *Window) Border(ls, rs, ts, bs, tl, tr, bl, br int) {
    C.wborder((*C.WINDOW) (win), C.chtype(ls), C.chtype(rs), C.chtype(ts), C.chtype(bs), C.chtype(tl), C.chtype(tr), C.chtype(bl), C.chtype(br))
}

// Delete current window.
func (win *Window) Del () {
    C.delwin((*C.WINDOW) (win))
}

// Get windows sizes.
func (win *Window) Getmaxyx () (row, col int) {
    row = int(C.getmaxy((*C.WINDOW)(win)))
    col = int(C.getmaxx((*C.WINDOW)(win)))
    return row, col
}

// Get terminal size.
func Getmaxyx () (row, col int) {
    row = int(C.LINES)
    col = int(C.COLS)
    return row, col
}
