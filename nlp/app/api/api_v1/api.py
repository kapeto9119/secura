from fastapi import APIRouter

from app.api.api_v1.endpoints import anonymize

api_router = APIRouter()
api_router.include_router(anonymize.router, prefix="/anonymize", tags=["anonymization"]) 