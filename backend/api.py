from hashlib import blake2b
import random
import string
import click
from flask import Flask, jsonify, request
from flask_sqlalchemy import SQLAlchemy
from sqlalchemy import text
import sqlite3

app = Flask(__name__)

db_name = "test_db.db"

app.config["SQLALCHEMY_DATABASE_URI"] = f"sqlite:///{db_name}"
app.config["SQLALCHEMY_TRACK_MODIFATIONS"] = False

db = SQLAlchemy()

db.init_app(app)


class URL(db.Model):
    __tablename__ = "urls"
    id = db.Column(db.Integer, primary_key=True)
    url_hash = db.Column(db.String, unique=True)
    long_url = db.Column(db.String)
    short_url = db.Column(db.String)


with app.app_context():
    db.create_all()


links = [
    # {
    #     "key": "blabla",
    #     "long_url": "https://codingchallenges.fyi/challenges/challenge-url-shortener/",
    #     "short_url": "http://localhost:5000/blabla",
    # }
]


@app.route("/")
def testdb():
    for link in links:
        url = URL(
            url_hash=link["key"], long_url=link["long_url"], short_url=link["short_url"]
        )
        db.session.add(url)
    db.session.commit()

    return "<h1>It works.</h1>"


@app.route("/links", methods=["GET"])
def get_shorten_urls():
    urls = URL.query.all()

    for url in urls:
        print(url.url_hash)
    return jsonify(links), 200


@app.route("/shortner", methods=["POST"])
def create_shorten_url():
    if not request.json:
        return jsonify({"error": "Request body must be JSON"}), 400

    new_link = request.json

    url = new_link.get("url")
    key = blake2b(url.encode(), digest_size=4).hexdigest()

    url_hash = URL.query.filter_by(url_hash=key).first()

    if url_hash and url_hash.long_url != url:
        random_str = "".join(
            random.choice(string.ascii_letters + string.digits) for _ in range(4)
        )
        key = blake2b(f"{url}{random_str}".encode(), digest_size=4).hexdigest()
    elif url_hash:
        return (
            jsonify(
                {
                    "key": f"{key}",
                    "long_url": new_link["url"],
                    "short_url": f"http://localhost:5000/{key}",
                }
            ),
            409,
        )

    item = {
        "key": f"{key}",
        "long_url": new_link["url"],
        "short_url": f"http://localhost:5000/{key}",
    }
    links.append(item)
    return jsonify(item), 201


if __name__ == "__main__":
    app.run(debug=True)
