from django.shortcuts import render
from ninja import Router
from django.utils.safestring import mark_safe
from config.utils.embed import embed_video, embed_image
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
    url = f"https://m.facebook.com/{user}/photos/{a_id}/{link_id}"
    return render(request, "base.html", embed_image(url))


@image.get("/photo")
def get_image(request, fbid: str = None):
    url = f"https://m.facebook.com/photo.php?fbid={fbid}"
    return render(request, "base.html", embed_image(url))