'use strict';

import * as Functions from './functions.js';
import type * as Types from './types.js';

let test1: Types.Info = {
  userName: "Jon",
};

let test2: Types.Info = {
  userName: "",
  param: `Hello ${test1.userName}, TypeScript sandbox loaded.`
};

document.addEventListener('DOMContentLoaded', () => {  
  Functions.FuncName(test2.param ?? `Undefined Placeholder via nullish coalescing operator`);
  Functions.FuncName(test1.param ? test1.param : `Undefined Placeholder via ternary operator`);
  Functions.FuncName(test2.userName);
});
