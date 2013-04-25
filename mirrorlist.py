#!/usr/bin/env python
import argparse
from urllib.request import urlopen
import json


def get_mirrors():
    result = urlopen('https://www.archlinux.org/mirrors/status/json/')
    return result.read().decode('UTF-8')


def parse_json(data):
    result = json.loads(data)
    return [Mirror(x['protocol'], x['url'], x['country'],
    x['last_sync'], x['score']) for x in result['urls']]


def parse_args():
    desc = 'A small program to fetch ArchLinux mirrors.'
    parser = argparse.ArgumentParser(description=desc)
    parser.add_argument('-o', help='Output to file.')
    return parser.parse_args()


class Mirror:
    def __init__(self, protocol, url, country, last_sync, score):
        self.protocol = protocol
        self.url = url
        self.country = country
        self.last_sync = last_sync
        self.score = score

    def __str__(self):
        return "{} mirror at {} with score: {}".format(self.protocol,
        self.url, self.score)


if __name__ == '__main__':
    mirrors = get_mirrors()
    parsed = parse_json(mirrors)
    args = parse_args()
    parsed.sort(key=lambda mirror: mirror.score)
    for i in parsed:
        location = '$repo/os/$arch'
        print('# Score: {}, {}'.format(i.score, i.country))
        print("Server = {}{}\n".format(i.url, location))
