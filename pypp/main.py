import os

from scripts.clean import filter_data

if __name__ == "__main__":
    # Specify your input and output file names
    input_file = ".data/wo_data.json"
    output_file = ".data/cleaned_wo_data.json"

    # Ensures that the output directory exists
    output_dir = os.path.dirname(output_file)
    if not os.path.exists(output_dir):
        os.makedirs(output_dir)

    filter_data(
        input_filename=input_file,
        output_filename=output_file
    )
