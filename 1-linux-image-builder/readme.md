# Linux Image Builder

This script downloads and compiles a Linux kernel and busybox, creates an initramfs, creates a disk image, copies the kernel and initramfs to the disk image, and runs QEMU with the disk image as a virtual hard drive. The QEMU instance boots into the initramfs and prints "hello world" to the console. The script does not require any user input, except for the initial installation of necessary packages that may require superuser privileges. The script runs within the working directory and does not consume any other locations on the host file system.

## Usage
1. navigate to directory
2. Make the script executable with the following command:
    ```
    chmod +x script.sh
    ```
3. Launch script with root privileges
    ```
    sudo ./script.sh
    ```