"""Main file for running app
"""
import os
import webbrowser
from threading import Thread

from ttr.configs import LOGGER_FILE , DICTIONARY_SHORTCUT , GOOGLE_TRANSLATE_SHORTCUT
from ttr.configs import first_try , add_word_to_db
from ttr.parser import parse_cam_dictionary, parse_translater

import pyclip

from notifypy import Notify
from pystray import Icon , Menu as menu, MenuItem as item
from loguru import logger
from PIL import Image
from pynput import keyboard

count = 0
logger.add(LOGGER_FILE, format="{time:MMMM D, YYYY > HH:mm:ss} | {level} | {message}")


def current_folder():
    return os.path.dirname(os.path.realpath(__file__))

def on_activate():
    word = str(pyclip.paste(text=True))
    webbrowser.open("https://dictionary.cambridge.org/dictionary/english/{0}".format(
        word.lower())
    )
    logger.debug('{0} was googled'.format(word))
    print((word , parse_cam_dictionary , parse_translater))
    add_word_to_db((word , parse_cam_dictionary(word) , parse_translater(word)))

def open_translater():
    word = str(pyclip.paste(text=True))
    gtr_link = "https://translate.google.com/?hl=ru&sl=auto&tl=ru&text={0}&op=translate"
    webbrowser.open(gtr_link.format(
        word.replace(" " , "%20").lower())
    )
    logger.debug('{0} was googled'.format(word))
    add_word_to_db((word , parse_cam_dictionary(word) , parse_translater(word)))

def run_listener():
    print('| ----------------Listener invoked---------------- |')
    with keyboard.GlobalHotKeys({
        GOOGLE_TRANSLATE_SHORTCUT: open_translater,
        DICTIONARY_SHORTCUT: on_activate
    }) as listener:
        listener.join()

def on_clicked(icon):
    notification = Notify()
    notification.title = "Tray Translater"
    notification.icon  = os.path.join(current_folder() , "icon.png")
    notification.audio = os.path.join(current_folder( ), 'notify.wav')
    notification.message = "The app has been closed"
    icon.stop()
    notification.send(block=False)

def import_stats():
    pass

def run_icon(icon):
    icon.run()


def main():
    first_try()
    icon = Icon('TTR' , Image.open(os.path.join(current_folder() , 'icon.png')))
    current_menu = menu(item('Import stats' , lambda: import_stats()) , menu.SEPARATOR , item('Exit' , lambda: on_clicked(icon)))
    icon.menu = current_menu
    icon_listen = Thread(target = run_listener , daemon=True)

    icon_listen.start()
    run_icon(icon)

main()
