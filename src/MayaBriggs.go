package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	//extrovertCount := 0
	//introvertCount := 0
	//sensitiveCount := 0
	//intuitiveCount := 0
	//thinkerCount := 0
	//feelerCount := 0
	//judgingCount := 0
	//perspectiveCount := 0

	questions := []string{
		"1) A. Expend energy, enjoy groups.       B. Conserve energy, one-on-one",
		"2) A. Interpret literally.               B. Look for meaning and possibilities",
		"3) A. Logical, thinking, questioning.    B. Empathetic, feeling, accommodating",
		"4) A. Organized, orderly.                B. Flexible, adaptable",
		"5) A. More outgoing, think out loud.     B. More reserved, think to yourself.",
		"6) A. Practical, realistic, experiential.B. Imagination, innovative, theoretical.",
		"7) A. Candid, straight forward, frank.   B. Tactful, kind, encouraging.",
		"8) A. Plan, schedule.                    B. Unplanned, spontaneous",
		"9) A. Seek many tasks, public activities, interaction with others. B. Seek private, solitary activities with quiet to concentrate.",
		"10) A. Standard, usual, conventional.     B. Different, novel, unique.",
		"11) A. Firm, tend to criticize, hold the line. B. Gentle, tend to appreciate, conciliate.",
		"12) A. Regulated, structured.             B. Easygoing, live and let live.",
		"13) A. External, communicative, express yourself. B. Internal, reticent, keep to yourself.",
		"14) A. Focus on here-and-now.             B. Look to the future, global perspective, big picture",
		"15) A. Tough minded, just.                B. Tender-hearted, merciful",
		"16) A. Preparation, plan ahead.           B. Go with the flow, adapt as you go.",
		"17) A. Active, initiate.                  B. Reflective, deliberate",
		"18) A. Facts, things, what is.            B. Ideas, dreams, 'what could be', philosophical",
		"19) A. Matter of fact, issue oriented.    B. Sensitive, people-oriented, compassionate",
		"20) A. Control, govern.                   B. Latitude, freedom",
	}

	userResponses := make([]int, len(questions))

	fmt.Print("What's your name: ")
	scanner.Scan()
	userName := scanner.Text()

	for response := 0; response < len(questions); response++ {
		userResponses[response] = askQuestion(questions[response], scanner)
	}

	fmt.Printf("\n%s, your personality type is: ", userName)
	personalityType := checkPersonalityType(userResponses)
	fmt.Println(personalityType)
}

func askQuestion(question string, scanner *bufio.Scanner) int {
	fmt.Println(question)
	fmt.Print("A or B: ")
	scanner.Scan()
	response := strings.ToUpper(scanner.Text())
	if response == "A" {
		return 1
	} else if response == "B" {
		return 2
	}
	return 0
}

func checkPersonalityType(userResponses []int) string {
	var personalityType strings.Builder
	personalityType.WriteString(extrovertOrIntrovert(userResponses[0], userResponses[1]))
	personalityType.WriteString(sensitiveOrIntuitive(userResponses[2], userResponses[3]))
	personalityType.WriteString(thinkerOrFeeler(userResponses[4], userResponses[5]))
	personalityType.WriteString(judgingOrPerspective(userResponses[6], userResponses[7]))
	return personalityType.String()
}

func extrovertOrIntrovert(extrovertCount, introvertCount int) string {
	if extrovertCount < introvertCount {
		return "I"
	}
	return "E"
}

func sensitiveOrIntuitive(sensitiveCount, intuitiveCount int) string {
	if sensitiveCount < intuitiveCount {
		return "I"
	}
	return "S"
}

func thinkerOrFeeler(thinkerCount, feelerCount int) string {
	if thinkerCount < feelerCount {
		return "F"
	}
	return "T"
}

func judgingOrPerspective(judgingCount, perspectiveCount int) string {
	if judgingCount < perspectiveCount {
		return "P"
	}
	return "J"
}
