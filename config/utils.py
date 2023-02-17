facebook_url = "https://www.facebook.com/watch/"

def sanitizing_url(url: str) -> str:
    url_split = url.split("/")
    id = url_split[len(url_split)-1]
    if id[0] != "?":
        return facebook_url + "?v=" + id
    if len(id) > 19:
        return facebook_url + id[:19]
    return url