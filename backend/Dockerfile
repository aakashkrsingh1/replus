FROM python:3.10

WORKDIR /app

COPY . .

RUN pip install django requests

CMD ["python", "manage.py", "runserver", "0.0.0.0:8000"]