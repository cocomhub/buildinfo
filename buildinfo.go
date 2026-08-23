// Copyright 2026 The Cocomhub Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// Package buildinfo 提供可复用的程序二进制版本信息能力。
// dirty_info 不在此内嵌，由调用方通过字段注入。
package buildinfo

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"runtime"
)

// Info 是程序二进制版本信息。
type Info struct {
	Version    string
	Branch     string
	CommitID   string
	DirtyInfo  string
	BuiltAt    string
	ReleaseURL string
	GoVersion  string
	GOOS       string
	GOARCH     string
}

// 构建时由 -X 注入的包级变量。
var (
	Version    string
	Branch     string
	CommitID   string
	DirtyInfo  string
	BuiltAt    string
	ReleaseURL string
)

// Default 返回由包级变量填充的 Info（供 -X 注入使用）。
func Default() Info {
	i := New()
	i.Version = Version
	i.Branch = Branch
	i.CommitID = CommitID
	i.DirtyInfo = DirtyInfo
	i.BuiltAt = BuiltAt
	i.ReleaseURL = ReleaseURL
	return i
}

func New() Info {
	return Info{
		GoVersion: runtime.Version(),
		GOOS:      runtime.GOOS,
		GOARCH:    runtime.GOARCH,
	}
}

// DirtyID 返回 DirtyInfo 的 10 位 md5 摘要；空返回 "clean"。
func (i Info) DirtyID() string {
	if i.DirtyInfo == "" {
		return "clean"
	}
	h := md5.Sum([]byte(i.DirtyInfo))
	return fmt.Sprintf("%x", h)[:10]
}

func defaults(i Info) Info {
	if i.Version == "" {
		i.Version = "dev"
	}
	if i.CommitID == "" {
		i.CommitID = "unknown"
	}
	if i.Branch == "" {
		i.Branch = "unknown"
	}
	if i.BuiltAt == "" {
		i.BuiltAt = "unknown"
	}
	if i.GoVersion == "" {
		i.GoVersion = runtime.Version()
	}
	if i.GOOS == "" {
		i.GOOS = runtime.GOOS
	}
	if i.GOARCH == "" {
		i.GOARCH = runtime.GOARCH
	}
	return i
}

// PrintVersion 输出文本版本信息。
func (i Info) PrintVersion(w io.Writer) error {
	i = defaults(i)
	fmt.Fprintf(w, "Version:    %s\n", i.Version)
	fmt.Fprintf(w, "Branch:     %s\n", i.Branch)
	fmt.Fprintf(w, "DirtyID:    %s\n", i.DirtyID())
	fmt.Fprintf(w, "CommitID:   %s\n", i.CommitID)
	fmt.Fprintf(w, "Runtime:    %s %s/%s\n", i.GoVersion, i.GOOS, i.GOARCH)
	fmt.Fprintf(w, "BuiltAt:    %s\n", i.BuiltAt)
	fmt.Fprintf(w, "ReleaseURL: %s\n", i.ReleaseURL)
	return nil
}

// PrintVersionJSON 输出 JSON 版本信息。
func (i Info) PrintVersionJSON(w io.Writer) error {
	i = defaults(i)
	m := map[string]string{
		"Version":    i.Version,
		"Branch":     i.Branch,
		"DirtyID":    i.DirtyID(),
		"CommitID":   i.CommitID,
		"GoVersion":  i.GoVersion,
		"GOOS":       i.GOOS,
		"GOARCH":     i.GOARCH,
		"BuiltAt":    i.BuiltAt,
		"ReleaseURL": i.ReleaseURL,
	}
	return json.NewEncoder(w).Encode(m)
}
