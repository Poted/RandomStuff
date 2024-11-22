package main

import (
	"fmt"
)

func NewDiceThrows() {

	sentences := Sentences{
		[]string{"raz", "dwa", "trzy"},

		[]string{"one", "two", "three", "four", "five"},

		[]string{"a", "b", "c", "d",
			"e", "f", "g", "h", "i"},
	}

	sentences.Translate(
		translations{
			func(sentences []string, polski int) string {
				return fmt.Sprintf("1st: %v and %d", sentences, polski)
			},
			func(sentences []string, w int) string {
				return "[]int{}"
			},
			func(sentences []string, dieSides int) string {
				return fmt.Sprintf("%d", dieSides+20)
			},
		},
	)

}

type Sentences [][]string

type translations []func([]string, int) string

func (t *Sentences) Translate(dest translations) {

	xd := *t
	var a = 420

	for _, v := range xd {

		for _, d := range dest {
			arr := d(v, a)
			fmt.Printf("arr: %v\n", arr)
		}

	}

}
