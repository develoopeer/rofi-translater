#!/usr/bin/env bash
RASI="~/.config/ttr/assets/default.rasi"
CAM_TRANSLATE=true
LIBRE_TRANSLATE=true
LIBRE_TRANSLATE_URL="http://localhost:5000/translate" # Url for libre translate instance
LIBRE_TRANSLATE_TARGET="en" # Any language supported by libre translate #TODO
CAM_TRANSLATE_DICTIONARY="english" # Any language supported by #TODO
TRANSLATION_ORDER=1 # 1(Cambridge translate , Libre translate) , 2(Libre translate , Cambridge translate)

if [[ $LIBRE_TRANSLATE == true ]]; then
	RUN_STRING="ttr translate --order $TRANSLATION_ORDER --libre --libre_target $LIBRE_TRANSLATE_TARGET --cam_dict $CAM_TRANSLATE_DICTIONARY"
else
	RUN_STRING="ttr translate --order $TRANSLATION_ORDER --libre_target $LIBRE_TRANSLATE_TARGET --cam_dict $CAM_TRANSLATE_DICTIONARY"
fi
rofi -show Search -modes "Search:$RUN_STRING" -theme  $RASI -p Search
