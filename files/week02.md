 
不应该，遇到这个错误我们 DAO 层无法根据业务判断能否做降级处理，因为 DAO 相关组件肯定应该和业务相关组件解耦，我们应该想办法把相关信息打包并返回，然后上层业务根据需要来做降级处理还是打印到日志还是其他什么操作。但是应该是不能直接打包这个错误，因为这样会导致这个全局的错误被修改，最终不可用。

```go
package main

import (
	"database/sql"
	"fmt"
)

var myDAOErrNoRows = 6543213
var otherDAOError = 3153564

func MockDaoError() error {
	return nil
}

func Dao(query string) error {
	err := MockDaoError()

	if err == sql.ErrNoRows {
		return fmt.Errorf("%d + 一些堆栈信息", myDAOErrNoRows)
	} else if err != nil {
		return fmt.Errorf("%d + 一些堆栈信息", otherDAOError)
	}

	return nil
}
```
