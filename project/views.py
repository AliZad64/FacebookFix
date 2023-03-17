from typing import Optional, Any
from urllib.parse import unquote
import logging
from django.http import HttpRequest, HttpResponse

from django.shortcuts import render
import requests
import re
import logging
from bs4 import BeautifulSoup
from ninja import Router
from config.utils import sanitizing_url, embed_video
from project.schemas.video_schema import VideoOut, MessageOut
from django.utils.safestring import mark_safe
from lxml import html
from yt_dlp import YoutubeDL
logger = logging.getLogger(__name__)

FACEBOOK_URL = "https://www.facebook.com/watch/?v="
embed_controller = Router(tags=["Embed Controller"])

headers = {
    "User-Agent":"facebookexternalhit/1.1",
    "Accept":"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
    "Accept-Language":"en-us,en;q=0.5",
    "Sec-Fetch-Mode":"navigate",
    "authority": "www.facebook.com",
    "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
    "accept-language": "en-US,en;q=0.9",
    "cache-control": "max-age=0",
    "sec-fetch-mode": "navigate",
    "upgrade-insecure-requests": "1",
    "referer": "https://www.facebook.com/",
    "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
    "viewport-width": "1280"
}
# ----------------- video routes ----------------- #
@embed_controller.get("/watch")
def get_video(request, v: str = None):
    print(request.headers.get("User-Agent"))
    url = FACEBOOK_URL + v
    logging.info("URL: ")
    logging.info(url)
    response = requests.get(url, headers=headers)
    logger.info(response)
    logger.info(response.content)
    logger.info("hey whats up")
    logging.info("response: ")
    logging.info(response)
    return render(request, "base.html", embed_video(response))


@embed_controller.get("watch/{link_id}")
@embed_controller.get("reel/{link_id}")
def get_video_by_id(request, link_id: str, v: str = None):
    if v:
        url = FACEBOOK_URL + v
    else:
        url = f"https://www.facebook.com" + mark_safe(request.path)
    with YoutubeDL() as ydl:
        result = ydl.extract_info(url, download=False)
        for format in result["formats"]:
            if format["format_id"] == "hd":
                result["video"] = mark_safe(format["url"])
    result["url"] = url
    return render(request, "base.html", result)


@embed_controller.get("{user}/video/{link_id}")
def get_video_by_user(request, user: str, link_id: str):
    sanitized_url = sanitizing_url(f"https://mbasic.facebook.com/{user}/video/{link_id}")
    response = requests.get(sanitized_url)
    return render(request, "base.html", embed_video(response))


# ----------------- image routes ----------------- #
@embed_controller.get("{user}/photos/{a_id}/{link_id}")
def get_image(request, user: str, a_id: str, link_id: str):
    ctx = {}
    meta = {}
    url = f"https://www.facebook.com/{user}/photos/{a_id}/{link_id}"
    # Send a GET request to the post URL and parse the HTML using BeautifulSoup
    response = requests.get(url, headers=headers)
    tree = html.fromstring(response.content)
    for tag in tree.xpath('//meta'):
            meta[tag.get('property')] = mark_safe(tag.get('content'))
        # assign every meta tag to the context
    for key, value in meta.items():
        if key:
            ctx[key[3:]] = value
    return render(request, "base.html", ctx)

@embed_controller.get("/photo")
def get_image(request, fbid: str = None):
    ctx = {}
    url = f"https://www.facebook.com/photo.php?fbid={fbid}"
    # Send a GET request to the post URL and parse the HTML using BeautifulSoup
    response = requests.get(url)
    soup = BeautifulSoup(response.content, 'html.parser')

    # Extract the image URL from the meta tag with property="og:image"
    image_url = soup.find('meta', property='og:image')['content']
    print(image_url)
    ctx["image_url"] = image_url
    return render(request, "base.html", ctx)


@embed_controller.get("test", response={200: VideoOut})
def test_func(request):
    print("heya")
    return 200, {"video_url": "test"}

@embed_controller.get("page")
def test_page(request):
    response = requests.get("https://www.instagram.com/p/CnonizJpAoV/embed/captioned/")
    
    ctx = {}
    ctx["url"] = response.text
    return render(request, "test.html", ctx)

@embed_controller.get("test2")
def http_response(request: HttpRequest, response: HttpResponse):
    print(request.headers.get("User-Agent"))
    response.set_cookie("cookie", "delicious")
    # Set a header.
    response["X-Cookiemonster"] = "blue"
    return render(request, "test.html", {})

@embed_controller.get("test_facebook", response={200: VideoOut})
def test_facebook(request):
    history = ""
    url = FACEBOOK_URL + '1774563489582348'
    response = requests.get(url, headers=headers)
    for resp in response.history:
        history += resp.url + resp.text
    return 200 , {"video_url": response.text, "user_agent": history}

@embed_controller.get("test_instagram")
def test_instagram(request):
    url = "https://www.instagram.com/p/CpJeHYcLCVP/embed/captioned/"
    response = requests.get(url)
    return 200 , {"video_url": response.text}

@embed_controller.get("test_normal")
def test_normal(request):
    url = "https://www.facebook.com/watch/?v=1774563489582348"
    response = requests.get(url)
    return 200 , {"video_url": response.text}

@embed_controller.get("example")
def facebook_example(request):
    return render(request, "example.html")