import json
import os
import requests
from django.http import JsonResponse
from django.views.decorators.csrf import csrf_exempt

EXECUTOR_URL = os.getenv("EXECUTOR_URL", "http://executor:8081")
@csrf_exempt
def run_code(request):
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

            return JsonResponse(response.json())

        except requests.exceptions.RequestException as e:
            return JsonResponse({
                "error": "Executor service unavailable",
                "details": str(e)
            }, status=500)