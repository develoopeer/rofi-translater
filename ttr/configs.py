"""This file must not be complicated that much
"""
import os
import json

DICTIONARY_SHORTCUT = "<ctrl>+<alt>+h"
GOOGLE_TRANSLATE_SHORTCUT="<ctrl>+<alt>+]"
HOME_FOLDER=os.path.expanduser('~')
JSON_FILENAME=os.path.join(HOME_FOLDER , "ttr.json")
LOGGER_FILE = os.path.join(HOME_FOLDER , "logs.log")
PID_FOLDER = os.path.join("/run/user/{0}/ttr/".format(os.getuid()))
PID_FILE   = os.path.join(PID_FOLDER , "ttr.pid")

def first_try():
    if not os.path.exists(HOME_FOLDER):
        os.mkdir(HOME_FOLDER)
    if not os.path.exists(JSON_FILENAME):
        os.mknod(JSON_FILENAME)
        with open(JSON_FILENAME , 'w' , encoding='utf-8') as file:
            json.dump({"words": []} , file , indent=4)
    if not os.path.exists(PID_FOLDER):
        os.mkdir(PID_FOLDER)
    if not os.path.exists(PID_FILE):
        os.mknod(PID_FILE)


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
        json.dump(cur_json , file , ensure_ascii=False , indent=4)
        # ensure_ascii is necessary for utf-8 symbols

def is_running():
    with open(PID_FILE , 'r') as file:
        if len(file.read()) != 0:
            return True
    return False

def dump_pid(pid):
    with open(PID_FILE , "w") as file:
        file.write(str(pid))

def clean_pid():
    with open(PID_FILE , "w") as file:
        file.write("")

class AnotherInstanceIsRunning(Exception):
    pass
