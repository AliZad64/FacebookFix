import logging
import requests
from django.shortcuts import render
from bs4 import BeautifulSoup
from ninja import Router
from django.utils.safestring import mark_safe
from lxml import html
from config.utils.video import embed_video

logger = logging.getLogger(__name__)

FACEBOOK_URL = "https://www.facebook.com"
FACEBOOK_QUERY_URL = "https://www.facebook.com/watch/?v="
embed_controller = Router(tags=["Embed Controller"])

headers = {
    "User-Agent": "facebookexternalhit/1.1",
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
    "Accept-Language": "en-us,en;q=0.5",
    "Sec-Fetch-Mode": "navigate",
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
@embed_controller.get("watch/")
@embed_controller.get("watch")
def get_video_by_query_param(request, v: str):
    url = FACEBOOK_QUERY_URL + v
    return render(request, "base.html", embed_video(url))

@embed_controller.get("watch/{link_id}")
@embed_controller.get("reel/{link_id}")
def get_video_by_id(request, link_id: str,  v: str = None):
    print(request.path)
    if v:
        url = f"https://www.facebook.com/watch?v={link_id}"
    else:
        url = f"https://www.facebook.com" + mark_safe(request.path)
    return render(request, "base.html", embed_video(url))



@embed_controller.get("watch/{link_id}")
@embed_controller.get("watch/{link_id}/")
@embed_controller.get("reel/{link_id}")
def get_video_by_id(request, link_id: str):
    url = FACEBOOK_URL + mark_safe(request.path)
    return render(request, "base.html", embed_video(url))

@embed_controller.get("watch/{link_id}")
@embed_controller.get("reel/{link_id}")
def get_video_by_id(request, link_id: str,  v: str = None):
    print(request.path)
    if v:
        url = f"https://www.facebook.com/watch?v={link_id}"
    else:
        url = f"https://www.facebook.com" + mark_safe(request.path)
    return render(request, "base.html", embed_video(url))



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
