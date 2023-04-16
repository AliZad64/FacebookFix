from fastapi import APIRouter, Request
import httpx
from lxml import html
from html import unescape
from utils.constants import headers
from template import templates


async def embed_image(url: str) -> dict:
    meta = {}
    ctx = {}
    async with httpx.AsyncClient() as client:
        response = await client.get(url, headers=headers)
        response_content = response.content

    tree = html.fromstring(response_content)
    for tag in tree.xpath('//meta'):
        meta[tag.get('property')] = tag.get('content')
    # assign every meta tag to the context
    for key, value in meta.items():
        if key:
            ctx[key[3:]] = value
    ctx["url"] = url
    ctx["card"] = "summary_large_image"
    return ctx


image = APIRouter()


@image.get("/{user}/photos/{a_id}/{link_id}")
async def get_image(request: Request, user: str, a_id: str, link_id: str):
    url = f"https://m.facebook.com/{user}/photos/{a_id}/{link_id}"
    data = await embed_image(url)
    return templates.TemplateResponse("base.html", {"request": request, "data": data})


@image.get("/photo")
@image.get("/photo/")
async def get_image(request: Request, fbid: str = None):
    url = f"https://m.facebook.com/photo/?fbid={fbid}"
    data = await embed_image(url)
    return templates.TemplateResponse("base.html", {"request": request, "data": data})