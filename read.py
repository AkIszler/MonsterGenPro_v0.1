import os
import glob
import pandas as pd

def find_spreadsheet_files(directory):
    # Find all .xlsx files in the specified directory
    return glob.glob(os.path.join(directory, "*.xlsx"))

def read_spreadsheets_to_dict(directory, default_value=0):
    # Find all spreadsheet files in the directory
    spreadsheet_files = find_spreadsheet_files(directory)
    
    # Dictionary to store the data from each spreadsheet
    all_data = {}

    # Iterate over each file
    for file in spreadsheet_files:
        # Read the spreadsheet into a DataFrame
        df = pd.read_excel(file)
        
        # Fill NaN values with the default value
        df.fillna(default_value, inplace=True)
        
        # Convert the DataFrame to a list of dictionaries, one per row
        file_data = df.to_dict(orient='records')
        
        # Use the filename (without path) as the key
        filename = os.path.basename(file)
        all_data[filename] = file_data
    
    return all_data

# Directory to search for spreadsheet files (current directory in this case)
current_directory = os.getcwd()

# Get the dictionary containing data from all spreadsheets
all_data = read_spreadsheets_to_dict(current_directory)

# Print the data
for filename, data in all_data.items():
    print(f"Data from {filename}:")
    for record in data:
        print(record)
    print()


