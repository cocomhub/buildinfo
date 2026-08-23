// Copyright 2026 The Cocomhub Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package buildinfo

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestNewVersionCmd_PrintsVersion(t *testing.T) {
	var buf bytes.Buffer
	cmd := NewVersionCmd(Info{Version: "1.0.0"})
	cmd.SetOut(&buf)
	cmd.SetArgs([]string{})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
	if !strings.Contains(buf.String(), "1.0.0") {
		t.Errorf("output missing Version %q:\\n%s", "1.0.0", buf.String())
	}
}

func TestNewVersionCmd_JSON(t *testing.T) {
	var buf bytes.Buffer
	cmd := NewVersionCmd(Info{Version: "1.0.0", CommitID: "abc123"})
	cmd.SetOut(&buf)
	cmd.SetArgs([]string{"--json"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
	var m map[string]string
	if err := json.Unmarshal(buf.Bytes(), &m); err != nil {
		t.Fatalf("output not valid JSON: %v\\n%s", err, buf.String())
	}
	if m["Version"] != "1.0.0" {
		t.Errorf(`m["Version"] = %q, want "1.0.0"`, m["Version"])
	}
}
