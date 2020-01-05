from bs4 import BeautifulSoup
import requests
import csv
import sys
reload(sys)
sys.setdefaultencoding('utf-8')

location_list = []

class Stop:
    id = 0
    lat = 0
    lon = 0
    imageUrl = ""
    description = ""
    name = ""

    def __init__(self, id, name, description, imageUrl, lon, lat):
        self.id = id
        self.name = name
        self.description = description
        self.imageUrl = imageUrl
        self.lon = lon
        self.lat = lat

with open('wiki_links.csv') as csv_file:
    csv_reader = csv.reader(csv_file, delimiter=',')
    for row in csv_reader:
        try:
            page_link = row[1]

            if page_link.find("http") < 0:
                continue
            
            print("Scraping: ", row[1])
            STOP_ID = row[0]

            page_response = requests.get(page_link, timeout=100)
            if page_response is None:
                continue
            page_content = BeautifulSoup(page_response.content, "html.parser")
            info_table = page_content.find(id='wdinfobox')
            
            if info_table is None:
                continue

            image = info_table.find_all('img')
            if image is None or len(image) < 1:
                STOP_IMAGE_SRC = ""
            else:
                STOP_IMAGE_SRC = image[0]['src'].encode()

            ma_link = info_table.find_all('a', {"class": "mw-kartographer-map"})
            if len(ma_link) > 0:
                STOP_LAT = ma_link[0]['data-lat']
                STOP_LON = ma_link[0]['data-lon']
            else:
                STOP_LAT = ""
                STOP_LON = ""

            trs = page_content.find_all('tr')
            div = trs[0].find_all('div')
            if len(div) > 0:
                STOP_DESCRIPTION = div[0].text
            elif div == "" or div is None or len(div) == 0:
                STOP_DESCRIPTION = ""
            else:
                STOP_DESCRIPTION = div.text.encode()

            # captions = info_table.find(id='wdinfoboxcaption')
            # if captions is None:
            #     STOP_NAME = ""
            # else:
            STOP_NAME = row[1].split(':')[2].encode()

            s = Stop(STOP_ID, STOP_NAME, STOP_DESCRIPTION, STOP_IMAGE_SRC, STOP_LON, STOP_LAT)
            location_list.append(s)
        except:
            print("Something went wrong ", row)
        

with open('final_locations.csv', mode='w') as locations_file:
    loc_writer = csv.writer(locations_file, delimiter=',', quotechar='"', quoting=csv.QUOTE_MINIMAL)
    loc_writer.writerow(["Name", "Description", "ImageUrl", "Latitude", "Longitude"])

    for location in location_list:
        try:
            print("Writing:", location.id, location.name, location.description, location.imageUrl, location.lat, location.lon)
            loc_writer.writerow([location.id, location.name, location.description, location.imageUrl, location.lat, location.lon])
        except:
            print("Something went wrong during writing ", location)