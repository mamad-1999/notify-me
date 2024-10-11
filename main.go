package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const soundURL = "https://raw.githubusercontent.com/mamad-1999/notify-me/refs/heads/master/alarm.wav"
const soundDir = ".notify-me"
const soundFileName = "alarm.wav"

func main() {
	showHelp := flag.Bool("h", false, "Display help")
	flag.Parse()

	if *showHelp {
		displayHelp()
		return
	}

	// Ensure correct number of arguments
	if len(flag.Args()) < 2 {
		fmt.Println("Error: Invalid number of arguments.")
		displayHelp()
		return
	}

	// Parse input time and message
	alarmTime := flag.Arg(0)
	message := flag.Arg(1)

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

	// Ensure sound file exists, download if necessary
	soundPath := getSoundFilePath()
	if _, err := os.Stat(soundPath); os.IsNotExist(err) {
		err := downloadSound(soundPath)
		if err != nil {
			fmt.Println("Error downloading sound:", err)
			return
		}
	}

	// Sleep until alarm time
	duration := time.Until(alarm)
	fmt.Printf("Sleeping for %v until %v\n", duration, alarm)
	time.Sleep(duration)

	// Trigger notification and sound when the alarm goes off
	notify(message)
	playSound(soundPath)

	// Close and exit
	fmt.Println("Alarm completed!")
}

// displayHelp shows usage instructions and examples
func displayHelp() {
	fmt.Println("Usage: notify-me [HH:MM] \"message\"")
	fmt.Println("Alarm program that notifies the user with a message at the specified time.")
	fmt.Println("\nOptions:")
	fmt.Println("  -h               Display help")
	fmt.Println("\nExamples:")
	fmt.Println("  notify-me 14:30 \"Go to the gym\"")
	fmt.Println("  notify-me 09:00 \"Join the meeting\"")
	fmt.Println("  notify-me 1:15 \"Take a break\"")
	fmt.Println("\nThe program will also play a sound when the alarm triggers.")
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
func playSound(soundPath string) {
	cmd := exec.Command("aplay", soundPath)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error playing sound:", err)
	}
}

// getSoundFilePath returns the path to the alarm.wav file in the user's home directory
func getSoundFilePath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, soundDir, soundFileName)
}

// downloadSound downloads the sound file from the GitHub repository and saves it locally
func downloadSound(dest string) error {
	fmt.Println("Downloading sound file...")

	// Create directory if not exists
	err := os.MkdirAll(filepath.Dir(dest), os.ModePerm)
	if err != nil {
		return err
	}

	// Download the sound file
	resp, err := http.Get(soundURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Save the file
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Sound file downloaded and saved to:", dest)
	return nil
}
