package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Problem interface {
	Question() string
	Answer() int
	IsCorrect(int) bool
}

type problem struct {
	a      int
	b      int
	answer int
}

type QuizResult struct {
	Problems       int
	CorrectAnswers int
}

func NewProblem(a, b int) *problem {
	return &problem{
		a:      a,
		b:      b,
		answer: a + b,
	}
}

func NewRandomeProblem() *problem {
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(100)
	b := rand.Intn(100)
	return NewProblem(a, b)
}

func (p *problem) Question() string {
	left := strconv.Itoa(p.a)
	if p.a < 0 {
		left = "(" + left + ")"
	}
	right := strconv.Itoa(p.b)
	if p.b < 0 {
		right = "(" + right + ")"
	}
	return left + "+" + right
}

func (p *problem) Answer() string {
	return strconv.Itoa(p.answer)
}

func (p *problem) IsCorrect(answer int) bool {
	return p.answer == answer
}

func main() {
	ctx, cansel := context.WithTimeout(context.Background(), time.Second*10)
	defer cansel()

	result := new(QuizResult)
	Quiz(ctx, result)
	fmt.Println("")
	fmt.Printf("CorrectAnswers/Problems: %d/%d\n", result.CorrectAnswers, result.Problems)
}

func Quiz(ctx context.Context, result *QuizResult) {
	for i := 0; i < 100; i++ {
		p := NewRandomeProblem()
		fmt.Printf("Q-%02d: %s = ", i+1, p.Question())
		result.Problems++

		ansChan := make(chan int)
		go func() {
			var ans int
			fmt.Scanln(&ans)
			ansChan <- ans
		}()

		select {
		case <-ctx.Done():
			return
		case ans := <-ansChan:
			if p.IsCorrect(ans) {
				result.CorrectAnswers++
				fmt.Println("Correct!")
			} else {
				fmt.Println("Wrong!")
			}
		}
	}
}
