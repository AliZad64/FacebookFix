from fastapi import  APIRouter, Request
from yt_dlp import YoutubeDL
from utils.constants import FACEBOOK_URL, FACEBOOK_QUERY_URL
from template import templates

def embed_video(url: str) -> dict:
    try:
        with YoutubeDL() as ydl:
            result = ydl.extract_info(url, download=False)
            for video_format in result["formats"]:
                if video_format["format_id"] == "hd":
                    result["video"] = (video_format["url"])
        result["url"] = url
        result["card"] = "player"
    except:
        result = {"url": url}
    return result


video = APIRouter()


@video.get("/watch/")
@video.get("/watch")
def get_video_by_query_param(request: Request, v: str):
    url = FACEBOOK_QUERY_URL + v
    return templates.TemplateResponse("base.html", {"request": request, "data": (embed_video(url))})


@video.get("/watch/{link_id}")
@video.get("/reel/{link_id}")
def get_video_by_id(request: Request, link_id: str):
    url = FACEBOOK_URL + request.url.path
    result = embed_video(url)
    if "reel" in request.url.path:
        result["width"] = "640"
        result["height"] = "360"
    return templates.TemplateResponse("base.html", {"request": request, "data": result})


@video.get("{user}/video/{link_id}")
@video.get("{user}/videos/{link_id}")
def get_video_by_user(request: Request, user: str, link_id: str):
    url = FACEBOOK_URL + request.url.path
    print(url)
    return templates.TemplateResponse("base.html", {"request": request, "data": embed_video(url)})
