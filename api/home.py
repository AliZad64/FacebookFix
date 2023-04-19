from fastapi import APIRouter
import redis
home = APIRouter()

@home.get("/")
def main_page():
    return 200, {"message": "Hello World"}

    