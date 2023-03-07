from urllib.parse import unquote
import re
from django.utils.safestring import mark_safe
from lxml import html

facebook_url = "https://www.facebook.com/watch/"


def get_url_id(url: str) -> str:
    url_split = url.split("/")
    if url_split[len(url_split) - 1] == "":
        return url_split[len(url_split) - 2]
    return url_split[len(url_split) - 1]


def sanitizing_url(url: str) -> str:
    id = get_url_id(url)
    if len(id) > 19 and id[0] != "w":
        return facebook_url + id[:19]
    if len(id) > 19 and id[0] == "w":
        return facebook_url + id[6:]
    if id[0] != "?":
        return facebook_url + "?v=" + id
    return url


def embed_video(response) -> dict:
    ctx = {}
    meta = {}
    if 'video_redirect' in response.text:
        reel = re.search(r'href\=\"\/video\_redirect\/\?src\=(.*?)\"', response.text)
        video = unquote(reel.group(1)).replace(";", "&")
        ctx["video"] = mark_safe(video)
        tree = html.fromstring(response.content)
        for tag in tree.xpath('//meta'):
            meta[tag.get('property')] = mark_safe(tag.get('content'))
        # assign every meta tag to the context
    for key, value in meta.items():
        if key and key != 'og:image':
            ctx[key[3:]] = value
    return ctx
