from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI()

class PaymentConfirmation(BaseModel):
    order_id: str
    amount: int
    status: str

@app.post("/confirm-payment")
def confirm_payment(data: PaymentConfirmation):
    print(f"Received payment confirmation: {data}")

    # Simulate decision
    if data.amount >= 10000:
        status = "APPROVED"
    else:
        status = "REJECTED"

    return {"status": status}
