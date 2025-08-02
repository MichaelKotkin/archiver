package cmd

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const packedExtension = "vls"

var ErrEmptyPath = errors.New("path to file is not specified or empty")

var vlsCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack files using varaible-length code",
	Run:   pack,
}

func pack(_ *cobra.Command, args []string) {

	if len(args) == 0 || args[0] == "" {
		HandlerErr(ErrEmptyPath)
	}

	filePath := args[0]

	r, err := os.Open(filePath)

	if err != nil {
		HandlerErr(err)
	}

	data, err := io.ReadAll(r)
	if err != nil {
		HandlerErr(err)
	}

	packed := "" + string(data)

	err = os.WriteFile(packedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		HandlerErr(err)
	}
}

func packedFileName(path string) string {
	fileName := filepath.Base(path)
	ext := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(fileName, ext)
	return baseName + "." + packedExtension
}

func init() {
	packCmd.AddCommand(vlsCmd)
}
