package helpers

import "fmt"

func FormatDuration(duration int) string {
	if duration < 60 {
		return "Just now"
	} else if duration < 3600 {
		minutes := duration / 60
		return fmt.Sprintf("%d minutes ago", minutes)
	} else if duration < 86400 {
		hours := duration / 3600
		return fmt.Sprintf("%d hours ago", hours)
	} else {
		days := duration / 86400
		return fmt.Sprintf("%d days ago", days)
	}
}
