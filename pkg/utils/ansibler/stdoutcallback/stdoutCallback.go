package stdoutcallback

import (
	"io"
	"os"

	"yunion.io/x/kubecomps/pkg/utils/ansibler/stdoutcallback/results"
)

// StdoutCallbackResultsFunc defines a function which manages ansible's stdout callbacks. The function expects and string for prefixing output lines, a reader that receives the data to be wrote and a writer that defines where to write the data comming from reader
type StdoutCallbackResultsFunc func(string, io.Reader, io.Writer) error

const (
	// AnsibleStdoutCallbackEnv
	AnsibleStdoutCallbackEnv = "ANSIBLE_STDOUT_CALLBACK"
)

const (
	// DebugStdoutCallback formatted stdout/stderr output
	DebugStdoutCallback = "debug"
	// DefaultStdoutCallback default ansible screen output
	DefaultStdoutCallback = "default"
	// DenseStdoutCallback minimal stdout output
	DenseStdoutCallback = "dense"
	// JSONStdoutCallback ansible screen output as json
	JSONStdoutCallback = "json"
	// MinimalStdoutCallback minmal ansible screen output
	MinimalStdoutCallback = "minimal"
	// NullStdoutCallback don't display stuff to screen
	NullStdoutCallback = "null"
	// OnelineStdoutCallback oneline ansible screen output
	OnelineStdoutCallback = "oneline"
	// StderrStdoutCallback splits output, sending failed tasks to stderr
	StderrStdoutCallback = "stderr"
	// TimerStdoutCallback adds time to play stats
	TimerStdoutCallback = "timer"
	// YamlStdoutCallback yamlized ansible screen output
	YamlStdoutCallback = "yaml"
)

// AnsibleStdoutCallbackToJSON sets the stdout callback to json
func AnsibleStdoutCallbackSetEnv(callback string) {

	if callback != DebugStdoutCallback &&
		callback != DefaultStdoutCallback &&
		callback != DenseStdoutCallback &&
		callback != JSONStdoutCallback &&
		callback != MinimalStdoutCallback &&
		callback != NullStdoutCallback &&
		callback != OnelineStdoutCallback &&
		callback != StderrStdoutCallback &&
		callback != TimerStdoutCallback &&
		callback != YamlStdoutCallback {
		callback = DefaultStdoutCallback
	}

	os.Setenv(AnsibleStdoutCallbackEnv, callback)
}

// GetResultsFunc return a func which manages the stdout callback results
func GetResultsFunc(callback string) StdoutCallbackResultsFunc {

	switch callback {
	case JSONStdoutCallback:
		return results.JSONStdoutCallbackResults
	default:
		return results.DefaultStdoutCallbackResults
	}
}
