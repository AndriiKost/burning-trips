import csv
import psycopg2
import os

with open('final_locations.csv') as csvfile:
    readCSV = csv.reader(csvfile, delimiter=',')
    for row in readCSV:
        if row:
            row0 = row[0]
            row1 = row[1]
            row2 = row[2]
            row3 = row[3]
            row4 = row[4]
            row5 = row[5]

            if not row4:
                row4 = 0
            
            if not row5:
                row5 = 0

            conn = psycopg2.connect(dbname='burning_trips', user='burning_trips', password='burning_trips', host='burning-trips-dev.cmdrlis240ol.us-east-2.rds.amazonaws.com', port='5432')
            cur = conn.cursor()
            print("INSERTING VALUE %s,%s", row0, row1)
            cur.execute("INSERT INTO landmarks (id, name, description, image, latitude, longitude) VALUES (%s,%s,%s,%s,%s,%s)",
            [row0, row1, row2, row3, row4, row5])
            result = cur.execute("SELECT * FROM landmarks")
            print("Result ", result)
