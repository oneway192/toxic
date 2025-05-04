from fastapi import FastAPI, Request
from pydantic import BaseModel
from transformers import AutoTokenizer, T5ForConditionalGeneration
import torch

app = FastAPI()

tokenizer = AutoTokenizer.from_pretrained("ai-forever/ruT5-base")
model = T5ForConditionalGeneration.from_pretrained("s-nlp/ruT5-base-detox")

class TextIn(BaseModel):
    text: str

@app.get("/")
def read_root():
    return {"message": "API работает!"}

@app.post("/detox")
def detox_text(data: TextIn):
    input_ids = tokenizer.encode(data.text, return_tensors="pt")
    with torch.no_grad():
        output_ids = model.generate(input_ids, max_length=50)
    output_text = tokenizer.decode(output_ids[0], skip_special_tokens=True)
    return {"detoxed_text": output_text}