package main

import (
	"fmt"
	"log"
	"os"

	"github.com/arunmalligere/createTFRecords/createTFR"
	"github.com/arunmalligere/createTFRecords/readImage"
	"github.com/urfave/cli/v2"
)

func main() {
	fmt.Println("Create TF records ....")
	app := cli.NewApp()

	app.Commands = []*cli.Command{
		{
			Name: "createTFR",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "inDir", Usage: "Absolute path of the Input Directory that contains Images and Labels"},
				&cli.StringFlag{Name: "outDir", Usage: "Absolute path of the Output Directory for the TFRecord"},
				&cli.IntFlag{Name: "batchSize", Usage: "Number of files per batch", DefaultText: "25 files"},
			},
			Action: func(ctx *cli.Context) error {
				err := createTFR.Create(ctx.String("inDir"), ctx.String("outDir"), ctx.Int("batchSize"))
				if err != nil {
					log.Fatal(err)
					return cli.Exit(err, 1)
				}

				return nil
			},
		},
		{
			Name:  "readImage",
			Usage: "Read image data from TFRecords",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "inFile", Usage: "Absolute path of the input tfRecord"},
				&cli.StringFlag{Name: "outDir", Usage: "Absolute path of the Output Directory"},
				&cli.IntFlag{Name: "numOfRecs", Usage: "Num records to read", DefaultText: " 2 records"},
			},
			Action: func(ctx *cli.Context) error {
				err := readImage.Read(ctx.String("inFile"), ctx.String("outDir"), ctx.Int("numOfRecs"))
				if err != nil {
					log.Fatal(err)
					return cli.Exit(err, 1)
				}

				return nil
			},
		},
	}
	app.Run(os.Args)
}
