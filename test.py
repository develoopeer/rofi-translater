import sqlite3
import json
import re
CLEANR = re.compile('<.*?>') 

def cleanhtml(raw_html):
  cleantext = re.sub(CLEANR, '', raw_html)
  return cleantext

con = sqlite3.connect("collection.anki2")
cur = con.cursor()

cur.execute("""select flds from notes;""")

data = cur.fetchall()


def dump_it():
    print(len(data))
    dump_data = {}
    for some in data:
        dump_data[some[0].split('\x1f')[0]] = cleanhtml(some[0].split('\x1f')[4].split('&nbsp;')[0])

    with open('dump.json' , 'w') as file:
        json.dump(dump_data , file , indent=4)

dump_it()