FROM python:3.10-alpine3.21

WORKDIR /app

COPY . .

RUN pip install --no-cache-dir -r requirements.txt

EXPOSE 3333

CMD ["python", "main.py"]