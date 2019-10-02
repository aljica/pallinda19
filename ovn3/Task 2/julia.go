// Stefan Nilsson 2013-02-27

// This program creates pictures of Julia sets (en.wikipedia.org/wiki/Julia_set).
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type ComplexFunc func(complex128) complex128

var Funcs []ComplexFunc = []ComplexFunc{
	func(z complex128) complex128 { return z*z - 0.61803398875 },
	func(z complex128) complex128 { return z*z + complex(0, 1) },
	func(z complex128) complex128 { return z*z + complex(-0.835, -0.2321) },
	func(z complex128) complex128 { return z*z + complex(0.45, 0.1428) },
	func(z complex128) complex128 { return z*z*z + 0.400 },
	func(z complex128) complex128 { return cmplx.Exp(z*z*z) - 0.621 },
	func(z complex128) complex128 { return (z*z+z)/cmplx.Log(z) + complex(0.268, 0.060) },
	func(z complex128) complex128 { return cmplx.Sqrt(cmplx.Sinh(z*z)) + complex(0.065, 0.122) },
}

func main() {
	before := time.Now()
	wg := new(sync.WaitGroup)
	for n, fn := range Funcs {
		wg.Add(1)
		go func(n int, fn ComplexFunc) {
			err := CreatePng("picture-"+strconv.Itoa(n)+".png", fn, 1024)
			if err != nil {
				log.Fatal(err)
			}
			wg.Done()
		}(n, fn)
		/*err := CreatePng("picture-"+strconv.Itoa(n)+".png", fn, 1024)
		if err != nil {
			log.Fatal(err)
		}*/
	}
	wg.Wait()
	fmt.Println("time:", time.Now().Sub(before))
	// Time without parallel version: time: 20.68751716s
	// Time with parallel && 8 CPUs: time: 14.639395088s

	// Time with parallel version && 8 CPUs && Julia()s calls to
	// Iterate() have also been parallellized by up to 10 000 goroutines:
	// time: 4.279291974s
}

// CreatePng creates a PNG picture file with a Julia image of size n x n.
func CreatePng(filename string, f ComplexFunc, n int) (err error) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()
	err = png.Encode(file, Julia(f, n))
	return
}

// Julia returns an image of size n x n of the Julia set for f.
func Julia(f ComplexFunc, n int) image.Image {
	bounds := image.Rect(-n/2, -n/2, n/2, n/2)
	img := image.NewRGBA(bounds)
	s := float64(n / 4)
	allowedWorkUnits := 10000
	activeWorkUnits := 0
	for i := bounds.Min.X; i < bounds.Max.X; i++ {
		for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
			if activeWorkUnits < allowedWorkUnits {
				activeWorkUnits += 1
				go func(i, j int) {
					n := Iterate(f, complex(float64(i)/s, float64(j)/s), 256)
					r := uint8(0)
					g := uint8(0)
					b := uint8(n % 32 * 8)
					img.Set(i, j, color.RGBA{r, g, b, 255})
					activeWorkUnits -= 1
				}(i, j)
			} else {
				n := Iterate(f, complex(float64(i)/s, float64(j)/s), 256)
				r := uint8(0)
				g := uint8(0)
				b := uint8(n % 32 * 8)
				img.Set(i, j, color.RGBA{r, g, b, 255})
			}
		}
	}
	return img
}

// Iterate sets z_0 = z, and repeatedly computes z_n = f(z_{n-1}), n â‰¥ 1,
// until |z_n| > 2  or n = max and returns this n.
func Iterate(f ComplexFunc, z complex128, max int) (n int) {
	for ; n < max; n++ {
		if real(z)*real(z)+imag(z)*imag(z) > 4 {
			break
		}
		z = f(z)
	}
	return
}

func init() {
	numcpu := runtime.NumCPU()
	fmt.Println("Num CPU:", numcpu)
	runtime.GOMAXPROCS(numcpu) // Try to use all available CPUs.
}
