from urllib.parse import unquote

from django.shortcuts import render
import requests
import re
import requests
from bs4 import BeautifulSoup
from ninja import Router
from config.utils import sanitizing_url
from project.schemas.video_schema import VideoOut, MessageOut
from django.utils.safestring import mark_safe
from lxml import html

embed_controller = Router(tags=["Embed Controller"])



# main endpoint
@embed_controller.get("")
def get_video(request, url: str = None):
    ctx = {}
    meta = {}
    print(url)
    if "watch" or "reel" in url:
        sanitized_url = sanitizing_url(url)
        if "www" in sanitized_url:
            sanitized_url = sanitized_url.replace("www.", "mbasic.", 1)
        else:
            sanitized_url = sanitized_url.replace("web.", "mbasic.", 1)
        response = requests.get(sanitized_url)

        if 'video_redirect' in response.text:
            reel = re.search(r'href\=\"\/video\_redirect\/\?src\=(.*?)\"', response.text)
            video = unquote(reel.group(1)).replace(";", "&")
            ctx["video"] = mark_safe(video)
            tree = html.fromstring(response.content)
            for tag in tree.xpath('//meta'):
                meta[tag.get('property')] = mark_safe(tag.get('content'))
            # assign every meta tag to the context
            for key, value in meta.items():
                if key and key != 'og:image':
                    ctx[key[3:]] = value

    else:
        # Send a GET request to the post URL and parse the HTML using BeautifulSoup
        response = requests.get(url)
        soup = BeautifulSoup(response.content, 'html.parser')

        # Extract the image URL from the meta tag with property="og:image"
        image_url = soup.find('meta', property='og:image')['content']
        ctx["image_url"] = image_url
    return render(request, "base.html", ctx)

