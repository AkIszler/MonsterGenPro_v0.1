from flask import Flask, request, jsonify

app = Flask(__name__)

characters = ["test", "500"]

@app.route('/api/characters', methods=['POST'])
def receive_characters():
    global characters
    data = request.get_json()
    characters = data
    return jsonify({"message": "Characters received"}), 200

@app.route('/api/characters', methods=['GET'])
def get_characters():
    return jsonify(characters), 200

if __name__ == '__main__':
    app.run(port=5000)