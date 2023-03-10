#!/bin/bash

PKGNAME=lightalert
PKGDIR=/opt/$PKGNAME

cd cmd/LightAlert/ && CGO_ENABLED=0 go build -o ../../$PKGNAME .
cd ../../

umask 0022

mkdir -p $PKGDIR
cp $PKGNAME $PKGDIR/
cp configs/$PKGNAME.service $PKGDIR/
cp configs/$PKGNAME@.service $PKGDIR/
cp configs/install.sh $PKGDIR/

cd /opt
tar cvzf $PKGNAME-$1.tar.gz $PKGNAME
cd -
cp /opt/$PKGNAME-$1.tar.gz .