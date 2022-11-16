package cmd

import (
	"bytes"
	"github.com/fatih/color"
	"github.com/uvite/u8/lib/consts"
	"io"
	"os"
	"sync"
)

func getColor(noColor bool, attributes ...color.Attribute) *color.Color {
	if noColor {
		c := color.New()
		c.DisableColor()
		return c
	}

	c := color.New(attributes...)
	c.EnableColor()
	return c
}
func getBanner(noColor bool) string {
	c := getColor(noColor, color.FgCyan)
	return c.Sprint(consts.Banner())
}

// A writer that syncs writes with a mutex and, if the output is a TTY, clears before newlines.
type consoleWriter struct {
	rawOut *os.File
	writer io.Writer
	isTTY  bool
	mutex  *sync.Mutex

	// Used for flicker-free persistent objects like the progressbars
	persistentText func()
}

func (w *consoleWriter) Write(p []byte) (n int, err error) {
	origLen := len(p)
	if w.isTTY {
		// Add a TTY code to erase till the end of line with each new line
		// TODO: check how cross-platform this is...
		p = bytes.ReplaceAll(p, []byte{'\n'}, []byte{'\x1b', '[', '0', 'K', '\n'})
	}

	w.mutex.Lock()
	n, err = w.writer.Write(p)
	if w.persistentText != nil {
		w.persistentText()
	}
	w.mutex.Unlock()

	if err != nil && n < origLen {
		return n, err
	}
	return origLen, err
}
