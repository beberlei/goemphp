#!/usr/bin/env bash
set -e

PHP_VER=7.3
DIR=php-7.3.11
TAR=http://php.net/get/${DIR}.tar.bz2/from/this/mirror

WORK_DIR=php-srcs
mkdir -p $WORK_DIR
[ -d $WORK_DIR/$DIR ] && echo "FATAL: $DIR already existed. Please remove $DIR manually and run $0 again." && exit

pushd .
cd $WORK_DIR

wget -O php.tar.bz2 $TAR
tar jxvf php.tar.bz2
rm php.tar.bz2
cd $DIR 
./configure --enable-embed
make -j4
popd
echo $PHP_VER > php-version
unlink php-lib
ln -s $WORK_DIR/$DIR php-lib

echo "Congratulations!!!"
