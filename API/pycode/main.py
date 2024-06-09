from flask import Flask, request, jsonify
import os
import pandas as pd

app = Flask(__name__)

characters = ["test", "500"]

def get_path():
    path = "../characters.xlsx"
    # Check if the file exists
    if os.path.isfile(path):
        return path
    else:
        raise FileNotFoundError(f"The file {path} does not exist.")

@app.route('/api/characters', methods=['POST'])
def receive_characters():
    global characters
    data = request.get_json()
    characters = data
    return jsonify({"message": "Characters received"}), 200

@app.route('/api/characters', methods=['GET'])
def get_characters():
    try:
        file_path = get_path()
        df = pd.read_excel(file_path)
        characters_from_excel = df.to_dict(orient='records')
        return jsonify(characters_from_excel), 200
    except FileNotFoundError as e:
        return jsonify({"error": str(e)}), 404
if __name__ == '__main__':
    app.run(port=5000)