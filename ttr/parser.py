import bs4
import requests


def parse_cam_dictionary(word):
    if len(word.split(" ")) != 1:
        return "line is too long"
    headers = {'User-Agent': "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36"}
    url = "https://dictionary.cambridge.org/dictionary/english/{0}".format(word.lower())
    data = bs4.BeautifulSoup(requests.get(url , headers=headers).text , 'html.parser')
    return data.select(".def.ddef_d.db")[0].text

def parse_translater(word):
    headers = {'User-Agent': "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36" , "Content-Type": "text/plain"}
    url = "https://translator-api.glosbe.com/translateByLangWithScore?sourceLang=en&targetLang=ru"
    response = requests.post(url , data = word.encode('utf-8') , headers=headers)
    return response.json()['translation']
