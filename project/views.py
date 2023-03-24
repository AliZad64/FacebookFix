import requests
from django.shortcuts import render
from bs4 import BeautifulSoup
from ninja import Router
from django.utils.safestring import mark_safe
from lxml import html
from config.utils.video import embed_video
from config.utils.consts import FACEBOOK_URL, FACEBOOK_QUERY_URL, headers



video = Router(tags=["Video Routes"])
image = Router(tags=["Image Routes"])


# ----------------- video routes ----------------- #
@video.get("watch/")
@video.get("watch")
def get_video_by_query_param(request, v: str):
    url = FACEBOOK_QUERY_URL + v
    return render(request, "base.html", embed_video(url))


@video.get("watch/{link_id}")
@video.get("watch/{link_id}/")
@video.get("reel/{link_id}")
@video.get("reel/{link_id}/")
def get_video_by_id(request, link_id: str):
    url = FACEBOOK_URL + mark_safe(request.path)
    result = embed_video(url)
    if "reel" in request.path:
        result["width"] = "640"
        result["height"] = "360"
    return render(request, "base.html", result)


@video.get("{user}/video/{link_id}")
@video.get("{user}/video/{link_id}/")
@video.get("{user}/videos/{link_id}")
@video.get("{user}/videos/{link_id}/")
def get_video_by_user(request, user: str, link_id: str):
    url = FACEBOOK_URL + request.path
    return render(request, "base.html", embed_video(url))


# ----------------- image routes ----------------- #
@image.get("{user}/photos/{a_id}/{link_id}")
def get_image(request, user: str, a_id: str, link_id: str):
    ctx = {}
    meta = {}
    url = f"https://m.facebook.com/{user}/photos/{a_id}/{link_id}"
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


@image.get("/photo")
def get_image(request, fbid: str = None):
    ctx = {}
    meta = {}
    url = f"https://m.facebook.com/photo.php?fbid={fbid}"
    response = requests.get(url, headers=headers)
    tree = html.fromstring(response.content)
    for tag in tree.xpath('//meta'):
        meta[tag.get('property')] = mark_safe(tag.get('content'))
    # assign every meta tag to the context
    for key, value in meta.items():
        if key:
            ctx[key[3:]] = value
    return render(request, "base.html", ctx)
    # # Send a GET request to the post URL and parse the HTML using BeautifulSoup
    # response = requests.get(url)
    # soup = BeautifulSoup(response.content, 'html.parser')

    # # Extract the image URL from the meta tag with property="og:image"
    # image_url = soup.find('meta', property='og:image')['content']
    # print(image_url)
    # ctx["image_url"] = image_url