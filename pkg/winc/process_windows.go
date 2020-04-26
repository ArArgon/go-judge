package winc

import (
	"time"
	"unsafe"

	"github.com/criyle/go-judge/pkg/envexec"
	"github.com/criyle/go-sandbox/runner"
	"golang.org/x/sys/windows"
)

var _ envexec.Process = &process{}

type process struct {
	done     chan struct{}
	result   runner.Result
	hJob     windows.Handle
	hProcess windows.Handle
}

func (p *process) Done() <-chan struct{} {
	return p.done
}

func (p *process) Result() runner.Result {
	<-p.done
	return p.result
}

func (p *process) Usage() envexec.Usage {
	t, m, _ := getJobOjbectUsage(p.hJob)
	return envexec.Usage{
		Time:   t,
		Memory: m,
	}
}

func getJobOjbectUsage(hJob windows.Handle) (time.Duration, runner.Size, error) {
	basicInfo := new(JOBOBJECT_BASIC_ACCOUNTING_INFORMATION)
	if _, err := QueryInformationJobObject(hJob, JobObjectBasicAccountingInformation,
		uintptr(unsafe.Pointer(basicInfo)), uint32(unsafe.Sizeof(*basicInfo)), nil); err != nil {
		return 0, 0, err
	}
	t := time.Duration(basicInfo.TotalUserTime) * 100 // 100 nanosecond tick

	extendedLimitInfo := new(windows.JOBOBJECT_EXTENDED_LIMIT_INFORMATION)
	if _, err := QueryInformationJobObject(hJob, windows.JobObjectExtendedLimitInformation,
		uintptr(unsafe.Pointer(extendedLimitInfo)), uint32(unsafe.Sizeof(*extendedLimitInfo)), nil); err != nil {
		return 0, 0, err
	}
	m := runner.Size(extendedLimitInfo.PeakJobMemoryUsed)
	return t, m, nil
}
