package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	song := "Beauty queen of only eighteen, she had some trouble with herself.\nHe was always there to help her, she always belonged to someone else.\nI drove for miles and miles and wound up at your door.\nI've had you so many times but somehow I want more.\nI don't mind spending everyday out on your corner in the pouring rain.\nLook for a girl with a broken smile, ask her if she wants to stay a while and she will be loved."
	scanner := bufio.NewScanner(strings.NewReader(song))
	//scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
