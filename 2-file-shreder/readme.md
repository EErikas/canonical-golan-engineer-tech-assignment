# Shred Application
The solution is build around a function called `Shred()` which overwrites the file for 3 times before deleting it

## Running the Application
1. Clone the repository to your local machine
2. Open a terminal window and navigate to the directory where you cloned the repository
3. Run the following command to build the application:
    ```
    go build main.go
    ```
4. To run the application, use the following command:
   ```
    ./main /path/to/file
   ```
   where `/path/to/file` is the path to the file you want to shred. Note that the file will be deleted after it has been shredded.

## Testing the Application
1. Clone the repository to your local machine
2. Open a terminal window and navigate to the directory where you cloned the repository
3. Run the following command to run the tests:
    ```
    go test
    ```

## Notes About the Solution
The `Shred()` function can be useful in situations where you need to securely delete a file that may contain sensitive information, such as passwords or financial data. By overwriting the file multiple times with random data before deleting it, you can make it much more difficult for anyone to recover the original contents of the file.

One advantage of using this function is that it's a simple and convenient way to securely delete a file. You don't need to manually overwrite the file or use a separate tool to securely delete it. Just calling the `Shred()` function with the file path as an argument will take care of everything for you.

However, one potential drawback of using this function is that it's not foolproof. While overwriting the file multiple times with random data makes it much harder to recover the original contents, it's still theoretically possible for someone with enough knowledge and resources to recover some or all of the data. Additionally, if the file is stored on a solid-state drive (SSD) instead of a traditional hard drive, the function may not be as effective at securely deleting the data due to the way SSDs handle data storage.

Overall, the `Shred()` function can be a useful tool for securely deleting files in certain situations, but it's important to understand its limitations and use it in combination with other security measures when necessary.