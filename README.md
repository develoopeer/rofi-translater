# Tranlater app for rofi

## How it looks like

![Preview](docs/preview.gif)

## Requirements
- [rofi](https://github.com/davatorium/rofi)
- [dmenu](https://archlinux.org/packages/extra/x86_64/dmenu/)
- [libretranslate](https://github.com/LibreTranslate/LibreTranslate) (optional)

## Installation

### Manually
1) Build and install golang binary with
```bash
go install
```
2) Create `~/.config/ttr/` folder and copy launch script to it
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
### To install libretranslate
```bash
docker run -ti --rm -p 5000:5000 libretranslate/libretranslate
```
For more detailed info about installation process watch [documentation]( https://github.com/LibreTranslate/LibreTranslate?tab=readme-ov-file#install-and-run )

## Configuration
In `launch.sh` script you can edit:
```bash
RASI="assets/default.rasi" # Default rofi theme file
CAM_TRANSLATE=true
LIBRE_TRANSLATE=true
LIBRE_TRANSLATE_URL="http://localhost:5000/translate" # Url for libre translate instance
LIBRE_TRANSLATE_TARGET="en" # Any language supported by libre translate #TODO
CAM_TRANSLATE_DICTIONARY="english" # Any language supported by #TODO
TRANSLATION_ORDER=1 # 1(Cambridge translate , Libre translate) , 2(Libre translate , Cambridge translate)
```
