package mock

import "strings"

type Retriever struct {
	Contents string
}
// 必须实现interface内的方法
func (r *Retriever) Connect(host string) {
	panic("implement me")
}

func (r *Retriever) Get(url string) string  {
	return strings.Join([]string{r.Contents, url}, "-")
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}
