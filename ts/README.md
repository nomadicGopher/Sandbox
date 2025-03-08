# TypeScript Sandbox

This is a minimal TypeScript sandbox environment for learning purposes.

## Setup

### Config explanations

#### package.json

While linting and formatting are not necessary with IDEs like VS Code + Extensions, it is helpful to enforce standardization across development platforms. Minification is good for production use.

```jsonc
{
  "name": "sandbox",
  "version": "1.0.0",
  "description": "A sandbox project for TypeScript development",
  //"main": "dist/index.js", // Specifies the entry point for the application. Popular options include "index.js" or "lib/main.js".
  //"type": "module", // Defines the module system. Options are "module" for ES modules and "commonjs" for CommonJS modules.
  "scripts": {
    "build": "tsc --watch", // Command to compile TypeScript to JavaScript.
    "minify": "terser dist/**/*.js -o dist --compress --mangle", // Command using additional dependency to compress js file in dist.
    "start": "http-server -c-1 .", // Starts a web server disabling cache for a fresh file.
    //"start": "node dist/index.js", // Command to start the application. Can be customized to use tools like nodemon for auto-restart.
    "lint": "eslint 'src/**/*.{ts,tsx}'",
    "lint:fix": "eslint 'src/**/*.{ts,tsx}' --fix",
    "format": "prettier --write 'src/**/*.{ts,tsx,js,json,md}'",
    "check-deps": "depcheck"
  },
  // "dependencies": {
  //   "packageName": "VERSION"
  // },
  "devDependencies": {
    "typescript": "^4.0.0",
    "depcheck": "^1.4.7",
    "eslint": "^7.0.0",
    "eslint-config-prettier": "^8.0.0",
    "eslint-plugin-prettier": "^4.0.0",
    "http-server": "^14.1.1",
    "prettier": "^2.0.0",
    "terser": "^5.0.0"
    //,"@types/packageName": "VERSION" // If package is JS & not TS (doest't include type definitions), the use @types for IDE to recognize the types.
  }
}
```

#### tsconfig.json

```jsonc
{
  "compilerOptions": {
    "target": "ES6", // Target ECMAScript 6 for the output.
    "module": "ES6", // Use ES6 modules for enhanced performance in the browser. When using node.js use CommonJS.
    "outDir": "dist", // Output compiled JavaScript files to the dist directory.
    "rootDir": "src", // Root directory for TypeScript source files.
    "strict": true // Enable all strict type-checking options.
    // "allowJS": true // Allow JS files to exist in the rootDir (src)
  },
  "include": [
    "src/**/*" // Include all TypeScript files in the src directory and subdirectories.
  ]
}
```

## Steps

1. Ensure you have Node.js installed. You can download it from [nodejs.org](https://nodejs.org).
2. Install TypeScript globally using npm: `npm install -g typescript`
3. Initialize a new npm project: `npm init -y`
4. Install the necessary dependencies:
    ```sh
    npm install --save-dev typescript
    npm install --save-dev terser
    npm install --save-dev http-server
    npm install --save-dev eslint
    npm install --save-dev eslint-config-prettier
    npm install --save-dev eslint-plugin-prettier
    npm install --save-dev prettier
    npm install --save-dev depcheck
    ```
    When installing dependencies, npm will automatically add the latest version to your package.json. If you need to specify a particular version, you can do so by appending the version number to the package name, like this: `npm install --save-dev typescript@4.0.0`

    _If dependencies are already defined you can simply use `npm install` & optionally `npm update` for the latest versions._*
    _`npm prune` can be used to clean up node_modules if they are not listed as a core package in `packages.json` or are a dependency of those core packages._
5. Create a `tsconfig.json` file to configure the TypeScript compiler:
    ```json
    {
      "compilerOptions": {
        "target": "es5",
        "module": "commonjs",
        "strict": true,
        "esModuleInterop": true,
        "skipLibCheck": true,
        "forceConsistentCasingInFileNames": true
      }
    }
    ```

## Usage

1. Write your TypeScript code in `scripts.ts`.
2. Compile the TypeScript code to JavaScript: `npm start`
    * _Leaving the terminal running without using `ctr+c` will watch for code changes and automatically apply them._
3. Open localhost:port provided by the server details in the terminal