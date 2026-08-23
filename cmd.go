// Copyright 2026 The Cocomhub Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package buildinfo

import (
	"github.com/spf13/cobra"
)

// NewVersionCmd 创建 version 命令，输出程序版本信息。
func NewVersionCmd(i Info) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "显示程序版本信息",
		RunE: func(cmd *cobra.Command, args []string) error {
			if json, _ := cmd.Flags().GetBool("json"); json {
				return i.PrintVersionJSON(cmd.OutOrStdout())
			}
			return i.PrintVersion(cmd.OutOrStdout())
		},
	}
	cmd.Flags().Bool("json", false, "以 JSON 格式输出")
	return cmd
}
