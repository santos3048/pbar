# PBar
A simple package that creates a nice looking and easily manageable progress bar in Go
## Example
```go
import (
        "fmt"
        "pbar/pbar"
        "time"
)

func main() {
        b := pbar.Create("Doing some task", 400)
        go b.Print()
        i := 0
        for i <= 300 {
                time.Sleep(5 * time.Millisecond)
                b.Msg(fmt.Sprintf("Performing task #%d", i))
                b.Up()
                i++
        }
        b.Stop()
        time.Sleep(1 * time.Second)
        go b.Print()
        for i <= 400 {
                time.Sleep(50 * time.Millisecond)
                b.Msg(fmt.Sprintf("Performing task #%d", i))
                b.Up()
                i++
        }
        b.Finish("Performed all tasks succesfully!")
}
```

The output will be a bar that goes fast until 300/400, sleeps for one second and then continues slower. Its message can be easily updated
![image](https://github.com/santos3048/pbar/assets/154024116/cc8252c2-e2ff-49f8-a72b-2029f58acb02)

