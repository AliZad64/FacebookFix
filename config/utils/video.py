from django.utils.safestring import mark_safe
from yt_dlp import YoutubeDL



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