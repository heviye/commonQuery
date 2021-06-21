# commonQuery
封装Go原生查询方法

## 使用

```shell
go get github.com/heviye/commonQuery
```

```go
package main

import (
	"fmt"
	"github.com/heviye/commonQuery"
)

func main() {
	data, err := commonQuery.Query("SELECT A,B,c,d as dog FROM table_name")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	for _, m := range data {
		a := m.String("A")
		b := m.String("B")
		c := m.Int64("c")
		d := m.String("dog")
		
		fmt.Println(a,b,c,d)
    }
}
```
