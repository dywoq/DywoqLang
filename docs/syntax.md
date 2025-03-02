# Syntax:
The language must support types from C# like ```string```, ```byte```, ```short``` and etc.
All declarations must use `create` keyword.
  
## Function declaration:
```
create function() {
  return 0;
} must_return ulong;
```
  
Whether a function returns or not, function declaration can be shorten:
1. ```create function() => 0 must_return ulong;```
2. ```create function() => some_action();```

## Variable declaration:
```
create x: 2 as ulong;
create y: 3 as ulong;
```
  
To make variables apply multiple types, we can use ```mix_with``` keyword:
```
create x: 2 mix_with ulong, byte, uint;
```
  
## Typealias declaration:
```
create typealias: byte as int8;
create typealias: short as int16;
```
  
They can be used to variables:
```
create x: 2 as int8;
```
