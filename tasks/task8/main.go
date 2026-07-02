package main

import (
	"fmt"
	"strings"
)
func removeInactiveUsers(users map[string]bool) {
    for name, active := range users {
        if !active {
            delete(users, name)
        }
    }
}

func findMaxScore(scores map[string]int) (string, int) {
    maxName := ""
    maxScore := 0
    first := true

    for name, score := range scores {
        if first || score > maxScore {
            maxName = name
            maxScore = score
            first = false
        }
    }

    return maxName, maxScore
}

func groupByLength(words []string) map[int][]string {
    groups := make(map[int][]string)

    for _, word := range words {
        length := len(word)
        groups[length] = append(groups[length], word)
    }

    return groups
}

func main() {
	// Task 1: Count the occurrences of each word in a string
	fmt.Println("=== Word Count ===")
    text := "go is fun go is powerful go is simple"
    words := strings.Split(text, " ")

    count := make(map[string]int)

    for _, word := range words {
        count[word]++
    }

    fmt.Println(count)
	// Task 2: Check if a user has admin permissions
	fmt.Println("=== Admin Permissions Check ===")	
	permissions := map[string]bool{
        "admin": true,
        "user":  false,
    }

    if perm, exists := permissions["admin"]; exists && perm {
        fmt.Println("Доступ администратора разрешён")
    }

    if _, exists := permissions["guest"]; !exists {
        fmt.Println("Пользователь guest не найден")
    }


// Task 3: Remove inactive users from a map
	fmt.Println("=== Remove Inactive Users ===")
    users := map[string]bool{
        "Alice": true,
        "Bob":   false,
        "Eve":   true,
        "Dave":  false,
    }

    removeInactiveUsers(users)

    fmt.Println(users)  // map[Alice:true Eve:true]


// Task 4: Find the user with the highest score
	fmt.Println("=== Highest Score ===")
    scores := map[string]int{
        "Alice": 95,
        "Bob":   88,
        "Eve":   92,
    }

    name, score := findMaxScore(scores)
    fmt.Printf("Highest score: %s with %d\n", name, score)
// Task 5: Nested maps to store student grades
	fmt.Println("=== Student Grades ===")
	 grades := map[string]map[string]int{
        "Alice": {
            "Math":    90,
            "Physics": 85,
        },
        "Bob": {
            "Math":    78,
            "Physics": 92,
        },
    }

    for student, subjects := range grades {
        fmt.Printf("%s's grades:\n", student)
        for subject, grade := range subjects {
            fmt.Printf("  %s: %d\n", subject, grade)
        }
    }

// Task 6: Group words by their length
	//words := []string{"go", "is", "fun", "and", "fast", "simple"}
	fmt.Println("=== Grouping Words by Length ===")
    grouped := groupByLength(words)

    for length, group := range grouped {
        fmt.Printf("Length %d: %v\n", length, group)
    }
}


