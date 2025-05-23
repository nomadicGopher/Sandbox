## Summary
The purpose of this project is to serve as an example of how to read 278 data into memory, transform some aspect of it, and then write the data out to a file in JSON format.

## Supplementary Files
* **Companion guide sourced from**: [https://www.cms.gov/files/document/esmd-x12n-278-companion-guide.pdf](https://www.cms.gov/files/document/esmd-x12n-278-companion-guide.pdf)
    * _Saved at: `.vscode/esMD X12N 278 Companion Guide AR2024.10.v18.0.pdf`_
* `/vscode/launch.json`: Debug launcher for VS Code to trace variables & set breakpoints.
* `/vscode/tasks.json`: Configuration to clean up test files, which is run as part of launch.json.
* `.vscode/test.sh`: Simpler method of testing without needing to debug (includes test file cleanup).
* **Sample input file**: `.vscode/hhpcr_x12n278.txt`
    * _Sourced from the companion guide._
* **Log file**: `inputFilePath/baseName + "_" + dateTime + ".log"`
* **Output file**: `inputFilePath/baseName + "_transformed_" + dateTime + ".json"`