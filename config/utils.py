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