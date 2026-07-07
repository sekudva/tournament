package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/sekudva/strategika/internal/service"
)

const (
	duelFile       = "results/duel_log.txt"
	roundRobinFile = "results/roundrobin_log.txt"
	trialFile      = "results/trial_log.txt"
	circulaireFile = "results/circulaire_log.txt"
	arenaFile      = "results/arena_log.txt"
	ecosystemFile  = "results/ecosystem_log.txt"
)

func handleDuel(rounds int, noise float64) {
	fmt.Println("\n=== DUEL ===")
	fmt.Println("Choose 2 strategies:")

	a1 := selectOne(selectGroup())
	a2 := selectOne(selectGroup())

	fmt.Printf("\nDuel: %s VS %s\n", a1.Name, a2.Name)

	if err := service.RunDuel(a1, a2, rounds, noise, duelFile); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	Results(duelFile)
}

func handleRoundRobin(rounds int, noise float64) {
	fmt.Println("\n=== ROUND-ROBIN ===")
	fmt.Println("Choose many strategies:")

	agents := selectEach(selectGroup())

	noSelf := true
	if !Quick {
		noSelf = readInt("Allow self-games? 0 - NO, 1 - YES [NO]: ", 0, 1, 0) == 0
	}

	fmt.Printf("\nRound Robin: %d agents, self-games: %v\n", len(agents), !noSelf)

	if err := service.RunRoundRobin(agents, rounds, noise, noSelf, roundRobinFile); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	Results(roundRobinFile)
}

func handleTrial(rounds int, noise float64) {
	fmt.Println("\n=== TRIAL ===")
	fmt.Println("Choose one leader and group:")

	fmt.Println("\nChoose leader (or victim):")
	leader := selectOne(selectGroup())

	fmt.Println("Choose group:")
	group := selectEach(selectGroup())

	fmt.Printf("\nTrial: %s VS %d agents\n", leader.Name, len(group))

	if err := service.RunTrial(leader, group, rounds, noise, trialFile); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	Results(trialFile)
}

func handleCirculaire(rounds int, noise float64) {
	fmt.Println("\n=== CIRCULAIRE ===")
	fmt.Println("Choose many strategies, many LIDERS and one GROUP:")

	fmt.Println("\nChoose leaders:")
	leaders := selectEach(selectGroup())

	fmt.Println("Choose group:")
	group := selectEach(selectGroup())

	fmt.Printf("\nCirculaire: %d leaders VS %d agents\n", len(leaders), len(group))

	if err := service.RunCirculaire(leaders, group, rounds, noise, circulaireFile); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	Results(circulaireFile)
}

func handleArena(rounds int, noise float64) {
	fmt.Println("\n=== ARENA ===")
	fmt.Println("Choose many strategies:")

	agents := selectEach(selectGroup())

	fmt.Printf("\nArena: %d agents\n", len(agents))

	if err := service.RunArena(agents, rounds, noise, arenaFile); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	Results(arenaFile)
}

func handleEcosystem(rounds int, noise float64) {
	fmt.Println("\n=== ECOSYSTEM ===")
	fmt.Println("Choose many strategies:")

	agents := selectEach(selectGroup())

	threshold := 0
	if !Quick {
		threshold = readInt("Death threshold (scores) [0]: ", -100000, 100000, 0)
	}

	fmt.Printf("\nEcosystem: %d agents, death threshold: %d\n", len(agents), threshold)

	if err := service.RunEcosystem(agents, rounds, noise, threshold, ecosystemFile); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	Results(ecosystemFile)
}

func handleInfo() {
	data, err := os.ReadFile("info.txt")
	if err != nil {
		fmt.Println("Info file not found.")
		return
	}
	fmt.Println(strings.TrimSpace(string(data)))
}

func Results(name string) {
	fmt.Println("")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("Results in generated file: %s\n", name)
	fmt.Println(strings.Repeat("=", 60))
	fmt.Print("\nWARNING: This file will override previous version\nif you run simulation with same Simulation Mode!\n")
	fmt.Print("If you want to save result, rename previous file.\n")
}
