package fileTypes

func GetFileExtension(fileType string) string {
	switch fileType {
	case Html:
		return ".html"
	case Css:
		return ".css"
	case Javascript:
		return ".js"
	default:
		return ""
	}
}

const Html = "text/html"
const Css = "text/css"
const Javascript = "application/javascript"
