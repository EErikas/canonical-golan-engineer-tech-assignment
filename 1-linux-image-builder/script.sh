#!/bin/bash

# Define variables
IMAGE_SIZE=1G
KERNEL_VERSION=5.10.25
KERNEL_URL=https://cdn.kernel.org/pub/linux/kernel/v5.x/linux-${KERNEL_VERSION}.tar.xz
BUSYBOX_VERSION=1.34.1
BUSYBOX_URL=https://busybox.net/downloads/busybox-${BUSYBOX_VERSION}.tar.bz2

# Install necessary packages
sudo apt-get update
sudo apt-get install -y qemu-system-x86 nasm xorriso build-essential

# Create a working directory
mkdir -p linux-image
cd linux-image

# Download and extract kernel source code
wget $KERNEL_URL
tar -xf linux-${KERNEL_VERSION}.tar.xz
cd linux-${KERNEL_VERSION}

# Compile kernel
make defconfig
make -j$(nproc)
make bzImage

# Download and extract busybox source code
cd ..
wget $BUSYBOX_URL
tar -xf busybox-${BUSYBOX_VERSION}.tar.bz2
cd busybox-${BUSYBOX_VERSION}

# Compile busybox and install to a temporary directory
make defconfig
make -j$(nproc)
make CONFIG_PREFIX=../rootfs install

# Create initramfs
cd ..
mkdir initramfs
cd initramfs
mkdir -p {bin,sbin,etc,proc,sys,usr/{bin,sbin}}
cp -a ../rootfs/* .

# Create init script
cat > init << EOF
#!/bin/sh
mount -t proc none /proc
mount -t sysfs none /sys
mount -t devtmpfs none /dev
echo "Hello, world!"
/bin/sh
EOF
chmod +x init

# Create cpio archive
find . | cpio --quiet -H newc -o > ../initramfs.cpio
cd ..

# Create disk image
dd if=/dev/zero of=linux-image.img bs=1 count=0 seek=$IMAGE_SIZE
mkfs.ext4 linux-image.img

# Mount disk image to a temporary directory
mkdir -p rootfs
sudo mount -o loop linux-image.img rootfs

# Copy kernel and initramfs to the disk image
sudo mkdir -p rootfs/boot
sudo cp linux-${KERNEL_VERSION}/arch/x86_64/boot/bzImage rootfs/boot/
sudo cp initramfs.cpio rootfs/boot/

# Unmount disk image and remove temporary directories
sudo umount rootfs
rm -rf rootfs

# Run qemu
qemu-system-x86_64 \
    -m 512M \
    -kernel boot/bzImage \
    -initrd boot/initramfs.cpio \
    -append "console=ttyS0 root=/dev/sda rw" \
    -nographic
