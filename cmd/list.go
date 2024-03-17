// ------------------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 bomctl authors
// SPDX-FileName: cmd/list.go
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// ------------------------------------------------------------------------
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// ------------------------------------------------------------------------
package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"

	"github.com/bomctl/bomctl/internal/pkg/db"
	"github.com/bomctl/bomctl/internal/pkg/utils"
)

func listCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List SBOM documents in local cache",
		Long:    "List SBOM documents in local cache",
		Run: func(cmd *cobra.Command, args []string) {
			logger = utils.NewLogger("list")

			documents, err := db.GetAllDocuments()
			if err != nil {
				logger.Error(err)
				os.Exit(1)
			}

			headers := []string{"ID", "Alias", "Name", "Version", "# Nodes"}
			rows := [][]string{}

			for _, document := range documents {
				rows = append(rows, []string{
					document.Metadata.Id,
					"",
					document.Metadata.Name,
					document.Metadata.Version,
					fmt.Sprint(len(document.NodeList.Nodes)),
				})
			}

			fmt.Printf("\n%s\n\n", table.New().
				Headers(headers...).
				Rows(rows...).
				Width(80).
				BorderTop(false).
				BorderBottom(false).
				BorderLeft(false).
				BorderRight(false).
				BorderHeader(true).
				String(),
			)
		},
	}

	return listCmd
}
