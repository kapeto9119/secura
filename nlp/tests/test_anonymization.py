import pytest
from app.services.anonymization import AnonymizationService


@pytest.fixture
def anonymization_service():
    return AnonymizationService()


async def test_empty_text(anonymization_service):
    """Test that empty text returns empty anonymized text and no entities."""
    anonymized_text, entities = await anonymization_service.anonymize("")
    assert anonymized_text == ""
    assert entities == []


async def test_no_pii(anonymization_service):
    """Test that text with no PII is returned unchanged."""
    text = "This text does not contain any personal information."
    anonymized_text, entities = await anonymization_service.anonymize(text)
    assert anonymized_text == text
    assert entities == []


async def test_basic_pii(anonymization_service):
    """Test that basic PII is properly anonymized."""
    text = "My name is John Smith and my email is john.smith@example.com."
    anonymized_text, entities = await anonymization_service.anonymize(text)
    
    # The anonymized text should not contain the original PII
    assert "John Smith" not in anonymized_text
    assert "john.smith@example.com" not in anonymized_text
    
    # Should have found at least 2 entities
    assert len(entities) >= 2
    
    # Check entity types
    entity_types = [entity.type for entity in entities]
    assert "PERSON" in entity_types
    assert "EMAIL_ADDRESS" in entity_types


async def test_phone_numbers(anonymization_service):
    """Test that phone numbers are properly anonymized."""
    text = "Call me at 555-123-4567 or (123) 456-7890."
    anonymized_text, entities = await anonymization_service.anonymize(text)
    
    # The anonymized text should not contain the original phone numbers
    assert "555-123-4567" not in anonymized_text
    assert "(123) 456-7890" not in anonymized_text
    
    # Should have found at least 2 entities
    assert len(entities) >= 2
    
    # Check entity types
    phone_entities = [entity for entity in entities if entity.type == "PHONE_NUMBER"]
    assert len(phone_entities) >= 1 