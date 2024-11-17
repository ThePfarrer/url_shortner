from hashlib import blake2b
from flask import Flask, jsonify, request

app = Flask(__name__)


links = [
    {
        "key": "blabla",
        "long_url": "https://codingchallenges.fyi/challenges/challenge-url-shortener/",
        "short_url": "http://localhost:5000/blabla",
    }
]


@app.route("/links", methods=["GET"])
def get_shorten_urls():
    return jsonify(links), 200


@app.route("/shortner", methods=["POST"])
def create_shorten_url():
    if not request.json:
        return jsonify({"error": "Request body must be JSON"}), 400

    new_link = request.json

    url = new_link.get("url")
    key = blake2b(url.encode(), digest_size=4).hexdigest()

    for link in links:
        if key in link["key"]:
            if url == link["long_url"]:
                return jsonify(link), 409

    item = {
        "key": f"{key}",
        "long_url": new_link["url"],
        "short_url": f"http://localhost:5000/{key}",
    }
    links.append(item)
    return jsonify(item), 201


if __name__ == "__main__":
    app.run(debug=True)
