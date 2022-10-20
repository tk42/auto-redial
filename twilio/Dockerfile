FROM tiangolo/uwsgi-nginx-flask:latest

WORKDIR /app
COPY . .
RUN pip install -r requirements.txt
EXPOSE 5050
CMD ["python", "main.py"]
