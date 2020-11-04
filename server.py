from dotenv import load_dotenv
import os
from flask import Flask, Blueprint, render_template

client_bp = Blueprint('client_app', __name__,
                      url_prefix='',
                      static_url_path='',
                      static_folder='dist',
                      template_folder='dist',
                      )

# os.chdir(os.path.join(os.path.dirname(__file__), 'dist'))

dotenv_path = os.path.join(os.path.dirname(
    __file__), '.env')  # Path to .env file
load_dotenv(dotenv_path)


@client_bp.route('/', defaults={'path': ''})
@client_bp.route('/<path:path>')
def index(path):
    return render_template('index.html')


@client_bp.errorhandler(404)
def page_not_found(e):
    return render_template('index.html')


app = Flask(__name__)

app.register_blueprint(client_bp)

app.host = '0.0.0.0'
app.port = int(os.getenv("PORT", 5000))  # or 5000
app.debug = "on"
if __name__ == '__main__':
    app.run()
