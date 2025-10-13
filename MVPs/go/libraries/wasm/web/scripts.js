"use-strict";

const go = new Go();
let wasmModule,
  wasmInstance,
  wasmReady = false;

WebAssembly.instantiateStreaming(fetch("go.wasm"), go.importObject)
  .then((wasm) => {
    wasmModule = wasm.module;
    wasmInstance = wasm.instance;
    go.run(wasmInstance);
    wasmReady = true;
    console.log("js: WASM is ready");
  })
  .catch((err) => {
    console.error(err);
  });

(async () => {
  while (!wasmReady) {
    await new Promise((resolve) => setTimeout(resolve, 10));
  }
  let statement = globalThis.helloDom();
  alert(statement)
  wasmInstance = await WebAssembly.instantiate(wasmModule, go.importObject); // Re-instantiate the module to reset its state
})();
