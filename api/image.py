from fastapi import APIRouter, Request
import requests
from lxml import html
from html import unescape
from utils.constants import headers
from template import templates


def embed_image(url: str) -> dict:
    meta = {}
    ctx = {}
    response = requests.get(url, headers=headers)
    tree = html.fromstring(response.content)
    for tag in tree.xpath('//meta'):
        meta[tag.get('property')] = unescape(tag.get('content'))
    # assign every meta tag to the context
    for key, value in meta.items():
        if key:
            ctx[key[3:]] = value
    ctx["url"] = url
    ctx["card"] = "summary_large_image"
    return ctx


image = APIRouter()


@image.get("/{user}/photos/{a_id}/{link_id}")
def get_image(request: Request, user: str, a_id: str, link_id: str):
    url = f"https://m.facebook.com/{user}/photos/{a_id}/{link_id}"
    return templates.TemplateResponse("base.html",{"request": request, "data": embed_image(url)})


@image.get("/photo")
@image.get("/photo/")
def get_image(request: Request, fbid: str = None):
    url = f"https://m.facebook.com/photo/?fbid={fbid}"
    return templates.TemplateResponse("base.html", {"request": request, "data": embed_image(url)})