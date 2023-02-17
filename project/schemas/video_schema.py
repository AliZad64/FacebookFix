from ninja import Schema


class MessageOut(Schema):
    message: str


class VideoOut(Schema):
    video_url: str
