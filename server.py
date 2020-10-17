from flask import Flask, render_template
import os


app = Flask(__name__,
            static_folder=os.path.join('./dist'),
            template_folder=os.path.join('./dist'))


@app.route('/', defaults={'path': ''})
@app.route('/<path:path>')
def index(path):
    return render_template('index.html')


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
