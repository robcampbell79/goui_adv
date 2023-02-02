package main

import(
	"fmt"
	"log"
	_"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jroimartin/gocui"
)

func storyColor(s string, c string) string {
	if c == "HiMagenta" {
		 return "\033[35;1m"+s+"\033[0m"
	} else {
		return "\033[37;1m"+s+"\033[0m"
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("v1", 0, 0, maxX-30, maxY-10); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Story"
		v.Wrap = true
		// var a int = 5
		// var b int = 1
		// fmt.Fprintf(v, "\033[3%d;%dmJust some random text\033[0m", a, b)
		hiMagenta := storyColor("Just some random words", "HiMagenta")
		fmt.Fprintf(v, hiMagenta)
	}

	if v, err := g.SetView("v2", maxX-30, 0, maxX-1, maxY/2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Inventory"
	}

	if v, err := g.SetView("v3", maxX-30, maxY/2, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Stats"
	}

	if v, err := g.SetView("v4", 0, maxY-10, maxX-50, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Actions"
		v.Highlight = true
		//v.SelFgColor = gocui.ColorGreen
		v.Editable = true

		if _, err := g.SetCurrentView("v4"); err != nil {
			return err
		}
	}

	if v, err := g.SetView("v5", maxX-50, maxY-10, maxX-30, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Location"
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func action(g *gocui.Gui, v *gocui.View) error {
	vbuf := v.ViewBuffer()
	word := strings.TrimSpace(vbuf)

	if word == "next" {

		v.Clear()
		v.SetCursor(0, 0)

		v, err := g.View("v1")
		if err != nil {
			return err
		}

		v.Clear()
		
		hiMagenta := storyColor("And now we have moved beyond", "HiMagenta")
		fmt.Fprintf(v, hiMagenta)

	} else {

		v.Clear()
		v.SetCursor(0, 0)

		v, err := g.View("v1")
		if err != nil {
			return err
		}

		v.Clear()

		hiMagenta := storyColor("\nOops, I messed up", "HiMagenta")
		fmt.Fprintf(v, hiMagenta)
		fmt.Fprintf(v, "\n"+vbuf)
	}

	return nil
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}

	defer g.Close()

	g.Highlight = true
	g.Cursor = true

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("v4", gocui.KeyEnter, gocui.ModNone, action); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}