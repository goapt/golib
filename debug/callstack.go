package debug

import (
	"runtime"
)

type callStack struct {
	Func   string
	File   string
	LineNo int
}

func Callstack(skipFrames int) callStack {
	pc, file, lineno, _ := runtime.Caller(skipFrames)
	f := runtime.FuncForPC(pc)
	return callStack{f.Name(), file, lineno}
}

// Stack gets the call stack
func Stack(calldepth int) []byte {
	var (
		e             = make([]byte, 1<<16) // 64k
		nbytes        = runtime.Stack(e, false)
		ignorelinenum = 2*calldepth + 1
		count         = 0
		startIndex    = 0
	)
	for i := range e {
		if e[i] == '\n' {
			count++
			if count == ignorelinenum {
				startIndex = i + 1
			}
		}
	}
	return e[startIndex:nbytes]
}