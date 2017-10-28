package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

type test struct {
	sourceWords []string
}

func NewTest() *test {
	t := new(test)
	t.sourceWords = getWords()
	return t
}

func main() {
	test := NewTest()
	test.run()
}

func (t *test) run() {
	fmt.Println("Typing start!!! 60sec")
	correct := 0
	for index, word := range t.sourceWords {
		fmt.Println("")
		fmt.Printf("No.%v Please input: %v\n>> ", index+1, word)
		start := time.Now().UnixNano()
		var input string
		fmt.Scan(&input)
		end := time.Now().UnixNano()
		if input == word {
			correct++
			fmt.Print("correct! ")
		} else {
			fmt.Println("wrong... ")
		}
		t := fmt.Sprint(end - start)
		integer := t[:len(t)-9]
		fractional := t[len(t)-9 : len(t)-7]

		fmt.Printf("%v.%v sec\n", integer, fractional)
	}
	fmt.Printf("Result: %v/%v\n", correct, len(t.sourceWords))
}

func getWords() []string {
	num := getNumber()
	cnt, _ := countLine()
	lines := getRandomLines(cnt, num)
	var words []string
	file, _ := os.Open("/usr/share/dict/words")
	defer file.Close()
	sc := bufio.NewScanner(file)
	for i := 1; sc.Scan(); i++ {
		if len(lines) == 0 {
			break
		}
		if i == lines[0] {
			words = append(words, sc.Text())
			// delete first element
			lines = lines[1:len(lines)]
		}
	}
	return words
}

func getNumber() int {
	fmt.Print("Number of words (default: 10) : ")
	var num string
	sc := bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		num = sc.Text()
	}
	n, err := strconv.Atoi(num)
	if err != nil {
		return 10
	}
	return n
}

func countLine() (int, error) {
	file, _ := os.Open("/usr/share/dict/words")
	defer file.Close()
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}
	for {
		c, err := file.Read(buf)
		count += bytes.Count(buf[:c], lineSep)
		switch {
		case err == io.EOF:
			return count, nil
		case err != nil:
			return count, err
		}
	}
}

func getRandomLines(cnt int, num int) []int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	var lines []int
	for i := 0; i < num; i++ {
		lines = append(lines, r.Intn(cnt-2)+1)
	}
	sort.Ints(lines)
	return lines
}
