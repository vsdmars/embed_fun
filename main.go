package main

import (
	"bytes"
	_ "embed"
	"io"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

//go:embed melodyx.mp3
var mp3Bytes []byte

func main() {
	f := io.NopCloser(bytes.NewReader(mp3Bytes))
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		println("shit, open file error!")
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
