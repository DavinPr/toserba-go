# Description

This is including common function where mostly used by golang developer

# Features

## Pointer

### To Pointer 
This will convert any type to pointer
```
import (
    fmt
    
    github.com/DavinPr/toserba-go/common
)

func main() {
    ptr := common.ToPtr("Pass any value")
    ...
}
```

### From Pointer
this will return the value of pointer and return default value of the type when pointer is nil
```
import (
    fmt
    
    github.com/DavinPr/toserba-go/common
)

func sample(name *string) {
    ptr := common.FromPtr(sample)
}
```