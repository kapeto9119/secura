import logging
from dataclasses import dataclass
from typing import List, Tuple
import spacy
from presidio_analyzer import AnalyzerEngine, RecognizerRegistry
from presidio_anonymizer import AnonymizerEngine
from presidio_anonymizer.entities import RecognizerResult

from app.core.config import settings

logger = logging.getLogger(__name__)

@dataclass
class Entity:
    type: str
    text: str
    start_offset: int
    end_offset: int


class AnonymizationService:
    def __init__(self):
        """Initialize the anonymization service with necessary models and components."""
        try:
            self.nlp = spacy.load(settings.SPACY_MODEL)
            logger.info(f"Loaded spaCy model: {settings.SPACY_MODEL}")
        except OSError:
            logger.warning(f"SpaCy model {settings.SPACY_MODEL} not found. Downloading...")
            spacy.cli.download(settings.SPACY_MODEL)
            self.nlp = spacy.load(settings.SPACY_MODEL)
            logger.info(f"Downloaded and loaded spaCy model: {settings.SPACY_MODEL}")
        
        # Initialize Microsoft Presidio components
        registry = RecognizerRegistry()
        self.analyzer = AnalyzerEngine(registry=registry)
        self.anonymizer = AnonymizerEngine()
        
        logger.info("Anonymization service initialized")

    async def anonymize(self, text: str) -> Tuple[str, List[Entity]]:
        """
        Anonymize sensitive information in text.
        
        Args:
            text: The input text to anonymize
            
        Returns:
            A tuple containing (anonymized_text, list_of_entities)
        """
        if not text:
            return "", []
            
        # Analyze the text with Presidio
        analyzer_results = self.analyzer.analyze(
            text=text,
            language="en",
            entities=[
                "PERSON", "EMAIL_ADDRESS", "PHONE_NUMBER", "CREDIT_CARD", 
                "US_SSN", "US_BANK_NUMBER", "LOCATION", "DATE_TIME",
                "NRP", "UK_NHS", "IP_ADDRESS", "US_DRIVER_LICENSE",
                "US_ITIN", "US_PASSPORT"
            ],
            score_threshold=0.5
        )
        
        # Anonymize the text
        anonymized_result = self.anonymizer.anonymize(
            text=text,
            analyzer_results=analyzer_results
        )
        
        # Get the anonymized text
        anonymized_text = anonymized_result.text
        
        # Convert the results to our Entity model
        entities = []
        for item in analyzer_results:
            entity = Entity(
                type=item.entity_type,
                text=text[item.start:item.end],
                start_offset=item.start,
                end_offset=item.end
            )
            entities.append(entity)
            
        return anonymized_text, entities


# Dependency for FastAPI
_anonymization_service = None

def get_anonymization_service() -> AnonymizationService:
    """
    Get or create the anonymization service instance.
    This function serves as a dependency for FastAPI.
    """
    global _anonymization_service
    if _anonymization_service is None:
        _anonymization_service = AnonymizationService()
    return _anonymization_service 