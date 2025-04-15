from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
import spacy
import os

# Load spaCy model
try:
    nlp = spacy.load("en_core_web_sm")
except:
    nlp = spacy.load("en_core_web_lg")

app = FastAPI(
    title="Secura NLP Service",
    description="API for text anonymization and processing",
    version="0.1.0"
)

class TextRequest(BaseModel):
    text: str

class TextResponse(BaseModel):
    anonymized_text: str

@app.get("/")
def read_root():
    return {"status": "ok", "service": "Secura NLP Service"}

@app.get("/health")
def health_check():
    return {
        "status": "ok",
        "version": "0.1.0",
        "environment": os.getenv("ENVIRONMENT", "development")
    }

@app.post("/anonymize", response_model=TextResponse)
def anonymize_text(request: TextRequest):
    if not request.text:
        raise HTTPException(status_code=400, detail="Empty text provided")
    
    # Process the text with spaCy
    doc = nlp(request.text)
    
    # Anonymize sensitive entities
    anonymized = request.text
    for ent in reversed(doc.ents):
        if ent.label_ in ["PERSON", "ORG", "GPE", "LOC", "MONEY", "CARDINAL", "DATE"]:
            anonymized = anonymized[:ent.start_char] + f"[{ent.label_}]" + anonymized[ent.end_char:]
    
    return {"anonymized_text": anonymized}

if __name__ == "__main__":
    import uvicorn
    uvicorn.run("main:app", host="0.0.0.0", port=8000, reload=True) 