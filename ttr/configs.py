"""This file must not be complicated that much
"""
import os
import json

LOGGER_FILE = "logs.log"
DICTIONARY_SHORTCUT = "<ctrl>+<alt>+h"
GOOGLE_TRANSLATE_SHORTCUT="<ctrl>+<alt>+]"
HOME_FOLDER=os.path.expanduser('~')
JSON_FILENAME=os.path.join(HOME_FOLDER , "ttr.json")

def first_try():
    if not os.path.exists(HOME_FOLDER):
        os.mkdir(HOME_FOLDER)
    if not os.path.exists(JSON_FILENAME):
        os.mknod(JSON_FILENAME)
        with open(JSON_FILENAME , 'w' , encoding='utf-8') as file:
            json.dump({"words": []} , file)

def get_json():
    with open(JSON_FILENAME , 'r') as file:
        return json.load(file)

def add_word_to_db(word):
    """Word is tuple like object: english , dictionary , translation
    """
    cur_json = get_json()
    cur_json['words'].append({
        "en": word[0],
        "dic": word[1],
        "trans": word[2]
    })
    with open(JSON_FILENAME , 'w' , encoding='utf-8') as file:
        json.dump(cur_json , file , ensure_ascii=False)
        # ensure_ascii is necessary for utf-8 symbols