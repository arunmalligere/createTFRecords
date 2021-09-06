package createTFR

import (
	"fmt"
	_ "image/jpeg"
	"io/fs"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/recogni/tfutils"
	"gopkg.in/yaml.v2"
)

type Label struct {
	Version   int    `yaml:"version"`
	ID        string `yaml:"id"`
	Imagesize struct {
		W int `yaml:"w"`
		H int `yaml:"h"`
	} `yaml:"imagesize"`
	Boxes []struct {
		Type int `yaml:"type"`
		Rect struct {
			X1 float32 `yaml:"x1"`
			Y1 float32 `yaml:"y1"`
			X2 float32 `yaml:"x2"`
			Y2 float32 `yaml:"y2"`
		} `yaml:"rect"`
	} `yaml:"boxes"`
}

func Create(inDir, outDir string, maxBatchSize int) error {
	start := time.Now()
	if len(outDir) == 0 {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		outDir = cwd
	} else {
		err := os.MkdirAll(outDir, 0755)
		if err != nil {
			return err
		}
	}

	// Number of files processed per batch
	if maxBatchSize == 0 {
		maxBatchSize = 25
	}
	files, _ := ioutil.ReadDir(inDir)

	numOfFilesProcessed := 0
	numOfFiles := len(files)

	numOfBatches := int(math.Ceil(float64(numOfFiles / maxBatchSize)))
	outFile := filepath.Join(outDir, fmt.Sprintf("output_%s.tfrecord", time.Now().Format("20060102150405")))

	w, err := tfutils.NewWriter(outFile, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println("Start processing.....")

	// Creating too many goroutines can lead to resource crunch
	// Process the set of files in batches
	for i := 0; i <= numOfBatches; i++ {
		lowerBound := numOfFilesProcessed
		upperBound := numOfFilesProcessed + maxBatchSize

		if upperBound > numOfFiles {
			upperBound = numOfFiles
		}

		batchFiles := files[lowerBound:upperBound]
		numOfFilesProcessed += maxBatchSize

		processingErrorChan := make(chan error)
		processingDoneChan := make(chan int)
		processingErrors := make([]error, 0)

		go func() {
			for {
				select {
				case err := <-processingErrorChan:
					processingErrors = append(processingErrors, err)
				case <-processingDoneChan:
					close(processingErrorChan)
					close(processingDoneChan)
					return
				}
			}
		}()

		var fileProcessingWG sync.WaitGroup
		var mu sync.Mutex
		fileProcessingWG.Add(len(batchFiles))

		for idx, file := range batchFiles {
			go func(file fs.FileInfo, idx int) {
				defer fileProcessingWG.Done()

				inFilePath := filepath.Join(inDir, file.Name())

				if filepath.Ext(file.Name()) == ".yaml" {
					var label Label
					yamlFile, err := ioutil.ReadFile(inFilePath)
					if err != nil {
						processingErrorChan <- err
					}

					err = yaml.Unmarshal(yamlFile, &label)
					if err != nil {
						processingErrorChan <- err
					}

					// Used this hack to open the JPEG file while processing the yaml
					// file because the file names of the images and the labels have to be the same.
					// Also, took the liberty to assume the extension will only be
					// ".jpg" and not ".jpeg". A simple tweak in the logic is needed
					// incase the directory contains but ".jpg" and ".jpeg" extensions
					jpegImage := filepath.Join(inDir, label.ID+".jpg")
					jpegImageBs, err := ioutil.ReadFile(jpegImage)
					if err != nil {
						processingErrorChan <- err
					}

					mu.Lock()
					// Step1:: Create Feature Map
					fm := createTFRecordFeatureMap(label, jpegImageBs)

					// Step2:: Get the feature from the map created in Step1
					tfFeature, err := tfutils.GetFeaturesFromMap(fm)
					if err != nil {
						processingErrorChan <- err
					}

					// Step3:: Get the tfrecord string for the feature from Step2
					tfrString, err := tfutils.GetTFRecordStringForFeatures(tfFeature)
					if err != nil {
						processingErrorChan <- err
					}

					// Step4:: Write the record to a consolidated file
					err = w.WriteRecord(tfrString)
					if err != nil {
						processingErrorChan <- err
					}
					mu.Unlock()
				}
			}(file, idx)
		}

		fileProcessingWG.Wait()
		processingDoneChan <- 1

		for _, err := range processingErrors {
			log.Println("File processing errors", err)
		}
	}
	fmt.Println("Finished processing.....")
	fmt.Println("Output file:: ", outFile)

	fmt.Println("Number of files processed:: ", numOfFiles)
	elapsed := time.Since(start)
    log.Printf("TFRecord creation took %s", elapsed)
	
	return nil
}

func createTFRecordFeatureMap(label Label, imgData []byte) (fsm map[string]interface{}) {

	// Create new tf record.
	fsm = make(map[string]interface{}, 11)
	fsm["image/filename"] = label.ID + ".jpg"
	fsm["image/source_id"] = label.ID
	fsm["image/format"] = "jpeg"
	fsm["image/width"] = label.Imagesize.W
	fsm["image/height"] = label.Imagesize.H
	fsm["image/encoded"] = imgData

	xmins := make([]float32, len(label.Boxes))
	xmaxs := make([]float32, len(label.Boxes))
	ymins := make([]float32, len(label.Boxes))
	ymaxs := make([]float32, len(label.Boxes))
	labels := make([]int64, len(label.Boxes))

	for i, b := range label.Boxes {
		xmins[i] = b.Rect.X1
		ymins[i] = b.Rect.Y1
		xmaxs[i] = b.Rect.X2
		ymaxs[i] = b.Rect.Y2
		labels[i] = int64(b.Type)
	}

	fsm["image/object/bbox/xmin"] = xmins
	fsm["image/object/bbox/xmax"] = xmaxs
	fsm["image/object/bbox/ymin"] = ymins
	fsm["image/object/bbox/ymax"] = ymaxs
	fsm["image/object/class/label"] = labels
	return fsm
}
