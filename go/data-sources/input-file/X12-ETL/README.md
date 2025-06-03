## Summary
The purpose of this project is to serve as an example of how to read 278 data into memory, transform some aspect of it, and then write the data out to a file in JSON format.

## Input/Output Files
* **Sample input file**: `.vscode/hhpcr_x12n278.txt`
    * _Sourced from the companion guide & simplified to just the Envelope segments for demonstration purposes (ISA, GS, ST, Sample Transformed Body, SE, GE IEA)._
* **Log file**: `.vscode/inputFilePath/baseName + "_" + dateTime + ".log"`
* **Output file**: `.vscode/inputFilePath/baseName + "_transformed_" + dateTime + ".json"`

## Debugging Configurations & Test Script
* `.vscode/launch.json`: Debug launcher for VS Code to trace variables & set breakpoints.
* `.vscode/tasks.json`: Configuration to clean up test files, which is run as part of launch.json.
* `.vscode/test.sh`: Simpler method of testing without needing to debug (includes test file cleanup).

## Documentation Resource(s)
* **Companion guide sourced from**: [https://www.cms.gov/files/document/esmd-x12n-278-companion-guide.pdf](https://www.cms.gov/files/document/esmd-x12n-278-companion-guide.pdf)
    * _Saved to: `.vscode/esMD X12N 278 Companion Guide AR2024.10.v18.0.pdf`_
* [**X12 Tools & Resources**](https://github.com/Stedi/awesome-edi)
* [**Go Library for Struct Validation**](https://github.com/moov-io/x12) (Can make a PR or request for 278 data. Includes 835 & 837 structs.)