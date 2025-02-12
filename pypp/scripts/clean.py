import json


def filter_data(input_filename: str, output_filename: str) -> None:
    """
    Use this script to clean data from a JSON file. The input file should
    contain a list of objects, each with a "title" and "description" key.
    The output file will contain a list of objects with only the "title" and
    "description" keys.
    """
    # Read the JSON data from the input file
    with open(input_filename, "r", encoding="utf-8") as infile:
        data = json.load(infile)

    # Create a new list with only the title and description for each element
    cleaned = []
    for item in data:
        # Using .get() so that if the key is missing, it defaults to an empty
        # string
        cleaned_item = {
            "title": item.get("title", ""),
            "description": item.get("description", "")
        }
        cleaned.append(cleaned_item)

    # Write the cleaned data to the output file
    with open(output_filename, "w", encoding="utf-8") as outfile:
        json.dump(cleaned, outfile, indent=2, ensure_ascii=False)

    print(f"Cleaned data written to: {output_filename}")


def remove_stop_words() -> None:
    pass
