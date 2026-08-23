// Copyright 2026 The Cocomhub Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package buildinfo

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"strings"
	"testing"
)

func TestDirtyID_Clean(t *testing.T) {
	i := Info{DirtyInfo: ""}
	if got := i.DirtyID(); got != "clean" {
		t.Fatalf("DirtyID() = %q, want %q", got, "clean")
	}
}

func TestDirtyID_MD5Prefix(t *testing.T) {
	dirty := "xyz"
	want := fmt.Sprintf("%x", md5.Sum([]byte(dirty)))[:10]
	i := Info{DirtyInfo: dirty}
	if got := i.DirtyID(); got != want {
		t.Fatalf("DirtyID() = %q, want %q", got, want)
	}
}

func TestPrintVersion_ContainsKeys(t *testing.T) {
	var buf bytes.Buffer
	i := Info{}
	if err := i.PrintVersion(&buf); err != nil {
		t.Fatalf("PrintVersion() error = %v", err)
	}
	out := buf.String()
	for _, key := range []string{"Version", "CommitID", "DirtyID", "BuiltAt"} {
		if !strings.Contains(out, key) {
			t.Errorf("output missing key %q:\n%s", key, out)
		}
	}
	// 空字段应回落到 default 的默认值
	if !strings.Contains(out, "dev") || !strings.Contains(out, "unknown") {
		t.Errorf("output should contain default fallbacks (dev/unknown):\n%s", out)
	}
}
