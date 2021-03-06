package pg

import (
	"github.com/spf13/cobra"
	"github.com/wal-g/tracelog"
	"github.com/wal-g/wal-g/internal"
	"github.com/wal-g/wal-g/internal/asm"
)

const walReceiveShortDescription = "Receive WAL stream with postgres Streaming Replication Protocol and push to storage"

// walReceiveCmd represents the walReceive command
var walReceiveCmd = &cobra.Command{
	Use:   "wal-receive",
	Short: walReceiveShortDescription,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		uploader, err := internal.ConfigureWalUploader()
		tracelog.ErrorLogger.FatalOnError(err)

		archiveStatusManager, err := internal.ConfigureArchiveStatusManager()
		if err == nil {
			uploader.ArchiveStatusManager = asm.NewDataFolderASM(archiveStatusManager)
		} else {
			tracelog.ErrorLogger.PrintError(err)
			uploader.ArchiveStatusManager = asm.NewNopASM()
		}
		internal.HandleWALReceive(uploader)
	},
}

func init() {
	cmd.AddCommand(walReceiveCmd)
}
