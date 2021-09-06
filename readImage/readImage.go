package readImage

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/recogni/tfutils"
)

func Read(inFile, outDir string, numOfRecs int) error {
	if len(inFile) == 0 {
		return errors.New("please provide an input file")
	}

	err := os.MkdirAll(outDir, 0755)
	if err != nil {
		return err
	}

	r, err := tfutils.NewReader([]string{inFile}, nil)
	if err != nil {
		return err
	}

	records := 0
	if numOfRecs == 0 {
		numOfRecs = 2
	}
	for {
		rbs, err := r.ReadRecord()
		if err == io.EOF {
			return err
		}

		rec, err := tfutils.GetFeatureMapFromTFRecord(rbs)
		if err != nil {
			return err
		}

		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		outFile := filepath.Join(cwd, "test.jpg")
		imgByte := []byte{}
		

		for k, v := range rec.Feature {
			if k == "image/filename" {
				outFile = filepath.Join(outDir, string(v.GetBytesList().Value[0]))
			}
			if k == "image/encoded" {
				imgByte = v.GetBytesList().Value[0]
			}
		}

		img, _, err := image.Decode(bytes.NewReader(imgByte))
		if err != nil {
			log.Fatalln(err)
		}

		out, _ := os.Create(outFile)
		defer out.Close()

		var opts jpeg.Options
		opts.Quality = 1

		err = jpeg.Encode(out, img, &opts)
		if err != nil {
			log.Println(err)
		}

		fmt.Println(outFile)

		records++

		if records == numOfRecs {
			break
		}
	}
	return nil
}
