#!/usr/bin/env bash

# Mock out textmate dialog application to test internals:
# This script mimicks the textmate dialog application such that returned output
# will be formatted correctly for a specific test case. If the script
# receives input that is not expected for this testcase, the script will return
# a non-zero status to make the test case fail.

EXPECT_POPUP_CMD='popup --suggestions ({display = "choice 1;"; insert = "you picked choice 1";},)'
EXPECT_IMAGES_CMD='images --register { test = "/tmp/test.png"; test2 = "/tmp/test2.pdf"; }'

case $1 in
"popup")
	echo "Popup command provided" 1>&2
	if [ "$*" == "${EXPECT_POPUP_CMD}" ]; then
		echo -n "choice 1;you picked choice 1"
		exit 0
	fi
	echo "invalid command" 1>&2
	exit 1
	;;
"images")
	echo "Images command provided: $@" 1>&2
	if [ "$*" == "${EXPECT_IMAGES_CMD}" ]; then
		exit 0
	fi
	echo "invalid command" 1>&2
	exit 1
	;;
*)
	echo "Unexpected subcommand: $@"
	exit 1
	;;
esac
