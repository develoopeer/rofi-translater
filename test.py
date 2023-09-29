import sqlite3
import json
import re
import zipfile
import base64
import os
import sys

CLEANR = re.compile('<.*?>') 

def cleanhtml(raw_html):
	cleantext = re.sub(CLEANR, '', raw_html)
	return cleantext


def parse_apkg(apkg_file):
	dump_data = {}
	archive = zipfile.ZipFile(apkg_file, 'r')
	media_file = json.loads(archive.read('media'))
	archive.extract('collection.anki2')

	con = sqlite3.connect('collection.anki2')
	cur = con.cursor()
	cur.execute("""select flds from notes;""")
	data = cur.fetchall()
	for elem in data:
		need = elem[0].split('\x1f').index([item for item in elem[0].split('\x1f') if item.startswith('[sound:')][0])
		string_name = elem[0].split('\x1f')[need].replace('sound:' , '').replace('[' , '').replace(']' , '').replace('&nbsp' , '')
		if string_name[-1] == ';':
			string_name = string_name[:-1]
		if len(string_name) != 0:
			value =  list(media_file.keys())[
				list(media_file.values()).index(string_name)
			]
		current_data = []
		current_file = base64.b64encode(archive.read(value))
		print(type(current_file))
		current_data.append(current_file.decode())
		current_data.append(cleanhtml(elem[0].split('\x1f')[4].split('&nbsp;')[0]))
		dump_data[elem[0].split('\x1f')[0]] = current_data

	with open('dump.json' , 'w') as file:
		json.dump(dump_data , file , indent=4)
	os.remove('collection.anki2')


def read_media():
	archive = zipfile.ZipFile('Cambridge_Vocabulary_for_IELTS_Advanced.apkg', 'r')
	media_file = archive.read('447')
	b = base64.b64encode(media_file)
	print(b)


parse_apkg('Cambridge_Vocabulary_for_IELTS_Advanced.apkg')