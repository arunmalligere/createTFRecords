# Description

Given a directory that contains pairs of images encoded as JPG and a YAML file containing labeling data, the program iterates over this directory, reads the image and label data of each pair and creates a TFRecord out of it.

# Create tfrecords

### Command help

```
> go run main.go createTFR -h
Create TF records ....
NAME:
   main createTFR - 

USAGE:
   main createTFR [command options] [arguments...]

OPTIONS:
   --inDir value      Absolute path of the Input Directory that contains Images and Labels
   --outDir value     Absolute path of the Output Directory for the TFRecord
   --batchSize value  Number of files per batch (default: 25 files)
   --help, -h         show help (default: false)
```

### Command format

```
> go run main.go createTFR --inDir <absolute path="" of="" the="" input="" directory="">--outDir <absolute path="" of="" the="" output="" directory="">--batchSize</absolute></absolute>
```

### Sample Result

```
> ./createTFRecords createTFR --inDir /home/inputImagesLabels --outDir /home/tfrOutput --batchSize 100
Create TF records ....
Start processing.....
Finished processing.....
Output file::  /home/tfrOutput/output_20210906204603.tfrecord
Number of files processed::  3202
2021/09/06 20:46:04 TFRecord creation took 1.776270685s
```

# Read Image from tfrecord

### Command help

USAGE:

```
> go run main.go readImage -h
Create TF records ....
NAME:
   createTFRecords readImage - Read image data from TFRecords

USAGE:
   createTFRecords readImage [command options] [arguments...]

OPTIONS:
   --inFile value     Absolute path of the input tfRecord
   --outDir value     Absolute path of the Output Directory
   --numOfRecs value  Num records to read (default:  2 records)
   --help, -h         show help (default: false)
```

### Command format

```
> go run main.go readImage --inFile <Absolute path of the input tfRecord> --outDir <Absolute path of the Output Directory> --numOfRecs <Num records to read>
```

### Sample result

```
> /createTFRecords readImage --inFile /home/tfrOutput/output_20210906204603.tfrecord --outDir /home/readOutput --numOfRecs 15
Create TF records ....
/home/readOutput/002025e55d5990edb85ff94a8d75a7f67cb7694ee2dbd2edcfcc0e02391f4faa.jpg
/home/readOutput/079ef518af25d053a1c0025675df1283bab70eda6bcfa35eb65ef0dd030384f0.jpg
/home/readOutput/008dee526472e7408b34efe5e54cb6408afb99133406cbe2ae9bca2fe5a7ca3d.jpg
/home/readOutput/003d9e774f3833b532222e1682c2b628acbfbbc223567d9c69d2ed696484aebe.jpg
/home/readOutput/0099e6c5f9ea32fe82b1a14f8cac0674f20091b7c222ca0cca215594dc433614.jpg
/home/readOutput/03d9ae66f761322d1e07cdf63027840b829f2850a6bec4b8fc2b0665a98b7084.jpg
/home/readOutput/008a717315eb91c6f3adbd08a24d2b7361d2b0d61365f19c3d3992154062ec25.jpg
/home/readOutput/079b1f91020be1164d5cab82590ba430eb9e038669c0f658ba2989853a9758d0.jpg
/home/readOutput/04842872878131c596ad616bb0ca776ffeed34a1346e81bb5a59fe955cf7ff71.jpg
/home/readOutput/05903d19305c4bdeebde02ba06d34c9f8c1586e287b7c704693cb31a91904fca.jpg
/home/readOutput/069095b42873667158d1cda7250245ab60a93198e399745df9b7a06a11d82c14.jpg
/home/readOutput/076c71720b7248ad8ecd2a84c4e0fe96ee19729d54cf8278def60c79bf7627ad.jpg
/home/readOutput/0154d991c586e7e6814b90bc92298748a77f26a61aff88ceb87a166a4b040a42.jpg
/home/readOutput/06986c2f7090217ff80170c015031fa04383075a06d27ed77cee4a2905e9d61c.jpg
/home/readOutput/026281e09a13097c2fe3acf848e88a79e5e040ba00b31991bb214cf801744c51.jpg
```