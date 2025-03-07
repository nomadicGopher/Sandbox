'use strict';
let userName = "Jon", param = `Hello ${userName}, TypeScript sandbox loaded.` // ES6 interpolate literal
;
// @ts-ignore
function initWithTsIgnore(localParam) {
    console.log(localParam);
}
document.addEventListener('DOMContentLoaded', () => {
    initWithTsIgnore(param);
});
