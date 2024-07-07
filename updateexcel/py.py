import requests
import json

url = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"
payload = json.dumps({
	"app_id": "cli_a603ea18833ad00e",
	"app_secret": "oTJQknTuwUUywRvmtDyv6edKD3aCnxPi"
})


headers = {
  'Content-Type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)
print(response.text)