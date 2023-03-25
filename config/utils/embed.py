import requests
from config.utils.consts import headers
from django.utils.safestring import mark_safe
from yt_dlp import YoutubeDL
from lxml import html


def embed_video(url: str) -> dict:
    try:
        with YoutubeDL() as ydl:
            result = ydl.extract_info(url, download=False)
            for video_format in result["formats"]:
                if video_format["format_id"] == "hd":
                    result["video"] = mark_safe(video_format["url"])
        result["url"] = url
        result["card"] = "player"
    except:
        result = {"url": url}
    return result


def embed_image(url: str) -> dict:
    meta = {}
    ctx = {}
    response = requests.get(url, headers=headers)
    tree = html.fromstring(response.content)
    for tag in tree.xpath('//meta'):
        meta[tag.get('property')] = mark_safe(tag.get('content'))
    # assign every meta tag to the context
    for key, value in meta.items():
        if key:
            ctx[key[3:]] = value
    return ctx
