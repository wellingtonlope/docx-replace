# docx-replace
## About
A simple command-line project to use a .csv database to create some .docx documents based on a template.

## How to build
```bash
go build
```

## How to use
You will need two files a .docx that will be the template and a .csv that will be the database.

### About the template
The template must contain identifiers that will be replaced by the data contained in the CSV.

### About the databse
The database in csv, in its first line must contain the identifiers (the ones that were inserted in the template) and the following lines must contain the data that will be inserted in the template.

### Example
Inside the test_data folder is a sample template and database.
