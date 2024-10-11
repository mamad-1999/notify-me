package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Ensure correct number of arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: notify-me HH:MM \"message\"")
		return
	}

	// Parse input time
	alarmTime := os.Args[1]
	message := os.Args[2]

	// Normalize time format to HH:MM, handling single-digit hours (e.g., 1:23)
	if !strings.Contains(alarmTime, ":") || len(alarmTime) < 4 || len(alarmTime) > 5 {
		fmt.Println("Invalid time format. Use HH:MM or H:MM")
		return
	}

	timeParts := strings.Split(alarmTime, ":")
	hourStr, minuteStr := timeParts[0], timeParts[1]

	// Convert to integer values
	hour, err := strconv.Atoi(hourStr)
	if err != nil || hour < 0 || hour > 23 {
		fmt.Println("Invalid hour value")
		return
	}

	minute, err := strconv.Atoi(minuteStr)
	if err != nil || minute < 0 || minute > 59 {
		fmt.Println("Invalid minute value")
		return
	}

	// Get current time and calculate duration until alarm
	now := time.Now()
	alarm := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())
	if alarm.Before(now) {
		// If alarm time is before the current time, set it for the next day
		alarm = alarm.Add(24 * time.Hour)
	}

	// Sleep until alarm time
	duration := time.Until(alarm)
	fmt.Printf("Sleeping for %v until %v\n", duration, alarm)
	time.Sleep(duration)

	// Trigger notification and sound when the alarm goes off
	notify(message)
	playSound()

	// Close and exit
	fmt.Println("Alarm completed!")
}

// Notify sends a desktop notification using the notify-send command
func notify(message string) {
	cmd := exec.Command("notify-send", "Alarm", message)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error sending notification:", err)
	}
}

// playSound plays a short sound using aplay or another system sound tool
func playSound() {
	cmd := exec.Command("aplay", "/home/mohammad/Videos/go/notify-me/alarm.wav") // or your preferred sound file
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error playing sound:", err)
	}
}
