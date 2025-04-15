import os
from pydantic import BaseSettings
from dotenv import load_dotenv

# Load environment variables from .env file
load_dotenv()

class Settings(BaseSettings):
    # API settings
    API_V1_STR: str = "/api/v1"
    PROJECT_NAME: str = "Secura Anonymization Service"
    
    # Environment settings
    ENVIRONMENT: str = os.getenv("ENVIRONMENT", "development")
    DEBUG: bool = ENVIRONMENT == "development"
    
    # NLP model settings
    SPACY_MODEL: str = os.getenv("SPACY_MODEL", "en_core_web_lg")
    
    # Logging settings
    LOG_LEVEL: str = os.getenv("LOG_LEVEL", "info")
    
    class Config:
        case_sensitive = True


settings = Settings() 