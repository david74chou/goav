package main

import (
	"github.com/david74chou/goav/avcodec"
	"github.com/david74chou/goav/avdevice"
	"github.com/david74chou/goav/avfilter"
	"github.com/david74chou/goav/avformat"
	"github.com/david74chou/goav/avutil"
	"github.com/david74chou/goav/swresample"
	"github.com/david74chou/goav/swscale"
	"log"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
