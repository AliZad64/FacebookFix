from urllib.parse import unquote

from django.shortcuts import render
import requests
import re
from ninja import Router
from config.utils import sanitizing_url
from project.schemas.video_schema import VideoOut, MessageOut

embed_controller = Router(tags=["Embed Controller"])


# main endpoint
@embed_controller.get("", response={200: VideoOut, 400: MessageOut})
def get_video(request, url: str = None):
    sanitized_url = sanitizing_url(url)
    if "www" in sanitized_url:
        sanitized_url = sanitized_url.replace("www", "mbasic", 1)
    else:
        sanitized_url = sanitized_url.replace("web", "mbasic", 1)
    response = requests.get(sanitized_url.replace("www", "mbasic"))

    if 'video_redirect' in response.text:
        reel = re.search(r'href\=\"\/video\_redirect\/\?src\=(.*?)\"', response.text)
        video_url = unquote(reel.group(1)).replace(";", "&")
        return {"video_url": video_url}
    else:
        return {"message": "Video not found"}
