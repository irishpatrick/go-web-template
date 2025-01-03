package util

import "fmt"

type Handler interface {
	Prefix() string
}

func Get(h Handler, url string) string {
	return build(h, "GET", url)
}

func Post(h Handler, url string) string {
	return build(h, "POST", url)
}

func Put(h Handler, url string) string {
	return build(h, "PUT", url)
}

func Patch(h Handler, url string) string {
	return build(h, "PATCH", url)
}

func Delete(h Handler, url string) string {
	return build(h, "DELETE", url)
}

func build(h Handler, header, url string) string {
	if url[0] == '/' {
		url = url[1:]
	}
	if url[0] != '/' {
		url += "/"
	}
	return fmt.Sprintf("%s %s%s", header, h.Prefix(), url)
}
