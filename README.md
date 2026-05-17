# Go 语言完整教程

按编号顺序学习，每个 `.go` 文件含详细中文注释，可独立运行。

## 目录说明

| 文件夹 | 难度 | 课程编号 |
|--------|------|----------|
| `chapter01_basics/` | 基础 | 01–20 |
| `chapter02_intermediate/` | 进阶 | 21–33 |
| `chapter03_advanced/` | 高级 | 34–45 |

> 文件夹以 `chapter` + 序号 + 难度名命名，符合 Go import 路径规范（仅 ASCII）。

## 设计约定

| 约定 | 说明 |
|------|------|
| 文件命名 | `01_hello_world.go` … 按全局学习顺序编号 |
| 可独立运行 | `go run ./chapter01_basics/01_hello_world.go` |
| 子目录例外 | `chapter02_intermediate/greet`、`chapter02_intermediate/testing`、`chapter03_advanced/buildtags` |

同一目录内多个 `package main` 文件不能整目录 `go build`，请每次 **`go run` 单个文件**。

---

## 完整目录结构

```
go-tutorial/
├── go.mod
├── README.md
├── chapter01_basics/           # 基础 01–20
│   ├── 01_hello_world.go … 20_closure.go
├── chapter02_intermediate/     # 进阶 21–33
│   ├── 21_methods.go … 33_json_fileio.go
│   ├── greet/greet.go
│   └── testing/calc.go, calc_test.go, errors.go
└── chapter03_advanced/         # 高级 34–45
    ├── 34_goroutine.go … 45_project_layout.go
    └── buildtags/normal.go, debug.go
```

---

## 如何运行

```bash
cd go-tutorial

go run ./chapter01_basics/01_hello_world.go
go run ./chapter02_intermediate/30_packages.go
go test ./chapter02_intermediate/testing/ -v
go run ./chapter03_advanced/34_goroutine.go
go run -tags debug ./chapter03_advanced/44_build_tags.go
```

---

## 学习顺序

按文件编号 **01 → 45** 依次学习：先 `chapter01_basics`，再 `chapter02_intermediate`，最后 `chapter03_advanced`。各文件顶部有「本课目标」与知识点说明。
