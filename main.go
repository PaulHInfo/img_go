package imgproc

import (
	"fmt"
	"imgproc/filter"
	"imgproc/task"
)

func main() {
	var f filter.Filter = filter.Grayscale{}
	t := task.NewWaitGrpTask("./imgs/", "output", f)
	fmt.Println(t)
}
