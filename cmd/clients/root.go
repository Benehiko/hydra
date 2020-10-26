/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package clients

import (
	"github.com/ory/hydra/cmd/cli"
	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var packageFlags = pflag.NewFlagSet("clients package flags", pflag.ContinueOnError)

// clientsCmd represents the clients command
var clientsCmd = &cobra.Command{
	Use:   "clients <command>",
	Short: "Manage OAuth 2.0 Clients",
}

func init() {
	cmdx.RegisterFormatFlags(packageFlags)
	cli.RegisterConnectionFlags(packageFlags)
	cli.RegisterFakeTLSTermination(packageFlags)

	clientsCmd.PersistentFlags().AddFlagSet(packageFlags)
}

func RegisterCommandRecursive(parent *cobra.Command) {
	parent.AddCommand(clientsCmd)

	clientsCmd.AddCommand(newCreateCmd())
	clientsCmd.AddCommand(newDeleteCmd())
	clientsCmd.AddCommand(newGetCmd())
	clientsCmd.AddCommand(newImportCmd())
	clientsCmd.AddCommand(newListCmd())
	clientsCmd.AddCommand(newUpdateCmd())
	clientsCmd.AddCommand(newValidateCmd())
}