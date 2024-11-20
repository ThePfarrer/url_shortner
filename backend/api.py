import random
import string
from hashlib import blake2b

from flask import Flask, redirect
from flask_restful import reqparse, fields, Resource, marshal_with, Api
from flask_sqlalchemy import SQLAlchemy

app = Flask(__name__)
api = Api(app)

db_name = "test_db.db"

app.config["SQLALCHEMY_DATABASE_URI"] = f"sqlite:///{db_name}"
app.config["SQLALCHEMY_TRACK_MODIFATIONS"] = False

db = SQLAlchemy()

db.init_app(app)


class URLModel(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    key = db.Column(db.String, unique=True)
    long_url = db.Column(db.String)
    short_url = db.Column(db.String)

    def __repr__(self):
        return f"URL(key= {self.key}, long_url={self.long_url}, short_url={self.short_url})"


url_args = reqparse.RequestParser()
url_args.add_argument("url", type=str, required=True, help="URL cannot be blank")

url_fields = {
    "key": fields.String,
    "long_url": fields.String,
    "short_url": fields.String,
}


class URLs(Resource):
    @marshal_with(url_fields)
    def get(self):
        result = URLModel.query.all()
        return result

    @marshal_with(url_fields)
    def post(self):
        args = url_args.parse_args()

        long_url = args["url"]
        key = blake2b(long_url.encode(), digest_size=4).hexdigest()
        url = URLModel.query.filter_by(key=key).first()

        if url and url.long_url != long_url:
            random_str = "".join(
                random.choice(string.ascii_letters + string.digits) for _ in range(4)
            )
            key = blake2b(f"{long_url}{random_str}".encode(), digest_size=4).hexdigest()
        elif url:
            return url, 409

        short_url = f"http://localhost:5000/{key}"

        url = URLModel(key=key, long_url=long_url, short_url=short_url)

        db.session.add(url)
        db.session.commit()
        return url, 201


class URL(Resource):
    @marshal_with(url_fields)
    def get(self, key):
        result = URLModel.query.filter_by(key=key).first()
        return result

    @marshal_with(url_fields)
    def delete(self, key):
        result = URLModel.query.filter_by(key=key).first()
        return result


api.add_resource(URLs, "/api/urls/")
api.add_resource(URL, "/api/urls/<string:key>")


@app.route("/")
def home():
    return "<h1>It works.</h1>"


@app.route("/<key>")
def url_redirect(key):
    result = URLModel.query.filter_by(key=key).first()
    if result:
        long_url = result.long_url
        return redirect(long_url, 302)
    return {"error": {"message": "URL not found"}}, 404


if __name__ == "__main__":
    app.run(debug=True)
