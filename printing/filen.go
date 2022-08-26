package printing

type PrintSettings struct {
	snyggPrint bool
	colorPrint bool
}

func New(snyggPrint bool, colorPrint bool) *PrintSettings {
	return &PrintSettings{
		snyggPrint: snyggPrint,
		colorPrint: colorPrint,
	}
}
