package templates

var (
	RootPath = "templates/"
	HtmlPath = "html/"
)

func GetRootPath() (path string) {
	path = RootPath
	return
}

func GetHTMLPath() (path string) {
	path = RootPath + HtmlPath
	return
}
