import logging
import requests
from django.shortcuts import render
from bs4 import BeautifulSoup
from ninja import Router
from django.utils.safestring import mark_safe
from lxml import html
from config.utils.video import embed_video
from config.utils.consts import FACEBOOK_URL, FACEBOOK_QUERY_URL, headers

logger = logging.getLogger(__name__)


embed_controller = Router(tags=["Embed Controller"])


# ----------------- video routes ----------------- #
@embed_controller.get("watch/")
@embed_controller.get("watch")
def get_video_by_query_param(request, v: str):
    url = FACEBOOK_QUERY_URL + v
    return render(request, "base.html", embed_video(url))


@embed_controller.get("watch/{link_id}")
@embed_controller.get("watch/{link_id}/")
@embed_controller.get("reel/{link_id}")
@embed_controller.get("reel/{link_id}/")
def get_video_by_id(request, link_id: str):
    url = FACEBOOK_URL + mark_safe(request.path)
    result = embed_video(url)
    if "reel" in request.path:
        result["width"] = "640"
        result["height"] = "360"
    return render(request, "base.html", result)


@embed_controller.get("{user}/video/{link_id}")
@embed_controller.get("{user}/video/{link_id}/")
@embed_controller.get("{user}/videos/{link_id}")
@embed_controller.get("{user}/videos/{link_id}/")
def get_video_by_user(request, user: str, link_id: str):
    url = FACEBOOK_URL + request.path
    return render(request, "base.html", embed_video(url))


# ----------------- image routes ----------------- #
@embed_controller.get("{user}/photos/{a_id}/{link_id}")
def get_image(request, user: str, a_id: str, link_id: str):
    url = f"https://www.facebook.com/{user}/photos/{a_id}/{link_id}"
    # Send a GET request to the post URL and parse the HTML using BeautifulSoup
    response = requests.get(url, headers=headers)
    tree = html.fromstring(response.content)
    meta = {
        tag.get('property'): mark_safe(tag.get('content'))
        for tag in tree.xpath('//meta')
    }
    ctx = {key[3:]: value for key, value in meta.items() if key}
    return render(request, "base.html", ctx)


@embed_controller.get("/photo")
def get_image(request, fbid: str = None):
    url = f"https://www.facebook.com/photo.php?fbid={fbid}"
    # Send a GET request to the post URL and parse the HTML using BeautifulSoup
    response = requests.get(url)
    soup = BeautifulSoup(response.content, 'html.parser')

    # Extract the image URL from the meta tag with property="og:image"
    image_url = soup.find('meta', property='og:image')['content']
    print(image_url)
    ctx = {"image_url": image_url}
    return render(request, "base.html", ctx)
