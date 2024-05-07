package main

import (
    "strconv"
    "fmt"
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    "os"
    "net/http"
    "log"
)

var palette = []color.Color{
    color.RGBA{0x00, 0xff, 0x00, 0xff}, 
    color.RGBA{0xff, 0x00, 0x00, 0xff},
    color.RGBA{0x00, 0x00, 0xff, 0xff},
    color.White, 
    color.Black}

const (
    whiteIndex = 0
    blackIndex = 1
)

const (
    cycles = 5
    res = 0.001
    size = 300
    nframes = 64
    delay = 8
)


func main() {
    if (len(os.Args) > 1 && os.Args[1] == "web") {
        http.HandleFunc("/", displayGif)
        log.Fatal(http.ListenAndServe("localhost:8000", nil))
        return
    }
    lissajous(os.Stdout)
}

func setCycles(w http.ResponseWriter, r *http.Request) {
    var newCycles string
    var result int
    var err error
    newCycles = r.URL.Query().Get("cycles")
    result, err = strconv.Atoi(newCycles)
    fmt.Fprintf(w, "cycles set to %q\n", newCycles)
}

func displayGif(w http.ResponseWriter, r *http.Request) {
    lissajous(w)
}

func lissajous(out io.Writer) {
    freq := rand.Float64() * 3.0
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.02
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < cycles*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 3)
        }
        phase += 1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)
}
