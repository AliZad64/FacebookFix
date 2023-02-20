from urllib.parse import unquote

from django.shortcuts import render
import requests
import re
import requests
from bs4 import BeautifulSoup
from ninja import Router
from config.utils import sanitizing_url
from project.schemas.video_schema import VideoOut, MessageOut

embed_controller = Router(tags=["Embed Controller"])


# main endpoint
@embed_controller.get("", response={200: MessageOut, 400: MessageOut})
def get_video(request, url: str = None):
    ctx = {}
    if "watch" or "reel" in url:
        sanitized_url = sanitizing_url(url)
        if "www" in sanitized_url:
            sanitized_url = sanitized_url.replace("www", "mbasic", 1)
        else:
            sanitized_url = sanitized_url.replace("web", "mbasic", 1)
        response = requests.get(sanitized_url.replace("www", "mbasic"))

        if 'video_redirect' in response.text:
            reel = re.search(r'href\=\"\/video\_redirect\/\?src\=(.*?)\"', response.text)
            video_url = unquote(reel.group(1)).replace(";", "&")
            ctx["video_url"] = video_url
    else:
        # Send a GET request to the post URL and parse the HTML using BeautifulSoup
        response = requests.get(url)
        soup = BeautifulSoup(response.content, 'html.parser')

        # Extract the image URL from the meta tag with property="og:image"
        image_url = soup.find('meta', property='og:image')['content']
        ctx["image_url"] = image_url
    return render(request, "base.html", ctx)

