# EasyConfig

```go
package main

import "fmt"
import "github.com/mimicloud/easyconfig"

var config = struct {
    Some string `json:"some"`
}{}

func main() {
    // Parse config
    easyconfig.Parse("./config.json", &config)

    // Print config
    fmt.Printf("Some: %s\n", config.Some)

    // Change config
    config.Some = "no"

    // Save config
    easyconfig.Save("./config.json", &config)
}
```