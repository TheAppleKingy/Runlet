package textdata

type runnerTable map[string]string

var Runners = runnerTable{
	"py": "py_runner:50051",
	"go": "go_runner:50051",
	"js": "js_runner:50051",
}
