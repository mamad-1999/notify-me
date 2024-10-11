<h4 align="center">A Go script to set an alarm with notifications.</h4>
<p align="center">
  <a href="#installation"><img src="https://img.shields.io/badge/Install-blue?style=for-the-badge" alt="Install"></a>
  <a href="#usage"><img src="https://img.shields.io/badge/Usage-green?style=for-the-badge" alt="Usage"></a>
  <a href="#contributing"><img src="https://img.shields.io/badge/Contributing-yellow?style=for-the-badge" alt="Contributing"></a>
</p>

## Installation

```bash
go install github.com/mamad-1999/notify-me@latest
```
> [!IMPORTANT]
> #### Requirement packeage
> 
> The script uses the `notify-send` command for desktop notifications and `aplay` to play a sound. Make sure these utilities are installed on your system:
> ```bash
> sudo apt install libnotify-bin alsa-utils
> ```

When you first run the script, it will automatically download the sound file from the GitHub repository, so there is no need for manual configuration.

## Usage

```bash
notify-me [HH:MM] "Your reminder message"
```

Examples:

```bash
notify-me 14:30 "Go to the gym"
notify-me 1:15 "Take a break"
notify-me 09:00 "Join the meeting"
```

> [!TIP]  
> The script can be run in the background using `nohup`:
> ```bash
> nohup notify-me 14:30 "Go to the gym" &
> ```


## Contributing

Contributions to `notify-me` are welcome! If you encounter bugs or have suggestions for improvements, feel free to open issues or submit pull requests on the GitHub repository.
