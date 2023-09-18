## Program 2 - Baby Steps
### Victor Marchesi
### Description:

In this program we will try and show some advantages of Go when compared to C++ and Python. Every high level language has its strengths, and I think Golang's can be seen in its similarity to both aforementioned languages, but more so in its library availability. Not as strong as Python's but more so than C++'s. I would like us to create Go's version of a class, that will do some rudimentary image manipulation.


### Files

|    #    | Files    | Description                      |
| :---: | -------- | -------------------------------- |
|    1    | [imagemod](./imagemod/) | P02 goLang mascot workspace folder |
|    2    | [imageManipulator](./imageManipulator) | The imageManipulator.go package folder. |
|    3    | [imageManipulator.go](./imageManipulator/imageManipulator.go) | The mascot go package that has a function that returns the best mascot |
|    4    | [go.mod](./imagemod/go.mod) | .mod file where goLang knows where to pull the package from |
|    5    | [go.sum](./imagemod/go.sum) | .sum file that pulls where the package dependencies come from |
|    6    | [main.go](./imagemod/main.go) | The main.go file that runs can modify .png files using the fogleman.gg libraries |
|    7    | [mustangs_edited.png](./imagemod/mustangs_edited.png) | The mustangs_edited.png file created by our first version of the go program |
|    8    | [mustangs.png](./imagemod/mustangs.png) | The mustangs.png file created by our second version with the added constructor of the go program |




### Instructions

- This program requres the "fmt" and "rsc.io" libraries

### Example Command

- None for now. Can run a test case in the mascot_test.go file
