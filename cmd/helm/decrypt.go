/*
Copyright The Helm Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"io"
    "os"
    "io/ioutil"
    "log"

	"github.com/spf13/cobra"
	"k8s.io/helm/pkg/helm"
    "k8s.io/helm/pkg/chartutil"
)

const decryptDesc = `
    Decrypt string
`

type decryptCmd struct {
	out        io.Writer
	client     helm.Interface
	decrypt_filepath   string
}

func newDecryptCmd(c helm.Interface, out io.Writer) *cobra.Command {
	decrypt_data := &decryptCmd{
		client: c,
		out:    out,
	}

	cmd := &cobra.Command{
		Use:   "decrypt",
		Short: "decrypt string",
		Long:  decryptDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
            decrypt_data.decrypt_filepath = args[0]
			return decrypt_data.run()
		},
	}
	f := cmd.Flags()
	settings.AddFlagsTLS(f)

	// set defaults from environment
	settings.InitTLS(f)

	return cmd
}

func (e *decryptCmd) run() error {
    decrypt_content := readDecryptFile(e.decrypt_filepath)
    decrypt_content = chartutil.String_trim(decrypt_content)
    decrypt_out, err := chartutil.Decrypt(settings.PasswordFile, decrypt_content)
    return err
}

func readDecryptFile(filepath string) string{
    file, err := os.Open(filepath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    b, err := ioutil.ReadAll(file)

    return string(b)
}
