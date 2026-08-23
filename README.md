# buildinfo

可复用的程序二进制版本信息库：内置版本字段、默认值回落、文本/JSON 输出，并提供 cobra `version` 子命令。

## 快速开始

```bash
go get github.com/cocomhub/buildinfo
```

在 `main.go` 中注入：

```go
package main

import (
	"os"

	"github.com/cocomhub/buildinfo"
)

func main() {
	i := buildinfo.Default()
	_ = i.PrintVersion(os.Stdout)      // 文本
	_ = i.PrintVersionJSON(os.Stdout) // JSON
}
```

构建时用 `-ldflags` 覆盖 `buildinfo` 包级变量：

```bash
go build -ldflags "-X github.com/cocomhub/buildinfo.Version=v1.0.0 -X github.com/cocomhub/buildinfo.Branch=main -X github.com/cocomhub/buildinfo.CommitID=abc123 -X github.com/cocomhub/buildinfo.BuiltAt=2026-08-23T00:00:00Z" .
```

大程序二进制通过 `Version`、`CommitID`、`BuiltAt`、`DirtyInfo` 等包级变量注入，运行时用 `Default()` 组装 `Info` 并输出。

## 代码示例

```go
import (
	"bytes"

	"github.com/cocomhub/buildinfo"
)

// 组装：未注入就回落默认值（dev / unknown / 运行时 GOOS·GOARCH）
i := buildinfo.Info{
	Version:   "dev",
	CommitID:  "unknown",
	DirtyInfo: "modified-file.txt",
}

// DirtyID：对 DirtyInfo 取 md5 前缀（空则 "clean"）
id := i.DirtyID() // fmt.Sprintf("%x", md5.Sum(...))[:10]

var buf bytes.Buffer
_ = i.PrintVersion(&buf) // 文本输出，含 Version/Branch/DirtyID/CommitID/Runtime/BuiltAt/ReleaseURL
_ = i.PrintVersionJSON(&buf) // JSON map[string]string 输出

// cobra 子命令： --json 切换输出
cmd := buildinfo.NewVersionCmd(i)
cmd.SetArgs([]string{"--json"})
_ = cmd.Execute()
```

## 导出 API

| 类型/函数 | 说明 |
|---|---|
| `Info` | 版本信息结构体（`Version` `Branch` `CommitID` `DirtyInfo` `BuiltAt` `ReleaseURL` `GoVersion` `GOOS` `GOARCH`） |
| `New() Info` | 构造运行时默认 `Info`（自动填 `GoVersion`/`GOOS`/`GOARCH`） |
| `Default() Info` | 由包级变量（`-X` 注入）填充的 `Info` |
| `(Info) DirtyID() string` | `DirtyInfo` 的 10 位 md5 摘要；空返回 `"clean"` |
| `(Info) PrintVersion(w io.Writer) error` | 文本版本输出 |
| `(Info) PrintVersionJSON(w io.Writer) error` | JSON 版本输出 |
| `NewVersionCmd(i Info) *cobra.Command` | `version` 子命令（`--json` flag） |

包级可注入变量（仅 `-X` 使用，不打算运行时修改）：`Version`、`Branch`、`CommitID`、`DirtyInfo`、`BuiltAt`、`ReleaseURL`。

## License

[Apache-2.0](LICENSE)
