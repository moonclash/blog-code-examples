from flask import Flask, session, jsonify
from flask_socketio import SocketIO

app = Flask(__name__)
app.secret_key = '1234xadas'
socketio = SocketIO(app)

def comp_helper(_dict):
  return {k : v for (k, v) in _dict.items()}

@app.route('/')
def index():
  session['start'] = True
  return 'hello'

@app.route('/cherries')
def cherries():
  session['fruit'] = 'cherries'
  return jsonify({'response': 'cherries', 'session': comp_helper(session)})

@app.route('/blueberries')
def blueberries():
  session['fruit'] = 'blueberries'
  return jsonify({'response': 'blueberries', 'session': comp_helper(session)})

@app.route('/destroy-session')
def destroy_session():
  session.clear()
  return jsonify({'response': 'session destroyed', 'session': comp_helper(session)})

if __name__ == '__main__':
  socketio.run(app, host='0.0.0.0', port=8000, debug=True)
