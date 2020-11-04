from dotenv import load_dotenv
import os
from flask import Flask, Blueprint, render_template

# client_bp = Blueprint('client_app', __name__,
#                       url_prefix='',
#                       static_url_path='',
#                       static_folder='./dist/',
#                       template_folder='./dist/',
#                       )

# os.chdir(os.path.join(os.path.dirname(__file__), 'dist'))

dotenv_path = os.path.join(os.path.dirname(
    __file__), '.env')  # Path to .env file
load_dotenv(dotenv_path)

app = Flask(__name__, static_folder='dist',
            template_folder='dist')
# app.register_blueprint(client_bp)


@app.route('/', defaults={'path': ''})
@app.route('/<path:path>')
def index(path):
    return render_template('index.html')


app.host = '0.0.0.0'
app.port = int(os.getenv("PORT", 5000))  # or 5000
app.debug = "on"
if __name__ == '__main__':
    app.run()
