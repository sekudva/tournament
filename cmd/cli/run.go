package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/sekudva/strategika/internal/service"
)

// Run — основной хендлер, запускает меню и диспатчит режимы.
func Run() error {
	service.Quick = Quick
	service.Silent = Silent

	if err := os.MkdirAll("results", 0755); err != nil {
		return fmt.Errorf("cannot create results directory: %w", err)
	}

	for {
		mode := MainMenu()
		if mode == 0 {
			return nil
		}
		var rounds int
		var noise float64

		if Quick || mode == 7 {
			rounds = 200
			noise = 0.0
		} else {
			rounds = readInt("Number of rounds [200]: ", 1, 10000, 200)
			noise = readFloat("Channel noise (0.0-1.0) [0.0]: ", 0.0, 1.0, 0.0)
		}

		switch mode {
		case 1:
			handleDuel(rounds, noise)
		case 2:
			handleRoundRobin(rounds, noise)
		case 3:
			handleTrial(rounds, noise)
		case 4:
			handleCirculaire(rounds, noise)
		case 5:
			handleArena(rounds, noise)
		case 6:
			handleEcosystem(rounds, noise)
		case 7:
			handleInfo()
			pause()
			continue
		}

		pause()
	}
}

func pause() {
	fmt.Println("\n⏳ Return to menu in 5 seconds...")
	time.Sleep(5 * time.Second)
}
