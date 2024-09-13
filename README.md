# Tranlater app for rofi

## How it looks like

![Preview](docs/preview.gif)

## Requirements
- rofi
- dmenu

## Installation

### Manually
1) Build and install golang binary with
```bash
go build . -o ttr-cli
go install
```
2) Create ~/.config/ttr/ folder and copy launch script to it
```bash
mkdir -p ~/.config/ttr/
cp launch.sh ~/.config/ttr/
```
3) Inside of your sxhkd configuration add keybinding
```bash
super + t
	~/.config/ttr/launch.sh
```

### From Package Manager
Arch
```bash
yay -S rofi-translater
```
Ubuntu
```bash
sudo apt-get install rofi-translater
```

## Configuration
In `launch.sh` script you can edit:
```bash
RASI="assets/default.rasi" # Default rofi theme file

```
