import logging
from fastapi import APIRouter, Depends, HTTPException
from pydantic import BaseModel

from app.services.anonymization import AnonymizationService, get_anonymization_service

router = APIRouter()
logger = logging.getLogger(__name__)

class AnonymizeRequest(BaseModel):
    text: str

class EntityModel(BaseModel):
    type: str
    text: str
    start_offset: int
    end_offset: int

class AnonymizeResponse(BaseModel):
    anonymized_text: str
    entities: list[EntityModel]

@router.post("", response_model=AnonymizeResponse)
async def anonymize_text(
    request: AnonymizeRequest,
    anonymization_service: AnonymizationService = Depends(get_anonymization_service)
):
    """
    Anonymize sensitive information in text.
    
    This endpoint detects and masks personally identifiable information (PII)
    in the provided text, replacing it with appropriate tokens.
    """
    try:
        logger.info("Processing anonymization request")
        anonymized_text, entities = await anonymization_service.anonymize(request.text)
        
        # Convert entities to response model
        response_entities = [
            EntityModel(
                type=entity.type,
                text=entity.text,
                start_offset=entity.start_offset,
                end_offset=entity.end_offset
            )
            for entity in entities
        ]
        
        logger.info(f"Anonymization completed. Found {len(entities)} entities.")
        return AnonymizeResponse(
            anonymized_text=anonymized_text,
            entities=response_entities
        )
    except Exception as e:
        logger.error(f"Error during anonymization: {str(e)}")
        raise HTTPException(
            status_code=500,
            detail=f"Error during anonymization: {str(e)}"
        ) 