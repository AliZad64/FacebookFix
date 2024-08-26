import json

from fastapi import APIRouter, Request
import httpx
from lxml import html
from html import unescape
from utils.constants import headers
from template import templates
from api.app import app
async def embed_image(url: str, id: str) -> dict:
    meta = {}
    ctx = {}
    # redi = app.state.redis
    # result = await redi.get(url)
    # if result:
    #     return json.loads(result)
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
    # await redi.set(id, json.dumps(ctx))
    return ctx


image = APIRouter()


@image.get("/{user}/photos/{a_id}/{link_id}")
async def get_image(request: Request, user: str, a_id: str, link_id: str):
    url = f"https://m.facebook.com/{user}/photos/{a_id}/{link_id}"
    data = await embed_image(url, link_id)
    return templates.TemplateResponse("base.html", {"request": request, "data": data})


@image.get("/photo")
@image.get("/photo/")
async def get_image(request: Request, fbid: str = None):
    url = f"https://m.facebook.com/photo/?fbid={fbid}"
    data = await embed_image(url, fbid)
    return templates.TemplateResponse("base.html", {"request": request, "data": data})