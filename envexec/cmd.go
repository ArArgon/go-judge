package envexec

import (
	"context"
	"time"

	"github.com/criyle/go-sandbox/runner"
)

// Size represent data size in bytes
type Size = runner.Size

// RunnerResult represent process finish result
type RunnerResult = runner.Result

// Cmd defines instruction to run a program in container environment
type Cmd struct {
	Environment Environment

	// file contents to copyin before exec
	CopyIn map[string]File

	// exec argument, environment
	Args []string
	Env  []string

	// Files for the executing command
	Files []File
	TTY   bool // use pty as input / output

	// resource limits
	TimeLimit         time.Duration
	MemoryLimit       Size
	StackLimit        Size
	ExtraMemoryLimit  Size
	OutputLimit       Size
	ProcLimit         uint64
	CPURateLimit      float64
	StrictMemoryLimit bool

	// Waiter is called after cmd starts and it should return
	// once time limit exceeded.
	// return true to as TLE and false as normal exits (context finished)
	Waiter func(context.Context, Process) bool

	// file names to copyout after exec
	CopyOut    []CmdCopyOutFile
	CopyOutMax Size // file size limit

	// CopyOutDir specifies a dir to dump all /w contnet
	CopyOutDir string
}

// CmdCopyOutFile defines the file to be copy out after cmd execution
type CmdCopyOutFile struct {
	Name     string // Name is the file out to copyOut
	Optional bool   // Optional ignores the file if not exists
}

// Result defines the running result for single Cmd
type Result struct {
	Status Status

	ExitStatus int

	Error string // error

	Time    time.Duration
	RunTime time.Duration
	Memory  Size // byte

	// Files stores copy out files
	Files map[string][]byte
}
