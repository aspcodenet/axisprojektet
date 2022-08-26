package printing

import (
	"fmt"

	"github.com/cheynewallace/tabby"
)

func (ps *PrintSettings) OurPrint(txt string) {
	if !ps.SnyggPrint {
		fmt.Println(txt)
		return
	}
	t := tabby.New()
	t.AddHeader("jkjfg", "4234234")
	t.AddLine(txt, "yes")
	t.Print()
}
