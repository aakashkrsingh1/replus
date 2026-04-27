import json
import os
import requests
from django.http import JsonResponse
from django.views.decorators.csrf import csrf_exempt
import time

EXECUTOR_URL = os.getenv("EXECUTOR_URL", "http://executor:8081")
@csrf_exempt
def run_code(request):
    start=time.time()
    if request.method == 'OPTIONS':
        return JsonResponse({}, status=200)
        
    if request.method == 'POST':
        data = json.loads(request.body)

        try:
            response = requests.post(
                f"{EXECUTOR_URL}/execute",
                json={
                    "code": data.get("code"),
                    "language": data.get("language")
                },
                timeout=(3,15)
            )
            end= time.time()

            executor_output= response.json();
            return JsonResponse({
                    "output": executor_output.get("output", ""),
                    "time": f"{(end - start)*1000:.2f} ms",
                    "status": "success" if response.status_code == 200 else "error"
                })

        except requests.exceptions.RequestException as e:
            return JsonResponse({
                "error": "Executor service unavailable",
                "details": str(e)
            }, status=500)