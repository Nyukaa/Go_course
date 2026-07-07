// Optional values with nil interfaces Use nil pointers to represent optional values:
package main

import "fmt"
type Config struct {
    Timeout *int  // nil means "use default"
    Debug   *bool // nil means "use default"
}

func NewConfig() Config {
    return Config{}
}

func (c *Config) SetTimeout(seconds int) {
    c.Timeout = &seconds
}

func (c *Config) GetTimeout() int {
    if c.Timeout == nil {
        return 30  // default
    }
    return *c.Timeout
}

func main() {
    config := NewConfig()
    fmt.Println(config.GetTimeout())  // 30 (default)

    config.SetTimeout(60)
    fmt.Println(config.GetTimeout())  // 60
}

//Strategy pattern Use interfaces to swap behavior at runtime:

// type SortStrategy interface {
//     Sort([]int) []int
// }

// type BubbleSort struct{}

// func (b BubbleSort) Sort(data []int) []int {
//     // Bubble sort implementation
//     return data
// }

// type QuickSort struct{}

// func (q QuickSort) Sort(data []int) []int {
//     // Quick sort implementation
//     return data
// }

// type Sorter struct {
//     strategy SortStrategy
// }

// func (s *Sorter) SetStrategy(strategy SortStrategy) {
//     s.strategy = strategy
// }

// func (s *Sorter) Sort(data []int) []int {
//     return s.strategy.Sort(data)
// }

// func main() {
//     data := []int{5, 2, 8, 1, 9}

//     sorter := Sorter{}

//     sorter.SetStrategy(BubbleSort{})
//     result1 := sorter.Sort(data)

//     sorter.SetStrategy(QuickSort{})
//     result2 := sorter.Sort(data)
// }
