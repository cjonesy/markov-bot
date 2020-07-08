package cli

import (
	"log"

	"github.com/cjonesy/markov-bot/pkg/markov-bot/bot"
	"github.com/cjonesy/markov-bot/pkg/markov-bot/markov"
	"github.com/spf13/cobra"
)

var (
	token        string
	corpusPath   string
	prefixLength int
	maxResponse  int
)

func init() {
	startCmd.Flags().StringVarP(
		&token,
		"token",
		"t",
		"",
		"slack bot token",
	)
	startCmd.MarkFlagRequired("token")
	startCmd.Flags().StringVarP(
		&corpusPath,
		"corpus-path",
		"c",
		"",
		"the full path to the corpus file",
	)
	startCmd.MarkFlagRequired("corpus-path")
	startCmd.Flags().IntVarP(
		&prefixLength,
		"prefix-length",
		"p",
		3,
		"the length of the prefix for keys in the markov chain, default: 3",
	)
	startCmd.Flags().IntVarP(
		&maxResponse,
		"max-response",
		"r",
		50,
		"the max number of the words the bot will respond with, default: 50",
	)
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the bot",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func start() {
	var exit = make(chan bool)
	log.Print("Starting Bot")
	chain := markov.NewChain(prefixLength)

	// Loading the Corpus can take a long time so we do it in the background
	go chain.Load(corpusPath)

	b, _ := bot.New(token, chain, maxResponse)
	go b.Start()

	// Block until something kills the process
	<-exit
}
