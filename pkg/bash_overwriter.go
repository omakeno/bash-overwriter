package bashoverwriter

import (
	"fmt"
	"strings"
)

type BashOverwriter struct {
	Row int
}

func GetBashoverwriter() BashOverwriter {
	return BashOverwriter{Row: 0}
}

func (w *BashOverwriter) Write(p []byte) (int, error) {
	w.Print(string(p))
	return len(p), nil
}

func (w *BashOverwriter) Print(out string) {
	Row := strings.Count(out, "\n")
	if w.Row > 0 {
		out = "\033[" + fmt.Sprint(w.Row) + "A\033[0;K" + out
	}
	if w.Row > Row {
		out = strings.Repeat("\033[\n", w.Row-Row) + out
	} else {
		w.Row = Row
	}
	fmt.Print(out)
}
