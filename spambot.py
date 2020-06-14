import requests
import discord
import asyncio
from config import *
import sys
import asyncio
import pyautogui
from bs4 import BeautifulSoup

print('Press Ctrl-C to quit.')
try:
    while True:
        x, y = pyautogui.position()
        positionStr = 'X: ' + str(x).rjust(4) + ' Y: ' + str(y).rjust(4)
        print("position is :={}".format(positionStr))
except KeyboardInterrupt:
     print('\nDone.')

client = discord.Client()
token = sys.argv[1]

# Get and print the mouse coordinates.

# permision integer - discord = 1811950599
