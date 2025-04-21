# TypeScript Sandbox

This is a minimal TypeScript sandbox environment for learning purposes.

## Setup

1. Ensure you have Node.js installed. You can download it from [nodejs.org](https://nodejs.org).
2. Initialize a new npm project: `npm init -y`
3. Setup a package.json. While linting and formatting are not necessary with IDEs like VS Code + Extensions, it is helpful to enforce standardization across development platforms. Minification is good for production use.

    ```jsonc
    {
      "name": "sandbox",
      "version": "1.0.0",
      "description": "A sandbox project for TypeScript development",
      //"main": "js/main.js", // Specifies the entry point for the application. Popular options include "index.js" or "lib/main.js".
      "type": "module", // Defines the module system. module or commonjs (default)
      "scripts": {
        "clean": "rm -rf js",
        "build": "npm run clean && tsc",
        "dev": "npm run clean && tsc --watch",
        "test": "http-server -c-1 .",
        "minify": "terser dist/**/*.js -o dist --compress --mangle",
        "lint": "eslint 'src/**/*.{ts,tsx}'",
        "lint:fix": "eslint 'src/**/*.{ts,tsx}' --fix",
        "format": "prettier --write 'src/**/*.{ts,tsx,js,json,md}'",
        "check-deps": "depcheck",
        "test": "jest",
        "test:watch": "jest --watch",
        "test:coverage": "jest --coverage",
        "docs": "jsdoc -c jsdoc.json"
      },
      "devDependencies": {
        //,"@types/packageName": "VERSION", // If package is JS & not TS (doest't include type definitions), the use @types for IDE to recognize the types.
        "typescript": "^5.8.2",
        "eslint": "^9.22.0",
        "http-server": "^14.1.1",
        "depcheck": "^1.4.7",
        "eslint-config-prettier": "^10.1.1",
        "eslint-plugin-prettier": "^5.2.3",
        "prettier": "^3.5.3",
        "terser": "^5.39.0",
        "jest": "^29.7.0",
        "jsdoc": "^4.0.4"
      }
    }
    ```

4. Setup a `tsconfig.json` file to configure the TypeScript compiler:

    ```json
    {
      "compilerOptions": {
        "target": "es6", //  This option specifies the JavaScript version that TypeScript should compile to.
        "module": "es6", // This option defines the module system that TypeScript should use when generating JavaScript code.
        "strict": true, // Enable all strict type-checking options.
        "rootDir": "ts", // Root directory for TypeScript source files.
        "outDir": "js" // Output compiled JavaScript files to the dist directory.
        // "allowJS": true, // Allow JS files to exist in the rootDir (src)
        //"esModuleInterop": true,
        //"skipLibCheck": true,
        //"forceConsistentCasingInFileNames": true
      },
      "include": ["ts/**/*"]
    }
    ```

5. Install the necessary dependencies with `npm install` or individually like:

    ```sh
    npm install --save-dev typescript
    npm install --save-dev eslint
    npm install --save-dev http-server
    ```

    When installing dependencies, npm will automatically add the latest version to your package.json. If you need to specify a particular version, you can do so by appending the version number to the package name, like this: `npm install --save-dev typescript@4.0.0`

    _If dependencies are already defined you can simply use `npm install` & optionally `npm update` for the latest versions._\*
    _`npm prune` can be used to clean up node_modules if they are not listed as a core package in `packages.json` or are a dependency of those core packages._

## Usage

1. Write your TypeScript code in `main.ts`.
2. Compile the TypeScript code to JavaScript: `npm `
3. Open localhost:port provided by the server details in the terminal

## Notes

- `//@ts-ignore`: Place this 1 line before the code which should ignore TypeScript rules. This should be avoided.
- to run scripts found in package.json use npm run ...
  - The tsc command can be used to build typescript files independently

## Resources

- [typescriptlang.org > Basic Types](https://www.typescriptlang.org/docs/handbook/basic-types.html)
- [TS Docs > Handbook (TS in 5 minutes)](https://www.typescriptlang.org/docs/handbook/typescript-in-5-minutes.html)
- [Playground](https://www.typescriptlang.org/play/?#code/PTAEHUFMBsGMHsC2lQBd5oBYoCoE8AHSAZVgCcBLA1UABWgEM8BzM+AVwDsATAGiwoBnUENANQAd0gAjQRVSQAUCEmYKsTKGYUAbpGF4OY0BoadYKdJMoL+gzAzIoz3UNEiPOofEVKVqAHSKymAAmkYI7NCuqGqcANag8ABmIjQUXrFOKBJMggBcISGgoAC0oACCbvCwDKgU8JkY7p7ehCTkVDQS2E6gnPCxGcwmZqDSTgzxxWWVoASMFmgYkAAeRJTInN3ymj4d-jSCeNsMq-wuoPaOltigAKoASgAywhK7SbGQZIIz5VWCFzSeCrZagNYbChbHaxUDcCjJZLfSDbExIAgUdxkUBIursJzCFJtXydajBBCcQQ0MwAUVWDEQC0gADVHBQGNJ3KAALygABEAAkYNAMOB4GRonzFBTBPB3AERcwABS0+mM9ysygc9wASmCKhwzQ8ZC8iHFzmB7BoXzcZmY7AYzEg-Fg0HUiQ58D0Ii8fLpDKZgj5SWxfPADlQAHJhAA5SASPlBFQAeS+ZHegmdWkgR1QjgUrmkeFATjNOmGWH0KAQiGhwkuNok4uiIgMHGxCyYrA4PCCJSAA)
