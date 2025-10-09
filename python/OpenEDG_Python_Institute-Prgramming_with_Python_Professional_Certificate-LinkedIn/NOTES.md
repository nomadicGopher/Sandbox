## Jupyter Lab
* To run:
    * Open a termnal session in a root folder
    * Run `Jupyter Lab`
* _When a cell's margin is clicked and the color is blue, it is in command mode_
* _When a cell's inner box is clicked and the color is green, it is in edit mode_
* Enter a new cell above (command mode): `a`
* Enter a new cell bellow:
    * When in command mode: `b`
    * When in edit mode: `Shift + Enter`
* Delete and existing cell (command mode): Type `d` twice
* Markdown mode to type text (command mode): `m`
    * _supports equations_

## Python
## Functions
* print()
* var.append(value)

## Variables & Types
* To acquire the type of an object, you can use `type(var)`
* vraiable names more common to use `_`
### complex (imaginary numbers)
are the square root of negative numbers. In algebra commonly represented as i, in python are represented as `j`. 
### str (string)
* can use `'` (more common) or `"`
* **concatenation**: `'1' + '1'` == '11'
### int (whole numbers)
### float (decimal)
### bool (True / False)

## Data Structures
### Lists (slice)
* type == list
* starts at 0
* reserves aditional memory to accmodate for appended items
* **samples**:
    * `sample_list1 = []`
    * `sample_list2 = [1,2,3,4]`
    * `sample_list3 = [sample_list2,true,'yes']`
* `len(sample_list)` == 4
* when comparing, ordering of values matters
    * `[1,2] == [1,2]` == true
    * `[1,2] == [2,1]` == false
### Sets
* smilar to lists but reduced to only unique values
* type == set
* **sample**: sample_set = {1,1,2,2} == {1, 2} where len() == 2
* when comparing, ordering of values doesn't matters. Its looking for the inclusion of rather than only equals
    * `{1,2} == {1,2}` == true
    * `{1,2} == {2,1}` == true
    * `{1,2} == {3,2,1,0}` == true    
### Tuples
* similar to lists but cannot be appended to and values cannot be modified
    * used in return values
* saves memory because only required amout of memory is allocated
* type == tuple
* **samples**:
    * `sample_tuple1 = ()`
    * `sample_tuple2 = (1,2,3,4)`
    * `sample_tuple = (sample_list2,true,'yes')`
* `len(sample_tuple2)` == 4
* when comparing, ordering of values matters
    * `(1,2) == (1,2)` == true
    * `(1,2) == (2,1)` == false
### Dictionaries
* **sample**:
```
sample_dict = {
    'name1': 'value',
    'name2': 2,
    'name1': 'overwrites name 1 value'
}
```
sample_dict['name1'] == 'overwrites name 1 value'

## Classes
* Begin with upper case letter