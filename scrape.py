#!/usr/bin/python

import requests
from bs4 import BeautifulSoup
import csv

year = 2021
url = 'https://admeter.usatoday.com/results/'
uri = '%s%d' % (url, year)

response = requests.get(uri)
soup = BeautifulSoup(response.text, "html.parser")
commercials = soup.find(id='commercials-list')

with open('%d.csv' % year, 'w') as csvfile:

    csvwriter = csv.writer(csvfile)

    for index, entry in enumerate(commercials.find_all('article', class_='commercial-block collapsible-block filterable__item commercial-block--results')):    
        link = entry.find('a', class_='commercial-block__video-title')['href']
        title = entry.find('a', class_='commercial-block__video-title')['title']
        brand = entry['data-advertiser']
        air_time = entry['data-quarter']

        score = entry.find('dd', class_="average-score__num")
        avg_ranking = score.text.strip()

        csvwriter.writerow([year, brand, title, index+1, avg_ranking, air_time, link])    