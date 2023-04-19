from fastapi import APIRouter
import redis
home = APIRouter()

@home.get("/")
def main_page():
    return 200, {"message": "Hello World"}

@home.get("/redis")
def get_key():
    r = redis.Redis()
    val = r.get("key")
    return 200, {"message": val}
    