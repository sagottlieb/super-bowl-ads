#!/usr/bin/python

import requests
from bs4 import BeautifulSoup
import csv

year = 2018
url = 'https://admeter.usatoday.com/results/'
uri = '%s%d' % (url, year)

response = requests.get(uri)
soup = BeautifulSoup(response.text, "html.parser")
content = soup.find(id='content')

with open('%d.csv' % year, 'w') as csvfile:

    csvwriter = csv.writer(csvfile)

    for index, entry in enumerate(content.find_all('article', class_='ranking')):    
        link = entry.find('a', class_='ranking_link')['href']
        
        ranking_parts = entry.find('dl').find_all('dd')
        [avg_ranking, air_time] = [x.text for x in ranking_parts]

        title = entry.find('h2', class_='ranking_title').text.strip()

        brand = entry.find('span', class_='ranking_advertiser').text

        csvwriter.writerow([year, brand, title, index+1, avg_ranking, air_time, link])    