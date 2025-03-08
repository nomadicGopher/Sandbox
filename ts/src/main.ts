'use strict';

import { Test } from './interfaces';
import { initWithTsIgnore } from './functions';

let test1 : Test = {
  userName: "Jon",
};

let test2 : Test = {
  userName: "", // Required
  param: `Hello ${test1.userName}, TypeScript sandbox loaded.` // ES6 interpolate literal. Backtics are required to use. Seprated into a second object because test1.name needs to be instanciated before it can be used in another object.
};

document.addEventListener('DOMContentLoaded', () => {  
  for (let i = 1; i <= 2; i++) {
   console.log(`Iteration ${i}`)
  }

  initWithTsIgnore(test2.param ?? `Undefined Placeholder via nullish coalescing operator`); // provide a default value when dealing with null or undefined. It checks if the left-hand side operand is null or undefined, and if so, it returns the right-hand side operand. Otherwise, it returns the left-hand side operand. Use when you want to provide a fallback value for null or undefined. Only considers null and undefined as nullish values.
  initWithTsIgnore(test2.param ? test2.param : `Undefined Placeholder via ternary operator`); // shorthand for an if-else statement. It evaluates a condition and returns one of two values based on whether the condition is true or false. formatted as (condition ? expr1 : expr2) where condition is the expression to evaluate, expr1 is the expression to return if condition is true and expr2 is the expression to return if condition is false. Use when you need to choose between two values based on a condition. Considers all falsy values (false, 0, "", null, undefined, NaN).
});
