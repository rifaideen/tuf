package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func main() {
	n := flag.Int("n", 3, "number of rows")
	flag.Parse()

	Pattern("Rectangular Star", RectangularStarPattern, n)

	Pattern("Right Angled Triangle Star", RightAngledTriangle, n)

	Pattern("Right Angled Number Pyramid - I", RightAngledNumberPyramid, n)
	Pattern("Right Angled Number Pyramid - II", RightAngledNumberPyramid2, n)

	Pattern("Inverted Right Pyramid", InvertedRightPyramid, n)

	Pattern("Inverted Right Numbered Pyramid", InvertedNumberedRightPyramid, n)

	Pattern("Pyramid Star", Pyramid, n)

	Pattern("Inverted Star Pyramid", InvertedStarPyramid, n)

	Pattern("Diamond Star", DiamondStar, n)
	Pattern("Diamond Numbered - Incremental", DiamondNumberedIncremental, n)
	Pattern("Diamond Numbered - Constant", DiamondNumberedConstant, n)
	Pattern("Diamond Star - Combined with Existing", DiamondStarCombined, n)

	Pattern("Half Diamond Star", HalfDiamondStar, n)

	Pattern("Binary Number Triangle", BinaryNumberTriangle, n)

	Pattern("Number Crown", NumberCrown, n)
}

func DiamondStarOld(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Printf(" ")
		}

		for j := 0; j < 2*i+1; j++ {
			fmt.Printf("*")
		}

		for j := 0; j < n-i-1; j++ {
			fmt.Printf(" ")
		}

		fmt.Println()
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf(" ")
		}

		for j := 0; j < 2*i+1; j++ {
			fmt.Printf("*")
		}

		for j := 0; j < n-i-1; j++ {
			fmt.Printf(" ")
		}

		fmt.Println()
	}
}

func PyramidOld(n int) {
	fmt.Println("Pyramid Pattern:")

	a := 1
	diff := 2

	// 4, 3, 2
	for i := 0; i < n; i++ {
		space := n - i - 1
		fmt.Print(strings.Repeat(" ", space))

		stars := a + (diff * i)

		fmt.Print(strings.Repeat("*", stars))

		fmt.Println()
	}
}

func InvertedPyramidOld(n int) {
	fmt.Println("Inverted Pyramid Pattern:")
	a := 1
	space := 0
	diff := 2

	for i := n; i >= 0; i-- {
		ap := a + (i * diff)

		fmt.Print(strings.Repeat(" ", space))

		for j := 0; j < ap; j++ {
			// if j == 0 {
			// }

			fmt.Print("*")
		}

		fmt.Println()

		space++
	}
}

func Pattern(name string, fn func(n int), n *int) {
	black := color.New(color.FgMagenta, color.Bold)
	black.Printf("%s Pattern:\n\n", name)

	fn(*n)

	fmt.Println()
}
