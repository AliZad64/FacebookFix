from fastapi import FastAPI
from api.video import video
from api.image import image


app = FastAPI()


app.include_router(video)
app.include_router(image)