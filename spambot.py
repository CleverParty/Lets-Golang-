import requests
import discord
import asyncio
from config import *
import sys
import asyncio
from bs4 import BeautifulSoup

client = discord.Client()
token = sys.argv[1]


@client.event
@asyncio.coroutine()
def on_ready():
    while not client.is_closed:
        html = requests.get('https://evilinsult.com/generate_insult.php?lang=en&type=json')
        soup = BeautifulSoup(html, "html.parser")
        actual = soup.find('h1')
        print(actual.text)
        print(soup)
        client.send_message(discord.Object(id=DiscordChannel), insult_text.text)

        asyncio.sleep(1)  # Changes how fast the messages are posted. (Anything under 0.7 tends to break it


print(token)
client.run(token, bot=False)