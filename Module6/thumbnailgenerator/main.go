package main

import (
	"io/ioutil"
	"log"

	"github.com/disintegration/imaging"
)

func main() {
	var srcfiles []string
	files, err := ioutil.ReadDir("./img/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		srcfiles = append(srcfiles, f.Name())
	}
	makeThumbs(srcfiles)

}

func makeThumbs(files []string) {
	ch := make(chan struct{})
	for _, f := range files {
		go func(fname string) {
			image_file, err := imaging.Open("./img/" + fname)
			if err != nil {
				log.Fatalf("Failed to open image: %v", err)
			}
			dstImage := imaging.Resize(image_file, 128, 128, imaging.Lanczos)

			//Save the resulting image as JPEG
			err = imaging.Save(dstImage, "./thumbs/"+fname+"_tmb.jpeg")
			if err != nil {
				log.Fatalf("Failed to save image: %v", err)
			}

			ch <- struct{}{}
		}(f)
	}
	//waiting while all goroutines will be finished
	for range files {
		<-ch
	}
}
