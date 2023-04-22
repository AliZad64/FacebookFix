from api.app import app
from api.video import video
from api.image import image
from api.home import home


app.include_router(home)
app.include_router(video)
app.include_router(image)