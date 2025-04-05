package imgproc

import (
	"fmt"
	"imgproc/filter"
	"imgproc/task"
	"time"
)

func main() {
	var f filter.Filter = filter.Grayscale{}
	t := task.NewWaitGrpTask("./imgs/", "output", f)
	start := time.Now()
	t.Process()
	elapsed := time.Since(start)

	fmt.Println(elapsed)
}
