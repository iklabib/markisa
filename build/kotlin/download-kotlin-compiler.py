import os
from pathlib import Path
import sys
import json
import hashlib
import urllib.request

REPO = "JetBrains/kotlin"

api_url = f"https://api.github.com/repos/{REPO}/releases/latest"
response = urllib.request.urlopen(api_url)
release_info = json.loads(response.read())

latest_version: str = release_info["tag_name"]
print(f"Latest Kotlin version: {latest_version}")

# target = f"kotlin-native-linux-x86_64-{latest_version.lstrip('v')}.tar.gz"
target = f"kotlin-compiler-{latest_version.lstrip('v')}.zip"
asset = None
for item in release_info["assets"]:
    if item["name"] == target:
        asset = item
        break
else:
    print("No matching asset found")
    sys.exit()


os.makedirs("artifacts", exist_ok=True)

download_url = asset["browser_download_url"]
filename = asset["name"]
file_path = Path("artifacts") / filename
sha256_expected = None

hash_asset = None
for item in release_info["assets"]:
    if item["name"] == filename + ".sha256":
        hash_asset = item
        break
else:
    print("No hash asset found")
    sys.exit()

resp = urllib.request.urlopen(hash_asset["browser_download_url"])
sha256_expected = resp.read().decode()

os.system(f"aria2c --checksum=sha-256={sha256_expected} -c -j 8 -s 8 -x 8 --dir=artifacts {download_url}")