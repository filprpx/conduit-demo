from fastapi import FastAPI


app = FastAPI(title="Conduit Demo API")


@app.get("/")
def health():
    return {"application": "Conduit demo API", "status": "ok"}
