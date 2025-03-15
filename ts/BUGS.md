**packages.json > "start": "npm run build && node dist/main.js",**

```
  CORP+jvojtush@APPF5F6NLC MINGW64 ~/go/src/github.com/nomadicGopher/Sandbox/ts ((angular)learn_angular)
$ npm run start

> sandbox@1.0.0 start
> npm run build && node dist/main.js


> sandbox@1.0.0 build
> tsc

node:internal/modules/esm/resolve:275
    throw new ERR_MODULE_NOT_FOUND(
          ^

Error [ERR_MODULE_NOT_FOUND]: Cannot find module 'C:\Users\jvojtush\go\src\github.com\nomadicGopher\Sandbox\ts\dist\functions' imported from C:\Users\jvojtush\go\src\github.com\nomadicGopher\Sandbox\ts\dist\main.js
    at finalizeResolution (node:internal/modules/esm/resolve:275:11)
    at moduleResolve (node:internal/modules/esm/resolve:932:10)
    at defaultResolve (node:internal/modules/esm/resolve:1056:11)
    at ModuleLoader.defaultResolve (node:internal/modules/esm/loader:654:12)
    at #cachedDefaultResolve (node:internal/modules/esm/loader:603:25)
    at ModuleLoader.resolve (node:internal/modules/esm/loader:586:38)
    at ModuleLoader.getModuleJobForImport (node:internal/modules/esm/loader:242:38)
    at ModuleJob._link (node:internal/modules/esm/module_job:135:49) {
  code: 'ERR_MODULE_NOT_FOUND',
  url: 'file:///C:/Users/jvojtush/go/src/github.com/nomadicGopher/Sandbox/ts/dist/functions'
}

Node.js v22.13.1
```