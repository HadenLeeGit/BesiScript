# BesiScript

BesiScript is a programming language for the game *Besiege* 

BesiScript abstracts Automation Blocks in Besiege, making it easy for you to call them through functions with parameters. And directly compiled to a game save file `.bsg` through the compiler.

## How to use it?

You can save your BesiScript code in `input.txt`, put it in the same folder as the compiler, run the compiler, and the compiled `.bsg` file will appear in the same folder

Or, Use the terminal, followed by the files to be compiled as parameters.

```bash
BesiScriptComplier.exe ./mydev.txt
```



# Documentation

1. Initialization Functions
2. Comment
3. Variable
4. Logic Functions
5. Instrument Functions



## Initialization Functions

All initialization functions are **optional**, and **default values** will be used for undefined initialization functions.

You can use them to personalize your initialization parameters.

```jsx
name(projectName)

// example
name(Aircraft 1)

// Default values if you don't use it
name(output)
```

Use the name function to name your project and compiled file name.



```jsx
init(X, Y, Z)

// example
init(1, 2, 1)

// Default values if you don't use it
init(0, 5, 0)
```

Use this function to define the initial position of blocks.



```jsx
baseblock(true)

// Default values if you don't use it
baseblock(false)
```

Define whether to generate basic blocks.



```jsx
logicinit(X, Y, Z)

// Default values if you don't use it
logicinit(0.1,0.1,0.1)
```

Define the scale of LogicGate Block



```jsx
instrumentinit(X, Y, Z)

// Default values if you don't use it
instrumentinit(1,1,1)
```

Define the scale of Instrument Blocks (sensor, timer, altimeter, speedometer, anglemeter).



## Comment

```c
// single line comment

name(demo) // name the project

/*
	multi-line
	comment
*/
```

Comments are not compiled, making it easier for you to code.



## Variable

```jsx
// structure
$varableName = "message" //for message
$varableName = A //for keyboard keys

// Example
$emulate = "move" // single message
$actions = "up;down;left;right" // multi message

$inputKey = Z //single key
$control = W+A+S // multi keys, Max is 3 keys

//use
and($control, $actions, $emulate, true)
```

Use variables to define the message or collections of messages.
As well as keys and collections of keys.



## Logic Functions

```jsx
and(INPUT_A, INPUT_B, EMULATE, TOGGLE_MODE)

or(INPUT_A, INPUT_B, EMULATE, TOGGLE_MODE)

nor(INPUT_A, INPUT_B, EMULATE, TOGGLE_MODE)

nand(INPUT_A, INPUT_B, EMULATE, TOGGLE_MODE)

xor(INPUT_A, INPUT_B, EMULATE, TOGGLE_MODE)

xnor(INPUT_A, INPUT_B, EMULATE, TOGGLE_MODE)

random(INPUT, EMULATE, TOGGLE_MODE)

srlatch(INPUT_A, INPUT_B, EMULATE)

dlatch(INPUT_A, INPUT_B, EMULATE)

counter(INPUT_A, INPUT_B, EMULATE)

edge(INPUT, EMULATE, INVERTED)

not(INPUT, EMULATE, TOGGLE_MODE)

// example
$btn = A + B
and($btn, C, "AAA", true)
or($btn, "ON", Z, false)
edge(W+A, "MSG", true)
```



## Instrument Functions



### Sensor Block

```jsx
// automatic sensor
autosensor(DISTANCE, RADIUS, EMULATE, REVERSE_MODE, IGNORE_STATIC)

// Use key to detect 
sensor(DISTANCE, RADIUS, EMULATE, REVERSE_MODE, IGNORE_STATIC, ACTIVATE, HOLD_TO_ACTIVATE)

// example
$msg = "OUTPUT"
autosensor(5, 0.5, $msg, true, false)
sensor(5, 0.5, $msg, true, false, T, false)
```



### Timer Block

```jsx
// automatic timer
autotimer(WAIT, TIME, EMULATE, LOOP_MODE)

// hold to run timer
htrtimer(WAIT, TIME, EMULATE, LOOP_MODE, ACTIVATE)

// Use key to activate
timer(WAIT, TIME, EMULATE, LOOP_MODE, ACTIVATE, CAN_STOP)

// example
autotimer(5, 5, "OUTPUT1;OUTPUT2", true)
htrtimer(5, 5, "OUTPUT", true, A)
timer(5, 5, A+B, true, B, true)
```



### Altimeter Block

```jsx
// automatic
autoalti(HEIGHT, EMULATE, REVERSE_MODE)

// Use key to activate
alti(HEIGHT, EMULATE, REVERSE_MODE, ACTIVATE, HOLD_TO_ACTIVATE)

// example
autoalti(5, "OUTPUT", false)
alti(5, "OUTPUT", false, B, false)
```



### Speedometer Block

```jsx
// automatic
autospeed(SPEED, EMULATE, REVERSE_MODE)

// Use key to activate
speed(SPEED, EMULATE, REVERSE_MODE, ACTIVATE, HOLD_TO_ACTIVATE)

// example
autospeed(5, "EMULATE", false)
speed(5, "EMULATE", false, B, false)
```



### Anglemeter Block

```jsx
// automatic
autoangle(START_ANGLE, END_ANGLE, EMULATE, REVERSE_MODE)

// Use key to activate
angle(START_ANGLE, END_ANGLE, EMULATE, REVERSE_MODE, ACTIVATE, HOLD_TO_ACTIVATE)

// example
autoangle(45, -45, "EMULATE", false)
angle(45, -45, "EMULATE", false, B, false)
```