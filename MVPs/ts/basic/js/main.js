'use strict';
import * as Functions from './functions.js';
let test1 = {
    userName: "Jon",
};
let test2 = {
    userName: "",
    param: `Hello ${test1.userName}, TypeScript sandbox loaded.`
};
document.addEventListener('DOMContentLoaded', () => {
    var _a;
    Functions.FuncName((_a = test2.param) !== null && _a !== void 0 ? _a : `Undefined Placeholder via nullish coalescing operator`);
    Functions.FuncName(test1.param ? test1.param : `Undefined Placeholder via ternary operator`);
    Functions.FuncName(test2.userName);
});
