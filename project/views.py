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

embed_controller = Router(tags=["Embed Controller"])


# main endpoint
@embed_controller.get("", response={200: MessageOut, 400: MessageOut})
# @embed_controller.get("/reel/{id}", response={200: VideoOut, 400: MessageOut})
def get_video(request, url: str = None):
    print(request.path_info)
    print(request.get_full_path_info())
    print(request.get_full_path())

    ctx = {}
    # url = "https://www.facebook.com/reel/" + id
    print(url)
    if "watch" or "reel" in url:
        sanitized_url = sanitizing_url(url)
        if "www" in sanitized_url:
            sanitized_url = sanitized_url.replace("www.", "mbasic.", 1)
        else:
            sanitized_url = sanitized_url.replace("web.", "mbasic.", 1)
        print(sanitized_url)
        response = requests.get(sanitized_url)

        if 'video_redirect' in response.text:
            reel = re.search(r'href\=\"\/video\_redirect\/\?src\=(.*?)\"', response.text)
            video = unquote(reel.group(0)).replace(";", "&")
            ctx["video"] = mark_safe(video)
    else:
        # Send a GET request to the post URL and parse the HTML using BeautifulSoup
        response = requests.get(url)
        soup = BeautifulSoup(response.content, 'html.parser')

        # Extract the image URL from the meta tag with property="og:image"
        image_url = soup.find('meta', property='og:image')['content']
        ctx["image_url"] = image_url
    return render(request, "base.html", ctx)

