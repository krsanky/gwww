#!/bin/ksh

. /etc/rc.d/rc.subr

D=/home/wise/go/src/oldcode.org/gow

daemon="/home/wise/go/bin/gow"

gow_user="wise"



rc_cmd $1
