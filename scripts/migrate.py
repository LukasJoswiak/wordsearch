# Converts data from old database format to new database format and inserts
# into the database.

import csv
import datetime
import json
import mysql.connector
import random
import requests
import time

def connect():
    conn = None
    try:
        conn = mysql.connector.connect(host='localhost',
                                       database='wordsearch',
                                       user='root',
                                       password='password')
    except mysql.connector.Error as e:
        print(e)

    return conn


def generate_url():
    return random.randint(1000000000, 9999999999)


def convert_datetime(dt):
    return datetime.datetime.strptime(dt, '%Y-%m-%d %H:%M:%S')


# Given a data string in the old database format, converts it into
# the new format and returns it as a string.
def format_data(data):
    for i in range(0, len(data)):
        # Remove None types and empty strings from the data
        data[i] = ''.join(list(filter(None, data[i])))

    data = (','.join(data)).lower()
    return data


# Inserts a puzzle with the given data into the database, given a puzzle with
# the passed in URL does not already exist. Returns true on success.
def insert_puzzle(cursor, url, view_url, data, puzzle_type, datetime):
    try:
        data = json.loads(data)
    except:
        return False
    data = format_data(data)

    select_puzzle = ("SELECT url "
                     "FROM puzzles "
                     "WHERE url = %s")
    cursor.execute(select_puzzle, (url,))
    results = cursor.fetchall()
    if len(results) > 0:
        print(f'row with url {url} already exists, skipping...')
        return False

    insert_puzzle = ("INSERT INTO puzzles "
                     "(url, view_url, data, type, datetime) "
                     "VALUES (%s, %s, %s, %s, %s)")

    puzzle_data = (url, view_url, data, puzzle_type, datetime)

    cursor.execute(insert_puzzle, puzzle_data)
    row_id = cursor.lastrowid
    conn.commit()
    print(f'inserted puzzle with url {url} at row {row_id}')
    return True


if __name__ == '__main__':
    conn = connect()
    cursor = conn.cursor()

    print('parsing csv...')
    with open('puzzles.csv') as csv_file:
        csv_reader = csv.reader(csv_file, delimiter=',')
        rows = list(csv_reader)
        urls = {}

        # Change this value to skip some rows.
        skip_rows = 0
        rows = rows[skip_rows:]

        for row in rows:
            url = row[5]
            urls[url] = urls.get(url, 0) + 1

        line_count = skip_rows
        for row in rows:
            line_count += 1
            row_id, width, height, data, words, url, share, puzzle_type, dt = row

            urls[url] -= 1
            # Ignore rows that are duplicates. In the old system, duplicate URLs
            # caused both the new and old row to be overwritten, so only keep
            # the latest row when copying data to new database.
            if urls[url] != 0:
                print(f'url {url} is a duplicate, skipping...')
                continue

            # Generate shareable URLs for puzzles that don't currently have
            # them.
            if share == '0':
                share = generate_url()

            dt = convert_datetime(dt)

            if len(data) > 0:
                success = insert_puzzle(cursor, url, share, data, puzzle_type, dt)
                if not success:
                    continue

                if len(words) > 0:
                    try:
                        words = json.loads(words)
                    except:
                        continue
                    words = [word.lower().strip() for word in words]
                    # Remove duplicate words.
                    words = list(dict.fromkeys(words))

                    word_data = {}
                    for i, word in enumerate(words):
                        if len(word) > 255:
                            continue
                        word_data[f'Words.{i}.ExistingWord'] = ''
                        word_data[f'Words.{i}.Word'] = word

                    word_url = f'http://localhost:8080/puzzle/{url}/words'
                    r = requests.post(url=word_url, data=word_data)
                    if r.status_code != 200:
                        print(f'non-200 status code {r.status_code}, exiting...')
                        break
                    print(f'successfully inserted {len(words)} words for url {url} (row {line_count})')

