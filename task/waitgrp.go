package task

import "imgproc/filter"

type WaitGrpTask struct {
	dirCtx
	Filter filter.Filter
}

func NewWaitGrpTask(srcDir, dstDir string, Filter filter.Filter) Tasker {
	return &WaitGrpTask{
		Filter: filter,
		dirCtx: dirCtx{
			srcDir: srcDir,
			DstDir: DstDir,
			files:  BuildFileList(srcDir),
		},
	}
}
