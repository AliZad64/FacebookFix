import asyncio
import json
from fastapi import  APIRouter, Request
from yt_dlp import YoutubeDL
from utils.constants import FACEBOOK_URL, FACEBOOK_QUERY_URL
from template import templates
from api.app import app

async def embed_video(url: str, id: str) -> dict:
    ctx = {}
    redi = app.state.redis
    result = await redi.get(id)
    if result:
        return json.loads(result)
    with YoutubeDL() as ydl:
        result = await asyncio.to_thread(ydl.extract_info, url, False)
        for video_format in result["formats"]:
            if video_format["format_id"] == "hd":
                result["video"] = (video_format["url"])
            elif video_format["format_id"] == "sd":
                result["video"] = (video_format["url"])
        ctx = {
            "title": result["title"],
            "description": result["description"],
            "width": result["width"],
            "height": result["height"],
            "url": url,
            "card": "player",
            "video": result["video"],
        }
        await redi.set(id, json.dumps(ctx))
    return result


video = APIRouter()


@video.get("/watch/")
@video.get("/watch")
async def get_video_by_query_param(request: Request, v: str):
    url = FACEBOOK_QUERY_URL + v
    data = await embed_video(url, v)
    return templates.TemplateResponse("base.html", {"request": request, "data": data})


@video.get("/watch/{link_id}")
@video.get("/reel/{link_id}")
async def get_video_by_id(request: Request, link_id: str):
    url = FACEBOOK_URL + request.url.path
    result = await embed_video(url, link_id)
    if "reel" in request.url.path:
        result["width"] = "640"
        result["height"] = "360"
    return templates.TemplateResponse("base.html", {"request": request, "data": result})


@video.get("/{user}/video/{link_id}")
@video.get("/{user}/videos/{link_id}")
async def get_video_by_user(request: Request, user: str, link_id: str):
    url = FACEBOOK_URL + request.url.path
    data = await embed_video(url, link_id)
    return templates.TemplateResponse("base.html", {"request": request, "data": data})
