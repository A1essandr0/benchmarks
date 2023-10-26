from typing import Optional

from pydantic import BaseModel


class Event(BaseModel):
    source: Optional[str] = None
    event_name: Optional[str] = None
    event_status: Optional[str] = None
    created: Optional[str] = None
    payout: Optional[str] = None

