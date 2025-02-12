# Pypp

Python data pre-processor

## Development

### Python

```shell
# If necessary, upgrade your version of pip.
# If python3 does not work here try: py or python,
# and same goes for any of the other scripts going
# forward.
python3 -m pip install --upgrade pip
# Install virtualenv package
python3 -m pip install virtualenv
```

### Virtual environment

```shell
# Once the virtual environment has been sourced
# you can deactivate it or quit it by running:
deactivate
```

#### Unix/MacOS

```shell
# Create the virtual environment
python3 -m venv .venv &&
# Activate the environment
source .venv/bin/activate
```

#### Windows

```shell
# Create the virtual environment
python -m venv .venv &&
# Activate the environment
.venv\Scripts\activate.bat
# Alternatively run this if you are using Powershell
# myenv\Scripts\Activate.ps1
```

### Dependencies

```shell
# Make sure the virtual environment has been sourced,
# before running!
pip install -r requirements.txt
```

```shell
# Add installed dependencies to requirements.txt
pip freeze > requirements.txt
```

### Usage

```shell
# Run program using:
python3 main.py
```
