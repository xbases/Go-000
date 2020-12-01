学习笔记

```bash
go run cmd/error_wraps.go
no such table: nodes
dblayer: get all nodes err
Week02/internal/dblayer.(*DBORM).GetNodes
        /Users/cw/Data/projects/Go-000/Week02/internal/dblayer/orm.go:31
Week02/internal/handler.(*Handler).GetNodes
        /Users/cw/Data/projects/Go-000/Week02/internal/handler/handler.go:37
main.main
        /Users/cw/Data/projects/Go-000/Week02/cmd/error_wraps.go:14
runtime.main
        /usr/local/go/src/runtime/proc.go:204
runtime.goexit
        /usr/local/go/src/runtime/asm_amd64.s:1374
[]
```