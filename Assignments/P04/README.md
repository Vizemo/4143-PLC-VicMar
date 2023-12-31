## Program 4 - Concurrent Image Downloader
### Victor Marchesi
### Description:

The goal of this assignment is to understand and implement basic concurrency in Go. You will write a program that concurrently downloads a set of images from given URLs and saves them to disk. By comparing the time taken to download images sequentially vs. concurrently, you will observe the benefits of concurrency for I/O-bound tasks.


### Files

#### 1.x is specifying the files are in a folder like 2.x is in another folder, etc.

|    #    | Files | Description |
|  :---:  | -------- | -------------------------------- |
|    1    | [main.go](./main.go) | P04 goLang Concurrent Image Downloader main driver |
|    2    | [image_1.jpg](./image_1.jpg) | P04 picture 1 that are royalty free from Dr. Griffin |
|    3    | [image_2.jpg](./image_2.jpg) | P04 picture 2 that are royalty free from Dr. Griffin |
|    4    | [image_3.jpg](./image_3.jpg) | P04 picture 3 that are royalty free from Dr. Griffin |
|    5    | [image_4.jpg](./image_4.jpg) | P04 picture 4 that are royalty free from Dr. Griffin |
|    6    | [image_5.jpg](./image_5.jpg) | P04 picture 5 that are royalty free off of [https://www.pexels.com/](https://www.pexels.com/) |
|    7    | [image_6.jpg](./image_6.jpg) | P04 picture 6 that are royalty free off of [https://www.pexels.com/](https://www.pexels.com/) |
|    8    | [image_7.jpg](./image_7.jpg) | P04 picture 7 that are royalty free off of [https://www.pexels.com/](https://www.pexels.com/) |
|    9    | [image_8.jpg](./image_8.jpg) | P04 picture 8 that are royalty free off of [https://www.pexels.com/](https://www.pexels.com/) |
|    10   | [image_9.jpg](./image_9.jpg) | P04 picture 9 that are royalty free off of [https://www.pexels.com/](https://www.pexels.com/) |


### Instructions

- This program requres the fmt, io, net/http, os, path/filepath, time, packages
- To run the program path over to the folder/repository and run the command "go run main.go"

### Speed Sequential vs. Concurrent

- Sequential download took: 1.38 seconds
- Concurrent download took: 213.86 milliseconds