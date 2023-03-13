from typing import Optional, Any
from urllib.parse import unquote

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


FACEBOOK_URL = "https://mbasic.facebook.com/watch/?v="
embed_controller = Router(tags=["Embed Controller"])


# ----------------- video routes ----------------- #
@embed_controller.get("/watch")
def get_video(request, v: str = None):
    url = FACEBOOK_URL + v
    logging.info("URL: ")
    logging.info(url)
    response = requests.get(url)
    logging.info("response: ")
    logging.info(response)
    return render(request, "base.html", embed_video(response))


@embed_controller.get("watch/{link_id}")
@embed_controller.get("reel/{link_id}")
def get_video_by_id(request, link_id: str, v: str = None):
    if v:
        sanitized_url = FACEBOOK_URL + v
    else:
        sanitized_url = sanitizing_url(f"https://mbasic.facebook.com/watch/{link_id}")
    response = requests.get(sanitized_url)
    return render(request, "base.html", embed_video(response))


@embed_controller.get("{user}/video/{link_id}")
def get_video_by_user(request, user: str, link_id: str):
    sanitized_url = sanitizing_url(f"https://mbasic.facebook.com/{user}/video/{link_id}")
    response = requests.get(sanitized_url)
    return render(request, "base.html", embed_video(response))


# ----------------- image routes ----------------- #
@embed_controller.get("photo/{user}/photos/{a_id}/{link_id}")
def get_image(request, user: str, a_id: str, link_id: str):
    ctx = {}
    url = f"https://www.facebook.com/{user}/photos/{a_id}/{link_id}"
    # Send a GET request to the post URL and parse the HTML using BeautifulSoup
    response = requests.get(url)
    soup = BeautifulSoup(response.content, 'html.parser')

    # Extract the image URL from the meta tag with property="og:image"
    image_url = soup.find('meta', property='og:image')['content']
    ctx["image_url"] = image_url
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
    ctx["image_url"] = image_url
    return render(request, "base.html", ctx)


@embed_controller.get("test", response={200: VideoOut})
def test_func(request):
    print("heya")
    return 200, {"video_url": "test"}

@embed_controller.get("page")
def test_page(request):
    ctx = {}
    ctx["url"] = "www.facebook.com"
    return render(request, "test.html", ctx)