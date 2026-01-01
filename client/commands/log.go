package commands

import (
	"fmt"
	"time"

	"github.com/adk-saugat/stash/models"
)

type LogCommand struct{}

func (c *LogCommand) Name() string        { return "log" }
func (c *LogCommand) Description() string { return "Show store history" }

func (c *LogCommand) Run(args []string) error {
	stores, err := models.GetAllStores()
	if err != nil {
		return fmt.Errorf("could not read stores: %w", err)
	}

	if len(stores) == 0 {
		fmt.Println("No stores yet.")
		return nil
	}

	fmt.Println()
	fmt.Printf("Store History (%d stores)\n", len(stores))
	fmt.Println("─────────────────────────")
	fmt.Println()
	for i, store := range stores {
		shortId := store.StoreId[:8]
		fmt.Printf("● %s - %s (%s) <%s>\n", shortId, store.Message, timeAgo(store.Date), store.Author)

		// Print connector line (except for last item)
		if i < len(stores)-1 {
			fmt.Println("│")
		}
	}
	fmt.Println()

	return nil
}

func timeAgo(t time.Time) string {
	diff := time.Since(t)

	if diff < time.Minute {
		return "just now"
	} else if diff < time.Hour {
		return fmt.Sprintf("%d minutes ago", int(diff.Minutes()))
	} else if diff < 24*time.Hour {
		return fmt.Sprintf("%d hours ago", int(diff.Hours()))
	} else {
		return fmt.Sprintf("%d days ago", int(diff.Hours()/24))
	}
}

