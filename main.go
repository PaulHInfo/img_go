package imgproc

import (
	"fmt"
	"imgproc/task"
)

func main() {
	f := task.BuildFileList("./imgs/")
	fmt.Println(f)

}
