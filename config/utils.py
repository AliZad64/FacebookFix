facebook_url = "https://www.facebook.com/watch/"

def sanitizing_url(url: str) -> str:
    url_split = url.split("/")
    if url_split[len(url_split)-1] == "":
        id = url_split[len(url_split)-2]
    else:
        id = url_split[len(url_split)-1]
    if id[0] != "?":
        return facebook_url + "?v=" + id
    if len(id) > 19 and id[0] != "w":
        return facebook_url + id[:19]
    return url